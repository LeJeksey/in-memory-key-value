package compute

import (
	"errors"
	"strings"
)

// query = set_command | get_command | del_command
//
// set_command = "SET" argument argument
// get_command = "GET" argument
// del_command = "DEL" argument
// argument    = punctuation | letter | digit { punctuation | letter | digit }
//
// punctuation = "*" | "/" | "_"
// letter      = "a" | ... | "z" | "A" | ... | "Z"
// digit       = "0" | ... | "9"
//
// Examples:
// SET weather_2_pm cold_moscow_weather
// GET /etc/nginx/config.yaml
// DEL user_****

type state int

const (
	startState state = iota
	wordState
	spaceState
	errorState
	statesCount
)

type event int

const (
	letterEvent event = iota
	spaceEvent
	unknownEvent
	eventsCount
)

var (
	ErrInvalidQuery = errors.New("invalid query")
	ErrEmptyQuery   = errors.New("empty query")
)

type transition struct {
	jump     func(byte) state
	reaction func()
}

type stateMachine struct {
	transitions [statesCount][eventsCount]transition
	state       state

	tokens []string

	currToken strings.Builder
}

func newStateMachine() *stateMachine {
	m := &stateMachine{
		state: startState,
	}

	m.transitions = [statesCount][eventsCount]transition{
		startState: {
			letterEvent: {jump: m.appendLetterJump},
			spaceEvent:  {jump: m.skipWhiteSpaceJump},
		},
		wordState: {
			letterEvent: {jump: m.appendLetterJump},
			spaceEvent:  {jump: m.skipWhiteSpaceJump, reaction: m.saveToken},
		},
		spaceState: {
			letterEvent: {jump: m.appendLetterJump},
			spaceEvent:  {jump: m.skipWhiteSpaceJump},
		},
		errorState: {},
	}

	return m
}

func (m *stateMachine) parse(query string) ([]string, error) {
	for i := 0; i < len(query); i++ {
		b := query[i]

		e := m.getEvent(b)
		if e == unknownEvent {
			m.state = errorState
			break
		}

		t := m.transitions[m.state][e]

		m.state = t.jump(b)
		if t.reaction != nil {
			t.reaction()
		}
	}

	if m.state == startState {
		return nil, ErrEmptyQuery
	}
	if m.state == errorState {
		return nil, ErrInvalidQuery
	}

	// Save the last token
	if m.state == wordState {
		m.saveToken()
	}

	return m.tokens, nil
}

func (m *stateMachine) getEvent(b byte) event {
	switch {
	case isLetter(b):
		return letterEvent
	case isSpace(b):
		return spaceEvent
	default:
		return unknownEvent
	}
}

func isLetter(b byte) bool {
	return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' ||
		b >= '0' && b <= '9' ||
		b == '*' || b == '/' || b == '_'
}

func isSpace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\n'
}

func (m *stateMachine) skipWhiteSpaceJump(_ byte) state {
	return spaceState
}

func (m *stateMachine) appendLetterJump(b byte) state {
	m.currToken.WriteByte(b)
	return wordState
}

func (m *stateMachine) saveToken() {
	m.tokens = append(m.tokens, m.currToken.String())
	m.currToken.Reset()
}
