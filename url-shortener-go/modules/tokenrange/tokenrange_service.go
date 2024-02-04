package tokenrange

import (
	"os"
	"strconv"
	"url-shortener-go/storage/models"

	"gorm.io/gorm"
)

type TokenRangeService struct {
	DB *gorm.DB
}

type ITokenRangeReponse struct {
	Start int64 `json:"empid"`
	End   int64 `json:"firstname"`
}

func (rs *TokenRangeService) GetNextRange(service string) *models.TokenRange {
	var maxEndRange int64
	rs.DB.Model(&models.TokenRange{}).Select("MAX(end)").Scan(&maxEndRange)

	tokenRangeSize, _ := strconv.ParseInt(os.Getenv("TOKEN_RANGE_SIZE"), 10, 64)
	start := maxEndRange + 1
	end := start + tokenRangeSize

	tokenRange := &models.TokenRange{
		Service: service,
		Start:   start,
		End:     end,
	}

	rs.DB.Create(tokenRange)

	return tokenRange
}
