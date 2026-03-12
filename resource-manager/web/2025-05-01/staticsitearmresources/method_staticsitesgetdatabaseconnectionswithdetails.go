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

type StaticSitesGetDatabaseConnectionsWithDetailsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DatabaseConnection
}

type StaticSitesGetDatabaseConnectionsWithDetailsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DatabaseConnection
}

type StaticSitesGetDatabaseConnectionsWithDetailsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *StaticSitesGetDatabaseConnectionsWithDetailsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// StaticSitesGetDatabaseConnectionsWithDetails ...
func (c StaticSiteARMResourcesClient) StaticSitesGetDatabaseConnectionsWithDetails(ctx context.Context, id StaticSiteId) (result StaticSitesGetDatabaseConnectionsWithDetailsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &StaticSitesGetDatabaseConnectionsWithDetailsCustomPager{},
		Path:       fmt.Sprintf("%s/showDatabaseConnections", id.ID()),
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
		Values *[]DatabaseConnection `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// StaticSitesGetDatabaseConnectionsWithDetailsComplete retrieves all the results into a single object
func (c StaticSiteARMResourcesClient) StaticSitesGetDatabaseConnectionsWithDetailsComplete(ctx context.Context, id StaticSiteId) (StaticSitesGetDatabaseConnectionsWithDetailsCompleteResult, error) {
	return c.StaticSitesGetDatabaseConnectionsWithDetailsCompleteMatchingPredicate(ctx, id, DatabaseConnectionOperationPredicate{})
}

// StaticSitesGetDatabaseConnectionsWithDetailsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c StaticSiteARMResourcesClient) StaticSitesGetDatabaseConnectionsWithDetailsCompleteMatchingPredicate(ctx context.Context, id StaticSiteId, predicate DatabaseConnectionOperationPredicate) (result StaticSitesGetDatabaseConnectionsWithDetailsCompleteResult, err error) {
	items := make([]DatabaseConnection, 0)

	resp, err := c.StaticSitesGetDatabaseConnectionsWithDetails(ctx, id)
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

	result = StaticSitesGetDatabaseConnectionsWithDetailsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
