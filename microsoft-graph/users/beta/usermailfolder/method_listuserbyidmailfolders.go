package usermailfolder

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

type ListUserByIdMailFoldersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.MailFolderCollectionResponse
}

type ListUserByIdMailFoldersCompleteResult struct {
	Items []models.MailFolderCollectionResponse
}

// ListUserByIdMailFolders ...
func (c UserMailFolderClient) ListUserByIdMailFolders(ctx context.Context, id UserId) (result ListUserByIdMailFoldersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/mailFolders", id.ID()),
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

// ListUserByIdMailFoldersComplete retrieves all the results into a single object
func (c UserMailFolderClient) ListUserByIdMailFoldersComplete(ctx context.Context, id UserId) (ListUserByIdMailFoldersCompleteResult, error) {
	return c.ListUserByIdMailFoldersCompleteMatchingPredicate(ctx, id, models.MailFolderCollectionResponseOperationPredicate{})
}

// ListUserByIdMailFoldersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserMailFolderClient) ListUserByIdMailFoldersCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.MailFolderCollectionResponseOperationPredicate) (result ListUserByIdMailFoldersCompleteResult, err error) {
	items := make([]models.MailFolderCollectionResponse, 0)

	resp, err := c.ListUserByIdMailFolders(ctx, id)
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

	result = ListUserByIdMailFoldersCompleteResult{
		Items: items,
	}
	return
}
