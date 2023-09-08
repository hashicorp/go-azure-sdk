package meoutlooktaskfoldertask

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

type ListMeOutlookTaskFolderByIdTasksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OutlookTaskCollectionResponse
}

type ListMeOutlookTaskFolderByIdTasksCompleteResult struct {
	Items []models.OutlookTaskCollectionResponse
}

// ListMeOutlookTaskFolderByIdTasks ...
func (c MeOutlookTaskFolderTaskClient) ListMeOutlookTaskFolderByIdTasks(ctx context.Context, id MeOutlookTaskFolderId) (result ListMeOutlookTaskFolderByIdTasksOperationResponse, err error) {
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

// ListMeOutlookTaskFolderByIdTasksComplete retrieves all the results into a single object
func (c MeOutlookTaskFolderTaskClient) ListMeOutlookTaskFolderByIdTasksComplete(ctx context.Context, id MeOutlookTaskFolderId) (ListMeOutlookTaskFolderByIdTasksCompleteResult, error) {
	return c.ListMeOutlookTaskFolderByIdTasksCompleteMatchingPredicate(ctx, id, models.OutlookTaskCollectionResponseOperationPredicate{})
}

// ListMeOutlookTaskFolderByIdTasksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeOutlookTaskFolderTaskClient) ListMeOutlookTaskFolderByIdTasksCompleteMatchingPredicate(ctx context.Context, id MeOutlookTaskFolderId, predicate models.OutlookTaskCollectionResponseOperationPredicate) (result ListMeOutlookTaskFolderByIdTasksCompleteResult, err error) {
	items := make([]models.OutlookTaskCollectionResponse, 0)

	resp, err := c.ListMeOutlookTaskFolderByIdTasks(ctx, id)
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

	result = ListMeOutlookTaskFolderByIdTasksCompleteResult{
		Items: items,
	}
	return
}
