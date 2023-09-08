package meactivity

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

type ListMeActivitiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UserActivityCollectionResponse
}

type ListMeActivitiesCompleteResult struct {
	Items []models.UserActivityCollectionResponse
}

// ListMeActivities ...
func (c MeActivityClient) ListMeActivities(ctx context.Context) (result ListMeActivitiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/activities",
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
		Values *[]models.UserActivityCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeActivitiesComplete retrieves all the results into a single object
func (c MeActivityClient) ListMeActivitiesComplete(ctx context.Context) (ListMeActivitiesCompleteResult, error) {
	return c.ListMeActivitiesCompleteMatchingPredicate(ctx, models.UserActivityCollectionResponseOperationPredicate{})
}

// ListMeActivitiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeActivityClient) ListMeActivitiesCompleteMatchingPredicate(ctx context.Context, predicate models.UserActivityCollectionResponseOperationPredicate) (result ListMeActivitiesCompleteResult, err error) {
	items := make([]models.UserActivityCollectionResponse, 0)

	resp, err := c.ListMeActivities(ctx)
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

	result = ListMeActivitiesCompleteResult{
		Items: items,
	}
	return
}
