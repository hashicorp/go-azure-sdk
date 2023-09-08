package useroutlooktaskgrouptaskfolder

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

type ListUserByIdOutlookTaskGroupByIdTaskFoldersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OutlookTaskFolderCollectionResponse
}

type ListUserByIdOutlookTaskGroupByIdTaskFoldersCompleteResult struct {
	Items []models.OutlookTaskFolderCollectionResponse
}

// ListUserByIdOutlookTaskGroupByIdTaskFolders ...
func (c UserOutlookTaskGroupTaskFolderClient) ListUserByIdOutlookTaskGroupByIdTaskFolders(ctx context.Context, id UserOutlookTaskGroupId) (result ListUserByIdOutlookTaskGroupByIdTaskFoldersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/taskFolders", id.ID()),
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
		Values *[]models.OutlookTaskFolderCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdOutlookTaskGroupByIdTaskFoldersComplete retrieves all the results into a single object
func (c UserOutlookTaskGroupTaskFolderClient) ListUserByIdOutlookTaskGroupByIdTaskFoldersComplete(ctx context.Context, id UserOutlookTaskGroupId) (ListUserByIdOutlookTaskGroupByIdTaskFoldersCompleteResult, error) {
	return c.ListUserByIdOutlookTaskGroupByIdTaskFoldersCompleteMatchingPredicate(ctx, id, models.OutlookTaskFolderCollectionResponseOperationPredicate{})
}

// ListUserByIdOutlookTaskGroupByIdTaskFoldersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserOutlookTaskGroupTaskFolderClient) ListUserByIdOutlookTaskGroupByIdTaskFoldersCompleteMatchingPredicate(ctx context.Context, id UserOutlookTaskGroupId, predicate models.OutlookTaskFolderCollectionResponseOperationPredicate) (result ListUserByIdOutlookTaskGroupByIdTaskFoldersCompleteResult, err error) {
	items := make([]models.OutlookTaskFolderCollectionResponse, 0)

	resp, err := c.ListUserByIdOutlookTaskGroupByIdTaskFolders(ctx, id)
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

	result = ListUserByIdOutlookTaskGroupByIdTaskFoldersCompleteResult{
		Items: items,
	}
	return
}
