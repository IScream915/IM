package repo

import (
	"IM/internal/repo/models"
	"context"
	"gorm.io/gorm"
)

type FriendShip interface {
	Create(ctx context.Context, friend *models.FriendShip) error
	Update(ctx context.Context, friend *models.FriendShip) error
	FindByUserId(ctx context.Context, id uint64) ([]*models.FriendShip, error)
	FindAllFriends(ctx context.Context, id uint64) ([]*models.FriendShip, error)
	Transaction(ctx context.Context, f func(tx *gorm.DB) error) error
	WithTx(tx *gorm.DB) FriendShip
}

func NewFriendShip(client *gorm.DB) FriendShip {
	return &friendShip{client: client}
}

type friendShip struct {
	client *gorm.DB
}

func (obj *friendShip) Create(ctx context.Context, friend *models.FriendShip) error {
	return obj.client.WithContext(ctx).Create(friend).Error
}

func (obj *friendShip) Update(ctx context.Context, friend *models.FriendShip) error {
	return obj.client.WithContext(ctx).Model(&models.FriendShip{}).Where("id = ?", friend.ID).Updates(friend).Error
}

func (obj *friendShip) FindByUserId(ctx context.Context, id uint64) ([]*models.FriendShip, error) {
	friendShips := make([]*models.FriendShip, 0)
	err := obj.client.WithContext(ctx).Where("id = ?", id).Find(&friendShips).Error
	if err != nil {
		return nil, err
	}
	return friendShips, nil
}

func (obj *friendShip) FindAllFriends(ctx context.Context, id uint64) ([]*models.FriendShip, error) {
	friendShips := make([]*models.FriendShip, 0)
	// SQL: SELECT *
	//	FROM t_friend_ships
	//	WHERE status = '1'
	//	AND (user_id = 1
	//    	OR friend_id = 1);
	err := obj.client.WithContext(ctx).
		Where("status = ?", "1").
		Where(obj.client.WithContext(ctx).
			Where("user_id = ?", id).
			Or("friend_id = ?", id)).Find(&friendShips).Error
	if err != nil {
		return nil, err
	}
	return friendShips, nil
}

func (obj *friendShip) Transaction(ctx context.Context, f func(tx *gorm.DB) error) error {
	tx := obj.client.WithContext(ctx).Begin()
	if err := f(tx); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (obj *friendShip) WithTx(tx *gorm.DB) FriendShip {
	return &friendShip{client: tx}
}
