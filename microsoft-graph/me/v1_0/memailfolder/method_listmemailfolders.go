package memailfolder

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMeMailFoldersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.MailFolderCollectionResponse
}

type ListMeMailFoldersCompleteResult struct {
	Items []models.MailFolderCollectionResponse
}

// ListMeMailFolders ...
func (c MeMailFolderClient) ListMeMailFolders(ctx context.Context) (result ListMeMailFoldersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/mailFolders",
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
		Values *[]models.MailFolderCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeMailFoldersComplete retrieves all the results into a single object
func (c MeMailFolderClient) ListMeMailFoldersComplete(ctx context.Context) (ListMeMailFoldersCompleteResult, error) {
	return c.ListMeMailFoldersCompleteMatchingPredicate(ctx, models.MailFolderCollectionResponseOperationPredicate{})
}

// ListMeMailFoldersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeMailFolderClient) ListMeMailFoldersCompleteMatchingPredicate(ctx context.Context, predicate models.MailFolderCollectionResponseOperationPredicate) (result ListMeMailFoldersCompleteResult, err error) {
	items := make([]models.MailFolderCollectionResponse, 0)

	resp, err := c.ListMeMailFolders(ctx)
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

	result = ListMeMailFoldersCompleteResult{
		Items: items,
	}
	return
}
