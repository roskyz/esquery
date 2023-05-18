package parser

import "errors"

var (
	ErrInvalidStart      = errors.New("invalid start token")
	ErrInvalidEnd        = errors.New("invalid end token")
	ErrInvalidState      = errors.New("invalid state token")
	ErrInvalidTransition = errors.New("invalid transition token")
)

const operatorExpr = "expr"

var (
	startDict = map[string]bool{
		operatorExpr: true,
		operatorLP:   true,
		operatorNOT:  true,
	}
	endDict = map[string]bool{
		operatorExpr: true,
		operatorRP:   true,
	}
	stateTransition = map[string]map[string]bool{
		operatorLP: {
			operatorLP:   true,
			operatorNOT:  true,
			operatorExpr: true,
		},
		operatorRP: {
			operatorRP:  true,
			operatorAND: true,
			operatorOR:  true,
		},
		operatorNOT: {
			operatorLP:   true,
			operatorNOT:  true,
			operatorExpr: true,
		},
		operatorAND: {
			operatorLP:   true,
			operatorNOT:  true,
			operatorExpr: true,
		},
		operatorOR: {
			operatorLP:   true,
			operatorNOT:  true,
			operatorExpr: true,
		},
		operatorExpr: {
			operatorRP:  true,
			operatorAND: true,
			operatorOR:  true,
		},
	}
)

func operatorMapper(op string) string {
	switch op {
	case operatorLP:
		return op
	case operatorRP:
		return op
	case operatorNOT:
		return op
	case operatorAND:
		return op
	case operatorOR:
		return op
	}
	return operatorExpr
}
