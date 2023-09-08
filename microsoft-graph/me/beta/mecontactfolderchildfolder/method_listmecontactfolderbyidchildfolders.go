package mecontactfolderchildfolder

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

type ListMeContactFolderByIdChildFoldersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ContactFolderCollectionResponse
}

type ListMeContactFolderByIdChildFoldersCompleteResult struct {
	Items []models.ContactFolderCollectionResponse
}

// ListMeContactFolderByIdChildFolders ...
func (c MeContactFolderChildFolderClient) ListMeContactFolderByIdChildFolders(ctx context.Context, id MeContactFolderId) (result ListMeContactFolderByIdChildFoldersOperationResponse, err error) {
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
		Values *[]models.ContactFolderCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeContactFolderByIdChildFoldersComplete retrieves all the results into a single object
func (c MeContactFolderChildFolderClient) ListMeContactFolderByIdChildFoldersComplete(ctx context.Context, id MeContactFolderId) (ListMeContactFolderByIdChildFoldersCompleteResult, error) {
	return c.ListMeContactFolderByIdChildFoldersCompleteMatchingPredicate(ctx, id, models.ContactFolderCollectionResponseOperationPredicate{})
}

// ListMeContactFolderByIdChildFoldersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeContactFolderChildFolderClient) ListMeContactFolderByIdChildFoldersCompleteMatchingPredicate(ctx context.Context, id MeContactFolderId, predicate models.ContactFolderCollectionResponseOperationPredicate) (result ListMeContactFolderByIdChildFoldersCompleteResult, err error) {
	items := make([]models.ContactFolderCollectionResponse, 0)

	resp, err := c.ListMeContactFolderByIdChildFolders(ctx, id)
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

	result = ListMeContactFolderByIdChildFoldersCompleteResult{
		Items: items,
	}
	return
}
