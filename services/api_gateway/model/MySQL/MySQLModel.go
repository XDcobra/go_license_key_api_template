package MySQL

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:30;not null" json:"name"`
	Age   int    `gorm:"not null" json:"age"`
	Email string `gorm:"size:30" json:"email"`
}
