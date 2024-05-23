package applicationtypeversion

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByApplicationTypesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ApplicationTypeVersionResource
}

type ListByApplicationTypesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ApplicationTypeVersionResource
}

// ListByApplicationTypes ...
func (c ApplicationTypeVersionClient) ListByApplicationTypes(ctx context.Context, id ApplicationTypeId) (result ListByApplicationTypesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/versions", id.ID()),
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
		Values *[]ApplicationTypeVersionResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByApplicationTypesComplete retrieves all the results into a single object
func (c ApplicationTypeVersionClient) ListByApplicationTypesComplete(ctx context.Context, id ApplicationTypeId) (ListByApplicationTypesCompleteResult, error) {
	return c.ListByApplicationTypesCompleteMatchingPredicate(ctx, id, ApplicationTypeVersionResourceOperationPredicate{})
}

// ListByApplicationTypesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ApplicationTypeVersionClient) ListByApplicationTypesCompleteMatchingPredicate(ctx context.Context, id ApplicationTypeId, predicate ApplicationTypeVersionResourceOperationPredicate) (result ListByApplicationTypesCompleteResult, err error) {
	items := make([]ApplicationTypeVersionResource, 0)

	resp, err := c.ListByApplicationTypes(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = ListByApplicationTypesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
