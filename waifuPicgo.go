package waifuPicgo

import (
	"encoding/json"
        "errors"
	"fmt"
	"io"
	"net/http"
)

func sfw(category  string)(err error){
        var waifu struct {
		 string `json:"url"`
	}
        URL := fmt.Sprintf("https://api.waifu.pics/sfw/",  category)
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &waifu)
	if err != nil {
		return nil, err
	}
	return waifu.img, nil
}

func nsfw(category  string)(err error){
        var waifu struct {
		img string `json:"url"`
	}
        URL := fmt.Sprintf("https://api.waifu.pics/nsfw/",  category)
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &waifu)
	if err != nil {
		return nil, err
	}
	return waifu.img, nil
}
