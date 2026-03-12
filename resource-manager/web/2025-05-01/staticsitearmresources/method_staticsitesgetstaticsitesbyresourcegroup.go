package staticsitearmresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StaticSitesGetStaticSitesByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]StaticSiteARMResource
}

type StaticSitesGetStaticSitesByResourceGroupCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []StaticSiteARMResource
}

type StaticSitesGetStaticSitesByResourceGroupCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *StaticSitesGetStaticSitesByResourceGroupCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// StaticSitesGetStaticSitesByResourceGroup ...
func (c StaticSiteARMResourcesClient) StaticSitesGetStaticSitesByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (result StaticSitesGetStaticSitesByResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &StaticSitesGetStaticSitesByResourceGroupCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.Web/staticSites", id.ID()),
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
		Values *[]StaticSiteARMResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// StaticSitesGetStaticSitesByResourceGroupComplete retrieves all the results into a single object
func (c StaticSiteARMResourcesClient) StaticSitesGetStaticSitesByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId) (StaticSitesGetStaticSitesByResourceGroupCompleteResult, error) {
	return c.StaticSitesGetStaticSitesByResourceGroupCompleteMatchingPredicate(ctx, id, StaticSiteARMResourceOperationPredicate{})
}

// StaticSitesGetStaticSitesByResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c StaticSiteARMResourcesClient) StaticSitesGetStaticSitesByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate StaticSiteARMResourceOperationPredicate) (result StaticSitesGetStaticSitesByResourceGroupCompleteResult, err error) {
	items := make([]StaticSiteARMResource, 0)

	resp, err := c.StaticSitesGetStaticSitesByResourceGroup(ctx, id)
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

	result = StaticSitesGetStaticSitesByResourceGroupCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
