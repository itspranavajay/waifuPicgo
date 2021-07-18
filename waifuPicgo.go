package waifuPicgo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SFW struct {
	Url string                 `json:"url"`
}

type NSFW struct {
	Url string                 `json:"url"`
}

func sfw(Category string) (*SFW, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.waifu.pics/sfw/%v", Category))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var sFw = new(Response)
	json.NewDecoder(resp.Body).Decode(&sFw)
	return &sFw.Data.SFW, nil

func nsfw(Category string) (*NSFW, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.waifu.pics/nsfw/%v", Category))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var nsFw = new(Response)
	json.NewDecoder(resp.Body).Decode(&nsFw)
	return &nsFw.Data.NSFW, nil
