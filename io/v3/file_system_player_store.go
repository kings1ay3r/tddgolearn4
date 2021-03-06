package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

type League []Player

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}

func (f *FileSystemPlayerStore) GetLeague() League {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}
func (f *FileSystemPlayerStore) GetPlayerScore(name string) (wins int) {
	player := f.GetLeague().Find(name)
	if player != nil {
		return player.Wins
	}
	return
}
func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.Find(name)
	if player != nil {
		player.Wins++
	}
	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)
}
