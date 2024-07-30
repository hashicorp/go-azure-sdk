package sitecontenttypebasetype

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

type ListSiteContentTypeBaseTypesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ContentType
}

type ListSiteContentTypeBaseTypesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ContentType
}

type ListSiteContentTypeBaseTypesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteContentTypeBaseTypesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteContentTypeBaseTypes ...
func (c SiteContentTypeBaseTypeClient) ListSiteContentTypeBaseTypes(ctx context.Context, id GroupIdSiteIdContentTypeId) (result ListSiteContentTypeBaseTypesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSiteContentTypeBaseTypesCustomPager{},
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
		Values *[]beta.ContentType `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteContentTypeBaseTypesComplete retrieves all the results into a single object
func (c SiteContentTypeBaseTypeClient) ListSiteContentTypeBaseTypesComplete(ctx context.Context, id GroupIdSiteIdContentTypeId) (ListSiteContentTypeBaseTypesCompleteResult, error) {
	return c.ListSiteContentTypeBaseTypesCompleteMatchingPredicate(ctx, id, ContentTypeOperationPredicate{})
}

// ListSiteContentTypeBaseTypesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteContentTypeBaseTypeClient) ListSiteContentTypeBaseTypesCompleteMatchingPredicate(ctx context.Context, id GroupIdSiteIdContentTypeId, predicate ContentTypeOperationPredicate) (result ListSiteContentTypeBaseTypesCompleteResult, err error) {
	items := make([]beta.ContentType, 0)

	resp, err := c.ListSiteContentTypeBaseTypes(ctx, id)
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

	result = ListSiteContentTypeBaseTypesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
