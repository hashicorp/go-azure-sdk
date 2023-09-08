package meemployeeexperiencelearningcourseactivity

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMeEmployeeExperienceLearningCourseActivitiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.LearningCourseActivityCollectionResponse
}

type ListMeEmployeeExperienceLearningCourseActivitiesCompleteResult struct {
	Items []models.LearningCourseActivityCollectionResponse
}

// ListMeEmployeeExperienceLearningCourseActivities ...
func (c MeEmployeeExperienceLearningCourseActivityClient) ListMeEmployeeExperienceLearningCourseActivities(ctx context.Context) (result ListMeEmployeeExperienceLearningCourseActivitiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/employeeExperience/learningCourseActivities",
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]models.LearningCourseActivityCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeEmployeeExperienceLearningCourseActivitiesComplete retrieves all the results into a single object
func (c MeEmployeeExperienceLearningCourseActivityClient) ListMeEmployeeExperienceLearningCourseActivitiesComplete(ctx context.Context) (ListMeEmployeeExperienceLearningCourseActivitiesCompleteResult, error) {
	return c.ListMeEmployeeExperienceLearningCourseActivitiesCompleteMatchingPredicate(ctx, models.LearningCourseActivityCollectionResponseOperationPredicate{})
}

// ListMeEmployeeExperienceLearningCourseActivitiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeEmployeeExperienceLearningCourseActivityClient) ListMeEmployeeExperienceLearningCourseActivitiesCompleteMatchingPredicate(ctx context.Context, predicate models.LearningCourseActivityCollectionResponseOperationPredicate) (result ListMeEmployeeExperienceLearningCourseActivitiesCompleteResult, err error) {
	items := make([]models.LearningCourseActivityCollectionResponse, 0)

	resp, err := c.ListMeEmployeeExperienceLearningCourseActivities(ctx)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListMeEmployeeExperienceLearningCourseActivitiesCompleteResult{
		Items: items,
	}
	return
}
