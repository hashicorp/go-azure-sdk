package userprofileeducationalactivity

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

type ListUserByIdProfileEducationalActivitiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.EducationalActivityCollectionResponse
}

type ListUserByIdProfileEducationalActivitiesCompleteResult struct {
	Items []models.EducationalActivityCollectionResponse
}

// ListUserByIdProfileEducationalActivities ...
func (c UserProfileEducationalActivityClient) ListUserByIdProfileEducationalActivities(ctx context.Context, id UserId) (result ListUserByIdProfileEducationalActivitiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/profile/educationalActivities", id.ID()),
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
		Values *[]models.EducationalActivityCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdProfileEducationalActivitiesComplete retrieves all the results into a single object
func (c UserProfileEducationalActivityClient) ListUserByIdProfileEducationalActivitiesComplete(ctx context.Context, id UserId) (ListUserByIdProfileEducationalActivitiesCompleteResult, error) {
	return c.ListUserByIdProfileEducationalActivitiesCompleteMatchingPredicate(ctx, id, models.EducationalActivityCollectionResponseOperationPredicate{})
}

// ListUserByIdProfileEducationalActivitiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserProfileEducationalActivityClient) ListUserByIdProfileEducationalActivitiesCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.EducationalActivityCollectionResponseOperationPredicate) (result ListUserByIdProfileEducationalActivitiesCompleteResult, err error) {
	items := make([]models.EducationalActivityCollectionResponse, 0)

	resp, err := c.ListUserByIdProfileEducationalActivities(ctx, id)
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

	result = ListUserByIdProfileEducationalActivitiesCompleteResult{
		Items: items,
	}
	return
}
