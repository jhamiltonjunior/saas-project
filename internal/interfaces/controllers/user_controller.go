package controller

import (
    "my-saas-app/internal/application/services"
    "my-saas-app/internal/domain/entities"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type UserController struct {
    userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
    return &UserController{userService: userService}
}

func (c *UserController) GetUserByID(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    user, err := c.userService.GetUserByID(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    ctx.JSON(http.StatusOK, user)
}

func (c *UserController) CreateUser(ctx *gin.Context) {
    var user entities.User
    if err := ctx.ShouldBindJSON(&user); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    if err := c.userService.CreateUser(&user); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    var user entities.User
    if err := ctx.ShouldBindJSON(&user); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    user.ID = id
    if err := c.userService.UpdateUser(&user); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
        return
    }

    ctx.JSON(http.StatusOK, user)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    if err := c.userService.DeleteUser(id); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}