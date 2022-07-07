package delivery

import (
	"api-2/src/helper"
	"api-2/src/model"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type authDelivery struct {
	userUsecase model.UserUsecase
}

type AuthDelivery interface {
	Mount(group *gin.RouterGroup)
}

func NewDelivery(uc model.UserUsecase) AuthDelivery {
	return &authDelivery{
		userUsecase: uc,
	}
}

func (d *authDelivery) Mount(group *gin.RouterGroup) {
	group.POST("register", d.StoreUserHandler)
	group.POST("login", d.LoginHanler)
	group.GET("me", d.GetCurrentUserHanlder)
}

func (d *authDelivery) GetCurrentUserHanlder(c *gin.Context) {
	token := c.Request.Header.Get("token")

	id, errValToken := helper.ValidateToken(token)
	if errValToken != nil {
		c.JSON(401, gin.H{
			"message": "not authorized",
		})
		return
	}

	user := &model.User{}
	err := d.userUsecase.FindUserById(user, id)

	if err != nil {
		c.JSON(401, gin.H{
			"message": "not authorized",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "authorized",
			"data":    user,
		})
	}
}

func (d *authDelivery) StoreUserHandler(c *gin.Context) {
	validate := validator.New()
	req := model.UserRequest{}
	c.Bind(&req)

	errValidation := validate.Struct(req)
	if errValidation != nil {
		c.JSON(400, gin.H{
			"message": "failed request",
			"error":   errValidation.Error(),
		})
		return
	}

	user := model.User{
		Email:    req.Email,
		Password: req.Password,
	}

	errEmailExist := d.userUsecase.FindUserByEmailUsecase(&user)
	if errEmailExist == nil {
		c.JSON(400, gin.H{
			"message": "failed, email already exist",
		})
		return
	}

	err := d.userUsecase.RegisterUsecase(&user)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "failed",
			"error":   err.Error(),
		})
	} else {
		token, _ := helper.GenerateToken(int(user.ID))

		c.JSON(201, gin.H{
			"message": "success",
			"token":   token,
		})
	}
}

func (d *authDelivery) LoginHanler(c *gin.Context) {
	validate := validator.New()
	req := &model.UserRequest{}

	c.Bind(req)

	// add validation
	errValidation := validate.Struct(req)

	if errValidation != nil {
		c.JSON(400, gin.H{
			"message": "failed request",
			"error":   errValidation.Error(),
		})
		return
	}

	user := model.User{
		Email:    req.Email,
		Password: req.Password,
	}

	err := d.userUsecase.FindUserByEmailUsecase(&user)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "email not found",
			"error":   err.Error(),
		})
	} else {
		errCompare := d.userUsecase.ComparePasswordUsecase(user.Password, req.Password)

		if errCompare != nil {
			c.JSON(400, gin.H{
				"message": "password not match",
				"error":   err.Error(),
			})
		} else {
			token, _ := helper.GenerateToken(int(user.ID))

			c.JSON(200, gin.H{
				"message": "success",
				"token":   token,
			})
		}
	}
}
