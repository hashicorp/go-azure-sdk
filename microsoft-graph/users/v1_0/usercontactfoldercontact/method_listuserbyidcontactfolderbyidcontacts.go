package usercontactfoldercontact

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

type ListUserByIdContactFolderByIdContactsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ContactCollectionResponse
}

type ListUserByIdContactFolderByIdContactsCompleteResult struct {
	Items []models.ContactCollectionResponse
}

// ListUserByIdContactFolderByIdContacts ...
func (c UserContactFolderContactClient) ListUserByIdContactFolderByIdContacts(ctx context.Context, id UserContactFolderId) (result ListUserByIdContactFolderByIdContactsOperationResponse, err error) {
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

// ListUserByIdContactFolderByIdContactsComplete retrieves all the results into a single object
func (c UserContactFolderContactClient) ListUserByIdContactFolderByIdContactsComplete(ctx context.Context, id UserContactFolderId) (ListUserByIdContactFolderByIdContactsCompleteResult, error) {
	return c.ListUserByIdContactFolderByIdContactsCompleteMatchingPredicate(ctx, id, models.ContactCollectionResponseOperationPredicate{})
}

// ListUserByIdContactFolderByIdContactsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserContactFolderContactClient) ListUserByIdContactFolderByIdContactsCompleteMatchingPredicate(ctx context.Context, id UserContactFolderId, predicate models.ContactCollectionResponseOperationPredicate) (result ListUserByIdContactFolderByIdContactsCompleteResult, err error) {
	items := make([]models.ContactCollectionResponse, 0)

	resp, err := c.ListUserByIdContactFolderByIdContacts(ctx, id)
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

	result = ListUserByIdContactFolderByIdContactsCompleteResult{
		Items: items,
	}
	return
}
