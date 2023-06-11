package controllers

import (
	"context"
	"cts-alerts/models"
	"cts-alerts/services"
	"fmt"
	"log"
	"net/http"

	talentpb "cloud.google.com/go/talent/apiv4beta1/talentpb"
	"github.com/gin-gonic/gin"
)

func SendEmailAlerts(c *gin.Context) {
	agentId := c.PostForm("agent_id")
	email := c.PostForm("email")
	fmt.Println("agentID", agentId)
	//fetch job role according to agentID and then pass it to SearchJobsRequest
	var ctsJobs []models.CTSJobs
	ctx := context.Background()
	projectID, jobClient, err := GetJobServiceClient(c)
	if err != nil {
		log.Printf("AddJobsOnCloud: Failed to create job client: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed while creating new job service",
		})
		return
	}
	defer jobClient.Close()
	//pass jobCategory according to user's role
	req := &talentpb.SearchJobsRequest{
		Parent: fmt.Sprintf("projects/%s", projectID),
		JobQuery: &talentpb.JobQuery{
			JobCategories: []talentpb.JobCategory{
				6, //BUSINESS_OPERATIONS
				8, //COMPUTER_AND_IT
			},
		},
		OrderBy:  "annualized_base_compensation desc",
		PageSize: 5,
	}
	// Execute the list jobs request.
	resp, err := jobClient.SearchJobs(ctx, req)
	if err != nil {
		fmt.Printf("Failed to list jobs: %v\n", err)
		return
	}
	for _, job := range resp.GetMatchingJobs() {
		fmt.Printf("Job name: %s\n", job.Job.GetTitle())
		fmt.Printf("Company name: %s\n", job.Job.GetCompany())
		fmt.Printf("Requisition ID: %s\n", job.Job.RequisitionId)
		fmt.Printf("Title: %s\n", job.Job.GetTitle())
		fmt.Printf("Description: %s\n", job.Job.Description)
		fmt.Printf("Job posting URL: %s\n", job.Job.ApplicationInfo.Uris[0])
		fmt.Println("------------------------------------")
		ctsJobs = append(ctsJobs, models.CTSJobs{
			JobTitle:    job.Job.Title,
			Description: job.Job.Description,
			JobID:       job.Job.RequisitionId,
			JobURL:      job.Job.ApplicationInfo.Uris[0],
		})
	}
	pData := map[string]interface{}{
		"jobs": ctsJobs,
	}
	requestData := services.Personalizations{
		To:          []string{email},
		DynamicData: pData,
		TemplateID:  "d-51c3d100bcbb4088a5dc53f8f2ee6e44",
	}
	services.ComposeDynamicTemplateEmail(requestData)
}

// searchForAlerts searches for jobs with email alert set which could receive
// updates later if search result updates.
// func SendEmailAlerts(c *gin.Context) {
// 	ctx := context.Background()

// 	client, err := google.DefaultClient(ctx, talent.CloudPlatformScope)
// 	if err != nil {
// 		return nil, fmt.Errorf("google.DefaultClient: %w", err)
// 	}
// 	// Create the jobs service client.
// 	service, err := talent.New(client)
// 	if err != nil {
// 		return nil, fmt.Errorf("talent.New: %w", err)
// 	}

// 	parent := "projects/" + projectID
// 	req := &talent.SearchJobsRequest{
// 		// Make sure to set the RequestMetadata the same as the associated
// 		// search request.
// 		RequestMetadata: &talent.RequestMetadata{
// 			// Make sure to hash your userID.
// 			UserId: "HashedUsrId",
// 			// Make sure to hash the sessionID.
// 			SessionId: "HashedSessionId",
// 			// Domain of the website where the search is conducted.
// 			Domain: "www.googlesample.com",
// 		},
// 		// Set the search mode to a regular search.
// 		SearchMode: "JOB_SEARCH",
// 	}
// 	if companyName != "" {
// 		req.JobQuery = &talent.JobQuery{
// 			CompanyNames: []string{companyName},
// 		}
// 	}

// 	resp, err := service.Projects.Jobs.SearchForAlert(parent, req).Do()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to search for jobs with alerts: %w", err)
// 	}

// 	fmt.Fprintln(w, "Jobs:")
// 	for _, j := range resp.MatchingJobs {
// 		fmt.Fprintf(w, "\t%q\n", j.Job.Name)
// 	}

// 	return resp, nil
// }
