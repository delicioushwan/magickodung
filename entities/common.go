package entities

import "time"

type Common struct {
	CreatedAt    time.Time `gorm:"autoUpdateTime:milli"`
  UpdatedAt    time.Time `gorm:"autoCreateTime"`
}
