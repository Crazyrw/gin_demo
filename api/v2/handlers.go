package v2

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 单个文件上传-application/octet-stream
// @Tags 文件上传
// @Accept mpfd
// @Produce json
// @Param file formData file true "文件"
// @Success 200 {object} v2.Files
// @Router /upload [post]
func SingleUpload(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	//获取文件中的内容
	data := make([]byte, 1<<20)
	n, err := file.Read(data)
	if err != nil {
		if err != io.EOF {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err})
			return
		}
	}
	// 保存文件列表中
	id++
	f := &File{
		ID:      id,
		Name:    fileHeader.Filename,
		Len:     int(fileHeader.Size),
		Content: data[:n],
	}
	files.Files = append(files.Files, f)
	files.Len++
	// 保存到磁盘中
	dst := fmt.Sprintf("E:/%s", fileHeader.Filename)
	c.SaveUploadedFile(fileHeader, dst)
	// 返回数据
	c.JSON(http.StatusOK, files)
}

// @Summary 多个文件上传
// @Tags 文件上传
// @Accept mpfd
// @Produce json
// @Param files formData file true "文件"
// @Success 200 {object} v2.Files
// @Router /multiUpload [post]
func MultiUpload(c *gin.Context) {
	form, err := c.MultipartForm()
	log.Println(form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	fileHeaders := form.File["files"] //这里面的参数是name  不是类型
	log.Println(len(fileHeaders))     //输出为0 说明没有获取到数据
	for _, fileHeader := range fileHeaders {
		log.Println("fileName: ", fileHeader.Filename)
		//打开文件
		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err,
			})
			return
		}
		//获取数据
		data := make([]byte, 8<<20)
		n, err := file.Read(data)
		if err != nil {
			if err != io.EOF {
				c.JSON(http.StatusBadRequest, gin.H{"msg": err})
				return
			}
		}
		//保存到文件列表中
		id++
		f := &File{
			ID:      id,
			Name:    fileHeader.Filename,
			Len:     int(fileHeader.Size),
			Content: data[:n],
		}
		files.Files = append(files.Files, f)
		files.Len++
		//将文件保存到磁盘中
		dst := fmt.Sprintf("E:/%s", fileHeader.Filename)
		c.SaveUploadedFile(fileHeader, dst)
	}
	c.JSON(http.StatusOK, files)
}

// @Summary HTTP重定向
// @Description 当访问/testRedirect的时候,直接转到百度的界面
// @Tags 重定向
// @Router /testRedirect [get]
func RedirectRequest(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/")
}

// @Summary 返回一个json
// @Tags 重定向
// @Produce json
// @Router /returnJson [get]
func ReturnJson(c *gin.Context) {
	name := c.MustGet("name").(string) //从钩子中取值
	c.JSON(http.StatusOK, gin.H{"name": name})
}
