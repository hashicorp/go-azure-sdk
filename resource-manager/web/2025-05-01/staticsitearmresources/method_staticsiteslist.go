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

type StaticSitesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]StaticSiteARMResource
}

type StaticSitesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []StaticSiteARMResource
}

type StaticSitesListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *StaticSitesListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// StaticSitesList ...
func (c StaticSiteARMResourcesClient) StaticSitesList(ctx context.Context, id commonids.SubscriptionId) (result StaticSitesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &StaticSitesListCustomPager{},
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

// StaticSitesListComplete retrieves all the results into a single object
func (c StaticSiteARMResourcesClient) StaticSitesListComplete(ctx context.Context, id commonids.SubscriptionId) (StaticSitesListCompleteResult, error) {
	return c.StaticSitesListCompleteMatchingPredicate(ctx, id, StaticSiteARMResourceOperationPredicate{})
}

// StaticSitesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c StaticSiteARMResourcesClient) StaticSitesListCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate StaticSiteARMResourceOperationPredicate) (result StaticSitesListCompleteResult, err error) {
	items := make([]StaticSiteARMResource, 0)

	resp, err := c.StaticSitesList(ctx, id)
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

	result = StaticSitesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
