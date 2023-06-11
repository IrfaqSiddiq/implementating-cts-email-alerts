package controllers

import (
	"context"
	"cts-alerts/utility"
	"log"

	talent "cloud.google.com/go/talent/apiv4beta1"
	"github.com/gin-gonic/gin"
)

func GetJobServiceClient(c *gin.Context) (string, *talent.JobClient, error) {
	projectID, err := utility.GetCloudProjectID()
	if err != nil {
		log.Println("GetJobServiceClient: failed while fetching project id with err", err)
		return "", &talent.JobClient{}, err
	}
	ctx := context.Background()
	jobClient, err := talent.NewJobClient(ctx)
	if err != nil {
		log.Printf("Failed to create job client: %v\n", err)
		return "", &talent.JobClient{}, err
	}
	return projectID, jobClient, nil
}
