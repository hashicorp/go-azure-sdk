package iscsipaths

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByPrivateCloudOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]IscsiPath
}

type ListByPrivateCloudCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []IscsiPath
}

// ListByPrivateCloud ...
func (c IscsiPathsClient) ListByPrivateCloud(ctx context.Context, id PrivateCloudId) (result ListByPrivateCloudOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/iscsiPaths", id.ID()),
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
		Values *[]IscsiPath `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByPrivateCloudComplete retrieves all the results into a single object
func (c IscsiPathsClient) ListByPrivateCloudComplete(ctx context.Context, id PrivateCloudId) (ListByPrivateCloudCompleteResult, error) {
	return c.ListByPrivateCloudCompleteMatchingPredicate(ctx, id, IscsiPathOperationPredicate{})
}

// ListByPrivateCloudCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IscsiPathsClient) ListByPrivateCloudCompleteMatchingPredicate(ctx context.Context, id PrivateCloudId, predicate IscsiPathOperationPredicate) (result ListByPrivateCloudCompleteResult, err error) {
	items := make([]IscsiPath, 0)

	resp, err := c.ListByPrivateCloud(ctx, id)
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

	result = ListByPrivateCloudCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
