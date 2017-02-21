package model

type PMmoveset struct{
	PMID int 
	Name string 
	Moves []Move
	ItemID int
	AbilityID int
	NatureID int
	Evs []int
	Comments string
}
type Move struct{
	AltmovesID []int 
}