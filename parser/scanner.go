package parser

import (
	"errors"
	"strings"
)

var (
	ErrUnbalancedParentheses = errors.New("unbalanced parentheses")
	ErrUnexpectedOperator    = errors.New("unexpected operator")
	ErrUnbalancedOperand     = errors.New("unbalanced operand")
)

type scanner struct {
	query         string
	tokens        []string
	operatorStack *sizedStack[*operator]
	operandStack  *sizedStack[*node]
}

func NewScanner(query string) *scanner {
	return &scanner{
		query:         query,
		tokens:        make([]string, 0, 64),
		operatorStack: newStack[*operator](64),
		operandStack:  newStack[*node](64),
	}
}

func (s *scanner) ParseAndValid() error {
	s.query = strings.ReplaceAll(s.query, "!", "NOT")
	s.query = strings.ReplaceAll(s.query, "&&", "AND")
	s.query = strings.ReplaceAll(s.query, "||", "OR")

	s.query = strings.ReplaceAll(s.query, "(", " ( ")
	s.query = strings.ReplaceAll(s.query, ")", " ) ")
	s.query = strings.ReplaceAll(s.query, "NOT", " NOT ")
	s.query = strings.ReplaceAll(s.query, "AND", " AND ")
	s.query = strings.ReplaceAll(s.query, "OR", " OR ")

	for _, item := range strings.Split(s.query, " ") {
		if item == "" {
			continue
		}
		if isOP(item) || isValidExpr(item) {
			s.tokens = append(s.tokens, item)
			continue
		}
		return ErrInvalidToken
	}
	if len(s.tokens) == 0 {
		return ErrEmptyToken
	}
	return newStateMachine(s.tokens).IsValid()
}

func (s *scanner) Scan() (*node, error) {
	tokensLen := len(s.tokens)
	for i := 0; i < tokensLen; {
		token := s.tokens[i]

		if isOP(token) {
			operator := newOperator(token)
			if s.operatorStack.IsEmpty() || operator.IsOperatorLP() {
				s.operatorStack.Push(operator)
				i++
				continue
			}

			if operator.IsOperatorRP() {
				if s.operatorStack.IsEmpty() {
					return nil, ErrUnbalancedParentheses
				}

				top := s.operatorStack.Peep()
				if top.IsOperatorLP() {
					s.operatorStack.Pop()
					i++
					continue
				}

				if err := s.operateOnTop(); err != nil {
					return nil, err
				}
				continue
			}

			top := s.operatorStack.Peep()
			if top.IsOperatorLP() {
				s.operatorStack.Push(operator)
				i++
			} else if operator.priority > top.priority {
				if err := s.operateOnTop(); err != nil {
					return nil, err
				}
			} else {
				s.operatorStack.Push(operator)
				i++
			}
		} else {
			s.operandStack.Push(newNonOpNode(token))
			i++
		}
	}

	for !s.operatorStack.IsEmpty() {
		if err := s.operateOnTop(); err != nil {
			return nil, err
		}
	}

	if s.operandStack.Length() != 1 {
		return nil, ErrUnbalancedOperand
	}

	return s.operandStack.Pop(), nil
}

func (s *scanner) operateOnTop() error {
	operator := s.operatorStack.Pop()

	if (operator.IsOperatorNOT() ||
		operator.IsOperatorAND() ||
		operator.IsOperatorOR()) && (s.operandStack.Length() < operator.operandNum) {
		return ErrUnbalancedParentheses
	}

	switch true {
	case operator.IsOperatorNOT():
		operand := s.operandStack.Pop()
		s.operandStack.Push(Not(operand))
	case operator.IsOperatorAND():
		rightOperand := s.operandStack.Pop()
		leftOperand := s.operandStack.Pop()
		s.operandStack.Push(And(leftOperand, rightOperand))
	case operator.IsOperatorOR():
		rightOperand := s.operandStack.Pop()
		leftOperand := s.operandStack.Pop()
		s.operandStack.Push(Or(leftOperand, rightOperand))
	default:
		return ErrUnexpectedOperator
	}

	return nil
}
