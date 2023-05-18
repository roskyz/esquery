package parser

type stateMachine struct {
	tokens     []string
	transition map[string]map[string]bool
	start      map[string]bool
	end        map[string]bool
	mapper     func(op string) string
}

func newStateMachine(tokens []string) *stateMachine {
	return &stateMachine{
		tokens:     tokens,
		transition: stateTransition,
		start:      startDict,
		end:        endDict,
		mapper:     operatorMapper,
	}
}

func (sm *stateMachine) IsValid() error {
	tokensLen := len(sm.tokens)
	if tokensLen == 0 {
		return nil
	}

	var current int

	startToken := sm.getToken(current)
	if !sm.start[startToken] {
		return ErrInvalidStart
	}

	endToken := sm.getToken(tokensLen - 1)
	if !sm.end[endToken] {
		return ErrInvalidEnd
	}

	next := current + 1
	for current < tokensLen {
		token := sm.getToken(current)
		nextAllowedToken, ok := sm.transition[token]
		if !ok {
			return ErrInvalidState
		}

		if next == tokensLen {
			return nil
		}

		if !nextAllowedToken[sm.getToken(next)] {
			return ErrInvalidTransition
		}
		current++
		next++
	}

	return nil
}

func (sm *stateMachine) getToken(idx int) string {
	return sm.mapper(sm.tokens[idx])
}
