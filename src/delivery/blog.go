package delivery

import (
	"api-2/src/helper"
	"api-2/src/middleware"
	"api-2/src/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type blogDelivery struct {
	blogUsecase model.BlogUsecase
	userUsecase model.UserUsecase
}

type BlogDelivery interface {
	Mount(group *gin.RouterGroup)
}

func NewBlogDelivery(ub model.BlogUsecase, uu model.UserUsecase) BlogDelivery {
	return &blogDelivery{
		blogUsecase: ub,
		userUsecase: uu,
	}
}

func (d *blogDelivery) Mount(group *gin.RouterGroup) {
	group.POST("", middleware.Authorize(), d.CreateHanlder)
}

func (d *blogDelivery) CreateHanlder(c *gin.Context) {
	token := c.Request.Header.Get("token")
	req := model.BlogRequest{}
	c.Bind(&req)

	id, errValToken := helper.ValidateToken(token)

	if errValToken != nil {
		c.JSON(401, gin.H{
			"message": "unauthorized",
		})
		return
	}

	userId, _ := strconv.ParseUint(id, 10, 32)

	blog := &model.Blog{
		Title:  req.Title,
		Body:   req.Body,
		UserID: uint(userId),
	}
	err := d.blogUsecase.CreateUsecase(blog)

	if err != nil {
		c.JSON(400, gin.H{
			"status":  "failed",
			"message": "failed",
			"data":    nil,
		})
	} else {
		c.JSON(201, gin.H{
			"status":  "success",
			"message": "success create blog",
			"data":    blog,
		})
	}
}
