package useroutlooktaskfoldertask

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

type ListUserByIdOutlookTaskFolderByIdTaskByIdCompletesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OutlookTask
}

type ListUserByIdOutlookTaskFolderByIdTaskByIdCompletesCompleteResult struct {
	Items []models.OutlookTask
}

// ListUserByIdOutlookTaskFolderByIdTaskByIdCompletes ...
func (c UserOutlookTaskFolderTaskClient) ListUserByIdOutlookTaskFolderByIdTaskByIdCompletes(ctx context.Context, id UserOutlookTaskFolderTaskId) (result ListUserByIdOutlookTaskFolderByIdTaskByIdCompletesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/complete", id.ID()),
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
		Values *[]models.OutlookTask `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdOutlookTaskFolderByIdTaskByIdCompletesComplete retrieves all the results into a single object
func (c UserOutlookTaskFolderTaskClient) ListUserByIdOutlookTaskFolderByIdTaskByIdCompletesComplete(ctx context.Context, id UserOutlookTaskFolderTaskId) (ListUserByIdOutlookTaskFolderByIdTaskByIdCompletesCompleteResult, error) {
	return c.ListUserByIdOutlookTaskFolderByIdTaskByIdCompletesCompleteMatchingPredicate(ctx, id, models.OutlookTaskOperationPredicate{})
}

// ListUserByIdOutlookTaskFolderByIdTaskByIdCompletesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserOutlookTaskFolderTaskClient) ListUserByIdOutlookTaskFolderByIdTaskByIdCompletesCompleteMatchingPredicate(ctx context.Context, id UserOutlookTaskFolderTaskId, predicate models.OutlookTaskOperationPredicate) (result ListUserByIdOutlookTaskFolderByIdTaskByIdCompletesCompleteResult, err error) {
	items := make([]models.OutlookTask, 0)

	resp, err := c.ListUserByIdOutlookTaskFolderByIdTaskByIdCompletes(ctx, id)
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

	result = ListUserByIdOutlookTaskFolderByIdTaskByIdCompletesCompleteResult{
		Items: items,
	}
	return
}
