package staticsitelinkedbackendarmresourceoperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StaticSitesGetLinkedBackendsForBuildOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]StaticSiteLinkedBackendARMResource
}

type StaticSitesGetLinkedBackendsForBuildCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []StaticSiteLinkedBackendARMResource
}

type StaticSitesGetLinkedBackendsForBuildCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *StaticSitesGetLinkedBackendsForBuildCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// StaticSitesGetLinkedBackendsForBuild ...
func (c StaticSiteLinkedBackendARMResourceOperationGroupClient) StaticSitesGetLinkedBackendsForBuild(ctx context.Context, id BuildId) (result StaticSitesGetLinkedBackendsForBuildOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &StaticSitesGetLinkedBackendsForBuildCustomPager{},
		Path:       fmt.Sprintf("%s/linkedBackends", id.ID()),
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
		Values *[]StaticSiteLinkedBackendARMResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// StaticSitesGetLinkedBackendsForBuildComplete retrieves all the results into a single object
func (c StaticSiteLinkedBackendARMResourceOperationGroupClient) StaticSitesGetLinkedBackendsForBuildComplete(ctx context.Context, id BuildId) (StaticSitesGetLinkedBackendsForBuildCompleteResult, error) {
	return c.StaticSitesGetLinkedBackendsForBuildCompleteMatchingPredicate(ctx, id, StaticSiteLinkedBackendARMResourceOperationPredicate{})
}

// StaticSitesGetLinkedBackendsForBuildCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c StaticSiteLinkedBackendARMResourceOperationGroupClient) StaticSitesGetLinkedBackendsForBuildCompleteMatchingPredicate(ctx context.Context, id BuildId, predicate StaticSiteLinkedBackendARMResourceOperationPredicate) (result StaticSitesGetLinkedBackendsForBuildCompleteResult, err error) {
	items := make([]StaticSiteLinkedBackendARMResource, 0)

	resp, err := c.StaticSitesGetLinkedBackendsForBuild(ctx, id)
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

	result = StaticSitesGetLinkedBackendsForBuildCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
