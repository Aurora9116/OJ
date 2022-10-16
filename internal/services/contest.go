package services

import (
	"hgoj/internal/dao"
	"hgoj/internal/models"
	"time"
)

type Contest struct {
}

// QueryContestByConId 通过contest_id查找一条记录
func (con *Contest) QueryContestByConId(contestID int32) (*models.Contest, error) {
	contestInfo := new(models.Contest)
	err := dao.Orm.Debug().First(contestInfo, contestID).Error

	return contestInfo, err

}

// GetDoingContest 比较结束时间得到正在进行中的比赛
func (con *Contest) GetDoingContest() (*[]*models.Contest, error) {
	contestInfo := new([]*models.Contest)
	err := dao.Orm.Debug().Where("end_time < ?", time.Now().Format("2006-01-02 15:04:05")).Find(contestInfo).Error
	return contestInfo, err
}
