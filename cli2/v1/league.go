package poker

import (
	"encoding/json"
	"fmt"
	"io"
)

/*
NewLeague Takes a JSON from IO reader and returns a player.
Passes error in case of an error with the JSON Decoder
*/
func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league %v", err)
	}
	return league, err
}
