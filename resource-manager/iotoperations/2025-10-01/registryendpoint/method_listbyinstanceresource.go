package registryendpoint

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByInstanceResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]RegistryEndpointResource
}

type ListByInstanceResourceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []RegistryEndpointResource
}

type ListByInstanceResourceCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByInstanceResourceCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByInstanceResource ...
func (c RegistryEndpointClient) ListByInstanceResource(ctx context.Context, id InstanceId) (result ListByInstanceResourceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByInstanceResourceCustomPager{},
		Path:       fmt.Sprintf("%s/registryEndpoints", id.ID()),
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
		Values *[]RegistryEndpointResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByInstanceResourceComplete retrieves all the results into a single object
func (c RegistryEndpointClient) ListByInstanceResourceComplete(ctx context.Context, id InstanceId) (ListByInstanceResourceCompleteResult, error) {
	return c.ListByInstanceResourceCompleteMatchingPredicate(ctx, id, RegistryEndpointResourceOperationPredicate{})
}

// ListByInstanceResourceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RegistryEndpointClient) ListByInstanceResourceCompleteMatchingPredicate(ctx context.Context, id InstanceId, predicate RegistryEndpointResourceOperationPredicate) (result ListByInstanceResourceCompleteResult, err error) {
	items := make([]RegistryEndpointResource, 0)

	resp, err := c.ListByInstanceResource(ctx, id)
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

	result = ListByInstanceResourceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
