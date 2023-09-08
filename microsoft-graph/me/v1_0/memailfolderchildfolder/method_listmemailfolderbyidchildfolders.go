package memailfolderchildfolder

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

type ListMeMailFolderByIdChildFoldersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.MailFolderCollectionResponse
}

type ListMeMailFolderByIdChildFoldersCompleteResult struct {
	Items []models.MailFolderCollectionResponse
}

// ListMeMailFolderByIdChildFolders ...
func (c MeMailFolderChildFolderClient) ListMeMailFolderByIdChildFolders(ctx context.Context, id MeMailFolderId) (result ListMeMailFolderByIdChildFoldersOperationResponse, err error) {
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

// ListMeMailFolderByIdChildFoldersComplete retrieves all the results into a single object
func (c MeMailFolderChildFolderClient) ListMeMailFolderByIdChildFoldersComplete(ctx context.Context, id MeMailFolderId) (ListMeMailFolderByIdChildFoldersCompleteResult, error) {
	return c.ListMeMailFolderByIdChildFoldersCompleteMatchingPredicate(ctx, id, models.MailFolderCollectionResponseOperationPredicate{})
}

// ListMeMailFolderByIdChildFoldersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeMailFolderChildFolderClient) ListMeMailFolderByIdChildFoldersCompleteMatchingPredicate(ctx context.Context, id MeMailFolderId, predicate models.MailFolderCollectionResponseOperationPredicate) (result ListMeMailFolderByIdChildFoldersCompleteResult, err error) {
	items := make([]models.MailFolderCollectionResponse, 0)

	resp, err := c.ListMeMailFolderByIdChildFolders(ctx, id)
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

	result = ListMeMailFolderByIdChildFoldersCompleteResult{
		Items: items,
	}
	return
}
