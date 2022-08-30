package main

import (
	"fmt"
	"interface/service"
	"sync"
)

func main() {
	var db []*service.User
	userSvc := service.NewUserService(db)
	names := []string{"inas", "thalia", "ricky", "fisa"}

	for _, n := range names {
		addNames := userSvc.Register(&service.User{Nama: n})
		fmt.Println(addNames)
	}

	output := userSvc.GetUser()
	fmt.Println("----------------Hasil Get User--------------")
	for _, v := range output {
		fmt.Println(v.Nama)
	}
}

func cetakNama(wg *sync.WaitGroup, nama string) {
	fmt.Println(nama)
	wg.Done()
}
