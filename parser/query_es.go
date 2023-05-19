package parser

import (
	"strings"
	"time"

	"github.com/olivere/elastic/v7"
)

func (n *node) GenEsQuery() elastic.Query {
	if n.IsNotNode() || n.IsAndNode() || n.IsOrNode() {
		query := elastic.NewBoolQuery()
		queryfilter := query.MustNot
		if n.IsAndNode() {
			queryfilter = query.Must
		} else if n.IsOrNode() {
			queryfilter = query.Should
		}
		for _, operand := range n.operands {
			subquery := operand.GenEsQuery()
			queryfilter(subquery)
		}
		return query
	}
	key, pattern, value := parseExpr(n.expr)
	switch pattern {
	case patternIN:
		items := strings.Split(value, ",")
		var terms = make([]interface{}, len(items))
		for i := 0; i < len(items); i++ {
			terms[i] = items[i]
		}
		return elastic.NewTermsQuery(key.keyword(), terms...)
	case patternMATCH:
		return elastic.NewMatchQuery(key.keyword(), value)
	case patternPREFIX:
		return elastic.NewPrefixQuery(key.keyword(), value)
	case patternSUFFIX:
		return elastic.NewRegexpQuery(key.keyword(), ".*"+value+"$")
	case patternREGEX:
		return elastic.NewRegexpQuery(key.keyword(), value)
	case patternEQ:
		return elastic.NewTermQuery(key.string(), value)
	case patternLT:
		return elastic.NewRangeQuery(key.string()).Lt(value)
	case patternGT:
		return elastic.NewRangeQuery(key.string()).Gt(value)
	case patternLTE:
		return elastic.NewRangeQuery(key.string()).Lte(value)
	case patternGTE:
		return elastic.NewRangeQuery(key.string()).Gte(value)
	case patternBOOL:
		return elastic.NewTermQuery(key.string(), value == "true")
	case patternBEFORE:
		t, _ := time.Parse("2006-01-02T15:04:05", value)
		return elastic.NewRangeQuery(key.string()).Lte(t.UTC())
	case patternAFTER:
		t, _ := time.Parse("2006-01-02T15:04:05", value)
		return elastic.NewRangeQuery(key.string()).Gte(t.UTC())
	case patternEXIST:
		return elastic.NewExistsQuery(key.string())
	}
	return nil
}
