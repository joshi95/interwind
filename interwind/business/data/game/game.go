package game

import "log"

type Store struct {
	log *log.Logger
}

func NewStore(log *log.Logger) *Store {
	return &Store{log: log}
}

func (s *Store) AllGames() []Game {
	return []Game{
		{
			ID: "1234",
		},
		{
			ID: "666666",
		},
	}
}
