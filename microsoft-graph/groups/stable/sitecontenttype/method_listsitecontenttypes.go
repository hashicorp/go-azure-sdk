package sitecontenttype

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

type ListSiteContentTypesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ContentType
}

type ListSiteContentTypesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ContentType
}

type ListSiteContentTypesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteContentTypesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteContentTypes ...
func (c SiteContentTypeClient) ListSiteContentTypes(ctx context.Context, id GroupIdSiteId) (result ListSiteContentTypesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSiteContentTypesCustomPager{},
		Path:       fmt.Sprintf("%s/contentTypes", id.ID()),
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

// ListSiteContentTypesComplete retrieves all the results into a single object
func (c SiteContentTypeClient) ListSiteContentTypesComplete(ctx context.Context, id GroupIdSiteId) (ListSiteContentTypesCompleteResult, error) {
	return c.ListSiteContentTypesCompleteMatchingPredicate(ctx, id, ContentTypeOperationPredicate{})
}

// ListSiteContentTypesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteContentTypeClient) ListSiteContentTypesCompleteMatchingPredicate(ctx context.Context, id GroupIdSiteId, predicate ContentTypeOperationPredicate) (result ListSiteContentTypesCompleteResult, err error) {
	items := make([]stable.ContentType, 0)

	resp, err := c.ListSiteContentTypes(ctx, id)
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

	result = ListSiteContentTypesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
