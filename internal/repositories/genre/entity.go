package genre_repository

import "time"

type Genre struct {
	ID        int64      `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string     `gorm:"column:name;unique"`
	CreatedAt *time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
