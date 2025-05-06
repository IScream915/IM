package services

import (
	"IM/internal/repo"
)

type FriendShip interface {
}

func NewFriendShip(repo repo.FriendShip) FriendShip {
	return &friendShip{repo: repo}
}

type friendShip struct {
	repo repo.FriendShip
}
