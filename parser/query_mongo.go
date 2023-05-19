package parser

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (n *node) GenMongoQuery() bson.D {
	if n.IsNotNode() || n.IsAndNode() || n.IsOrNode() {
		var query bson.D
		if n.IsNotNode() {
			return append(query, bson.E{Key: "$not", Value: n.operands[0].GenMongoQuery()})
		}

		var subqueries []bson.D
		var queryfilter string
		if n.IsNotNode() {
			queryfilter = "$not"
		} else if n.IsAndNode() {
			queryfilter = "$and"
		} else if n.IsOrNode() {
			queryfilter = "$or"
		}
		for _, operand := range n.operands {
			subqueries = append(subqueries, operand.GenMongoQuery())
		}
		query = append(query, bson.E{Key: queryfilter, Value: subqueries})
		return query
	}
	key, pattern, value := parseExpr(n.expr)
	switch pattern {
	case patternIN:
		return bson.D{{Key: key.string(), Value: bson.M{"$in": value}}}
	case patternMATCH:
		return bson.D{{Key: key.string(), Value: value}}
	case patternPREFIX:
		return bson.D{{Key: key.string(), Value: primitive.Regex{Pattern: "^" + value}}}
	case patternSUFFIX:
		return bson.D{{Key: key.string(), Value: primitive.Regex{Pattern: value + "$"}}}
	case patternREGEX:
		return bson.D{{Key: key.string(), Value: primitive.Regex{Pattern: value}}}
	case patternEQ:
		return bson.D{{Key: key.string(), Value: bson.M{"$eq": value}}}
	case patternLT:
		return bson.D{{Key: key.string(), Value: bson.M{"$lt": value}}}
	case patternGT:
		return bson.D{{Key: key.string(), Value: bson.M{"$gt": value}}}
	case patternLTE:
		return bson.D{{Key: key.string(), Value: bson.M{"$lte": value}}}
	case patternGTE:
		return bson.D{{Key: key.string(), Value: bson.M{"$gte": value}}}
	case patternBOOL:
		return bson.D{{Key: key.string(), Value: value == "true"}}
	case patternBEFORE:
		t, _ := time.Parse("2006-01-02T15:04:05", value)
		return bson.D{{Key: key.string(), Value: bson.M{"$lte": t.UTC()}}}
	case patternAFTER:
		t, _ := time.Parse("2006-01-02T15:04:05", value)
		return bson.D{{Key: key.string(), Value: bson.M{"$gte": t.UTC()}}}
	case patternEXIST:
		return bson.D{{Key: key.string(), Value: bson.M{"$exists": true}}}
	}
	return nil
}
