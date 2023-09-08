package meoutlooktask

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

type ListMeOutlookTasksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OutlookTaskCollectionResponse
}

type ListMeOutlookTasksCompleteResult struct {
	Items []models.OutlookTaskCollectionResponse
}

// ListMeOutlookTasks ...
func (c MeOutlookTaskClient) ListMeOutlookTasks(ctx context.Context) (result ListMeOutlookTasksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/outlook/tasks",
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
		Values *[]models.OutlookTaskCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeOutlookTasksComplete retrieves all the results into a single object
func (c MeOutlookTaskClient) ListMeOutlookTasksComplete(ctx context.Context) (ListMeOutlookTasksCompleteResult, error) {
	return c.ListMeOutlookTasksCompleteMatchingPredicate(ctx, models.OutlookTaskCollectionResponseOperationPredicate{})
}

// ListMeOutlookTasksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeOutlookTaskClient) ListMeOutlookTasksCompleteMatchingPredicate(ctx context.Context, predicate models.OutlookTaskCollectionResponseOperationPredicate) (result ListMeOutlookTasksCompleteResult, err error) {
	items := make([]models.OutlookTaskCollectionResponse, 0)

	resp, err := c.ListMeOutlookTasks(ctx)
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

	result = ListMeOutlookTasksCompleteResult{
		Items: items,
	}
	return
}
