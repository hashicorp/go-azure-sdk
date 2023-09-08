package usercontactfolderchildfoldercontact

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

type ListUserByIdContactFolderByIdChildFolderByIdContactsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ContactCollectionResponse
}

type ListUserByIdContactFolderByIdChildFolderByIdContactsCompleteResult struct {
	Items []models.ContactCollectionResponse
}

// ListUserByIdContactFolderByIdChildFolderByIdContacts ...
func (c UserContactFolderChildFolderContactClient) ListUserByIdContactFolderByIdChildFolderByIdContacts(ctx context.Context, id UserContactFolderChildFolderId) (result ListUserByIdContactFolderByIdChildFolderByIdContactsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/contacts", id.ID()),
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
		Values *[]models.ContactCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdContactFolderByIdChildFolderByIdContactsComplete retrieves all the results into a single object
func (c UserContactFolderChildFolderContactClient) ListUserByIdContactFolderByIdChildFolderByIdContactsComplete(ctx context.Context, id UserContactFolderChildFolderId) (ListUserByIdContactFolderByIdChildFolderByIdContactsCompleteResult, error) {
	return c.ListUserByIdContactFolderByIdChildFolderByIdContactsCompleteMatchingPredicate(ctx, id, models.ContactCollectionResponseOperationPredicate{})
}

// ListUserByIdContactFolderByIdChildFolderByIdContactsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserContactFolderChildFolderContactClient) ListUserByIdContactFolderByIdChildFolderByIdContactsCompleteMatchingPredicate(ctx context.Context, id UserContactFolderChildFolderId, predicate models.ContactCollectionResponseOperationPredicate) (result ListUserByIdContactFolderByIdChildFolderByIdContactsCompleteResult, err error) {
	items := make([]models.ContactCollectionResponse, 0)

	resp, err := c.ListUserByIdContactFolderByIdChildFolderByIdContacts(ctx, id)
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

	result = ListUserByIdContactFolderByIdChildFolderByIdContactsCompleteResult{
		Items: items,
	}
	return
}
