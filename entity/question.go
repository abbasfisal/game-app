package entity

type Question struct {
	ID             uint
	Question       string
	PossibleAnswer []string
	CorrectAnswer  string
	Difficulty     string
	Category       string
}
