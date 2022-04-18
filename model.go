package modelrole

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Role struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt
	UserID    uuid.UUID
	Role      RoleEnum
}

func Init(db *gorm.DB) {
	DB = db
}

func (r *Role) List() ([]Role, error) {
	list := []Role{}
	err := DB.Find(&list, r).Error

	return list, err
}

func (r *Role) Create() error {
	return DB.FirstOrCreate(&r, r).Error
}
