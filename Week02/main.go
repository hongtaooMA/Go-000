package main

import (
	"database/sql"
	"github.com/pkg/errors"
	"log"
)

func someSqlFunc(uuid string) error {
	return sql.ErrNoRows
}

func dao(uuid string) error {
	err := someSqlFunc(uuid)
	if err != nil {
		return errors.Wrap(err, "dao failed for uuid" + uuid)
	}
	return nil
}

func service(uuid string) error {
	return dao(uuid)
}

func main()  {
	err := service("24t2oivgf232ogvnhska")
	if err != nil {
		log.Printf("service failed, error: %+v", err)
	}
}




