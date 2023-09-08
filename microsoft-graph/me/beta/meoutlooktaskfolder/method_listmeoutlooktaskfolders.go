package meoutlooktaskfolder

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

type ListMeOutlookTaskFoldersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OutlookTaskFolderCollectionResponse
}

type ListMeOutlookTaskFoldersCompleteResult struct {
	Items []models.OutlookTaskFolderCollectionResponse
}

// ListMeOutlookTaskFolders ...
func (c MeOutlookTaskFolderClient) ListMeOutlookTaskFolders(ctx context.Context) (result ListMeOutlookTaskFoldersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/outlook/taskFolders",
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

// ListMeOutlookTaskFoldersComplete retrieves all the results into a single object
func (c MeOutlookTaskFolderClient) ListMeOutlookTaskFoldersComplete(ctx context.Context) (ListMeOutlookTaskFoldersCompleteResult, error) {
	return c.ListMeOutlookTaskFoldersCompleteMatchingPredicate(ctx, models.OutlookTaskFolderCollectionResponseOperationPredicate{})
}

// ListMeOutlookTaskFoldersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeOutlookTaskFolderClient) ListMeOutlookTaskFoldersCompleteMatchingPredicate(ctx context.Context, predicate models.OutlookTaskFolderCollectionResponseOperationPredicate) (result ListMeOutlookTaskFoldersCompleteResult, err error) {
	items := make([]models.OutlookTaskFolderCollectionResponse, 0)

	resp, err := c.ListMeOutlookTaskFolders(ctx)
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

	result = ListMeOutlookTaskFoldersCompleteResult{
		Items: items,
	}
	return
}
