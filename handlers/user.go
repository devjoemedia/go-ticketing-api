package handlers

import (
	"net/http"

	"github.com/devjoemedia/chitodopostgress/database"
	"github.com/devjoemedia/chitodopostgress/models"
	"github.com/devjoemedia/chitodopostgress/types"
	api_response "github.com/devjoemedia/chitodopostgress/types/response"
	"github.com/devjoemedia/chitodopostgress/utils"
	"github.com/go-chi/chi/v5"
)

// GetUsers godoc
// @Summary      Get users with pagination and filters
// @Description  Retrieve users with optional pagination (page, size)
// @Tags         users
// @Security BearerAuth
// @Produce      json
// @Param        page        query    int    false  "Page number (default: 1)"
// @Param        size        query    int    false  "Page size (default: 10, max: 100)"
// @Success      200 {object} api_response.GetUsersResponse
// @Router       /api/v1/users [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	database.DB.Find(&users)

	utils.JSON(w, http.StatusOK, api_response.GetUsersResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Users retrieved successfully",
		Users:   users,
		Pagination: types.Pagination{
			Page:  1,
			Size:  10,
			Total: 10,
			Pages: 10,
		},
	})
}

// GetUser godoc
// @Summary      Get user by ID
// @Description  Fetch a specific user by its ID
// @Tags         users
// @Security BearerAuth
// @Produce      json
// @Param        id   path    int  true  "user ID"
// @Success      200  {object} api_response.GetUserResponse
// @Failure      400  {string} string     "Invalid ID"
// @Failure      404  {string} string     "User not found"
// @Router       /api/v1/users/{id} [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var user models.User
	if result := database.DB.First(&user, id); result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	utils.JSON(w, http.StatusOK, api_response.GetUserResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "User retrieved successfully",
		User:    &user,
	})
}

// UpdateUser godoc
// @Summary      Update an existing user
// @Description  Update a user item by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id     path    int  true  "User ID"
// @Param        body   body    models.UpdateUserRequest  true  "User object"
// @Success      200    {object} api_response.UpdateUserResponse
// @Failure      400    {string} string  "Invalid JSON"
// @Failure      404    {string} string  "User not found"
// @Router       /api/v1/users/{id} [put]
// @Router       /api/v1/users/{id} [patch]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var user models.User
	if result := database.DB.First(&user, id); result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	utils.JSON(w, http.StatusOK, api_response.GetUserResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "User retrieved successfully",
		User:    &user,
	})
}

// DeleteUser godoc
// @Summary      Delete user by ID
// @Description  Delete a specific user by its ID
// @Tags         users
// @Security BearerAuth
// @Produce      json
// @Param        id   path    int  true  "User ID"
// @Success      200  {object} api_response.DeleteUserResponse
// @Failure      400  {string} string             "Invalid ID"
// @Failure      404  {string} string             "User not found"
// @Router       /api/v1/users/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	database.DB.Delete(&models.User{}, id)

	response := api_response.DeleteUserResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "User deleted successfully",
	}

	utils.JSON(w, http.StatusOK, response)
}
