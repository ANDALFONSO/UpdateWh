package service

import (
	"UPDATE_WH/entity"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/mercadolibre/fury_go-core/pkg/rusty"
)

var (
	BaseUrlFlex   = "https://internal-api.mercadolibre.com/shipping/selfservice/adoption/search"
	BaseUrlWh     = "https://internal-api.mercadolibre.com/shipping/working-hours"
	BaseUrlUpdate = "https://internal-api.mercadolibre.com/test/shipping/working-hours/services/"
	SITES         = []string{"MLA", "MLB", "MLC", "MLU", "MLM", "MCO", "MPE"}
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
	fmt.Println("Tamaño:", len(workingHours))
	serviceBySite := fecthServiceFlex(u.Client, ctx)
	idsToProcess := sortDataProcessed(workingHours, serviceBySite)
	return processData(u.Client, ctx, process, idsToProcess)
}

func fectServiceWH(Client RustyClient, ctx context.Context) []entity.RequestWh {
	fmt.Println("Trayendo servicios WH")
	cl, err := rusty.NewEndpoint(Client, BaseUrlWh, rusty.WithHeader("Content-Type", "application/json"))
	if err != nil {
		fmt.Println("Error:", err)
		return []entity.RequestWh{}
	}
	res, err := cl.Get(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return []entity.RequestWh{}
	}

	responseWh := entity.ResponseWh{
		/*Values: []entity.Values{
			entity.Values{
				Id: "1",
				Monday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Tuesday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Wednesday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Thursday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Friday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Saturday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Sunday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
			},
			entity.Values{
				Id: "2",
				Monday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Tuesday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Wednesday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Thursday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Friday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Saturday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Sunday: entity.Day{
					Ranges: []string{"0900-1800", "0700-0800"},
				},
			},
			entity.Values{
				Id: "3",
				Monday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Tuesday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Wednesday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Thursday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Friday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Saturday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Sunday: entity.Day{
					Ranges: []string{},
				},
			},
			entity.Values{
				Id: "4",
				Monday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Tuesday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Wednesday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Thursday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Friday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Saturday: entity.Day{
					Ranges: []string{"0900-1800", "0700-0800"},
				},
				Sunday: entity.Day{
					Ranges: []string{},
				},
			},
			entity.Values{
				Id: "5",
				Monday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Tuesday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Wednesday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Thursday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Friday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Saturday: entity.Day{
					Ranges: []string{},
				},
				Sunday: entity.Day{
					Ranges: []string{},
				},
			},
			entity.Values{
				Id: "6",
				Monday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Tuesday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Wednesday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Thursday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Friday: entity.Day{
					Ranges: []string{"0900-1800", "0700-0800"},
				},
				Saturday: entity.Day{
					Ranges: []string{},
				},
				Sunday: entity.Day{
					Ranges: []string{},
				},
			},
			entity.Values{
				Id: "7",
				Monday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Tuesday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Wednesday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Thursday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Friday: entity.Day{
					Ranges: []string{},
				},
				Saturday: entity.Day{
					Ranges: []string{},
				},
				Sunday: entity.Day{
					Ranges: []string{},
				},
			},
			entity.Values{
				Id: "8",
				Monday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Tuesday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Wednesday: entity.Day{
					Ranges: []string{},
				},
				Thursday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Friday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Saturday: entity.Day{
					Ranges: []string{},
				},
				Sunday: entity.Day{
					Ranges: []string{},
				},
			},
			entity.Values{
				Id: "9",
				Monday: entity.Day{
					Ranges: []string{"0700-1800"},
				},
				Tuesday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Wednesday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Thursday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Friday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Saturday: entity.Day{
					Ranges: []string{},
				},
				Sunday: entity.Day{
					Ranges: []string{},
				},
			},
			entity.Values{
				Id: "10",
				Monday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Tuesday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Wednesday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Thursday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Friday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Saturday: entity.Day{
					Ranges: []string{"0700-1800"},
				},
				Sunday: entity.Day{
					Ranges: []string{},
				},
			},
			entity.Values{
				Id: "11",
				Monday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Tuesday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Wednesday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Thursday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Friday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Saturday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Sunday: entity.Day{
					Ranges: []string{"0700-1800"},
				},
			},
			entity.Values{
				Id: "12",
				Monday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Tuesday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Wednesday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Thursday: entity.Day{
					Ranges: []string{"0700-1800"},
				},
				Friday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Saturday: entity.Day{
					Ranges: []string{"0900-1800"},
				},
				Sunday: entity.Day{
					Ranges: []string{"0700-1800"},
				},
			},
		},*/
	}
	json.Unmarshal(res.Body, &responseWh)

	var ids []entity.RequestWh
	for _, v := range responseWh.Values {
		sc := scheduler(v.Id, v.Monday, v.Tuesday, v.Wednesday, v.Thursday, v.Friday, v.Saturday, v.Sunday)
		ids = append(ids, sc)
	}

	return ids
}

func scheduler(id string, monday entity.Day, tuesday entity.Day, wednesday entity.Day, thursday entity.Day, friday entity.Day, saturday entity.Day, sunday entity.Day) entity.RequestWh {
	sc := entity.RequestWh{
		Id: id,
	}
	if len(monday.Ranges) != 1 || (len(monday.Ranges) == 1 && monday.Ranges[0] != "0900-1800") {
		sc.Monday = monday
	} else {
		sc.Monday = entity.Day{
			Ranges: []string{"0000-2359"},
		}
	}
	if len(tuesday.Ranges) != 1 || (len(tuesday.Ranges) == 1 && tuesday.Ranges[0] != "0900-1800") {
		sc.Tuesday = tuesday
	} else {
		sc.Tuesday = entity.Day{
			Ranges: []string{"0000-2359"},
		}
	}
	if len(wednesday.Ranges) != 1 || (len(wednesday.Ranges) == 1 && wednesday.Ranges[0] != "0900-1800") {
		sc.Wednesday = wednesday
	} else {
		sc.Wednesday = entity.Day{
			Ranges: []string{"0000-2359"},
		}
	}
	if len(thursday.Ranges) != 1 || (len(thursday.Ranges) == 1 && thursday.Ranges[0] != "0900-1800") {
		sc.Thursday = thursday
	} else {
		sc.Thursday = entity.Day{
			Ranges: []string{"0000-2359"},
		}
	}
	if len(friday.Ranges) != 1 || (len(friday.Ranges) == 1 && friday.Ranges[0] != "0900-1800") {
		sc.Friday = friday
	} else {
		sc.Friday = entity.Day{
			Ranges: []string{"0000-2359"},
		}
	}
	if len(saturday.Ranges) != 1 || (len(saturday.Ranges) == 1 && saturday.Ranges[0] != "0900-1800") {
		sc.Saturday = saturday
	} else {
		sc.Saturday = entity.Day{
			Ranges: []string{"0000-2359"},
		}
	}
	if len(sunday.Ranges) != 1 || (len(sunday.Ranges) == 1 && sunday.Ranges[0] != "0900-1800") {
		sc.Sunday = sunday
	} else {
		sc.Sunday = entity.Day{
			Ranges: []string{"0000-2359"},
		}
	}

	return sc
}

func fecthServiceFlex(Client RustyClient, ctx context.Context) map[string][]int {
	fmt.Println("Trayendo servicios flex")
	m := make(map[string][]int)
	var total int
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
		total = total + len(ids)
	}
	fmt.Println("Total:", total)
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

func sortDataProcessed(workingHours []entity.RequestWh, serviceBySite map[string][]int) []entity.RequestWh {
	fmt.Println("Buscando coincidencias..")
	var idsToProcess []entity.RequestWh
	var idsNotProcess []entity.RequestWh
	for _, wh := range workingHours {
		id, err := strconv.Atoi(wh.Id)
		if err != nil {
			fmt.Println("Error transformando id:", wh.Id)
			continue
		}
		var contain bool = true
		for _, site := range serviceBySite {
			if searchItem(site, id) {
				contain = false
			}
		}
		if contain {
			idsToProcess = append(idsToProcess, wh)
		} else {
			idsNotProcess = append(idsNotProcess, wh)
		}
	}
	fmt.Println("A procesar:", len(idsToProcess))
	fmt.Println("No procesar:", len(idsNotProcess))
	return idsToProcess
}

func searchItem(data []int, searchterm int) bool {
	for _, v := range data {
		if v == searchterm {
			return true
		}
	}
	return false
}

func processData(Client RustyClient, ctx context.Context, process bool, idsToProcess []entity.RequestWh) string {
	fmt.Println("Procesando datos..", process)
	var idsWithErrors []entity.RequestWh
	if process {
		for _, wh := range idsToProcess {
			if err := updateServiceWh(Client, ctx, wh); err != nil {
				idsWithErrors = append(idsWithErrors, wh)
			}
		}
	}
	fmt.Println("Total de id con errores:", len(idsWithErrors))
	fmt.Println("---------------END-------------------")
	if process && len(idsWithErrors) > 0 {
		fmt.Println("Errores:", idsWithErrors)
		return "Proceso terminado con errores"
	}
	return "Proceso terminado exitosamente!!"
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
