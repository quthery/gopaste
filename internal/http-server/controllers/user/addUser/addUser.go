package addUser

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	"gopaste/internal/gorm/models"
	"net/http"
)

type UserSaver interface {
	CreateUser(user *models.User) error
}

func New(service UserSaver) func(c *gin.Context) {
	return func(c *gin.Context) {
		validate := validator.New(validator.WithRequiredStructEnabled())
		var userModel models.User
		if err := c.ShouldBindJSON(&userModel); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := validate.Struct(userModel)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := service.CreateUser(&userModel); err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "User exists",
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
	}
}
