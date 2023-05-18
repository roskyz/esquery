package parser

import (
	"errors"
)

var (
	ErrEmptyToken   = errors.New("query has empty token")
	ErrInvalidToken = errors.New("token is invalid")
)

type node struct {
	operator *operator
	operands []*node
	expr     string
}

func newNode(op *operator, operands []*node, expr string) *node {
	return &node{
		operator: op, operands: operands, expr: expr,
	}
}

func newOpNode(op *operator) *node {
	return newNode(op, nil, "")
}

func newNonOpNode(expr string) *node {
	return newNode(nil, nil, expr)
}

func (n *node) IsNotNode() bool {
	return n.operator != nil && n.operator.IsOperatorNOT()
}

func (n *node) IsAndNode() bool {
	return n.operator != nil && n.operator.IsOperatorAND()
}

func (n *node) IsOrNode() bool {
	return n.operator != nil && n.operator.IsOperatorOR()
}

func (n *node) AppendOperands(operand ...*node) {
	n.operands = append(n.operands, operand...)
}

type operator struct {
	token      string
	priority   int
	operandNum int
}

func newOperator(token string) *operator {
	return &operator{
		token:      token,
		priority:   opPriorityDict[token],
		operandNum: opOperandDoct[token],
	}
}

func (op *operator) IsOperatorLP() bool {
	return op.token == operatorLP
}

func (op *operator) IsOperatorRP() bool {
	return op.token == operatorRP
}

func (op *operator) IsOperatorNOT() bool {
	return op.token == operatorNOT
}

func (op *operator) IsOperatorAND() bool {
	return op.token == operatorAND
}

func (op *operator) IsOperatorOR() bool {
	return op.token == operatorOR
}
