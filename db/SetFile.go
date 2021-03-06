package db

import (
	"errors"
	"github.com/number571/gopeer"
	"github.com/number571/hiddenlake/models"
	"github.com/number571/hiddenlake/settings"
)

func SetFile(user *models.User, file *models.File) error {
	id := GetUserId(user.Username)
	if id < 0 {
		return errors.New("User id undefined")
	}
	_, err := settings.DB.Exec(
		"INSERT INTO File (IdUser, Hash, Name, PathName, Size, Encr) VALUES ($1, $2, $3, $4, $5, $6)",
		id,
		file.Hash,
		gopeer.Base64Encode(
			gopeer.EncryptAES(
				user.Auth.Pasw,
				[]byte(file.Name),
			),
		),
		gopeer.Base64Encode(
			gopeer.EncryptAES(
				user.Auth.Pasw,
				[]byte(file.Path),
			),
		),
		file.Size,
		file.Encr,
	)
	if err != nil {
		panic("exec 'setfile' failed")
	}
	return nil
}
