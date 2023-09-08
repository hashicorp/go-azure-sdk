package usermailfolderchildfolder

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

type ListUserByIdMailFolderByIdChildFoldersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.MailFolderCollectionResponse
}

type ListUserByIdMailFolderByIdChildFoldersCompleteResult struct {
	Items []models.MailFolderCollectionResponse
}

// ListUserByIdMailFolderByIdChildFolders ...
func (c UserMailFolderChildFolderClient) ListUserByIdMailFolderByIdChildFolders(ctx context.Context, id UserMailFolderId) (result ListUserByIdMailFolderByIdChildFoldersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/childFolders", id.ID()),
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

// ListUserByIdMailFolderByIdChildFoldersComplete retrieves all the results into a single object
func (c UserMailFolderChildFolderClient) ListUserByIdMailFolderByIdChildFoldersComplete(ctx context.Context, id UserMailFolderId) (ListUserByIdMailFolderByIdChildFoldersCompleteResult, error) {
	return c.ListUserByIdMailFolderByIdChildFoldersCompleteMatchingPredicate(ctx, id, models.MailFolderCollectionResponseOperationPredicate{})
}

// ListUserByIdMailFolderByIdChildFoldersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserMailFolderChildFolderClient) ListUserByIdMailFolderByIdChildFoldersCompleteMatchingPredicate(ctx context.Context, id UserMailFolderId, predicate models.MailFolderCollectionResponseOperationPredicate) (result ListUserByIdMailFolderByIdChildFoldersCompleteResult, err error) {
	items := make([]models.MailFolderCollectionResponse, 0)

	resp, err := c.ListUserByIdMailFolderByIdChildFolders(ctx, id)
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

	result = ListUserByIdMailFolderByIdChildFoldersCompleteResult{
		Items: items,
	}
	return
}
