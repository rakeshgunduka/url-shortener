package event

import (
	"url-shortener-go/storage/models"

	"gorm.io/gorm"
)

type EventService struct {
	DB *gorm.DB
}

func (ss *EventService) StoreAnalytics(eventName EventName, userAgent string, OS string, DeviceBrand string, DeviceFamily string, DeviceModel string) {
	event := &models.Event{
		Name:         string(eventName),
		UserAgent:    userAgent,
		OSFamily:     OS,
		DeviceBrand:  DeviceBrand,
		DeviceFamily: DeviceFamily,
		DeviceModel:  DeviceModel,
	}
	ss.DB.Create(event)
}

func (ss *EventService) GetAnalytics() map[string]int64 {
	var clickCount, viewCount, submitCount int64

	ss.DB.Model(&models.Event{}).Select("COUNT(*)").Where("name = ?", "click").Scan(&clickCount)
	ss.DB.Model(&models.Event{}).Select("COUNT(*)").Where("name = ?", "view").Scan(&viewCount)
	ss.DB.Model(&models.Event{}).Select("COUNT(*)").Where("name = ?", "submit").Scan(&submitCount)

	analytics := map[string]int64{
		"clickEvents":  clickCount,
		"viewEvents":   viewCount,
		"submitEvents": submitCount,
	}

	return analytics
}
