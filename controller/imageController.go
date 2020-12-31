package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv13/global"
	"github.com/liuhongdi/digv13/pkg/result"
	"github.com/liuhongdi/digv13/pkg/validCheck"
	"github.com/liuhongdi/digv13/request"
	"strconv"
)

type ImageController struct{}

func NewImageController() ImageController {
	return ImageController{}
}
//上传单张图片
func (a *ImageController) UploadOne(c *gin.Context) {
	resultRes := result.NewResult(c)
	param := request.ArticleRequest{ID: validCheck.StrTo(c.Param("id")).MustUInt64()}
	valid, errs := validCheck.BindAndValid(c, &param)
	if !valid {
		resultRes.Error(400,errs.Error())
		return
	}

    //save image
	f, err := c.FormFile("f1s")
	//错误处理
	if err != nil {
		fmt.Println(err.Error())
		resultRes.Error(1,"图片上传失败")
		} else {
             //将文件保存至本项目根目录中
			  idstr:=strconv.FormatUint(param.ID, 10)
			  destImage := global.ArticleImageSetting.UploadDir+"/"+idstr+".jpg"
              err := c.SaveUploadedFile(f, destImage)
              if (err != nil){
              	  fmt.Println("save err:")
				  fmt.Println(err)
				  resultRes.Error(1,"图片保存失败")
			  } else {
			  	  imageUrl := global.ArticleImageSetting.ImageHost+"/static/ware/article/"+idstr+".jpg"
				  resultRes.Success(gin.H{"url":imageUrl})
			  }
      }
	return
}

//上传多张图片
func (a *ImageController) UploadMore(c *gin.Context) {
	resultRes := result.NewResult(c)
	param := request.ArticleRequest{ID: validCheck.StrTo(c.Param("id")).MustUInt64()}
	valid, errs := validCheck.BindAndValid(c, &param)
	if !valid {
		resultRes.Error(400,errs.Error())
		return
	}

	//save image
	form,err:=c.MultipartForm()
	files:=form.File["f1m"]
	         //错误处理
	 if err != nil {
		          resultRes.Error(1,"图片上传失败")
	             return
	 }
	idstr:=strconv.FormatUint(param.ID, 10)
	var images []string
	         for i,f:=range files{
	             //fmt.Println(f.Filename)
	         	istr := strconv.Itoa(i)
				 destImage := global.ArticleImageSetting.UploadDir+"/"+idstr+"_"+istr+".jpg"
	             c.SaveUploadedFile(f,destImage)
				 //return image url
				 imageUrl := global.ArticleImageSetting.ImageHost+"/static/ware/article/"+idstr+"_"+istr+".jpg"
				 images = append(images, imageUrl)
			 }
	         resultRes.Success(gin.H{"imagesurls":images})
	return
}