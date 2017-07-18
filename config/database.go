package config

import (
	"fmt"
	"log"

	r "gopkg.in/gorethink/gorethink.v3"
)

var session *r.Session

func InitDb() *r.Session {
	var err error
	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "pokemon",
	})

	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
		return session
	}

	res, err := r.Expr("Conexion correcta!").Run(session)
	if err != nil {
		log.Fatalln(err)
	}

	var response string
	err = res.One(&response)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(response)
	res.Close()
	return session
}
