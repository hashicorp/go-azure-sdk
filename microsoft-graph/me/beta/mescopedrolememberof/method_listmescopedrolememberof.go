package mescopedrolememberof

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

type ListMeScopedRoleMemberOfOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ScopedRoleMembershipCollectionResponse
}

type ListMeScopedRoleMemberOfCompleteResult struct {
	Items []models.ScopedRoleMembershipCollectionResponse
}

// ListMeScopedRoleMemberOf ...
func (c MeScopedRoleMemberOfClient) ListMeScopedRoleMemberOf(ctx context.Context) (result ListMeScopedRoleMemberOfOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/scopedRoleMemberOf",
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
		Values *[]models.ScopedRoleMembershipCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeScopedRoleMemberOfComplete retrieves all the results into a single object
func (c MeScopedRoleMemberOfClient) ListMeScopedRoleMemberOfComplete(ctx context.Context) (ListMeScopedRoleMemberOfCompleteResult, error) {
	return c.ListMeScopedRoleMemberOfCompleteMatchingPredicate(ctx, models.ScopedRoleMembershipCollectionResponseOperationPredicate{})
}

// ListMeScopedRoleMemberOfCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeScopedRoleMemberOfClient) ListMeScopedRoleMemberOfCompleteMatchingPredicate(ctx context.Context, predicate models.ScopedRoleMembershipCollectionResponseOperationPredicate) (result ListMeScopedRoleMemberOfCompleteResult, err error) {
	items := make([]models.ScopedRoleMembershipCollectionResponse, 0)

	resp, err := c.ListMeScopedRoleMemberOf(ctx)
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

	result = ListMeScopedRoleMemberOfCompleteResult{
		Items: items,
	}
	return
}
