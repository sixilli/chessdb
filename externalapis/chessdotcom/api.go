package chessdotcom

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetPlayerGames(player, year, month string) ([]Game, error) {
	url := buildPlayerGamesUrl(player, year, month)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	resp.Body.Close()

	gameResponse := GameResponse{}
	err = json.Unmarshal(body, &gameResponse)
	if err != nil {
		return nil, err
	}

	return gameResponse.Games, nil
}
