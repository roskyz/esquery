package parser

func Not(n *node) *node {
	if n.IsNotNode() {
		return n.operands[0]
	}

	notNode := newOpNode(newOperator(operatorNOT))
	notNode.AppendOperands(notNode)
	return notNode
}

func And(left, right *node) *node {
	andNode := newOpNode(newOperator(operatorAND))
	if left.IsAndNode() && right.IsAndNode() {
		andNode.AppendOperands(left.operands...)
		andNode.AppendOperands(right.operands...)
	} else if left.IsAndNode() {
		andNode.AppendOperands(left.operands...)
		andNode.AppendOperands(right)
	} else if right.IsAndNode() {
		andNode.AppendOperands(right.operands...)
		andNode.AppendOperands(left)
	} else {
		andNode.AppendOperands(left, right)
	}
	return andNode
}

func Or(left, right *node) *node {
	orNode := newOpNode(newOperator(operatorOR))
	if left.IsOrNode() && right.IsOrNode() {
		orNode.AppendOperands(left.operands...)
		orNode.AppendOperands(right.operands...)
	} else if left.IsOrNode() {
		orNode.AppendOperands(left.operands...)
		orNode.AppendOperands(right)
	} else if right.IsOrNode() {
		orNode.AppendOperands(right.operands...)
		orNode.AppendOperands(left)
	} else {
		orNode.AppendOperands(left, right)
	}
	return orNode
}
