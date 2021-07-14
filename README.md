### waifuPicgo

A powerful waifu.pics api wrapper for Go.

```
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

```

- [inspire code ](https://github.com/erizkiatama/gojikan)   
