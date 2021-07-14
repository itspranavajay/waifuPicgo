package test

import (
	"testing"

	"github.com/moezilla/waifuPicgo"
)

func sfw(w *testing.W) {
	_, err := waifuPicgo.sfw("waifu")
	if err != nil {
		w.Error(err)
	}
}

func nsfw(w *testing.W) {
	_, err := waifuPicgo.nsfw("waifu")
	if err != nil {
		w.Error(err)
	}
}

