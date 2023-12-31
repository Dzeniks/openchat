package chat

import "github.com/gin-gonic/gin"

func InitChat(r *gin.RouterGroup) {
	authGroup := r.Group("/chat")
	{
		authGroup.POST("/ChatCompletetion", ChatCompletetion)
		authGroup.POST("/ChatCompletetionTest", ChatCompletetionTest)
	}
}

func ChatCompletetion(r *gin.Context) {
	dummy := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec euismod, nisl vitae aliquam lacinia, nunc nisl luctus nunc, vitae aliquam nisl nunc eu nisl. Donec euismod, nisl vitae aliquam lacinia, nunc nisl luctus nunc, vitae aliquam nisl nunc eu nisl."
	r.JSON(200, gin.H{
		"message": "OK",
		"output":  dummy,
	})
}

func ChatCompletetionTest(r *gin.Context) {
	dummy := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec euismod, nisl vitae aliquam lacinia, nunc nisl luctus nunc, vitae aliquam nisl nunc eu nisl. Donec euismod, nisl vitae aliquam lacinia, nunc nisl luctus nunc, vitae aliquam nisl nunc eu nisl."
	r.JSON(200, gin.H{
		"message": "OK",
		"output":  dummy,
	})
}
