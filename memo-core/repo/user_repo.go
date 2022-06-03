package repo

import "memo-core/repo/entity"

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

type UserRepo struct {
}

// Exist 检测邮箱或用户名是否已经存在
func (r *UserRepo) Exist(username, email string) (bool, error) {
	var cnt int64
	if err := DB.Model(&entity.User{}).
		Where("username = ? OR email = ?", username, email).
		Count(&cnt).Error; err != nil {
		return false, err
	}
	return cnt > 0, nil
}
