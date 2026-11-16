package openapis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapabilitiesByLocationListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Capability
}

type CapabilitiesByLocationListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Capability
}

type CapabilitiesByLocationListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *CapabilitiesByLocationListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// CapabilitiesByLocationList ...
func (c OpenapisClient) CapabilitiesByLocationList(ctx context.Context, id LocationId) (result CapabilitiesByLocationListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &CapabilitiesByLocationListCustomPager{},
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
		Values *[]Capability `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// CapabilitiesByLocationListComplete retrieves all the results into a single object
func (c OpenapisClient) CapabilitiesByLocationListComplete(ctx context.Context, id LocationId) (CapabilitiesByLocationListCompleteResult, error) {
	return c.CapabilitiesByLocationListCompleteMatchingPredicate(ctx, id, CapabilityOperationPredicate{})
}

// CapabilitiesByLocationListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) CapabilitiesByLocationListCompleteMatchingPredicate(ctx context.Context, id LocationId, predicate CapabilityOperationPredicate) (result CapabilitiesByLocationListCompleteResult, err error) {
	items := make([]Capability, 0)

	resp, err := c.CapabilitiesByLocationList(ctx, id)
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

	result = CapabilitiesByLocationListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
