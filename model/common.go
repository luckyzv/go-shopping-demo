package model

import "time"

type PublishModel struct {
  PublishedAt time.Time `json:"publishedAt" gorm:"type:datetime(3);default:NULL;comment:'发布时间'"`
}
