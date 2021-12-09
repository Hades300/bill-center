package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollectionService_FetchCollectionByCode(t *testing.T) {
	s := NewCollectionService()
	code := "412423"
	ret, err := s.FetchCollectionByCode(code, ":")
	assert.NotNil(t, ret)
	assert.Nil(t, err)
	for _, r := range ret.ResultList {
		t.Log(r)
	}
}

func TestCollectionService_RegisterCollection(t *testing.T) {
	s := NewCollectionService()
	code := "412423"
	err := s.RegisterCollection(code, "0", ":", 3600*24*7)
	assert.Nil(t, err)
}
