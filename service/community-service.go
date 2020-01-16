package service

import (
	"errors"

	"github.com/ivohutasoit/alira/model"
	"github.com/ivohutasoit/alira/model/domain"
)

type CommunityService struct{}

func (s *CommunityService) Get(args ...interface{}) (map[interface{}]interface{}, error) {
	if len(args) < 1 {
		return nil, errors.New("not enough parameters")
	}
	id, ok := args[0].(string)
	if !ok {
		return nil, errors.New("plain text parameter not type string")
	}

	community := &domain.Community{}
	model.GetDatabase().First(community, "id = ?", id)
	if community.BaseModel.ID == "" {
		return nil, errors.New("invalid community")
	}

	var members []domain.CommunityMember
	model.GetDatabase().Find(&members, "community_id = ?", community.BaseModel.ID)

	return map[interface{}]interface{}{
		"community": community,
		"members":   members,
	}, nil
}
