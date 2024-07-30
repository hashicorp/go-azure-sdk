package sitelistcontenttypecolumnposition

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSiteListContentTypeColumnPositionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ColumnDefinition
}

type ListSiteListContentTypeColumnPositionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ColumnDefinition
}

type ListSiteListContentTypeColumnPositionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteListContentTypeColumnPositionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteListContentTypeColumnPositions ...
func (c SiteListContentTypeColumnPositionClient) ListSiteListContentTypeColumnPositions(ctx context.Context, id GroupIdSiteIdListIdContentTypeId) (result ListSiteListContentTypeColumnPositionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSiteListContentTypeColumnPositionsCustomPager{},
		Path:       fmt.Sprintf("%s/columnPositions", id.ID()),
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
		Values *[]stable.ColumnDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteListContentTypeColumnPositionsComplete retrieves all the results into a single object
func (c SiteListContentTypeColumnPositionClient) ListSiteListContentTypeColumnPositionsComplete(ctx context.Context, id GroupIdSiteIdListIdContentTypeId) (ListSiteListContentTypeColumnPositionsCompleteResult, error) {
	return c.ListSiteListContentTypeColumnPositionsCompleteMatchingPredicate(ctx, id, ColumnDefinitionOperationPredicate{})
}

// ListSiteListContentTypeColumnPositionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteListContentTypeColumnPositionClient) ListSiteListContentTypeColumnPositionsCompleteMatchingPredicate(ctx context.Context, id GroupIdSiteIdListIdContentTypeId, predicate ColumnDefinitionOperationPredicate) (result ListSiteListContentTypeColumnPositionsCompleteResult, err error) {
	items := make([]stable.ColumnDefinition, 0)

	resp, err := c.ListSiteListContentTypeColumnPositions(ctx, id)
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

	result = ListSiteListContentTypeColumnPositionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
