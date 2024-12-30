package texter

type Texter interface {
	Text() (Message, error)
}
