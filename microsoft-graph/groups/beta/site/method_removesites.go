package site

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoveSitesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Site
}

type RemoveSitesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Site
}

type RemoveSitesOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultRemoveSitesOperationOptions() RemoveSitesOperationOptions {
	return RemoveSitesOperationOptions{}
}

func (o RemoveSitesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RemoveSitesOperationOptions) ToOData() *odata.Query {
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

func (o RemoveSitesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type RemoveSitesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *RemoveSitesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// RemoveSites - Invoke action remove. Unfollow a user's site or multiple sites.
func (c SiteClient) RemoveSites(ctx context.Context, id beta.GroupId, input RemoveSitesRequest, options RemoveSitesOperationOptions) (result RemoveSitesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &RemoveSitesCustomPager{},
		Path:          fmt.Sprintf("%s/sites/remove", id.ID()),
		RetryFunc:     options.RetryFunc,
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
		Values *[]beta.Site `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// RemoveSitesComplete retrieves all the results into a single object
func (c SiteClient) RemoveSitesComplete(ctx context.Context, id beta.GroupId, input RemoveSitesRequest, options RemoveSitesOperationOptions) (RemoveSitesCompleteResult, error) {
	return c.RemoveSitesCompleteMatchingPredicate(ctx, id, input, options, SiteOperationPredicate{})
}

// RemoveSitesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteClient) RemoveSitesCompleteMatchingPredicate(ctx context.Context, id beta.GroupId, input RemoveSitesRequest, options RemoveSitesOperationOptions, predicate SiteOperationPredicate) (result RemoveSitesCompleteResult, err error) {
	items := make([]beta.Site, 0)

	resp, err := c.RemoveSites(ctx, id, input, options)
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

	result = RemoveSitesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
