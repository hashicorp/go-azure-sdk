package staticsitecustomdomainoverviewarmresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StaticSitesListStaticSiteCustomDomainsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]StaticSiteCustomDomainOverviewARMResource
}

type StaticSitesListStaticSiteCustomDomainsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []StaticSiteCustomDomainOverviewARMResource
}

type StaticSitesListStaticSiteCustomDomainsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *StaticSitesListStaticSiteCustomDomainsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// StaticSitesListStaticSiteCustomDomains ...
func (c StaticSiteCustomDomainOverviewARMResourcesClient) StaticSitesListStaticSiteCustomDomains(ctx context.Context, id StaticSiteId) (result StaticSitesListStaticSiteCustomDomainsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &StaticSitesListStaticSiteCustomDomainsCustomPager{},
		Path:       fmt.Sprintf("%s/customDomains", id.ID()),
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
		Values *[]StaticSiteCustomDomainOverviewARMResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// StaticSitesListStaticSiteCustomDomainsComplete retrieves all the results into a single object
func (c StaticSiteCustomDomainOverviewARMResourcesClient) StaticSitesListStaticSiteCustomDomainsComplete(ctx context.Context, id StaticSiteId) (StaticSitesListStaticSiteCustomDomainsCompleteResult, error) {
	return c.StaticSitesListStaticSiteCustomDomainsCompleteMatchingPredicate(ctx, id, StaticSiteCustomDomainOverviewARMResourceOperationPredicate{})
}

// StaticSitesListStaticSiteCustomDomainsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c StaticSiteCustomDomainOverviewARMResourcesClient) StaticSitesListStaticSiteCustomDomainsCompleteMatchingPredicate(ctx context.Context, id StaticSiteId, predicate StaticSiteCustomDomainOverviewARMResourceOperationPredicate) (result StaticSitesListStaticSiteCustomDomainsCompleteResult, err error) {
	items := make([]StaticSiteCustomDomainOverviewARMResource, 0)

	resp, err := c.StaticSitesListStaticSiteCustomDomains(ctx, id)
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

	result = StaticSitesListStaticSiteCustomDomainsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
