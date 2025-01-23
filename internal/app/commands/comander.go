package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"sync"
)

type ComRout struct {
	bot        *tgbotapi.BotAPI
	userStates map[int64][]string
	mu         sync.Mutex
}

func NewComRout(bot *tgbotapi.BotAPI) *ComRout {
	return &ComRout{
		bot:        bot,
		userStates: make(map[int64][]string),
	}
}
func (c *ComRout) pushState(userID int64, state string) {
	if c.userStates == nil {
		c.userStates = make(map[int64][]string)
	}
	c.userStates[userID] = append(c.userStates[userID], state)
}
func (c *ComRout) getState(userID int64) string {
	states := c.userStates[userID]
	lastState := states[len(states)-1]
	return lastState
}
func (c *ComRout) popState(userID int64) string {
	if c.userStates == nil || len(c.userStates[userID]) == 0 {
		return ""
	}
	states := c.userStates[userID]
	lastState := states[len(states)-1]
	c.userStates[userID] = states[:len(states)-1]
	return lastState
}
