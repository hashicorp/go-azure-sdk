package groupthreadpostinreplytoextension

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

type ListGroupByIdThreadByIdPostByIdInReplyToExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ExtensionCollectionResponse
}

type ListGroupByIdThreadByIdPostByIdInReplyToExtensionsCompleteResult struct {
	Items []models.ExtensionCollectionResponse
}

// ListGroupByIdThreadByIdPostByIdInReplyToExtensions ...
func (c GroupThreadPostInReplyToExtensionClient) ListGroupByIdThreadByIdPostByIdInReplyToExtensions(ctx context.Context, id GroupThreadPostId) (result ListGroupByIdThreadByIdPostByIdInReplyToExtensionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/inReplyTo/extensions", id.ID()),
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

// ListGroupByIdThreadByIdPostByIdInReplyToExtensionsComplete retrieves all the results into a single object
func (c GroupThreadPostInReplyToExtensionClient) ListGroupByIdThreadByIdPostByIdInReplyToExtensionsComplete(ctx context.Context, id GroupThreadPostId) (ListGroupByIdThreadByIdPostByIdInReplyToExtensionsCompleteResult, error) {
	return c.ListGroupByIdThreadByIdPostByIdInReplyToExtensionsCompleteMatchingPredicate(ctx, id, models.ExtensionCollectionResponseOperationPredicate{})
}

// ListGroupByIdThreadByIdPostByIdInReplyToExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupThreadPostInReplyToExtensionClient) ListGroupByIdThreadByIdPostByIdInReplyToExtensionsCompleteMatchingPredicate(ctx context.Context, id GroupThreadPostId, predicate models.ExtensionCollectionResponseOperationPredicate) (result ListGroupByIdThreadByIdPostByIdInReplyToExtensionsCompleteResult, err error) {
	items := make([]models.ExtensionCollectionResponse, 0)

	resp, err := c.ListGroupByIdThreadByIdPostByIdInReplyToExtensions(ctx, id)
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

	result = ListGroupByIdThreadByIdPostByIdInReplyToExtensionsCompleteResult{
		Items: items,
	}
	return
}
