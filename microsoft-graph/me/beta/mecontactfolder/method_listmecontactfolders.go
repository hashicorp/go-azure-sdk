package mecontactfolder

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

type ListMeContactFoldersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ContactFolderCollectionResponse
}

type ListMeContactFoldersCompleteResult struct {
	Items []models.ContactFolderCollectionResponse
}

// ListMeContactFolders ...
func (c MeContactFolderClient) ListMeContactFolders(ctx context.Context) (result ListMeContactFoldersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/contactFolders",
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
		Values *[]models.ContactFolderCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeContactFoldersComplete retrieves all the results into a single object
func (c MeContactFolderClient) ListMeContactFoldersComplete(ctx context.Context) (ListMeContactFoldersCompleteResult, error) {
	return c.ListMeContactFoldersCompleteMatchingPredicate(ctx, models.ContactFolderCollectionResponseOperationPredicate{})
}

// ListMeContactFoldersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeContactFolderClient) ListMeContactFoldersCompleteMatchingPredicate(ctx context.Context, predicate models.ContactFolderCollectionResponseOperationPredicate) (result ListMeContactFoldersCompleteResult, err error) {
	items := make([]models.ContactFolderCollectionResponse, 0)

	resp, err := c.ListMeContactFolders(ctx)
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

	result = ListMeContactFoldersCompleteResult{
		Items: items,
	}
	return
}
