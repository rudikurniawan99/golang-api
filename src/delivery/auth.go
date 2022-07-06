package delivery

import (
	"api-2/src/model"
	"log"

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
	group.POST("test", d.TestRequest)
}

func (d *authDelivery) StoreUserHandler(c *gin.Context) {
	req := model.UserRequest{}
	c.Bind(&req)

	user := model.User{
		Email:    req.Email,
		Password: req.Password,
	}

	err := d.userUsecase.RegisterUsecase(&user)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "failed",
			"data":    nil,
			"error":   err.Error(),
		})
	} else {
		c.JSON(201, gin.H{
			"message": "success",
			"data":    user,
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
			"data":    nil,
			"error":   err.Error(),
		})
	} else {
		errCompare := d.userUsecase.ComparePasswordUsecase(user.Password, req.Password)

		log.Println(errCompare)
		if errCompare != nil {

			c.JSON(400, gin.H{
				"message": "password not match",
				"data":    nil,
				"error":   err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"message": "success",
				"data":    user,
			})
		}
	}
}

func (d *authDelivery) TestRequest(c *gin.Context) {
	req := &model.UserRequest{}
	c.Bind(req)

	log.Printf("== %T", req)

	user := &model.User{
		Email: req.Email,
		// Password: []byte(req.Password),
	}

	log.Printf("== %T", *user)

	c.JSON(200, gin.H{
		"data": gin.H{
			"req":      req,
			"userData": user,
		},
	})
}
