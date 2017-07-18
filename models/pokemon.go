package models

type Pokemon struct {
	Id    string   `json:"id" binding:"required" gorethink:"id"`
	Name  string   `json:"name" binding:"required" form:"name" gorethink:"name"`
	Type  []string `json:"type" binding:"required" form:"type" gorethink:"type"`
	Photo string   `json:"photo" binding:"required" form:"photo" gorethink:"photo"`
}

type PokemonPost struct {
	Name  string   `json:"name" binding:"required" form:"name" gorethink:"name"`
	Type  []string `json:"type" form:"type" gorethink:"type"`
	Photo string   `json:"photo" form:"photo" gorethink:"photo"`
}


type Image struct {
    Filename    string `json:"Filename" form:"Filename" gorethink:"Filename"`
    ContentType string `json:"ContentType" form:"ContentType" gorethink:"ContentType"`
    Data        string `json:"Data" form:"Data" gorethink:"Data"`
    Size        int `json:"Size" form:"Size" gorethink:"Size"`
}