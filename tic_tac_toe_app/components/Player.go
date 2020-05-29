package components

//Player structure stores the name and assigned mark of a player
type Player struct {
	Name string
	Mark string
}

//NewPlayer will return the Structure Player after updating its values
func NewPlayer(name, mark string) *Player {
	return &Player{
		Name: name,
		Mark: mark,
	}
}
