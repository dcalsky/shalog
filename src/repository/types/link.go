package types

import "time"

type Link struct {
	Id        int64     `gorm:"column:id;primary_key" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Link      string    `gorm:"column:link" json:"link"`
	Icon      string    `gorm:"column:icon;" json:"icon"`
	State     int       `gorm:"column:state;" json:"state"`
	CreatedAt time.Time `gorm:"column:created_at;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;" json:"updated_at"`
}

func (p *Link) TableName() string {
	return "links"
}
