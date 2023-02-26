package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Binbiubiubiu/cnpmcore/repository"
)

func main() {
	var repo = &repository.PackageRepository{}
	// var newB = entity.CreatePackage(&entity.Package{
	// 	Scope: "hi",
	// 	Name:  "world",
	// })
	// // newB.ID = 48
	list, err := repo.FindPackage("hi", "world")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", list.AbbreviatedsDist)

	bs, _ := json.Marshal(list)
	var out bytes.Buffer
	json.Indent(&out, bs, "", "\t")
	fmt.Printf("student=%v\n", out.String())
}
