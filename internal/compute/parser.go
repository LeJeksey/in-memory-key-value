package compute

import "fmt"

type parserSM struct {
}

func NewParserSM() Parser {
	return &parserSM{}
}

func (p *parserSM) Parse(query string) (tokens []string, err error) {
	tokens, err = newStateMachine().parse(query)
	if err != nil {
		return nil, fmt.Errorf("failed to parse query: %w", err)
	}

	return tokens, nil
}
