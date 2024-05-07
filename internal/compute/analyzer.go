package compute

import "fmt"

type analyzer struct {
}

func NewAnalyzer() Analyzer {
	return &analyzer{}
}

func (a *analyzer) Analyze(tokens []string) (Query, error) {
	if len(tokens) == 0 {
		return nil, fmt.Errorf("empty query")
	}

	command := tokens[0]

	switch command {
	case "set":
		return a.analyzeSet(tokens)
	case "get":
		return a.analyzeGet(tokens)
	case "del":
		return a.analyzeDelete(tokens)
	default:
		return nil, fmt.Errorf("unknown command: %s", command)
	}
}

func (a *analyzer) analyzeSet(tokens []string) (Query, error) {
	if len(tokens) != 3 {
		return nil, fmt.Errorf("invalid set query: %v", tokens)
	}

	return &SetQuery{Key: tokens[1], Val: tokens[2]}, nil
}

func (a *analyzer) analyzeGet(tokens []string) (Query, error) {
	if len(tokens) != 2 {
		return nil, fmt.Errorf("invalid get query: %v", tokens)
	}

	return &GetQuery{Key: tokens[1]}, nil
}

func (a *analyzer) analyzeDelete(tokens []string) (Query, error) {
	if len(tokens) != 2 {
		return nil, fmt.Errorf("invalid delete query: %v", tokens)
	}

	return &DeleteQuery{Key: tokens[1]}, nil
}
