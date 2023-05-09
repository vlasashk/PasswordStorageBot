package storage

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func ConnectGorm() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("../../internal/storage/users.db"), &gorm.Config{})
	if err != nil {
		log.Panic("Failed to open the SQLite database.")
	}

	// Create the table from our struct.
	err = db.AutoMigrate(&UserStorage{}, &ServiceStorage{})
	if err != nil {
		log.Panic("Failed to migrate.")
	}
	return db
}

func AddUser(db *gorm.DB, ID int64) {
	user := UserStorage{
		UserID: ID,
	}
	db.FirstOrCreate(&user, &user)
}

func AddCredentials(db *gorm.DB, user User) {
	var storage ServiceStorage
	if db.First(&storage, "user_id = ? and name = ?", user.UserID, user.CurrServ).Error != nil {
		db.Save(&ServiceStorage{
			Name:     user.CurrServ,
			UserID:   user.UserID,
			Login:    user.ServiceName.Login,
			Password: user.ServiceName.Password,
		})
	}
}
func CheckServiceLimit(db *gorm.DB, userID int64) (res bool) {
	var count int64
	db.Model(&ServiceStorage{}).Where("user_id = ?", userID).Count(&count)
	if count < 20 {
		res = true
	}
	log.Println("User ", userID, " has ", count, " services")
	return
}

func ServiceExist(db *gorm.DB, userID int64, service string) (res bool) {
	var storage ServiceStorage
	if db.First(&storage, "user_id = ? and name = ?", userID, service).Error != nil {
		res = false
	} else {
		res = true
	}
	return
}

func GetServicesByUser(db *gorm.DB, userID int64) map[string]Service {
	services := make([]ServiceStorage, 0)
	if err := db.Where("user_id = ?", userID).Find(&services).Error; err != nil {
		return make(map[string]Service)
	}
	result := make(map[string]Service)
	for _, service := range services {
		result[service.Name] = Service{
			Login:    service.Login,
			Password: service.Password,
		}
	}
	return result
}

func GetService(db *gorm.DB, userID int64, find string) Service {
	var services ServiceStorage
	db.Where("user_id = ? and name =?", userID, find).Find(&services)
	var result Service
	result = Service{
		Login:    services.Login,
		Password: services.Password,
	}
	return result
}

func DeleteService(db *gorm.DB, userID int64, serviceName string) error {
	if err := db.Where("user_id = ? AND name = ?", userID, serviceName).Delete(&ServiceStorage{}).Error; err != nil {
		return err
	}
	return nil
}
