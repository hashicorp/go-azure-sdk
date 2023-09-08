package metransitivememberof

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

type ListMeTransitiveMemberOfOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObjectCollectionResponse
}

type ListMeTransitiveMemberOfCompleteResult struct {
	Items []models.DirectoryObjectCollectionResponse
}

// ListMeTransitiveMemberOf ...
func (c MeTransitiveMemberOfClient) ListMeTransitiveMemberOf(ctx context.Context) (result ListMeTransitiveMemberOfOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/transitiveMemberOf",
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
		Values *[]models.DirectoryObjectCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeTransitiveMemberOfComplete retrieves all the results into a single object
func (c MeTransitiveMemberOfClient) ListMeTransitiveMemberOfComplete(ctx context.Context) (ListMeTransitiveMemberOfCompleteResult, error) {
	return c.ListMeTransitiveMemberOfCompleteMatchingPredicate(ctx, models.DirectoryObjectCollectionResponseOperationPredicate{})
}

// ListMeTransitiveMemberOfCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeTransitiveMemberOfClient) ListMeTransitiveMemberOfCompleteMatchingPredicate(ctx context.Context, predicate models.DirectoryObjectCollectionResponseOperationPredicate) (result ListMeTransitiveMemberOfCompleteResult, err error) {
	items := make([]models.DirectoryObjectCollectionResponse, 0)

	resp, err := c.ListMeTransitiveMemberOf(ctx)
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

	result = ListMeTransitiveMemberOfCompleteResult{
		Items: items,
	}
	return
}
