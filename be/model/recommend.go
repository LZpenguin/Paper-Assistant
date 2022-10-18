package model

func (u *UserModel) Recommend(openid string) {
	// 用user数据推荐20篇文章

}

func (u *UserModel) Any20() ([]Paper, error) {
	// 随机20 篇paper

	var rs []Paper
	result := u.db.Model(&Paper{}).Order("random()").Limit(20).Find(&rs)

	return rs, result.Error
}
