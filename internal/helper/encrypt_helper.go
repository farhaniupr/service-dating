package helper

import (
	"github.com/farhaniupr/dating-api/package/library"
	"golang.org/x/crypto/bcrypt"
)

type EcncryptHelper struct {
	library.Database
}

func ModuleEncryptHelper(db library.Database, env library.Env) EcncryptHelper {
	return EcncryptHelper{
		Database: db,
	}
}

func (d EcncryptHelper) EncryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (d EcncryptHelper) CheckPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
