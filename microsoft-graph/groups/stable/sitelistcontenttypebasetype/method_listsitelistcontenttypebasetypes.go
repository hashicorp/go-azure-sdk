package sitelistcontenttypebasetype

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

type ListSiteListContentTypeBaseTypesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ContentType
}

type ListSiteListContentTypeBaseTypesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ContentType
}

type ListSiteListContentTypeBaseTypesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteListContentTypeBaseTypesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteListContentTypeBaseTypes ...
func (c SiteListContentTypeBaseTypeClient) ListSiteListContentTypeBaseTypes(ctx context.Context, id GroupIdSiteIdListIdContentTypeId) (result ListSiteListContentTypeBaseTypesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSiteListContentTypeBaseTypesCustomPager{},
		Path:       fmt.Sprintf("%s/baseTypes", id.ID()),
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
		Values *[]stable.ContentType `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteListContentTypeBaseTypesComplete retrieves all the results into a single object
func (c SiteListContentTypeBaseTypeClient) ListSiteListContentTypeBaseTypesComplete(ctx context.Context, id GroupIdSiteIdListIdContentTypeId) (ListSiteListContentTypeBaseTypesCompleteResult, error) {
	return c.ListSiteListContentTypeBaseTypesCompleteMatchingPredicate(ctx, id, ContentTypeOperationPredicate{})
}

// ListSiteListContentTypeBaseTypesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteListContentTypeBaseTypeClient) ListSiteListContentTypeBaseTypesCompleteMatchingPredicate(ctx context.Context, id GroupIdSiteIdListIdContentTypeId, predicate ContentTypeOperationPredicate) (result ListSiteListContentTypeBaseTypesCompleteResult, err error) {
	items := make([]stable.ContentType, 0)

	resp, err := c.ListSiteListContentTypeBaseTypes(ctx, id)
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

	result = ListSiteListContentTypeBaseTypesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
