package staticsiteuserprovidedfunctionapparmresourceoperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StaticSitesGetUserProvidedFunctionAppsForStaticSiteOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]StaticSiteUserProvidedFunctionAppARMResource
}

type StaticSitesGetUserProvidedFunctionAppsForStaticSiteCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []StaticSiteUserProvidedFunctionAppARMResource
}

type StaticSitesGetUserProvidedFunctionAppsForStaticSiteCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *StaticSitesGetUserProvidedFunctionAppsForStaticSiteCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// StaticSitesGetUserProvidedFunctionAppsForStaticSite ...
func (c StaticSiteUserProvidedFunctionAppARMResourceOperationGroupClient) StaticSitesGetUserProvidedFunctionAppsForStaticSite(ctx context.Context, id StaticSiteId) (result StaticSitesGetUserProvidedFunctionAppsForStaticSiteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &StaticSitesGetUserProvidedFunctionAppsForStaticSiteCustomPager{},
		Path:       fmt.Sprintf("%s/userProvidedFunctionApps", id.ID()),
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
		Values *[]StaticSiteUserProvidedFunctionAppARMResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// StaticSitesGetUserProvidedFunctionAppsForStaticSiteComplete retrieves all the results into a single object
func (c StaticSiteUserProvidedFunctionAppARMResourceOperationGroupClient) StaticSitesGetUserProvidedFunctionAppsForStaticSiteComplete(ctx context.Context, id StaticSiteId) (StaticSitesGetUserProvidedFunctionAppsForStaticSiteCompleteResult, error) {
	return c.StaticSitesGetUserProvidedFunctionAppsForStaticSiteCompleteMatchingPredicate(ctx, id, StaticSiteUserProvidedFunctionAppARMResourceOperationPredicate{})
}

// StaticSitesGetUserProvidedFunctionAppsForStaticSiteCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c StaticSiteUserProvidedFunctionAppARMResourceOperationGroupClient) StaticSitesGetUserProvidedFunctionAppsForStaticSiteCompleteMatchingPredicate(ctx context.Context, id StaticSiteId, predicate StaticSiteUserProvidedFunctionAppARMResourceOperationPredicate) (result StaticSitesGetUserProvidedFunctionAppsForStaticSiteCompleteResult, err error) {
	items := make([]StaticSiteUserProvidedFunctionAppARMResource, 0)

	resp, err := c.StaticSitesGetUserProvidedFunctionAppsForStaticSite(ctx, id)
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

	result = StaticSitesGetUserProvidedFunctionAppsForStaticSiteCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
