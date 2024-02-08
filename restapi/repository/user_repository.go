// user_repository.go
package repository

import (
	"errors"
	"restapi/model"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	// Initialize GORM with a MySQL database connection
	// Replace 'your-dsn' with your MySQL connection string
	dsn := "host=localhost port=5432 dbname=sal user=your password=password  connect_timeout=10"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// Auto migrate the User model to create the 'users' table
	db.AutoMigrate(&model.User{})
	newUser := model.User{
		Name: "John Doe",
		// Set other fields as needed
	}

	// Save the user to the database
	db.Create(&newUser)

	// Check for errors
	if db.Error != nil {
		panic(db.Error)
	}
}

func GetAllUsers() []model.User {
	var users []model.User
	db.Find(&users)
	return users
}

func GetUserByID(id int) (*model.User, error) {
	var user model.User
	result := db.First(&user, id)
	if result.Error != nil {
		return nil, errors.New("Utilisateur non trouvé")
	}
	return &user, nil
}

func CreateUser(user *model.User) error {
	result := db.Create(user)
	if result.Error != nil {
		return errors.New("Erreur lors de la création de l'utilisateur")
	}
	return nil
}

func UpdateUser(updatedUser *model.User) error {
	result := db.Save(updatedUser)
	if result.Error != nil {
		return errors.New("Erreur lors de la mise à jour de l'utilisateur")
	}
	return nil
}

func DeleteUser(id int) error {
	result := db.Delete(&model.User{}, id)
	if result.Error != nil {
		return errors.New("Erreur lors de la suppression de l'utilisateur")
	}
	return nil
}
