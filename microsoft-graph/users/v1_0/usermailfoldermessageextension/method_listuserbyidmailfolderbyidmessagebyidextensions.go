package usermailfoldermessageextension

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

type ListUserByIdMailFolderByIdMessageByIdExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ExtensionCollectionResponse
}

type ListUserByIdMailFolderByIdMessageByIdExtensionsCompleteResult struct {
	Items []models.ExtensionCollectionResponse
}

// ListUserByIdMailFolderByIdMessageByIdExtensions ...
func (c UserMailFolderMessageExtensionClient) ListUserByIdMailFolderByIdMessageByIdExtensions(ctx context.Context, id UserMailFolderMessageId) (result ListUserByIdMailFolderByIdMessageByIdExtensionsOperationResponse, err error) {
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

// ListUserByIdMailFolderByIdMessageByIdExtensionsComplete retrieves all the results into a single object
func (c UserMailFolderMessageExtensionClient) ListUserByIdMailFolderByIdMessageByIdExtensionsComplete(ctx context.Context, id UserMailFolderMessageId) (ListUserByIdMailFolderByIdMessageByIdExtensionsCompleteResult, error) {
	return c.ListUserByIdMailFolderByIdMessageByIdExtensionsCompleteMatchingPredicate(ctx, id, models.ExtensionCollectionResponseOperationPredicate{})
}

// ListUserByIdMailFolderByIdMessageByIdExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserMailFolderMessageExtensionClient) ListUserByIdMailFolderByIdMessageByIdExtensionsCompleteMatchingPredicate(ctx context.Context, id UserMailFolderMessageId, predicate models.ExtensionCollectionResponseOperationPredicate) (result ListUserByIdMailFolderByIdMessageByIdExtensionsCompleteResult, err error) {
	items := make([]models.ExtensionCollectionResponse, 0)

	resp, err := c.ListUserByIdMailFolderByIdMessageByIdExtensions(ctx, id)
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

	result = ListUserByIdMailFolderByIdMessageByIdExtensionsCompleteResult{
		Items: items,
	}
	return
}
