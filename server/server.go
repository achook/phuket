package server

import "phuket/database"

type Server struct {
	db *database.Database
}

func NewServer() (*Server, error) {
	d, err := database.NewDatabase()
	if err != nil {
		return nil, err
	}

	s := new(Server)
	s.db = d

	return s, nil
}
