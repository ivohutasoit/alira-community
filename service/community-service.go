package service

import (
	"errors"
	"fmt"
	"strings"

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

func (s *CommunityService) Create(args ...interface{}) (map[interface{}]interface{}, error) {
	if len(args) < 2 {
		return nil, errors.New("not enough parameters")
	}
	userid, ok := args[0].(string)
	if !ok {
		return nil, errors.New("plain text parameter not type string")
	}

	name, ok := args[1].(string)
	if !ok {
		return nil, errors.New("plain text parameter not type string")
	}

	community := &domain.Community{
		Name: strings.ToUpper(name),
	}
	model.GetDatabase().Create(community)

	creator := &domain.CommunityMember{
		CommunityID: community.BaseModel.ID,
		UserID:      userid,
		Creator:     true,
		Admin:       true,
		JoinBy:      "CREATION",
		Approved:    true,
	}
	model.GetDatabase().Create(creator)

	return map[interface{}]interface{}{
		"status":  "SUCCESS",
		"message": fmt.Sprintf("Community %s has been created successful", name),
	}, nil
}
