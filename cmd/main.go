package main

import (
	"log"
	"scheduler/internal/app"
)

func  main()  {
	err := app.Run()
	log.Fatal(err)


	/*var re = regexp.MustCompile(`{id}`)
	str := "users/{id}"
	//url := "/users/15"
	url := "/users"
	reg:= re.ReplaceAllString(str, `\d+`)

	isMatch, err := regexp.MatchString(reg, url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(url)
	fmt.Println(reg)
	fmt.Println(isMatch)*/
}
