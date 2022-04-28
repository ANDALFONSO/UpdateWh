package service

import (
	"UPDATE_WH/entity"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/mercadolibre/fury_go-core/pkg/rusty"
)

var (
	BaseUrlFlex   = "https://internal-api.mercadolibre.com/shipping/selfservice/adoption/search"
	BaseUrlWh     = "https://internal-api.mercadolibre.com/shipping/working-hours"
	BaseUrlUpdate = "https://internal-api.mercadolibre.com/test/shipping/working-hours/services/"
	SITES         = []string{"MLA", "MLB", "MLC", "MLU", "MLM", "MCO"}
)

type IUpdateService interface {
	UpdateWh(ctx context.Context, process bool) string
}

type RustyClient interface {
	Do(*http.Request) (*http.Response, error)
}

type UpdateService struct {
	Client RustyClient
}

func NewPS(client RustyClient) IUpdateService {
	return &UpdateService{
		client,
	}
}

func (u *UpdateService) UpdateWh(ctx context.Context, process bool) string {
	fmt.Println("--------------START------------------")
	workingHours := fectServiceWH(u.Client, ctx)
	serviceBySite := fecthServiceFlex(u.Client, ctx)
	idsToProcess := sortDataProcessed(workingHours, serviceBySite)
	return processData(u.Client, ctx, process, idsToProcess)
}

func fectServiceWH(Client RustyClient, ctx context.Context) []int {
	fmt.Println("Trallendo servicios WH")
	cl, err := rusty.NewEndpoint(Client, BaseUrlWh, rusty.WithHeader("Content-Type", "application/json"))
	if err != nil {
		fmt.Println("Error:", err)
		return []int{}
	}
	res, err := cl.Get(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return []int{}
	}

	responseWh := entity.ResponseWh{}
	json.Unmarshal(res.Body, &responseWh)

	var ids []int
	for _, v := range responseWh.Values {
		id, err := strconv.Atoi(v.Id)
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}
	return ids
}

func fecthServiceFlex(Client RustyClient, ctx context.Context) map[string][]int {
	fmt.Println("Trallendo servicios flex")
	m := make(map[string][]int)
	for _, site := range SITES {
		object := entity.Request{
			Type: "scroll",
			Query: entity.Query{
				AnyEquals: []entity.AnyEquals{
					{
						Path:   "services.status.id",
						Values: []string{"in", "pending"},
					},
				},
				Equals: []entity.Equals{
					{
						Path:  "is_test",
						Value: false,
					},
					{
						Path:  "site_id",
						Value: site,
					},
				},
			},
			Size:      100,
			ContextId: "",
		}
		ids := fetchIdsBySite(Client, ctx, site, object)
		fmt.Printf("site:%v  rows: %d\n", site, len(ids))
		m[site] = ids
	}
	return m
}

func fetchIdsBySite(Client RustyClient, ctx context.Context, site string, object entity.Request) []int {
	query := []entity.Response{}
	query1 := serviceFlexHttp(Client, ctx, object)
	tope := query1.Total / 100
	query = append(query, query1)

	for i := 0; i < tope; i++ {
		object.ContextId = query1.ContextId
		query2 := serviceFlexHttp(Client, ctx, object)
		query = append(query, query2)
	}
	var ids []int
	for _, q := range query {
		for _, d := range q.Documents {
			for _, s := range d.Services {
				ids = append(ids, s.Id)
			}
		}
	}
	return ids
}

func serviceFlexHttp(Client RustyClient, ctx context.Context, object interface{}) entity.Response {
	cl, err := rusty.NewEndpoint(Client, BaseUrlFlex, rusty.WithHeader("Content-Type", "application/json"))
	if err != nil {
		fmt.Println("Error:", err)
		return entity.Response{}
	}
	res, err := cl.Post(ctx, rusty.WithBody(object))
	if err != nil {
		fmt.Println("Error:", err)
		return entity.Response{}
	}
	response := entity.Response{}
	json.Unmarshal(res.Body, &response)
	return response
}

func sortDataProcessed(workingHours []int, serviceBySite map[string][]int) []int {
	fmt.Println("Buscando coincidencias..")
	var idsToProcess []int
	var idsNotProcess []int
	for _, id := range workingHours {
		for _, site := range serviceBySite {
			if !contains(site, id) {
				idsToProcess = append(idsToProcess, id)
			} else {
				idsNotProcess = append(idsNotProcess, id)
			}
		}
	}
	fmt.Println("A procesar:", len(idsToProcess))
	fmt.Println("No procesar:", len(idsNotProcess))
	return idsToProcess
}

func contains(id []int, searchterm int) bool {
	i := sort.SearchInts(id, searchterm)
	return i < len(id) && id[i] == searchterm
}

func processData(Client RustyClient, ctx context.Context, process bool, idsToProcess []int) string {
	fmt.Println("Procesando datos..", process)
	var idsWithErrors []int
	if process {
		for _, id := range idsToProcess {
			if err := processIds(Client, ctx, id); err != nil {
				idsWithErrors = append(idsWithErrors, id)
			}
		}
	}
	fmt.Println("Total de id con errores:", len(idsWithErrors))
	fmt.Println("---------------END-------------------")
	if process && len(idsWithErrors) > 0 {
		return "Proceso terminado con errores"
	}
	return "Proceso terminado exitosamente!!"
}

func processIds(Client RustyClient, ctx context.Context, id int) error {
	day := entity.Day{
		Ranges: []string{"0000-2359"},
	}
	object := entity.RequestWh{
		Id:        strconv.Itoa(id),
		Monday:    day,
		Tuesday:   day,
		Wednesday: day,
		Thursday:  day,
		Friday:    day,
		Saturday:  day,
		Sunday:    day,
	}

	return updateServiceWh(Client, ctx, object)
}

func updateServiceWh(Client RustyClient, ctx context.Context, object entity.RequestWh) error {
	cl, err := rusty.NewEndpoint(Client, BaseUrlUpdate+"/"+object.Id,
		rusty.WithHeader("Content-Type", "application/json"),
		rusty.WithHeader("X-Caller-Scopes", "admin"))
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	_, err = cl.Put(ctx, rusty.WithBody(object))
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	return nil
}
