package model

type Lists struct {
	Lists []List `json:"lists"`
}

type List struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
}
