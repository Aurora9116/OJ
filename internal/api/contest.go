package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
	"hgoj/internal/dhttp"
	"hgoj/internal/models"
	"hgoj/internal/services"
	"net/http"
)

func AddContest(ctx *gin.Context) {
	data := models.Contest{}
	err := ctx.ShouldBindBodyWith(&data, binding.JSON)
	if err != nil {
		erro := fmt.Sprintf("%s", err)
		ctx.JSON(http.StatusOK, gin.H{
			"msg": erro,
		})
		logrus.Info("[", err, "]")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "successful",
		"data": data,
	})
}

// GetDoingContest 得到所有正在进行中的比赛
func GetDoingContest(ctx *gin.Context) {
	Contest := new(services.Contest)
	contestInfo, err := Contest.GetDoingContest()
	if err != nil {
		fmt.Printf("[api.GetDoingContest] err:%v", err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":             dhttp.OK,
		"DoingContestData": contestInfo,
	})
}
