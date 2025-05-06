package models

const TableNameFriendShip = "t_friend_ships"

var FriendShipModel = &friendShipColumn{
	ID:        "id",
	UserId:    "user_id",
	FriendId:  "friend_id",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

type FriendShip struct {
	BaseModel
	UserId   uint64 `gorm:"column:user_id;type:bigint unsigned;not null;comment:用户ID" json:"user_id"`                   // 用户ID
	FriendId uint64 `gorm:"column:friend_id;type:bigint unsigned;not null;comment:好友ID" json:"friend_id"`               // 好友ID
	Status   string `gorm:"column:status;type:enum('0','1','2');not null;comment:请求状态：0-待确认，1-已接受，2-已拒绝" json:"status"` // 状态
}

func (*FriendShip) TableName() string { return TableNameFriendShip }

type friendShipColumn struct {
	ID        string
	UserId    string
	FriendId  string
	Status    string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}
