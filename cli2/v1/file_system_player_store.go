package poker

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {

	err := initPlayerDBFile(file)
	if err != nil {
		return nil, err
	}
	league, err := NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}
	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}, nil
}

func initPlayerDBFile(file *os.File) error {
	file.Seek(0, 0)
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}
	return nil
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
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) (wins int) {
	player := f.league.Find(name)
	if player != nil {
		return player.Wins
	}
	return
}

func (f *FileSystemPlayerStore) RecordWin(name string) {

	player := f.league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}
	f.database.Encode(f.league)
}
