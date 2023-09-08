package useremployeeexperiencelearningcourseactivity

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserByIdEmployeeExperienceLearningCourseActivitiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.LearningCourseActivityCollectionResponse
}

type ListUserByIdEmployeeExperienceLearningCourseActivitiesCompleteResult struct {
	Items []models.LearningCourseActivityCollectionResponse
}

// ListUserByIdEmployeeExperienceLearningCourseActivities ...
func (c UserEmployeeExperienceLearningCourseActivityClient) ListUserByIdEmployeeExperienceLearningCourseActivities(ctx context.Context, id UserId) (result ListUserByIdEmployeeExperienceLearningCourseActivitiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/employeeExperience/learningCourseActivities", id.ID()),
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

// ListUserByIdEmployeeExperienceLearningCourseActivitiesComplete retrieves all the results into a single object
func (c UserEmployeeExperienceLearningCourseActivityClient) ListUserByIdEmployeeExperienceLearningCourseActivitiesComplete(ctx context.Context, id UserId) (ListUserByIdEmployeeExperienceLearningCourseActivitiesCompleteResult, error) {
	return c.ListUserByIdEmployeeExperienceLearningCourseActivitiesCompleteMatchingPredicate(ctx, id, models.LearningCourseActivityCollectionResponseOperationPredicate{})
}

// ListUserByIdEmployeeExperienceLearningCourseActivitiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserEmployeeExperienceLearningCourseActivityClient) ListUserByIdEmployeeExperienceLearningCourseActivitiesCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.LearningCourseActivityCollectionResponseOperationPredicate) (result ListUserByIdEmployeeExperienceLearningCourseActivitiesCompleteResult, err error) {
	items := make([]models.LearningCourseActivityCollectionResponse, 0)

	resp, err := c.ListUserByIdEmployeeExperienceLearningCourseActivities(ctx, id)
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

	result = ListUserByIdEmployeeExperienceLearningCourseActivitiesCompleteResult{
		Items: items,
	}
	return
}
