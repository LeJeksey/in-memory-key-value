package compute

import "fmt"

type Storage interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

type Parser interface {
	Parse(query string) ([]string, error)
}

type Analyzer interface {
	Analyze(tokens []string) (Query, error)
}

type Computer struct {
	storage  Storage
	parser   Parser
	analyzer Analyzer
}

func NewComputer(storage Storage, parser Parser, analyzer Analyzer) *Computer {
	return &Computer{
		storage:  storage,
		parser:   parser,
		analyzer: analyzer,
	}
}

func (c *Computer) Compute(rawQuery string) (result string, err error) {
	tokens, err := c.parser.Parse(rawQuery)
	if err != nil {
		return "", fmt.Errorf("failed to parse query: %w", err)
	}

	query, err := c.analyzer.Analyze(tokens)
	if err != nil {
		return "", fmt.Errorf("failed to analyze query: %w", err)
	}

	result, err = query.Run(c.storage)
	if err != nil {
		return "", fmt.Errorf("failed to run query: %w", err)
	}

	return result, nil
}
