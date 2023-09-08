package usercontactfolderchildfoldercontactextension

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

type ListUserByIdContactFolderByIdChildFolderByIdContactByIdExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ExtensionCollectionResponse
}

type ListUserByIdContactFolderByIdChildFolderByIdContactByIdExtensionsCompleteResult struct {
	Items []models.ExtensionCollectionResponse
}

// ListUserByIdContactFolderByIdChildFolderByIdContactByIdExtensions ...
func (c UserContactFolderChildFolderContactExtensionClient) ListUserByIdContactFolderByIdChildFolderByIdContactByIdExtensions(ctx context.Context, id UserContactFolderChildFolderContactId) (result ListUserByIdContactFolderByIdChildFolderByIdContactByIdExtensionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/extensions", id.ID()),
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
		Values *[]models.ExtensionCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdContactFolderByIdChildFolderByIdContactByIdExtensionsComplete retrieves all the results into a single object
func (c UserContactFolderChildFolderContactExtensionClient) ListUserByIdContactFolderByIdChildFolderByIdContactByIdExtensionsComplete(ctx context.Context, id UserContactFolderChildFolderContactId) (ListUserByIdContactFolderByIdChildFolderByIdContactByIdExtensionsCompleteResult, error) {
	return c.ListUserByIdContactFolderByIdChildFolderByIdContactByIdExtensionsCompleteMatchingPredicate(ctx, id, models.ExtensionCollectionResponseOperationPredicate{})
}

// ListUserByIdContactFolderByIdChildFolderByIdContactByIdExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserContactFolderChildFolderContactExtensionClient) ListUserByIdContactFolderByIdChildFolderByIdContactByIdExtensionsCompleteMatchingPredicate(ctx context.Context, id UserContactFolderChildFolderContactId, predicate models.ExtensionCollectionResponseOperationPredicate) (result ListUserByIdContactFolderByIdChildFolderByIdContactByIdExtensionsCompleteResult, err error) {
	items := make([]models.ExtensionCollectionResponse, 0)

	resp, err := c.ListUserByIdContactFolderByIdChildFolderByIdContactByIdExtensions(ctx, id)
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

	result = ListUserByIdContactFolderByIdChildFolderByIdContactByIdExtensionsCompleteResult{
		Items: items,
	}
	return
}
