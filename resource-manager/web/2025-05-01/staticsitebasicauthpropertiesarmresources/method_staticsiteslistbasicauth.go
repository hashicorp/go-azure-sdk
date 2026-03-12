package staticsitebasicauthpropertiesarmresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StaticSitesListBasicAuthOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]StaticSiteBasicAuthPropertiesARMResource
}

type StaticSitesListBasicAuthCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []StaticSiteBasicAuthPropertiesARMResource
}

type StaticSitesListBasicAuthCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *StaticSitesListBasicAuthCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// StaticSitesListBasicAuth ...
func (c StaticSiteBasicAuthPropertiesARMResourcesClient) StaticSitesListBasicAuth(ctx context.Context, id StaticSiteId) (result StaticSitesListBasicAuthOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &StaticSitesListBasicAuthCustomPager{},
		Path:       fmt.Sprintf("%s/basicAuth", id.ID()),
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
		Values *[]StaticSiteBasicAuthPropertiesARMResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// StaticSitesListBasicAuthComplete retrieves all the results into a single object
func (c StaticSiteBasicAuthPropertiesARMResourcesClient) StaticSitesListBasicAuthComplete(ctx context.Context, id StaticSiteId) (StaticSitesListBasicAuthCompleteResult, error) {
	return c.StaticSitesListBasicAuthCompleteMatchingPredicate(ctx, id, StaticSiteBasicAuthPropertiesARMResourceOperationPredicate{})
}

// StaticSitesListBasicAuthCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c StaticSiteBasicAuthPropertiesARMResourcesClient) StaticSitesListBasicAuthCompleteMatchingPredicate(ctx context.Context, id StaticSiteId, predicate StaticSiteBasicAuthPropertiesARMResourceOperationPredicate) (result StaticSitesListBasicAuthCompleteResult, err error) {
	items := make([]StaticSiteBasicAuthPropertiesARMResource, 0)

	resp, err := c.StaticSitesListBasicAuth(ctx, id)
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

	result = StaticSitesListBasicAuthCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
