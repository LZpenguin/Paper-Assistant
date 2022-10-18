package model

import "gorm.io/gorm"

type Feedback struct {
	gorm.Model

	UserRefer string

	Content string `gorm:"uniqueIndex;not null;"`
}

func (u UserModel) FeedBack(openid, content string) error {
	result := u.db.Create(&Feedback{UserRefer: openid, Content: content})
	return result.Error
}
