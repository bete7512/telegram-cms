package repositories

import "github.com/bete7512/telegram-cms/models"



func (db *DB) FindAll() ([]models.User, error) {
	var users []models.User
	if err := db.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}


func (db *DB) FindByID(id int) (models.User, error) {
	var user models.User
	if err := db.db.First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}


func (db *DB) Create(user models.User) (models.User, error) {
	if err := db.db.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}



func (db *DB) Update(user models.User) (models.User, error) {
	if err := db.db.Save(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}


func (db *DB) Delete(id int) (bool, error) {
	if err := db.db.Delete(&models.User{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}



func (db *DB) Login(email string, password string) (models.User, error) {
	var user models.User
	if err := db.db.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (db *DB) FindByEmail(email string) (models.User, error) {
	var user models.User
	if err := db.db.Where("email = ?", email).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}







