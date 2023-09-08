package useroutlooktask

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

type ListUserByIdOutlookTasksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OutlookTaskCollectionResponse
}

type ListUserByIdOutlookTasksCompleteResult struct {
	Items []models.OutlookTaskCollectionResponse
}

// ListUserByIdOutlookTasks ...
func (c UserOutlookTaskClient) ListUserByIdOutlookTasks(ctx context.Context, id UserId) (result ListUserByIdOutlookTasksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/outlook/tasks", id.ID()),
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

// ListUserByIdOutlookTasksComplete retrieves all the results into a single object
func (c UserOutlookTaskClient) ListUserByIdOutlookTasksComplete(ctx context.Context, id UserId) (ListUserByIdOutlookTasksCompleteResult, error) {
	return c.ListUserByIdOutlookTasksCompleteMatchingPredicate(ctx, id, models.OutlookTaskCollectionResponseOperationPredicate{})
}

// ListUserByIdOutlookTasksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserOutlookTaskClient) ListUserByIdOutlookTasksCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.OutlookTaskCollectionResponseOperationPredicate) (result ListUserByIdOutlookTasksCompleteResult, err error) {
	items := make([]models.OutlookTaskCollectionResponse, 0)

	resp, err := c.ListUserByIdOutlookTasks(ctx, id)
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

	result = ListUserByIdOutlookTasksCompleteResult{
		Items: items,
	}
	return
}
