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

type AddSitesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Site
}

type AddSitesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Site
}

type AddSitesOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultAddSitesOperationOptions() AddSitesOperationOptions {
	return AddSitesOperationOptions{}
}

func (o AddSitesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AddSitesOperationOptions) ToOData() *odata.Query {
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

func (o AddSitesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type AddSitesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *AddSitesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AddSites - Invoke action add. Follow a user's site or multiple sites.
func (c SiteClient) AddSites(ctx context.Context, id stable.GroupId, input AddSitesRequest, options AddSitesOperationOptions) (result AddSitesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &AddSitesCustomPager{},
		Path:          fmt.Sprintf("%s/sites/add", id.ID()),
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
		Values *[]stable.Site `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AddSitesComplete retrieves all the results into a single object
func (c SiteClient) AddSitesComplete(ctx context.Context, id stable.GroupId, input AddSitesRequest, options AddSitesOperationOptions) (AddSitesCompleteResult, error) {
	return c.AddSitesCompleteMatchingPredicate(ctx, id, input, options, SiteOperationPredicate{})
}

// AddSitesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteClient) AddSitesCompleteMatchingPredicate(ctx context.Context, id stable.GroupId, input AddSitesRequest, options AddSitesOperationOptions, predicate SiteOperationPredicate) (result AddSitesCompleteResult, err error) {
	items := make([]stable.Site, 0)

	resp, err := c.AddSites(ctx, id, input, options)
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

	result = AddSitesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
