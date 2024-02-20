package entity

type Game struct {
	ID          uint
	Category    string
	QuestionIDs []uint
	PlayerIDs   []uint
}
