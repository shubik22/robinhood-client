package models

type Leaderboard struct {
	Entries []User `json:"entries"`
}
