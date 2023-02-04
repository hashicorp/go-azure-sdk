package privatelinkresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByServerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PrivateLinkResource
}

type ListByServerCompleteResult struct {
	Items []PrivateLinkResource
}

// ListByServer ...
func (c PrivateLinkResourcesClient) ListByServer(ctx context.Context, id ServerId) (result ListByServerOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/privateLinkResources", id.ID()),
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
		Values *[]PrivateLinkResource `json:"values"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByServerComplete retrieves all the results into a single object
func (c PrivateLinkResourcesClient) ListByServerComplete(ctx context.Context, id ServerId) (ListByServerCompleteResult, error) {
	return c.ListByServerCompleteMatchingPredicate(ctx, id, PrivateLinkResourceOperationPredicate{})
}

// ListByServerCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivateLinkResourcesClient) ListByServerCompleteMatchingPredicate(ctx context.Context, id ServerId, predicate PrivateLinkResourceOperationPredicate) (result ListByServerCompleteResult, err error) {
	items := make([]PrivateLinkResource, 0)

	resp, err := c.ListByServer(ctx, id)
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

	result = ListByServerCompleteResult{
		Items: items,
	}
	return
}
