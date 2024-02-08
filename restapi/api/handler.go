// api/handler.go
package api

import (
	"net/http"
	"restapi/model"
	"restapi/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUsers handles the HTTP GET request to fetch all users.
func GetUsers(c *gin.Context) {
	// Retrieve all users from the repository
	users := repository.GetAllUsers()

	// Respond with a JSON array of users and HTTP status 200 (OK)
	c.JSON(http.StatusOK, users)
}

// GetUser handles the HTTP GET request to fetch a specific user by ID.
func GetUser(c *gin.Context) {
	// Extract the user ID from the request parameters
	userID, _ := strconv.Atoi(c.Param("id"))

	// Retrieve the user from the repository based on the ID
	user, err := repository.GetUserByID(userID)

	// Check for errors during user retrieval
	if err != nil {
		// Respond with an internal server error if an error occurs
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération de l'utilisateur"})
		return
	}

	// Check if the user is not found
	if user == nil {
		// Respond with HTTP status 404 (Not Found) and an error message
		c.JSON(http.StatusNotFound, gin.H{"error": "Utilisateur non trouvé"})
		return
	}

	// Respond with the user details and HTTP status 200 (OK)
	c.JSON(http.StatusOK, user)
}

// CreateUser handles the HTTP POST request to create a new user.
func CreateUser(c *gin.Context) {
	// Declare a variable to hold the incoming JSON data
	var newUser model.User

	// Bind the JSON data from the request body to the newUser variable
	if err := c.BindJSON(&newUser); err != nil {
		// Respond with HTTP status 400 (Bad Request) and an error message if JSON data is invalid
		c.JSON(http.StatusBadRequest, gin.H{"error": "Données JSON invalides"})
		return
	}

	// Attempt to create the new user in the repository
	if err := repository.CreateUser(&newUser); err != nil {
		// Respond with HTTP status 500 (Internal Server Error) and an error message if creation fails
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création de l'utilisateur"})
		return
	}

	// Respond with the created user details and HTTP status 201 (Created)
	c.JSON(http.StatusCreated, newUser)
}

// UpdateUser handles the HTTP PUT request to update an existing user by ID.
func UpdateUser(c *gin.Context) {
	// Extract the user ID from the request parameters
	userID, _ := strconv.Atoi(c.Param("id"))

	// Declare a variable to hold the updated user data
	var updatedUser model.User

	// Bind the JSON data from the request body to the updatedUser variable
	if err := c.BindJSON(&updatedUser); err != nil {
		// Respond with HTTP status 400 (Bad Request) and an error message if JSON data is invalid
		c.JSON(http.StatusBadRequest, gin.H{"error": "Données JSON invalides"})
		return
	}

	// Set the ID of the updated user to the extracted user ID
	updatedUser.ID = uint(userID)

	// Attempt to update the user in the repository
	if err := repository.UpdateUser(&updatedUser); err != nil {
		// Respond with HTTP status 500 (Internal Server Error) and an error message if update fails
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la mise à jour de l'utilisateur"})
		return
	}

	// Respond with the updated user details and HTTP status 200 (OK)
	c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser handles the HTTP DELETE request to delete a user by ID.
func DeleteUser(c *gin.Context) {
	// Extract the user ID from the request parameters
	userID, _ := strconv.Atoi(c.Param("id"))

	// Attempt to delete the user from the repository
	if err := repository.DeleteUser(userID); err != nil {
		// Respond with HTTP status 500 (Internal Server Error) and an error message if deletion fails
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la suppression de l'utilisateur"})
		return
	}

	// Respond with HTTP status 200 (OK) and a success message
	c.JSON(http.StatusOK, gin.H{"message": "Utilisateur supprimé avec succès"})
}
