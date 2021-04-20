package data

import (
	"errors"

	"github.com/jinzhu/gorm"
)

//DBGetBooks return all the books store in database
func (handler *SQLHandler) DBGetUsers() ([]User, error) {

	Users := []User{}
	result := handler.DB.Debug().Find(&Users)
	if result.Error != nil {
		//Log.ErrorLog.Error(result.Error)
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("no row effected")
	}
	return Users, nil
}

//DBGetBookByID return just a single book correspond to BookID
func (handler *SQLHandler) DBGetUserByID(UserName string) (User, error) {
	user := User{}
	if result := handler.DB.Where("name = ?", UserName).First(&user); result.Error != nil || result.Error == gorm.ErrRecordNotFound {
		return User{}, result.Error
	}
	return user, nil
}
