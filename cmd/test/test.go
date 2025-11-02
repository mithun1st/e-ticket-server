package main

import (
	"e-ticket/internal/config"
	appdatabase "e-ticket/pkg/database"
	appenviroment "e-ticket/pkg/enviroment"
	"e-ticket/pkg/utils"
	"fmt"
	"log"
)

func main() {
	var tem int = 11

	var s1 *int = nil
	var s2 *int = &tem

	str := fmt.Sprintf("hello %s, hello1 %s", utils.NilToStrDB(s1), utils.NilToStrDB(s2))
	fmt.Println(str)
	return

	var alph []string = []string{"a", "b", "c"}
	var nums []int = []int{20, 30, 40}
	// fmt.Println(strings.Join(alph, "+"))

	// fmt.Println(nums[1])

	fmt.Println(nums, alph)
	fmt.Println(utils.JoinArray(nums))
	fmt.Println(utils.JoinArray(alph))
	return

	appenviroment.Set(appenviroment.Development)
	//* Load Config
	appConfig, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	//* Initialize Database
	db, err := appdatabase.NewPostgresDb(appConfig.Db.DbConnedtionString())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)

	// repo := authrepository.NewAuthRepository(db)
	// data, err := repo.FindUser("mithun.2121@yahoo.com", "", "87654321")
	// data, err := repo.FindCompaniesByOwner(1)
	// fmt.Println(err)
	// fmt.Println(data)
}
