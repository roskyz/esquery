package parser

import (
	"strings"
)

const DefaultStackSize = 64

const (
	operatorLP  = "("
	operatorRP  = ")"
	operatorNOT = "NOT"
	operatorAND = "AND"
	operatorOR  = "OR"
)

const (
	patternIN     = "in"
	patternMATCH  = "match"
	patternPREFIX = "prefix"
	patternSUFFIX = "suffix"
	patternREGEX  = "regex"
	patternEQ     = "eq"
	patternLT     = "lt"
	patternGT     = "gt"
	patternLTE    = "lte"
	patternGTE    = "gte"
	patternBOOL   = "bool"
	patternBEFORE = "before"
	patternAFTER  = "after"
	patternEXIST  = "exist"
)

var (
	opPriorityDict = map[string]int{
		operatorLP:  0,
		operatorRP:  0,
		operatorNOT: 1,
		operatorAND: 2,
		operatorOR:  3,
	}

	opOperandDoct = map[string]int{
		operatorLP:  0,
		operatorRP:  0,
		operatorNOT: 1,
		operatorAND: 2,
		operatorOR:  2,
	}

	patternDict = map[string]bool{
		patternIN:     true,
		patternMATCH:  true,
		patternPREFIX: true,
		patternSUFFIX: true,
		patternREGEX:  true,
		patternEQ:     true,
		patternLT:     true,
		patternGT:     true,
		patternLTE:    true,
		patternGTE:    true,
		patternBOOL:   true,
		patternBEFORE: true,
		patternAFTER:  true,
		patternEXIST:  true,
	}
)

func isOP(op string) bool { _, ok := opPriorityDict[op]; return ok }

func isValidPattern(pattern string) bool { return patternDict[pattern] }

func isValidExpr(token string) bool {
	kpv := strings.Split(token, "=")
	if len(kpv) != 2 {
		return false
	}
	kp := strings.Split(kpv[0], "__")
	if len(kp) < 1 || len(kp) > 2 || kp[0] == "" {
		return false
	}
	if len(kp) == 2 {
		return isValidPattern(kp[1])
	}
	return true
}
