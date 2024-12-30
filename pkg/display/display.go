package display

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/smrqdt/zutexter/pkg/texter"
)

type Display struct {
	URL url.URL
}

func (d *Display) Queue(m texter.Message) error {
	jsonMsg, err := json.Marshal(m)
	if err != nil {
		return err
	}
	data := url.Values{}
	data.Add("content", string(jsonMsg))
	queueURL, _ := url.JoinPath(d.URL.String(), "/queue")
	_, err = http.PostForm(queueURL, data)
	return err
}
