package utility

import (
	"errors"
	"os"
)

func GetCloudProjectID() (string, error) {
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	var err error
	if len(projectID) == 0 {
		err = errors.New("project id is missing")
		return "", err
	}
	return projectID, nil
}
