package chessdotcom

type GameResponse struct {
	Games []Game `json:"games"`
}

type Game struct {
	Url         string
	Pgn         string
	TimeControl string
	EndTime     uint64
	Rated       bool
	Fen         string
	Time_class  string
	Rules       string
	White       struct {
		Rating   int
		Result   string
		Id       string `json:"@id"`
		Username string
	}
	Black struct {
		Rating   int
		Result   string
		Id       string `json:"@id"`
		Username string
	}
}
