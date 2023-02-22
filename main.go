package MyTikTok

import (
	"MyTikTok/router"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	router.RoutersInit(engine)
}
