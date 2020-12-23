package model

type Stock interface {
	CalculateShareToBuy(positionSize int)
	ToString() []string
}
