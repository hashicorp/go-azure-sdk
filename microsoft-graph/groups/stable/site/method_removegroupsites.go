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

type RemoveGroupSitesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Site
}

type RemoveGroupSitesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Site
}

type RemoveGroupSitesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *RemoveGroupSitesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// RemoveGroupSites ...
func (c SiteClient) RemoveGroupSites(ctx context.Context, id GroupId, input RemoveGroupSitesRequest) (result RemoveGroupSitesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &RemoveGroupSitesCustomPager{},
		Path:       fmt.Sprintf("%s/sites/remove", id.ID()),
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

// RemoveGroupSitesComplete retrieves all the results into a single object
func (c SiteClient) RemoveGroupSitesComplete(ctx context.Context, id GroupId, input RemoveGroupSitesRequest) (RemoveGroupSitesCompleteResult, error) {
	return c.RemoveGroupSitesCompleteMatchingPredicate(ctx, id, input, SiteOperationPredicate{})
}

// RemoveGroupSitesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteClient) RemoveGroupSitesCompleteMatchingPredicate(ctx context.Context, id GroupId, input RemoveGroupSitesRequest, predicate SiteOperationPredicate) (result RemoveGroupSitesCompleteResult, err error) {
	items := make([]stable.Site, 0)

	resp, err := c.RemoveGroupSites(ctx, id, input)
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

	result = RemoveGroupSitesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
