// Package search implements a plugin to search on Google Custom Search.
package search

import (
	"log"
	"strings"

	"github.com/StalkR/goircbot/bot"
	"github.com/StalkR/goircbot/lib/google/search"
)

func find(e *bot.Event, key, cx string) {
	term := strings.TrimSpace(e.Args)
	if len(term) == 0 {
		return
	}
	r, err := search.Search(term, key, cx)
	if err != nil {
		log.Println("search:", err)
		return
	}
	if len(r.Items) == 0 {
		e.Bot.Privmsg(e.Target, "No result.")
		return
	}
	e.Bot.Privmsg(e.Target, r.Items[0].String())
}

// Register registers the plugin with a bot.
func Register(b bot.Bot, key, cx string) {
	b.Commands().Add("search", bot.Command{
		Help:    "search the Web with Google",
		Handler: func(e *bot.Event) { find(e, key, cx) },
		Pub:     true,
		Priv:    true,
		Hidden:  false})
}
