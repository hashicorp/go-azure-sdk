package sitelistcolumn

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

type ListSiteListColumnsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ColumnDefinition
}

type ListSiteListColumnsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ColumnDefinition
}

type ListSiteListColumnsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteListColumnsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteListColumns ...
func (c SiteListColumnClient) ListSiteListColumns(ctx context.Context, id GroupIdSiteIdListId) (result ListSiteListColumnsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSiteListColumnsCustomPager{},
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
		Values *[]beta.ColumnDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteListColumnsComplete retrieves all the results into a single object
func (c SiteListColumnClient) ListSiteListColumnsComplete(ctx context.Context, id GroupIdSiteIdListId) (ListSiteListColumnsCompleteResult, error) {
	return c.ListSiteListColumnsCompleteMatchingPredicate(ctx, id, ColumnDefinitionOperationPredicate{})
}

// ListSiteListColumnsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteListColumnClient) ListSiteListColumnsCompleteMatchingPredicate(ctx context.Context, id GroupIdSiteIdListId, predicate ColumnDefinitionOperationPredicate) (result ListSiteListColumnsCompleteResult, err error) {
	items := make([]beta.ColumnDefinition, 0)

	resp, err := c.ListSiteListColumns(ctx, id)
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

	result = ListSiteListColumnsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
