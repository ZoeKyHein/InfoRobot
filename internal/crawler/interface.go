package crawler

import "InfoRobot/models"

type Crawler interface {
	FetchData() (models.Data, error)
}
