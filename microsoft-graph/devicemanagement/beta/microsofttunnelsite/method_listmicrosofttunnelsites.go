package microsofttunnelsite

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

type ListMicrosoftTunnelSitesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MicrosoftTunnelSite
}

type ListMicrosoftTunnelSitesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MicrosoftTunnelSite
}

type ListMicrosoftTunnelSitesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListMicrosoftTunnelSitesOperationOptions() ListMicrosoftTunnelSitesOperationOptions {
	return ListMicrosoftTunnelSitesOperationOptions{}
}

func (o ListMicrosoftTunnelSitesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListMicrosoftTunnelSitesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListMicrosoftTunnelSitesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListMicrosoftTunnelSitesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMicrosoftTunnelSitesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMicrosoftTunnelSites - Get microsoftTunnelSites from deviceManagement. Collection of MicrosoftTunnelSite settings
// associated with account.
func (c MicrosoftTunnelSiteClient) ListMicrosoftTunnelSites(ctx context.Context, options ListMicrosoftTunnelSitesOperationOptions) (result ListMicrosoftTunnelSitesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListMicrosoftTunnelSitesCustomPager{},
		Path:          "/deviceManagement/microsoftTunnelSites",
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
		Values *[]beta.MicrosoftTunnelSite `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMicrosoftTunnelSitesComplete retrieves all the results into a single object
func (c MicrosoftTunnelSiteClient) ListMicrosoftTunnelSitesComplete(ctx context.Context, options ListMicrosoftTunnelSitesOperationOptions) (ListMicrosoftTunnelSitesCompleteResult, error) {
	return c.ListMicrosoftTunnelSitesCompleteMatchingPredicate(ctx, options, MicrosoftTunnelSiteOperationPredicate{})
}

// ListMicrosoftTunnelSitesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MicrosoftTunnelSiteClient) ListMicrosoftTunnelSitesCompleteMatchingPredicate(ctx context.Context, options ListMicrosoftTunnelSitesOperationOptions, predicate MicrosoftTunnelSiteOperationPredicate) (result ListMicrosoftTunnelSitesCompleteResult, err error) {
	items := make([]beta.MicrosoftTunnelSite, 0)

	resp, err := c.ListMicrosoftTunnelSites(ctx, options)
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

	result = ListMicrosoftTunnelSitesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
