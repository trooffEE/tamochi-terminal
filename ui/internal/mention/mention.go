package mention

import "errors"

type mention = int32

const (
	WelcomeMessage       mention = 0
	PleaseProvideMessage mention = 1
)

var mentions map[mention]string = map[mention]string{
	WelcomeMessage:       "Hello, Welcome to \"Tomachi\" - a place where beautiful creatures live",
	PleaseProvideMessage: "Please login or authorize In order to use it!",
}

func GetMention(mentionKey mention) (string, error) {
	value, ok := mentions[mentionKey]
	if !ok {
		return "", errors.New("Mention not found")
	}
	return value, nil
}
