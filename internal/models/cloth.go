package models

type Cloth struct {
	Id   int
	Name string
	Type string
}

type ClothInsertUpdate struct {
	Id     int
	Name   string
	TypeId int
}
