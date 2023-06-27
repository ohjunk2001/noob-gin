package controller

import (
	"goGinVue/common"
	"goGinVue/model"
	"goGinVue/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(ctx *gin.Context) {
	db := common.GetDB()
	// 获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	passwd := ctx.PostForm("passwd")
	//数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 442, "msg": "手机号必须为11位!"})
		return
	}
	if len(passwd) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 442, "msg": "密码要大于6位!"})
		return
	}
	if len(name) == 0 {
		name = utils.RandName(10)
	}
	//判断手机号是否存在
	if isTelephoneExits(db, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 442, "msg": "手机号已存在"})
		return
	}
	//创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Passwd:    passwd,
	}
	db.Create(&newUser)
	log.Println(name, telephone, passwd)
	//返回结果
	ctx.JSON(200, gin.H{
		"message": "注册成功",
	})
}

func isTelephoneExits(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
