package model

type Config struct {
	PlaceApiKey      string  `yaml:"placeApiKey"`
	ImgUrlPrefix     string  `yaml:"imgUrlPrefix"`
	BangsueLat       float64 `yaml:"bangsueLat"`
	BangsueLng       float64 `yaml:"bangsueLng"`
	ParamFoodType    string  `yaml:paramFoodType`
	PlaceBaseUrl     string  `yaml:placeBaseUrl`
	PlaceRelativeUrl string  `yaml:placeRelativeUrl`
}
