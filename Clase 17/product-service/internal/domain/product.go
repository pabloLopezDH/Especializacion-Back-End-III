package domain

type Product struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Price   int    `json:"price"`
	IconUrl string `json:"iconUrl"`
}
