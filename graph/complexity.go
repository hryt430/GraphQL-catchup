package graph

import "gql_server/internal"

type ComplexityRoot struct {
	Repository struct {
		Issues func(childComplexity int, after *string, before *string, first *int32, last *int32) int32
	}
}

func ComplexityConfig() internal.ComplexityRoot {
	var c internal.ComplexityRoot

	c.Repository.Issues = func(childComplexity int, after *string, before *string, first *int32, last *int32) int {
		var cnt int32
		switch {
		case first != nil && last != nil:
			if *first < *last {
				cnt = *last
			} else {
				cnt = *first
			}
		case first != nil && last == nil:
			cnt = *first
		case first == nil && last != nil:
			cnt = *last
		default:
			cnt = 1
		}
		return int(cnt) * childComplexity
	}
	return c
}
