package shorturl

import (
	"strconv"
	"time"
	"url-shortener-go/storage/models"
	"url-shortener-go/utils"

	"gorm.io/gorm"
)

type ShortUrlService struct {
	DB *gorm.DB
}

func (ss *ShortUrlService) GenerateUrlAlias(url string) string {
	tokenCounter := utils.GetTokenCounter()
	utils.IncrementTokenCounter()

	base62URL := utils.ConvertToBase62(strconv.FormatInt(tokenCounter, 10))

	hash := utils.GenerateHashWithSalt(base62URL, url)
	truncatedHash := hash[:7]

	created := time.Now()
	entry := models.UrlAliasMap{
		URL:     url,
		Alias:   truncatedHash,
		Created: created,
		Expiry:  created.Add(time.Hour * 24 * 7),
	}

	ss.DB.Create(&entry)

	return truncatedHash
}

func (ss *ShortUrlService) GetOriginalUrl(alias string) (models.UrlAliasMap, error) {
	var entry models.UrlAliasMap
	result := ss.DB.Model(&models.UrlAliasMap{}).Where("alias = ?", alias).First(&entry)
	if result.Error != nil {
		return models.UrlAliasMap{}, result.Error
	}

	return entry, nil
}

func (ss *ShortUrlService) GetShortUrl(url string) (string, error) {
	var urlMap models.UrlAliasMap
	result := ss.DB.Model(&models.UrlAliasMap{}).Where("url = ?", url).First(&urlMap)
	if result.Error != nil {
		return "", result.Error
	}
	return urlMap.Alias, nil
}

func (ss *ShortUrlService) GetShortUrls() ([]models.UrlAliasMap, error) {
	var urlMap []models.UrlAliasMap
	result := ss.DB.Model(&models.UrlAliasMap{}).Find(&urlMap)
	if result.Error != nil {
		return nil, result.Error
	}
	return urlMap, nil
}
