package followedsite

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

type AddFollowedSitesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Site
}

type AddFollowedSitesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Site
}

type AddFollowedSitesOperationOptions struct {
	Metadata *odata.Metadata
	Skip     *int64
	Top      *int64
}

func DefaultAddFollowedSitesOperationOptions() AddFollowedSitesOperationOptions {
	return AddFollowedSitesOperationOptions{}
}

func (o AddFollowedSitesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AddFollowedSitesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o AddFollowedSitesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type AddFollowedSitesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *AddFollowedSitesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AddFollowedSites - Invoke action add. Follow a user's site or multiple sites.
func (c FollowedSiteClient) AddFollowedSites(ctx context.Context, input AddFollowedSitesRequest, options AddFollowedSitesOperationOptions) (result AddFollowedSitesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &AddFollowedSitesCustomPager{},
		Path:          "/me/followedSites/add",
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

// AddFollowedSitesComplete retrieves all the results into a single object
func (c FollowedSiteClient) AddFollowedSitesComplete(ctx context.Context, input AddFollowedSitesRequest, options AddFollowedSitesOperationOptions) (AddFollowedSitesCompleteResult, error) {
	return c.AddFollowedSitesCompleteMatchingPredicate(ctx, input, options, SiteOperationPredicate{})
}

// AddFollowedSitesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c FollowedSiteClient) AddFollowedSitesCompleteMatchingPredicate(ctx context.Context, input AddFollowedSitesRequest, options AddFollowedSitesOperationOptions, predicate SiteOperationPredicate) (result AddFollowedSitesCompleteResult, err error) {
	items := make([]stable.Site, 0)

	resp, err := c.AddFollowedSites(ctx, input, options)
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

	result = AddFollowedSitesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
