package main

import (
	u "github.com/italorfeitosa/mongo-migrations-with-go/utils"
	"github.com/italorfeitosa/mongo-migrations-with-go/infra"
)

func main(){
	db := &infra.DB{
		URI: "mongodb://root:mongo@mongo:27017/",
	}
	db.Init()
}