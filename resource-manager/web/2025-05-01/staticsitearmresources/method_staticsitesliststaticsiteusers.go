package staticsitearmresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StaticSitesListStaticSiteUsersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]StaticSiteUserARMResource
}

type StaticSitesListStaticSiteUsersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []StaticSiteUserARMResource
}

type StaticSitesListStaticSiteUsersCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *StaticSitesListStaticSiteUsersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// StaticSitesListStaticSiteUsers ...
func (c StaticSiteARMResourcesClient) StaticSitesListStaticSiteUsers(ctx context.Context, id AuthProviderId) (result StaticSitesListStaticSiteUsersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &StaticSitesListStaticSiteUsersCustomPager{},
		Path:       fmt.Sprintf("%s/listUsers", id.ID()),
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
		Values *[]StaticSiteUserARMResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// StaticSitesListStaticSiteUsersComplete retrieves all the results into a single object
func (c StaticSiteARMResourcesClient) StaticSitesListStaticSiteUsersComplete(ctx context.Context, id AuthProviderId) (StaticSitesListStaticSiteUsersCompleteResult, error) {
	return c.StaticSitesListStaticSiteUsersCompleteMatchingPredicate(ctx, id, StaticSiteUserARMResourceOperationPredicate{})
}

// StaticSitesListStaticSiteUsersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c StaticSiteARMResourcesClient) StaticSitesListStaticSiteUsersCompleteMatchingPredicate(ctx context.Context, id AuthProviderId, predicate StaticSiteUserARMResourceOperationPredicate) (result StaticSitesListStaticSiteUsersCompleteResult, err error) {
	items := make([]StaticSiteUserARMResource, 0)

	resp, err := c.StaticSitesListStaticSiteUsers(ctx, id)
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

	result = StaticSitesListStaticSiteUsersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
