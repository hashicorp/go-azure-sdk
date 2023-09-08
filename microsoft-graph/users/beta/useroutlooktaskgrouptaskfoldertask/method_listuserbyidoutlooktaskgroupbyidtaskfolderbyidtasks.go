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

type ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTasksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OutlookTaskCollectionResponse
}

type ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTasksCompleteResult struct {
	Items []models.OutlookTaskCollectionResponse
}

// ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTasks ...
func (c UserOutlookTaskGroupTaskFolderTaskClient) ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTasks(ctx context.Context, id UserOutlookTaskGroupTaskFolderId) (result ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTasksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/tasks", id.ID()),
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

// ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTasksComplete retrieves all the results into a single object
func (c UserOutlookTaskGroupTaskFolderTaskClient) ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTasksComplete(ctx context.Context, id UserOutlookTaskGroupTaskFolderId) (ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTasksCompleteResult, error) {
	return c.ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTasksCompleteMatchingPredicate(ctx, id, models.OutlookTaskCollectionResponseOperationPredicate{})
}

// ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTasksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserOutlookTaskGroupTaskFolderTaskClient) ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTasksCompleteMatchingPredicate(ctx context.Context, id UserOutlookTaskGroupTaskFolderId, predicate models.OutlookTaskCollectionResponseOperationPredicate) (result ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTasksCompleteResult, err error) {
	items := make([]models.OutlookTaskCollectionResponse, 0)

	resp, err := c.ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTasks(ctx, id)
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

	result = ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTasksCompleteResult{
		Items: items,
	}
	return
}
