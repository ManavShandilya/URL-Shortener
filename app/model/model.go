package model

type Goly struct {
	Id       uint64 `json:"id" gorm:"primarykey"`
	Redirect string `json:"redirect" gorm:"not null"`
	Goly     string `json:"goly" gorm:"unique;not null"`
	Clicked  uint64 `json:"clicked"`
	Random   bool   `json:"random"`
}
