package usercontactfolderchildfolder

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

type ListUserByIdContactFolderByIdChildFoldersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ContactFolderCollectionResponse
}

type ListUserByIdContactFolderByIdChildFoldersCompleteResult struct {
	Items []models.ContactFolderCollectionResponse
}

// ListUserByIdContactFolderByIdChildFolders ...
func (c UserContactFolderChildFolderClient) ListUserByIdContactFolderByIdChildFolders(ctx context.Context, id UserContactFolderId) (result ListUserByIdContactFolderByIdChildFoldersOperationResponse, err error) {
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

// ListUserByIdContactFolderByIdChildFoldersComplete retrieves all the results into a single object
func (c UserContactFolderChildFolderClient) ListUserByIdContactFolderByIdChildFoldersComplete(ctx context.Context, id UserContactFolderId) (ListUserByIdContactFolderByIdChildFoldersCompleteResult, error) {
	return c.ListUserByIdContactFolderByIdChildFoldersCompleteMatchingPredicate(ctx, id, models.ContactFolderCollectionResponseOperationPredicate{})
}

// ListUserByIdContactFolderByIdChildFoldersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserContactFolderChildFolderClient) ListUserByIdContactFolderByIdChildFoldersCompleteMatchingPredicate(ctx context.Context, id UserContactFolderId, predicate models.ContactFolderCollectionResponseOperationPredicate) (result ListUserByIdContactFolderByIdChildFoldersCompleteResult, err error) {
	items := make([]models.ContactFolderCollectionResponse, 0)

	resp, err := c.ListUserByIdContactFolderByIdChildFolders(ctx, id)
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

	result = ListUserByIdContactFolderByIdChildFoldersCompleteResult{
		Items: items,
	}
	return
}
