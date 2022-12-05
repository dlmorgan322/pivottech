package main

type Character struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RosterRepsonse struct {
	Code            int    `json:"code"`
	Status          string `json:"status"`
	Copyright       string `json:"copyright"`
	AttributionText string `json:"attributionText"`
	AttributionHTML string `json:"attributionHTML"`
	Etag            string `json:"etag"`
	Data            struct {
		Offset     int `json:"offset"`
		Limit      int `json:"limit"`
		Total      int `json:"total"`
		Count      int `json:"count"`
		Characters []Character
	} `json:"data"`
}
