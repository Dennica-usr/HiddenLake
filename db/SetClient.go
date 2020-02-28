package db

import (
	"errors"
	"github.com/number571/gopeer"
	"github.com/number571/hiddenlake/models"
	"github.com/number571/hiddenlake/settings"
)

func SetClient(user *models.User, client *models.Client) error {
	if client.ThrowClient == nil {
		client.ThrowClient = client.Public
	}

	if gopeer.HashPublic(client.Public) != client.Hashname {
		return errors.New("hashname is not derived from the public key")
	}

	id := GetUserId(user.Username)
	if id < 0 {
		return errors.New("User id undefined")
	}

	_, err := settings.DB.Exec(
		"DELETE FROM Client WHERE IdUser=$1 AND Hashname=$2",
		id,
		client.Hashname,
	)
	if err != nil {
		panic("exec 'setclient.delete' failed")
	}

	_, err = settings.DB.Exec(
		"INSERT INTO Client (IdUser, Hashname, Address, PublicKey, ThrowClient, Certificate) VALUES ($1, $2, $3, $4, $5, $6)",
		id,
		client.Hashname,
		gopeer.Base64Encode(
			gopeer.EncryptAES(
				user.Auth.Pasw,
				[]byte(client.Address),
			),
		),
		gopeer.StringPublic(client.Public),
		gopeer.StringPublic(client.ThrowClient),
		client.Certificate,
	)
	if err != nil {
		panic("exec 'setclient.insert' failed")
	}

	return nil
}
