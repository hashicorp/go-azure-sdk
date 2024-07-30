package sitelistcontenttypecolumn

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

type ListSiteListContentTypeColumnsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ColumnDefinition
}

type ListSiteListContentTypeColumnsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ColumnDefinition
}

type ListSiteListContentTypeColumnsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteListContentTypeColumnsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteListContentTypeColumns ...
func (c SiteListContentTypeColumnClient) ListSiteListContentTypeColumns(ctx context.Context, id GroupIdSiteIdListIdContentTypeId) (result ListSiteListContentTypeColumnsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSiteListContentTypeColumnsCustomPager{},
		Path:       fmt.Sprintf("%s/columns", id.ID()),
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

// ListSiteListContentTypeColumnsComplete retrieves all the results into a single object
func (c SiteListContentTypeColumnClient) ListSiteListContentTypeColumnsComplete(ctx context.Context, id GroupIdSiteIdListIdContentTypeId) (ListSiteListContentTypeColumnsCompleteResult, error) {
	return c.ListSiteListContentTypeColumnsCompleteMatchingPredicate(ctx, id, ColumnDefinitionOperationPredicate{})
}

// ListSiteListContentTypeColumnsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteListContentTypeColumnClient) ListSiteListContentTypeColumnsCompleteMatchingPredicate(ctx context.Context, id GroupIdSiteIdListIdContentTypeId, predicate ColumnDefinitionOperationPredicate) (result ListSiteListContentTypeColumnsCompleteResult, err error) {
	items := make([]stable.ColumnDefinition, 0)

	resp, err := c.ListSiteListContentTypeColumns(ctx, id)
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

	result = ListSiteListContentTypeColumnsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
