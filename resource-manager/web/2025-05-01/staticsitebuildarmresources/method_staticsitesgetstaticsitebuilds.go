package staticsitebuildarmresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StaticSitesGetStaticSiteBuildsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]StaticSiteBuildARMResource
}

type StaticSitesGetStaticSiteBuildsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []StaticSiteBuildARMResource
}

type StaticSitesGetStaticSiteBuildsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *StaticSitesGetStaticSiteBuildsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// StaticSitesGetStaticSiteBuilds ...
func (c StaticSiteBuildARMResourcesClient) StaticSitesGetStaticSiteBuilds(ctx context.Context, id StaticSiteId) (result StaticSitesGetStaticSiteBuildsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &StaticSitesGetStaticSiteBuildsCustomPager{},
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
		Values *[]StaticSiteBuildARMResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// StaticSitesGetStaticSiteBuildsComplete retrieves all the results into a single object
func (c StaticSiteBuildARMResourcesClient) StaticSitesGetStaticSiteBuildsComplete(ctx context.Context, id StaticSiteId) (StaticSitesGetStaticSiteBuildsCompleteResult, error) {
	return c.StaticSitesGetStaticSiteBuildsCompleteMatchingPredicate(ctx, id, StaticSiteBuildARMResourceOperationPredicate{})
}

// StaticSitesGetStaticSiteBuildsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c StaticSiteBuildARMResourcesClient) StaticSitesGetStaticSiteBuildsCompleteMatchingPredicate(ctx context.Context, id StaticSiteId, predicate StaticSiteBuildARMResourceOperationPredicate) (result StaticSitesGetStaticSiteBuildsCompleteResult, err error) {
	items := make([]StaticSiteBuildARMResource, 0)

	resp, err := c.StaticSitesGetStaticSiteBuilds(ctx, id)
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

	result = StaticSitesGetStaticSiteBuildsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
