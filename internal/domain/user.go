package domain

type (
	User struct {
		ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
		Name      string `json:"name" gorm:"not null"`
		CreatedAt int64  `json:"created_at" gorm:"autoCreateTime:milli"`
		UpdatedAt int64  `json:"updated_at" gorm:"autoUpdateTime:milli"`
	}

	UserRequest struct {
		Name string `form:"name" binding:"required"`
	}

	GetUserRequest struct {
		Page int
		Size int
	}
)
