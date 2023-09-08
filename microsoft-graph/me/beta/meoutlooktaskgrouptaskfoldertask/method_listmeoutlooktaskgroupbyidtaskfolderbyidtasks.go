package meoutlooktaskgrouptaskfoldertask

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

type ListMeOutlookTaskGroupByIdTaskFolderByIdTasksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OutlookTaskCollectionResponse
}

type ListMeOutlookTaskGroupByIdTaskFolderByIdTasksCompleteResult struct {
	Items []models.OutlookTaskCollectionResponse
}

// ListMeOutlookTaskGroupByIdTaskFolderByIdTasks ...
func (c MeOutlookTaskGroupTaskFolderTaskClient) ListMeOutlookTaskGroupByIdTaskFolderByIdTasks(ctx context.Context, id MeOutlookTaskGroupTaskFolderId) (result ListMeOutlookTaskGroupByIdTaskFolderByIdTasksOperationResponse, err error) {
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

// ListMeOutlookTaskGroupByIdTaskFolderByIdTasksComplete retrieves all the results into a single object
func (c MeOutlookTaskGroupTaskFolderTaskClient) ListMeOutlookTaskGroupByIdTaskFolderByIdTasksComplete(ctx context.Context, id MeOutlookTaskGroupTaskFolderId) (ListMeOutlookTaskGroupByIdTaskFolderByIdTasksCompleteResult, error) {
	return c.ListMeOutlookTaskGroupByIdTaskFolderByIdTasksCompleteMatchingPredicate(ctx, id, models.OutlookTaskCollectionResponseOperationPredicate{})
}

// ListMeOutlookTaskGroupByIdTaskFolderByIdTasksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeOutlookTaskGroupTaskFolderTaskClient) ListMeOutlookTaskGroupByIdTaskFolderByIdTasksCompleteMatchingPredicate(ctx context.Context, id MeOutlookTaskGroupTaskFolderId, predicate models.OutlookTaskCollectionResponseOperationPredicate) (result ListMeOutlookTaskGroupByIdTaskFolderByIdTasksCompleteResult, err error) {
	items := make([]models.OutlookTaskCollectionResponse, 0)

	resp, err := c.ListMeOutlookTaskGroupByIdTaskFolderByIdTasks(ctx, id)
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

	result = ListMeOutlookTaskGroupByIdTaskFolderByIdTasksCompleteResult{
		Items: items,
	}
	return
}
