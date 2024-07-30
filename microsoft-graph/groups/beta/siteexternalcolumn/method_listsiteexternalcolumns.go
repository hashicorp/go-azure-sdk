package siteexternalcolumn

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSiteExternalColumnsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ColumnDefinition
}

type ListSiteExternalColumnsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ColumnDefinition
}

type ListSiteExternalColumnsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteExternalColumnsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteExternalColumns ...
func (c SiteExternalColumnClient) ListSiteExternalColumns(ctx context.Context, id GroupIdSiteId) (result ListSiteExternalColumnsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSiteExternalColumnsCustomPager{},
		Path:       fmt.Sprintf("%s/externalColumns", id.ID()),
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
		Values *[]beta.ColumnDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteExternalColumnsComplete retrieves all the results into a single object
func (c SiteExternalColumnClient) ListSiteExternalColumnsComplete(ctx context.Context, id GroupIdSiteId) (ListSiteExternalColumnsCompleteResult, error) {
	return c.ListSiteExternalColumnsCompleteMatchingPredicate(ctx, id, ColumnDefinitionOperationPredicate{})
}

// ListSiteExternalColumnsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteExternalColumnClient) ListSiteExternalColumnsCompleteMatchingPredicate(ctx context.Context, id GroupIdSiteId, predicate ColumnDefinitionOperationPredicate) (result ListSiteExternalColumnsCompleteResult, err error) {
	items := make([]beta.ColumnDefinition, 0)

	resp, err := c.ListSiteExternalColumns(ctx, id)
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

	result = ListSiteExternalColumnsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
