package flexibleservercapabilities

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerCapabilitiesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]FlexibleServerCapability
}

type ServerCapabilitiesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []FlexibleServerCapability
}

// ServerCapabilitiesList ...
func (c FlexibleServerCapabilitiesClient) ServerCapabilitiesList(ctx context.Context, id FlexibleServerId) (result ServerCapabilitiesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/capabilities", id.ID()),
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
		Values *[]FlexibleServerCapability `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ServerCapabilitiesListComplete retrieves all the results into a single object
func (c FlexibleServerCapabilitiesClient) ServerCapabilitiesListComplete(ctx context.Context, id FlexibleServerId) (ServerCapabilitiesListCompleteResult, error) {
	return c.ServerCapabilitiesListCompleteMatchingPredicate(ctx, id, FlexibleServerCapabilityOperationPredicate{})
}

// ServerCapabilitiesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c FlexibleServerCapabilitiesClient) ServerCapabilitiesListCompleteMatchingPredicate(ctx context.Context, id FlexibleServerId, predicate FlexibleServerCapabilityOperationPredicate) (result ServerCapabilitiesListCompleteResult, err error) {
	items := make([]FlexibleServerCapability, 0)

	resp, err := c.ServerCapabilitiesList(ctx, id)
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

	result = ServerCapabilitiesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
