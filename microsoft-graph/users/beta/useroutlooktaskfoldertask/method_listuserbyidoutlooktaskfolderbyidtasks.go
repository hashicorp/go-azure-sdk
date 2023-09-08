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

type ListUserByIdOutlookTaskFolderByIdTasksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OutlookTaskCollectionResponse
}

type ListUserByIdOutlookTaskFolderByIdTasksCompleteResult struct {
	Items []models.OutlookTaskCollectionResponse
}

// ListUserByIdOutlookTaskFolderByIdTasks ...
func (c UserOutlookTaskFolderTaskClient) ListUserByIdOutlookTaskFolderByIdTasks(ctx context.Context, id UserOutlookTaskFolderId) (result ListUserByIdOutlookTaskFolderByIdTasksOperationResponse, err error) {
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

// ListUserByIdOutlookTaskFolderByIdTasksComplete retrieves all the results into a single object
func (c UserOutlookTaskFolderTaskClient) ListUserByIdOutlookTaskFolderByIdTasksComplete(ctx context.Context, id UserOutlookTaskFolderId) (ListUserByIdOutlookTaskFolderByIdTasksCompleteResult, error) {
	return c.ListUserByIdOutlookTaskFolderByIdTasksCompleteMatchingPredicate(ctx, id, models.OutlookTaskCollectionResponseOperationPredicate{})
}

// ListUserByIdOutlookTaskFolderByIdTasksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserOutlookTaskFolderTaskClient) ListUserByIdOutlookTaskFolderByIdTasksCompleteMatchingPredicate(ctx context.Context, id UserOutlookTaskFolderId, predicate models.OutlookTaskCollectionResponseOperationPredicate) (result ListUserByIdOutlookTaskFolderByIdTasksCompleteResult, err error) {
	items := make([]models.OutlookTaskCollectionResponse, 0)

	resp, err := c.ListUserByIdOutlookTaskFolderByIdTasks(ctx, id)
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

	result = ListUserByIdOutlookTaskFolderByIdTasksCompleteResult{
		Items: items,
	}
	return
}
