package containerapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerAppPrivateLinkResourcesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PrivateLinkResource
}

type ContainerAppPrivateLinkResourcesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PrivateLinkResource
}

type ContainerAppPrivateLinkResourcesListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ContainerAppPrivateLinkResourcesListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ContainerAppPrivateLinkResourcesList ...
func (c ContainerAppsClient) ContainerAppPrivateLinkResourcesList(ctx context.Context, id ContainerAppId) (result ContainerAppPrivateLinkResourcesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ContainerAppPrivateLinkResourcesListCustomPager{},
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
		Values *[]PrivateLinkResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ContainerAppPrivateLinkResourcesListComplete retrieves all the results into a single object
func (c ContainerAppsClient) ContainerAppPrivateLinkResourcesListComplete(ctx context.Context, id ContainerAppId) (ContainerAppPrivateLinkResourcesListCompleteResult, error) {
	return c.ContainerAppPrivateLinkResourcesListCompleteMatchingPredicate(ctx, id, PrivateLinkResourceOperationPredicate{})
}

// ContainerAppPrivateLinkResourcesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ContainerAppsClient) ContainerAppPrivateLinkResourcesListCompleteMatchingPredicate(ctx context.Context, id ContainerAppId, predicate PrivateLinkResourceOperationPredicate) (result ContainerAppPrivateLinkResourcesListCompleteResult, err error) {
	items := make([]PrivateLinkResource, 0)

	resp, err := c.ContainerAppPrivateLinkResourcesList(ctx, id)
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

	result = ContainerAppPrivateLinkResourcesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
