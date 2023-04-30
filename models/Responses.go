package models

type GetAllListsResponse struct {
	Data []Product `json:"data"`
}

type GetAllWareHousesResponse struct {
	Data []WareHouse `json:"data"`
}
