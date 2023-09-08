package mechatpermissiongrant

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

type ListMeChatByIdPermissionGrantsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ResourceSpecificPermissionGrantCollectionResponse
}

type ListMeChatByIdPermissionGrantsCompleteResult struct {
	Items []models.ResourceSpecificPermissionGrantCollectionResponse
}

// ListMeChatByIdPermissionGrants ...
func (c MeChatPermissionGrantClient) ListMeChatByIdPermissionGrants(ctx context.Context, id MeChatId) (result ListMeChatByIdPermissionGrantsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/permissionGrants", id.ID()),
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
		Values *[]models.ResourceSpecificPermissionGrantCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeChatByIdPermissionGrantsComplete retrieves all the results into a single object
func (c MeChatPermissionGrantClient) ListMeChatByIdPermissionGrantsComplete(ctx context.Context, id MeChatId) (ListMeChatByIdPermissionGrantsCompleteResult, error) {
	return c.ListMeChatByIdPermissionGrantsCompleteMatchingPredicate(ctx, id, models.ResourceSpecificPermissionGrantCollectionResponseOperationPredicate{})
}

// ListMeChatByIdPermissionGrantsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeChatPermissionGrantClient) ListMeChatByIdPermissionGrantsCompleteMatchingPredicate(ctx context.Context, id MeChatId, predicate models.ResourceSpecificPermissionGrantCollectionResponseOperationPredicate) (result ListMeChatByIdPermissionGrantsCompleteResult, err error) {
	items := make([]models.ResourceSpecificPermissionGrantCollectionResponse, 0)

	resp, err := c.ListMeChatByIdPermissionGrants(ctx, id)
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

	result = ListMeChatByIdPermissionGrantsCompleteResult{
		Items: items,
	}
	return
}
