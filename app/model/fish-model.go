package model

type FishModel struct {
	Id          int      `json:"id"`
	Image       string   `json:"image"`
	Name        string   `json:"name"`
	Habitat     []string `json:"habitat"`
	Description string   `json:"description"`
	Price       uint64   `json:"price"`
}
