package containerappsbuilds

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ByContainerAppListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ContainerAppsBuildResource
}

type ByContainerAppListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ContainerAppsBuildResource
}

type ByContainerAppListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ByContainerAppListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ByContainerAppList ...
func (c ContainerAppsBuildsClient) ByContainerAppList(ctx context.Context, id ContainerAppId) (result ByContainerAppListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ByContainerAppListCustomPager{},
		Path:       fmt.Sprintf("%s/builds", id.ID()),
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
		Values *[]ContainerAppsBuildResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ByContainerAppListComplete retrieves all the results into a single object
func (c ContainerAppsBuildsClient) ByContainerAppListComplete(ctx context.Context, id ContainerAppId) (ByContainerAppListCompleteResult, error) {
	return c.ByContainerAppListCompleteMatchingPredicate(ctx, id, ContainerAppsBuildResourceOperationPredicate{})
}

// ByContainerAppListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ContainerAppsBuildsClient) ByContainerAppListCompleteMatchingPredicate(ctx context.Context, id ContainerAppId, predicate ContainerAppsBuildResourceOperationPredicate) (result ByContainerAppListCompleteResult, err error) {
	items := make([]ContainerAppsBuildResource, 0)

	resp, err := c.ByContainerAppList(ctx, id)
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

	result = ByContainerAppListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
