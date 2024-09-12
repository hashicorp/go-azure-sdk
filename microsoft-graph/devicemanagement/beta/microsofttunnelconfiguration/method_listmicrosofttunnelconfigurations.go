package microsofttunnelconfiguration

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

type ListMicrosoftTunnelConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MicrosoftTunnelConfiguration
}

type ListMicrosoftTunnelConfigurationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MicrosoftTunnelConfiguration
}

type ListMicrosoftTunnelConfigurationsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListMicrosoftTunnelConfigurationsOperationOptions() ListMicrosoftTunnelConfigurationsOperationOptions {
	return ListMicrosoftTunnelConfigurationsOperationOptions{}
}

func (o ListMicrosoftTunnelConfigurationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListMicrosoftTunnelConfigurationsOperationOptions) ToOData() *odata.Query {
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

func (o ListMicrosoftTunnelConfigurationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListMicrosoftTunnelConfigurationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMicrosoftTunnelConfigurationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMicrosoftTunnelConfigurations - Get microsoftTunnelConfigurations from deviceManagement. Collection of
// MicrosoftTunnelConfiguration settings associated with account.
func (c MicrosoftTunnelConfigurationClient) ListMicrosoftTunnelConfigurations(ctx context.Context, options ListMicrosoftTunnelConfigurationsOperationOptions) (result ListMicrosoftTunnelConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListMicrosoftTunnelConfigurationsCustomPager{},
		Path:          "/deviceManagement/microsoftTunnelConfigurations",
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
		Values *[]beta.MicrosoftTunnelConfiguration `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMicrosoftTunnelConfigurationsComplete retrieves all the results into a single object
func (c MicrosoftTunnelConfigurationClient) ListMicrosoftTunnelConfigurationsComplete(ctx context.Context, options ListMicrosoftTunnelConfigurationsOperationOptions) (ListMicrosoftTunnelConfigurationsCompleteResult, error) {
	return c.ListMicrosoftTunnelConfigurationsCompleteMatchingPredicate(ctx, options, MicrosoftTunnelConfigurationOperationPredicate{})
}

// ListMicrosoftTunnelConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MicrosoftTunnelConfigurationClient) ListMicrosoftTunnelConfigurationsCompleteMatchingPredicate(ctx context.Context, options ListMicrosoftTunnelConfigurationsOperationOptions, predicate MicrosoftTunnelConfigurationOperationPredicate) (result ListMicrosoftTunnelConfigurationsCompleteResult, err error) {
	items := make([]beta.MicrosoftTunnelConfiguration, 0)

	resp, err := c.ListMicrosoftTunnelConfigurations(ctx, options)
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

	result = ListMicrosoftTunnelConfigurationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
