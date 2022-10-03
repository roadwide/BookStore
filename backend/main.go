package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	_ "github.com/mattn/go-sqlite3" //只init，但是不调用里面的函数
)

var dbName string = "./sqlite.db"
var dbTalbe string = "UserInfo"
var bookTalbe string = "BookInfo"

type BookInfo struct {
	OrderId   string
	UserId    string
	BookName  string
	BookPrice float64
	PicURL    string
}

func main() {
	initSQL()
	r := gin.Default()
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.Use(cors.Default())          //允许CORS
	r.POST("/register", handleRegister)
	r.POST("/login", handleLogin)
	r.POST("/token", handleParseToken)
	r.POST("/upload", handleUpload)
	r.POST("/addBook", handleAddBook)
	r.GET("/getAllBookInfo", getAllBookInfo)
	r.GET("/getBookInfoByUserId", getBookInfoByUserId)
	r.Static("/img", "./img") //Static不会显示目录，StaticFS会显示文件夹里的目录(FS=File System)

	r.Run("0.0.0.0:8081") // 监听并在 0.0.0.0:8081 上启动服务

}

//处理前端发送来的注册表单
func handleRegister(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	if insertUser(username, password, email) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "UserId Exists",
		})
	}
}

//处理登录请求
func handleLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if isExistUser(username) && validPassWord(username, password) {
		c.JSON(200, gin.H{
			"message": "ok",
			"userID":  username,
			"token":   genToken(username, "mlg123"),
		})
	} else {
		c.JSON(200, gin.H{
			"message": "error",
		})
	}
}

//检测是否存在数据库文件，初始化数据库
func initSQL() {
	if _, err := os.Stat(dbName); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			fmt.Println("Database not exists, creating Database")
			db, err := sql.Open("sqlite3", dbName)
			if err != nil {
				log.Fatal(err)
			}
			// 创建用户表 和 书目表
			sqlStmt := fmt.Sprintf("create table %s (UserId text primary key, PassWord text, Email text);"+
				"create table %s (OrderId text primary key, UserId text, BookName text, BookPrice real, PicURL text);", dbTalbe, bookTalbe)
			_, err = db.Exec(sqlStmt)
			if err != nil {
				log.Printf("%q: %s\n", err, sqlStmt)
			}
		} else {
			// other error
			fmt.Println("initSQL error")
		}
	} else {
		//exist
		fmt.Println("Database exists")
	}
}

//连接数据库，返回连接句柄
func connectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err //出现问题，返回错误
	}
	return db, nil //没有问题，返回数据库连接
}

//插入一条新的用户数据, 如果已经存在用户名返回false
func insertUser(UserId string, PassWord string, Email string) (validUser bool) {
	validUser = !isExistUser(UserId) //如果不存在就是valid的
	if !validUser {
		return false
	}
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//插入数据
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare(fmt.Sprintf("insert into %s (UserId, PassWord, Email) values(?, ?, ?)", dbTalbe))
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() //最后关闭statement
	_, err = stmt.Exec(UserId, PassWord, Email)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
	return true
}

//判断用户名是否存在于数据库中
func isExistUser(UserId string) bool {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("select count(*) from UserInfo where UserId = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var count int
	err = stmt.QueryRow(UserId).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	if count == 0 {
		return false
	} else {
		return true
	}
}

//判断密码是否正确
func validPassWord(UserId string, PassWord string) bool {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("select PassWord from UserInfo where UserId = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var RealPass string
	err = stmt.QueryRow(UserId).Scan(&RealPass)
	if err != nil {
		log.Fatal(err)
	}
	if RealPass == PassWord {
		return true
	} else {
		return false
	}
}

//生成指定用户名的token
func genToken(UserId string, secret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": UserId,
		"exp":      time.Now().Unix() + 24*60*60*7, // 过期时间，内部会进行自动判断
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Println(err)
	}

	return tokenString
}

//解析token返回储存的用户名
func parseToken(tokenString string, secret string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		log.Println(err)
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["username"].(string), nil
	} else {
		return "", errors.New("error token")
	}
}

//解析token的http接口，返回用户名或错误
func handleParseToken(c *gin.Context) {
	token := c.PostForm("token")
	userID := c.PostForm("userID")
	username, err := parseToken(token, "mlg123")
	if err != nil || userID != username {
		c.JSON(200, gin.H{
			"message": "error",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "ok",
			"userID":  username,
		})
	}
}

// 上传图片
func handleUpload(c *gin.Context) {
	uuid := getUUID()

	// 单文件
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	extstring := path.Ext(file.Filename) // 获得文件后缀, 自带 .
	dst := "./img/" + uuid + extstring
	// 上传文件至指定的完整文件路径
	c.SaveUploadedFile(file, dst)
	urlPrefix := "http://127.0.0.1:8081/img/"
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
		"picURL":  urlPrefix + uuid + extstring,
		"uuid":    uuid,
	})
}

func getUUID() string {
	rand.Seed(time.Now().UnixNano())
	uuid := time.Now().Format("20060102150405") + fmt.Sprintf("%d", rand.Int31n(1000000))
	return uuid
}

// 处理新增书目信息
// 处理前端发送来的注册表单
func handleAddBook(c *gin.Context) {
	orderId := c.PostForm("orderId")
	userId := c.PostForm("userId")
	bookName := c.PostForm("bookName")
	bookPrice, _ := strconv.ParseFloat(c.PostForm("bookPrice"), 64)
	picURL := c.PostForm("picURL")
	isSuccess := addBook(orderId, userId, bookName, bookPrice, picURL)
	if isSuccess {
		c.JSON(200, gin.H{
			"message":   "ok",
			"orderId":   orderId,
			"bookName":  bookName,
			"bookPrice": bookPrice,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "wrong",
		})
	}

}

// 插入一条新的数目信息，插入成功返回true
func addBook(OrderId string, UserId string, BookName string, BookPrice float64, PicURL string) bool {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//插入数据
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare(fmt.Sprintf("insert into %s (OrderId, UserId, BookName, BookPrice, PicURL) values(?, ?, ?, ?, ?)", bookTalbe))
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() //最后关闭statement
	_, err = stmt.Exec(OrderId, UserId, BookName, BookPrice, PicURL)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
	return true
}

// 查询数据库获得数目信息
func queryBookInfo(isAll bool, UserId string) (ret []BookInfo) {
	// isAll 判断是否是查询所有用户
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var rows *sql.Rows
	if isAll {
		rows, err = db.Query(fmt.Sprintf("select OrderId, UserId, BookName, BookPrice, PicURL from %s", bookTalbe))
	} else {
		rows, err = db.Query(fmt.Sprintf("select OrderId, UserId, BookName, BookPrice, PicURL from %s where UserId = %s", bookTalbe, UserId))
	}
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var OrderId string
		var UserId string
		var BookName string
		var BookPrice float64
		var PicURL string
		err = rows.Scan(&OrderId, &UserId, &BookName, &BookPrice, &PicURL)
		if err != nil {
			log.Fatal(err)
		}
		data := BookInfo{
			OrderId:   OrderId,
			UserId:    UserId,
			BookName:  BookName,
			BookPrice: BookPrice,
			PicURL:    PicURL,
		}
		ret = append(ret, data)
	}
	return ret
}

// 获得所有书目信息
func getAllBookInfo(c *gin.Context) {
	ret := queryBookInfo(true, "")
	c.JSON(200, gin.H{
		"message": "ok",
		"data":    ret,
	})

}

// 获得指定用户的书目信息
func getBookInfoByUserId(c *gin.Context) {
	UserId := c.Query("UserId")
	if isExistUser((UserId)) {
		ret := queryBookInfo(false, UserId)
		c.JSON(200, gin.H{
			"message": "ok",
			"data":    ret,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "wrong",
		})
	}

}
