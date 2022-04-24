package database

import (
	"indah_cantik/config"
	// "indah_cantik/middleware"
	"indah_cantik/model"
)

func GetUser() (interface{}, error) {
	var user []model.User

	if e := config.DB.Find(&user).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func GetDetailUser(userId int) (interface{}, error) {
	var user model.User

	if e := config.DB.Find(&user, userId).Error; e != nil {
		return nil, e
	}
	return user, nil
}

// func LoginUser(user model.User) (model.User,error){
// 	var err error
// 	if err = config.DB.Where("email = ? password =? ",user.Email, user.Password).First(&user).Error;err != nil {
// 		return model.User{},err
// 	}
// 	user.Token,err = middleware.CreateToken(user.Email,user.Password)
// 	if err != nil {
// 		return model.User{} , err
// 	}
// 	if err := config.DB.Save(user).Error; err != nil {
// 		return model.User{}, err
// 	}
// 	return user, nil
// }

func Login(email, password string) (*model.User, error) {
	var user model.User

	if err := config.DB.Where("email =? AND Password =?", email, password).First(&user).Error; err != nil {
		return &model.User{}, err
	}
	return &user, nil
}
