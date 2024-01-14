package video

import (
	"douyin/v1/models"
	"douyin/v1/service/video"
	util2 "douyin/v1/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

// key value 判断是否是视频/图片
var (
	videoIndexMap = map[string]struct{}{
		".mp4":  {},
		".avi":  {},
		".wmv":  {},
		".flv":  {},
		".mpeg": {},
		".mov":  {},
	}
	pictureIndexMap = map[string]struct{}{
		".jpg": {},
		".bmp": {},
		".png": {},
		".svg": {},
	}
)

// PublishVideoHandler 发布视频，并截取一帧画面作为封面
func PublishVideoHandler(c *gin.Context) {
	//准备参数
	rawId, _ := c.Get("user_id")
	//判断是否是int64
	userId, ok := rawId.(int64)
	if !ok {
		PublishVideoError(c, "解析UserId出错")
		return
	}
	//form-data里拿出数据来
	title := c.PostForm("title")
	form, err := c.MultipartForm()
	if err != nil {
		PublishVideoError(c, err.Error())
		return
	}
	//支持多文件上传
	files := form.File["data"]
	for _, file := range files {
		suffix := filepath.Ext(file.Filename)    //得到后缀
		if _, ok := videoIndexMap[suffix]; !ok { //判断是否为视频格式
			PublishVideoError(c, "不支持的视频格式")
			continue
		}
		name := util2.NewFileName(userId) //根据userId得到唯一的文件名
		filename := name + suffix
		//加前缀
		savePath := filepath.Join("./static", filename)
		//保存文件api
		err = c.SaveUploadedFile(file, savePath)
		if err != nil {
			PublishVideoError(c, err.Error())
			continue
		}
		//截取一帧画面作为封面
		err = util2.SaveImageFromVideo(name, true)
		if err != nil {
			PublishVideoError(c, err.Error())
			continue
		}
		//数据库持久化
		err := video.PostVideo(userId, filename, name+util2.GetDefaultImageSuffix(), title)
		if err != nil {
			PublishVideoError(c, err.Error())
			continue
		}
		//json返回gin
		PublishVideoOk(c, file.Filename+"上传成功")
	}
}

func PublishVideoError(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, models.CommonResponse{StatusCode: 1,
		StatusMsg: msg})
}

func PublishVideoOk(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, models.CommonResponse{StatusCode: 0, StatusMsg: msg})
}
