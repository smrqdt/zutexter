package powertexter

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/smrqdt/zutexter/pkg/texter"
)

//go:embed query.json
var query string

type PowerTexter struct {
	GrafanaURL string
}

func (t *PowerTexter) Text() (m texter.Message, err error) {
	vals, err := t.query()
	if err != nil {
		return texter.Message{}, err
	}
	m.Type = texter.SCROLL
	var valueStrings []string
	for _, val := range vals {
		valueStrings = append(valueStrings, fmt.Sprintf("%v %.2f kW", val.Name, val.Value/1000))
	}
	m.Content = "Power Usage: " + strings.Join(valueStrings, " / ")
	return m, nil
}

func (t *PowerTexter) query() ([]PowerValues, error) {
	resp, err := http.Post(t.GrafanaURL, "application/json ", strings.NewReader(query))
	if err != nil {
		return nil, fmt.Errorf("Could not fetch Power stats: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Could not fetch Power stats: %w", err)
	}
	response := &ResponseSchema{}
	json.Unmarshal(body, response)

	var vals []PowerValues

	measurement, ok := response.Results["A"]
	if !ok {
		return nil, fmt.Errorf("No measurement 'A' found")
	}

	re := regexp.MustCompile(`PDU (\d+) \((.*)\)`)

	for _, frame := range measurement.Frames {
		var val PowerValues
		for _, field := range frame.Schema.Fields {
			if location, ok := field.Labels["location"]; ok {
				matches := re.FindStringSubmatch(location)
				val.Name = matches[2]
				pduID, err := strconv.Atoi(matches[1])
				if err != nil {
					return nil, err
				}
				val.PDU = pduID
				break
			}
		}
		val.Time = time.Unix(int64(frame.Data.Values[0][len(frame.Data.Values)-1])/1000, 0)
		val.Value = frame.Data.Values[1][len(frame.Data.Values)-1]
		vals = append(vals, val)
	}
	return vals, nil
}

type PowerValues struct {
	PDU   int
	Name  string
	Time  time.Time
	Value float64
}

// JSON Schemas
type ResponseSchema struct {
	Results map[string]ResultSchema
}

type ResultSchema struct {
	Status int
	Frames []FrameSchema
}

type FrameSchema struct {
	Schema SchemaSchema
	Data   DataSchema
}

type SchemaSchema struct {
	Name   string
	Fields []FieldSchema
}

type FieldSchema struct {
	Name   string
	Type   string
	Labels map[string]string
}

type DataSchema struct {
	Values [][]float64
}
