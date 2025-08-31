package main
import (
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context){
		ctx.HTML(200, "index.html", gin.H{})
	})

	// 向模板文件中傳值
	router.GET("/data", func(ctx *gin.Context){
		ctx.HTML(200, "data.html", gin.H{
			"data": "Hello go/gin World!"})
	})

	// 返回json格式數據
	router.GET("/json", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"result": "ok",
			"data": "Hello Go/Gin World.",
			"developer": "tsai",
		})
	})

	// POST提交, JSON返回(接收用戶表單)
	router.GET("/form", func(ctx *gin.Context){
		ctx.HTML(200, "form.html", gin.H{})
	})
	router.POST("/service", func(ctx *gin.Context){
		uname := ctx.PostForm("uname")
		ctx.JSON(200, gin.H{
			"result": "ok",
			"Hello": uname,
		})
	})

	router.Run(":8000")
}