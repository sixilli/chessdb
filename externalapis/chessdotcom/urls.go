package chessdotcom

import "fmt"

const urlBase = "https://api.chess.com"

// Request requires games/{YYYY}/{MM}
func buildPlayerGamesUrl(player, year, month string) string {
	return fmt.Sprintf("%s/pub/player/%s/games/%s/%s", urlBase, player, year, month)
}
