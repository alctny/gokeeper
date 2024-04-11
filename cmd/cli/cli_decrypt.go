package main

import (
	"fmt"

	"github.com/alctny/frieren-keeper/model"
	"github.com/urfave/cli/v2"
)

// decrypt all passwors, loginId, bind
func decryptAll(ctx *cli.Context) error {
	keyFile := ctx.Path("key")
	passwords := []model.Password{}
	tx := db.Where(&model.Password{Encrypt: 1}).Find(&passwords)
	if tx.Error != nil {
		return tx.Error
	}
	var err error
	passwords, err = model.DecryptPasswords(passwords, keyFile)
	if err != nil {
		return err
	}
	tx = db.Begin()
	for _, p := range passwords {
		result := tx.Save(&p)
		if result.Error != nil {
			return result.Error
		}
	}
	tx = tx.Commit()
	if tx.Error != nil {
		return tx.Error
	}
	fmt.Println("decrypt success")

	return nil
}
