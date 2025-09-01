package main

import (
  "fmt"
  "net/http"
  "strconv"

  "github.com/gin-gonic/gin"
)

var balance = 1000

func main() {
  // 設定為 release 模式（可選）
  // gin.SetMode(gin.ReleaseMode)
  
  router := gin.Default()
  
  // 添加根路徑處理
  router.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "銀行帳戶 API 服務運行中",
      "endpoints": gin.H{
        "balance":  "GET /balance/",
        "deposit":  "POST /deposit/:amount",
        "withdraw": "POST /withdraw/:amount",
      },
    })
  })
  
  router.GET("/balance/", getBalance)
  router.POST("/deposit/:input", deposit)     
  router.POST("/withdraw/:input", withdraw)   

  fmt.Println("伺服器啟動在 :8082 端口")
  router.Run(":8082")
}

func getBalance(context *gin.Context) {
  fmt.Printf("查詢餘額請求，當前餘額: %d\n", balance)
  var msg = "您的帳戶內有:" + strconv.Itoa(balance) + "元"
  context.JSON(http.StatusOK, gin.H{
    "amount":  balance,
    "status":  "ok",
    "message": msg,
  })
}

func deposit(context *gin.Context) {
  var status string
  var msg string

  input := context.Param("input")
  fmt.Printf("存款請求，輸入參數: %s\n", input)
  
  amount, err := strconv.Atoi(input)

  if err == nil {
    if amount <= 0 {
      amount = 0
      status = "failed"
      msg = "操作失敗，存款金額需大於0元！"
      fmt.Printf("存款失敗：金額 <= 0\n")
    } else {
      balance += amount
      status = "ok"
      msg = "已成功存款" + strconv.Itoa(amount) + "元"
      fmt.Printf("存款成功：%d元，新餘額：%d\n", amount, balance)
    }
  } else {
    amount = 0
    status = "failed"
    msg = "操作失敗，輸入有誤！"
    fmt.Printf("存款失敗：輸入轉換錯誤 - %v\n", err)
  }
  context.JSON(http.StatusOK, gin.H{
    "amount":  amount,
    "status":  status,
    "message": msg,
    "balance": balance,
  })
}

func withdraw(context *gin.Context) {
  var status string
  var msg string

  input := context.Param("input")
  fmt.Printf("提款請求，輸入參數: %s\n", input)
  
  amount, err := strconv.Atoi(input)

  if err == nil {
    if amount <= 0 {
      amount = 0
      status = "failed"
      msg = "操作失敗，提款金額需大於0元！"
      fmt.Printf("提款失敗：金額 <= 0\n")
    } else {
      if balance-amount < 0 {
        amount = 0
        status = "failed"
        msg = "操作失敗，餘額不足！"
        fmt.Printf("提款失敗：餘額不足，當前餘額：%d\n", balance)
      } else {
        balance -= amount
        status = "ok"
        msg = "成功提款" + strconv.Itoa(amount) + "元"
        fmt.Printf("提款成功：%d元，新餘額：%d\n", amount, balance)
      }
    }
  } else {
    amount = 0
    status = "failed"
    msg = "操作失敗，輸入有誤！"
    fmt.Printf("提款失敗：輸入轉換錯誤 - %v\n", err)
  }
  context.JSON(http.StatusOK, gin.H{
    "amount":  amount,
    "status":  status,
    "message": msg,
    "balance": balance,
  })
}
