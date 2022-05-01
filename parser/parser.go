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
