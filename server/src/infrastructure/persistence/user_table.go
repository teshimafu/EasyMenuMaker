package persistence

import (
	"gorm.io/gorm"

	"github.com/teshimafu/lazyPM/server/src/domain/entity"
	"github.com/teshimafu/lazyPM/server/src/domain/repository"
	value "github.com/teshimafu/lazyPM/server/src/domain/valueobject"
	"github.com/teshimafu/lazyPM/server/src/infrastructure/model"
)

type userTable struct {
	db *gorm.DB
}

func NewUserTable(db *gorm.DB) repository.IUserRepository {
	return &userTable{
		db: db,
	}
}

func (ut *userTable) FindByID(id *value.UserID) (*entity.User, error) {
	var user *model.User
	if err := ut.db.Where("user_id = ?", id.Buffer()).First(&user).Error; err != nil {
		return nil, err
	}
	return toUserEntity(user)
}

func (ut *userTable) FindByEmail(email *value.Email) (*entity.User, error) {
	var user *model.User
	if err := ut.db.Where("email = ?", email.Value()).First(&user).Error; err != nil {
		return nil, err
	}
	return toUserEntity(user)
}

func (ut *userTable) Find(input *entity.User) (*entity.User, error) {
	var user *model.User
	if err := ut.db.Where(
		"name = ?", input.Name().Value()).
		Or("email = ?", input.Email().Value()).
		First(&user).Error; err != nil || user == nil {
		return nil, err
	}
	return toUserEntity(user)
}

func (ut *userTable) FindAll() ([]*entity.User, error) {
	var users []*model.User
	if err := ut.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return toUserEntities(users)
}

func (ut *userTable) Create(user *entity.User) (*entity.User, error) {
	userModel := toUserModel(user)
	result := ut.db.Create(&userModel)
	if result.Error != nil {
		return nil, result.Error
	}
	return toUserEntity(userModel)
}

func toUserModel(user *entity.User) *model.User {
	return &model.User{
		UserID:   user.ID().Buffer(),
		Name:     user.Name().Value(),
		Email:    user.Email().Value(),
		Password: user.Password().Hash(),
	}
}

func toUserEntity(user *model.User) (*entity.User, error) {
	return entity.NewUser(string(user.UserID), user.Name, user.Email, user.Password)

}

func toUserEntities(users []*model.User) ([]*entity.User, error) {
	entities := make([]*entity.User, len(users))
	for i, user := range users {
		user, err := toUserEntity(user)
		if err != nil {
			return nil, err
		}
		entities[i] = user
	}
	return entities, nil

}
