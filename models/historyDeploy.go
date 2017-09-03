package models

import (
	"time"
)

//HistoryDeploy represent an item of history_deploy
type HistoryDeploy struct {
	ID        	int       `schema:"-"`
	Component   string    `schema:"component"`
	Version     string    `schema:"version"`
	Responsible string    `schema:"responsible"`
	Status      string    `schema:"status"`
	CreatedAt time.Time   `schema:"-"`
	UpdatedAt time.Time   `schema:"-"`
}

func (HistoryDeploy) TableName() string {
  return "history_deploy"
}