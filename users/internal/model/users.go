package model

import "github.com/alexedwards/argon2id"

type User struct {
	ID            string   `json:"id"`
	Username      string   `json:"username"`
	Password      password `json:"-"`
	Email         string   `json:"email"`
	EmailVerified bool     `json:"email_verified"`
	IsGuest       bool     `json:"is_guest"`
}

type password struct {
	plainText *string
	hash      *string
}

func (p *password) Matches(plainTextPassword *string) (bool, error) {
	return argon2id.ComparePasswordAndHash(*plainTextPassword, *p.hash)
}

func (p *password) Set(plainTextPassword string) error {
	params := &argon2id.Params{
		Memory:      12 * 1024,
		Iterations:  3,
		Parallelism: 1,
		SaltLength:  16,
		KeyLength:   32,
	}

	hash, err := argon2id.CreateHash(plainTextPassword, params)
	if err != nil {
		return err
	}

	p.hash = &hash
	p.plainText = &plainTextPassword
	return nil
}
