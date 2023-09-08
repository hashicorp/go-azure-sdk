package useroutlooktaskgrouptaskfoldertask

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

type ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTaskByIdCompletesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OutlookTask
}

type ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTaskByIdCompletesCompleteResult struct {
	Items []models.OutlookTask
}

// ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTaskByIdCompletes ...
func (c UserOutlookTaskGroupTaskFolderTaskClient) ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTaskByIdCompletes(ctx context.Context, id UserOutlookTaskGroupTaskFolderTaskId) (result ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTaskByIdCompletesOperationResponse, err error) {
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

// ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTaskByIdCompletesComplete retrieves all the results into a single object
func (c UserOutlookTaskGroupTaskFolderTaskClient) ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTaskByIdCompletesComplete(ctx context.Context, id UserOutlookTaskGroupTaskFolderTaskId) (ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTaskByIdCompletesCompleteResult, error) {
	return c.ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTaskByIdCompletesCompleteMatchingPredicate(ctx, id, models.OutlookTaskOperationPredicate{})
}

// ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTaskByIdCompletesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserOutlookTaskGroupTaskFolderTaskClient) ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTaskByIdCompletesCompleteMatchingPredicate(ctx context.Context, id UserOutlookTaskGroupTaskFolderTaskId, predicate models.OutlookTaskOperationPredicate) (result ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTaskByIdCompletesCompleteResult, err error) {
	items := make([]models.OutlookTask, 0)

	resp, err := c.ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTaskByIdCompletes(ctx, id)
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

	result = ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTaskByIdCompletesCompleteResult{
		Items: items,
	}
	return
}
