package models

type GetAllListsResponse struct {
	Data []ProductResponse `json:"data"`
}

type GetAllWareHousesResponse struct {
	Data []WareHouse `json:"data"`
}
