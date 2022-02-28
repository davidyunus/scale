package repository

import "github.com/scale/src/domain"

func remove(s []domain.Scale, i int) []domain.Scale {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
