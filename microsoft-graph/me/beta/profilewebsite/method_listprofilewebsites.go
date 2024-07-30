package profilewebsite

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

type ListProfileWebsitesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PersonWebsite
}

type ListProfileWebsitesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PersonWebsite
}

type ListProfileWebsitesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListProfileWebsitesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListProfileWebsites ...
func (c ProfileWebsiteClient) ListProfileWebsites(ctx context.Context) (result ListProfileWebsitesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListProfileWebsitesCustomPager{},
		Path:       "/me/profile/websites",
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
		Values *[]beta.PersonWebsite `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListProfileWebsitesComplete retrieves all the results into a single object
func (c ProfileWebsiteClient) ListProfileWebsitesComplete(ctx context.Context) (ListProfileWebsitesCompleteResult, error) {
	return c.ListProfileWebsitesCompleteMatchingPredicate(ctx, PersonWebsiteOperationPredicate{})
}

// ListProfileWebsitesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProfileWebsiteClient) ListProfileWebsitesCompleteMatchingPredicate(ctx context.Context, predicate PersonWebsiteOperationPredicate) (result ListProfileWebsitesCompleteResult, err error) {
	items := make([]beta.PersonWebsite, 0)

	resp, err := c.ListProfileWebsites(ctx)
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

	result = ListProfileWebsitesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
