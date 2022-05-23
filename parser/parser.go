package parser

import (
	"encoding/json"
)

type Parser struct{}

func (p *Parser) ParseResults(body []byte) (Results, error) {
	results := Results{}
	err := json.Unmarshal(body, &results)
	if err != nil {
		return Results{}, err
	}
	return results, nil
}

func (r *Results) StringifyResults() (string, error) {
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
