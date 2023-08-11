package models

type Game struct {
	ID          int     `gorm:"index:idx_id;primary_key" json:"-"`
	Name        string  `gorm:"not null;type:varchar(100)"`
	Description string  `gorm:"not null"`
	Image       *string `gorm:"type:varchar(250)" json:"-"`
	Cover       *string `gorm:"type:varchar(250)" json:"-"`
	Category    int     `gorm:"index:idx_category;not null"`
	CreatedAt   int64   `gorm:"index:idx_created_at;not null" json:"created_at"`
	UpdatedAt   int64   `gorm:"index:idx_updated_at;not null" json:"updated_at"`
}
