package components

//Player stores name and assigned mark of a player
type Player struct {
	Name string
	Mark string
}

//NewPlayer returns struct Player
func NewPlayer(name, mark string) *Player {
	return &Player{
		Name: name,
		Mark: mark,
	}
}
