package password

import (
	"encoding/base64"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string, error) {
	passwordBytes := []byte(password)
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)
	if err != nil {
		return "", errors.Wrap(err, "error hashing password")
	}
	hashedPassword := base64.StdEncoding.EncodeToString(hashedPasswordBytes)
	return hashedPassword, nil
}

func Validate(password string, hashedPassword string) (bool, error) {
	passwordBytes := []byte(password)
	hashedPasswordBytes, err := base64.StdEncoding.DecodeString(hashedPassword)
	if err != nil {
		return false, errors.Wrap(err, "error decoding password hash")
	}
	err = bcrypt.CompareHashAndPassword(hashedPasswordBytes, passwordBytes)
	switch err {
	case nil:
		return true, nil
	case bcrypt.ErrMismatchedHashAndPassword:
		return false, nil
	default:
		return false, errors.Wrap(err, "error validating password")
	}
}
