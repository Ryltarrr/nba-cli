package parser

type Results struct {
	Meta       Meta       `json:"meta"`
	Scoreboard Scoreboard `json:"scoreboard"`
}

type Meta struct {
	Version int    `json:"version"`
	Request string `json:"request"`
	Time    string `json:"time"`
}

type Scoreboard struct {
	GameDate   string `json:"gameDate"`
	LeagueName string `json:"leagueName"`
	Games      []Game `json:"games"`
}

type Game struct {
	GameId            string  `json:"gameId"`
	GameCode          string  `json:"gameCode"`
	GameStatus        int     `json:"gameStatus"`
	GameStatusText    string  `json:"gameStatusText"`
	Period            int     `json:"period"`
	GameTimeUTC       string  `json:"gameTimeUTC"`
	RegulationPeriods int     `json:"regulationPeriods"`
	SeriesGameNumber  string  `json:"seriesGameNumber"`
	SeriesText        string  `json:"seriesText"`
	GameLeaders       Leaders `json:"gameLeaders"`
	TeamLeaders       Leaders `json:"teamLeaders"`
	HomeTeam          Team    `json:"homeTeam"`
	AwayTeam          Team    `json:"awayTeam"`
}

type Leaders struct {
	HomeLeaders Player `json:"homeLeaders"`
	AwayLeaders Player `json:"awayLeaders"`
}

type Player struct {
	PersonId    int     `json:"personId"`
	Name        string  `json:"name"`
	PlayerSlug  string  `json:"playerSlug"`
	JerseyNum   string  `json:"jerseyNum"`
	Position    string  `json:"position"`
	TeamTricode string  `json:"teamTricode"`
	Points      float32 `json:"points"`
	Rebounds    float32 `json:"rebounds"`
	Assists     float32 `json:"assists"`
}

type Team struct {
	TeamId            int      `json:"teamId"`
	TeamName          string   `json:"teamName"`
	TeamCity          string   `json:"teamCity"`
	TeamTricode       string   `json:"teamTricode"`
	TeamSlug          string   `json:"teamSlug"`
	Wins              int      `json:"wins"`
	Losses            int      `json:"losses"`
	Score             int      `json:"score"`
	Seed              int      `json:"seed"`
	InBonus           bool     `json:"inBonus"`
	TimeoutsRemaining int      `json:"timeoutsRemaining"`
	Periods           []Period `json:"periods"`
}

type Period struct {
	Period     int    `json:"period"`
	PeriodType string `json:"periodType"`
	Score      int    `json:"score"`
}
