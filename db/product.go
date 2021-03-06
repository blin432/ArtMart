package db 

import (
	"log"
	
	pg	"github.com/go-pg/pg"
)

type Product struct {
	ID int `sql:"id,serial,pk"`
	Product_ID	string `json:"product_id" sql:"product_id"` 
	Title string `json:"title" sql:"title"`
	Product_Type string `json:"product_type" sql:"product_type"`
	Img   string `json:"img" sql:"img"`
	Price int `json:"price" sql:"price"`
	Artist string `json:"artist" sql:"artist"`
	Info string `json:"info" sql:"info"`
	Count int `json:"count" sql:"count"`
	Total int `json:"total" sql:"total"`
	User_ID int `json:"user_id" sql:"user_id"`

}

func (p *Product) AddToCart(db *pg.DB) error{
	
	defer db.Close()
	err := db.Insert(p)
	if err!= nil {
		log.Printf("Error: %v\n",err)
		return err
	}
	return nil
}

