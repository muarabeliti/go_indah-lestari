package database

import (
	"project/config"
	// "project/middleware"
	"project/model"
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



func Login(username string) (string, error) {
	var  err error

	if err = config.DB.Where("email =? AND Password =?", user.Email, password).First(&user).Error; err != nil {
		return &string{}, err
	}
	return &user, nil
}

//func Login(email, password string) (*model.User, error) {
//	var  model.User

//	if err := config.DB.Where("email =? AND Password =?", email, password).First(&user).Error; err != nil {
//		return &model.User{}, err
//	}
//	return &user, nil
/}
