package useroutlooktaskfolder

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

type ListUserByIdOutlookTaskFoldersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OutlookTaskFolderCollectionResponse
}

type ListUserByIdOutlookTaskFoldersCompleteResult struct {
	Items []models.OutlookTaskFolderCollectionResponse
}

// ListUserByIdOutlookTaskFolders ...
func (c UserOutlookTaskFolderClient) ListUserByIdOutlookTaskFolders(ctx context.Context, id UserId) (result ListUserByIdOutlookTaskFoldersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/outlook/taskFolders", id.ID()),
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

// ListUserByIdOutlookTaskFoldersComplete retrieves all the results into a single object
func (c UserOutlookTaskFolderClient) ListUserByIdOutlookTaskFoldersComplete(ctx context.Context, id UserId) (ListUserByIdOutlookTaskFoldersCompleteResult, error) {
	return c.ListUserByIdOutlookTaskFoldersCompleteMatchingPredicate(ctx, id, models.OutlookTaskFolderCollectionResponseOperationPredicate{})
}

// ListUserByIdOutlookTaskFoldersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserOutlookTaskFolderClient) ListUserByIdOutlookTaskFoldersCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.OutlookTaskFolderCollectionResponseOperationPredicate) (result ListUserByIdOutlookTaskFoldersCompleteResult, err error) {
	items := make([]models.OutlookTaskFolderCollectionResponse, 0)

	resp, err := c.ListUserByIdOutlookTaskFolders(ctx, id)
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

	result = ListUserByIdOutlookTaskFoldersCompleteResult{
		Items: items,
	}
	return
}
