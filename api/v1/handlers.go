package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary URL queryString参数解析
// @Description 参数为user的数据结构
// @Tags 测试
// @Accept mpfd
// @Produce json
// @Param userName query string true "人名" default("")
// @Param password query string true "密码" default("")
// @Success 200 {string} string "{"userName": "renwen", "password": "12345", "status": "OK"}"
// @Failure 400 {string} string "{"msg": "error"}"
// @Router /hello [get]
func HandleHello(c *gin.Context) {
	var user = new(User)
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "error",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"userName": user.UserName,
		"password": user.Password,
		"status":   "OK",
	})
}

// @Summary form参数解析
// @Description 发送表单数据(格式为User数据结构)，返回收集到的表单数据 json格式
// @Tags JSON
// @Accept mpfd
// @Produce json
// @Param userName formData string true "姓名" default("")
// @Param password formData string true "密码" default("")
// @Success 200 {string} string "{"userName": "renwen", "password": "12345", "status": "OK"}"
// @Failure 400 {string} string "{"msg": "error"}"
// @Router /formData [post]
func HandleForm(c *gin.Context) {
	user := User{}
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "error",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"userName": user.UserName,
		"password": user.Password,
		"status":   "OK",
	})
}

// @Summary 获取JSON的示例
// @Tags JSON
// @Description 输入User的数据结构参数, 返回JSON格式数据
// @Accept json
// @Produce json
// @Param param body map[string]interface{} true "需要上传的JSON"
// @Success 200 {object} map[string]interface{} "返回"
// @Router /json [post]
func GetJson(c *gin.Context) {
	user := User{}
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "error",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"userName": user.UserName,
		"password": user.Password,
		"status":   "OK",
	})
}
