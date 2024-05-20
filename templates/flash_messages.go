package templates

type Message struct {
	Type    string
	Message string
}

var FlashedMessages []Message

func FlashMessage(message Message) {
	FlashedMessages = append(FlashedMessages, message)
}

func GetFlashedMessages() []Message {
	defer func() {
		FlashedMessages = []Message{}
	}()
	return FlashedMessages
}
