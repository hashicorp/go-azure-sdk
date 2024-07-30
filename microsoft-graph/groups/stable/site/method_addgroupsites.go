package site

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

type AddGroupSitesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Site
}

type AddGroupSitesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Site
}

type AddGroupSitesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *AddGroupSitesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AddGroupSites ...
func (c SiteClient) AddGroupSites(ctx context.Context, id GroupId, input AddGroupSitesRequest) (result AddGroupSitesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &AddGroupSitesCustomPager{},
		Path:       fmt.Sprintf("%s/sites/add", id.ID()),
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
		Values *[]stable.Site `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AddGroupSitesComplete retrieves all the results into a single object
func (c SiteClient) AddGroupSitesComplete(ctx context.Context, id GroupId, input AddGroupSitesRequest) (AddGroupSitesCompleteResult, error) {
	return c.AddGroupSitesCompleteMatchingPredicate(ctx, id, input, SiteOperationPredicate{})
}

// AddGroupSitesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteClient) AddGroupSitesCompleteMatchingPredicate(ctx context.Context, id GroupId, input AddGroupSitesRequest, predicate SiteOperationPredicate) (result AddGroupSitesCompleteResult, err error) {
	items := make([]stable.Site, 0)

	resp, err := c.AddGroupSites(ctx, id, input)
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

	result = AddGroupSitesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
