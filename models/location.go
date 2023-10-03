package models

type Country struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	IsoCode2         string `json:"iso_code_2"`
	IsoCode3         string `json:"iso_code_3"`
	IsoNumberCode    int    `json:"iso_numeric_code"`
	AddressFormat    string `json:"address_format"`
	PostCodeRequired int8   `json:"postcode_required"`
	PhoneCode        int    `json:"phonecode"`
	Ordering         int    `json:"ordering"`
	Status           int8   `json:"status"`
}

type State struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type CityAndLocation struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}
