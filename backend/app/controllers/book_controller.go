package controllers

import (
	"backend/app/models"
	"backend/app/queries"
	"backend/pkg/utils"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/labstack/echo/v4"
)

func AddBook(c echo.Context) error {
	userID, err := Authorize(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, utils.FailResponse(err.Error(), nil))
	}
	req := &models.AddBookRequest{}
	err = c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse(err.Error(), nil))
	}
	if req.UserID != userID {
		return c.JSON(http.StatusUnauthorized, utils.FailResponse("Illegal Access", nil))
	}
	book, err := queries.DataBase.CreateBook(req.UserID, req.Name, req.PicURL, req.Price)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.FailResponse(err.Error(), nil))
	}
	return c.JSON(http.StatusOK, utils.SuccessResponse(book))
}

func GetBook(c echo.Context) error {
	req := &models.GetBookRequest{}
	err := c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse(err.Error(), nil))
	}

	if req.UserID == "" {
		books, err := queries.DataBase.GetAllBook()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.FailResponse(err.Error(), nil))
		}
		return c.JSON(http.StatusOK, utils.SuccessResponse(books))
	}
	books, err := queries.DataBase.GetBook(req.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.FailResponse(err.Error(), nil))
	}
	return c.JSON(http.StatusOK, utils.SuccessResponse(books))
}

func UploadIMG(c echo.Context) error {
	// 目前不需要验证就能上传图片，后面再改
	// 通过echo.Contxt实例的FormFile函数获取客户端上传的单个文件
	file, err := c.FormFile("file") //此处的file和前端定义的对应
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.FailResponse(err.Error(), nil))
	}
	// 先打开文件源
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.FailResponse(err.Error(), nil))
	}
	defer src.Close()
	// 下面创建保存路径文件 file.Filename 即上传文件的名字 创建upload文件夹
	uuid := utils.GetUUID()
	extstring := path.Ext(file.Filename) // 获得文件后缀, 自带 .
	dst, err := os.Create("img/" + uuid + extstring)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.FailResponse(err.Error(), nil))
	}
	defer dst.Close()

	// 下面将源拷贝到目标文件
	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.FailResponse(err.Error(), nil))
	}
	urlPrefix := "http://127.0.0.1:8081/img/"
	// PicURL首字母要大写，不然有问题。我记得好像是大写才是public。结构体名字没有首字母大写别的包无法调用
	return c.JSON(http.StatusOK, utils.SuccessResponse(struct{ PicURL string }{urlPrefix + uuid + extstring}))
}
