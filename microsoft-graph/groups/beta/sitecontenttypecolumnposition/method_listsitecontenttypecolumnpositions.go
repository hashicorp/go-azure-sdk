package sitecontenttypecolumnposition

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

type ListSiteContentTypeColumnPositionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ColumnDefinition
}

type ListSiteContentTypeColumnPositionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ColumnDefinition
}

type ListSiteContentTypeColumnPositionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteContentTypeColumnPositionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteContentTypeColumnPositions ...
func (c SiteContentTypeColumnPositionClient) ListSiteContentTypeColumnPositions(ctx context.Context, id GroupIdSiteIdContentTypeId) (result ListSiteContentTypeColumnPositionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSiteContentTypeColumnPositionsCustomPager{},
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
		Values *[]beta.ColumnDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteContentTypeColumnPositionsComplete retrieves all the results into a single object
func (c SiteContentTypeColumnPositionClient) ListSiteContentTypeColumnPositionsComplete(ctx context.Context, id GroupIdSiteIdContentTypeId) (ListSiteContentTypeColumnPositionsCompleteResult, error) {
	return c.ListSiteContentTypeColumnPositionsCompleteMatchingPredicate(ctx, id, ColumnDefinitionOperationPredicate{})
}

// ListSiteContentTypeColumnPositionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteContentTypeColumnPositionClient) ListSiteContentTypeColumnPositionsCompleteMatchingPredicate(ctx context.Context, id GroupIdSiteIdContentTypeId, predicate ColumnDefinitionOperationPredicate) (result ListSiteContentTypeColumnPositionsCompleteResult, err error) {
	items := make([]beta.ColumnDefinition, 0)

	resp, err := c.ListSiteContentTypeColumnPositions(ctx, id)
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

	result = ListSiteContentTypeColumnPositionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
