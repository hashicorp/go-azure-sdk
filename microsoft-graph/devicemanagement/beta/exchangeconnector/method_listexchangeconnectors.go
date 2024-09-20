package exchangeconnector

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

type ListExchangeConnectorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementExchangeConnector
}

type ListExchangeConnectorsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementExchangeConnector
}

type ListExchangeConnectorsOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListExchangeConnectorsOperationOptions() ListExchangeConnectorsOperationOptions {
	return ListExchangeConnectorsOperationOptions{}
}

func (o ListExchangeConnectorsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListExchangeConnectorsOperationOptions) ToOData() *odata.Query {
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
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
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

func (o ListExchangeConnectorsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListExchangeConnectorsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListExchangeConnectorsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListExchangeConnectors - Get exchangeConnectors from deviceManagement. The list of Exchange Connectors configured by
// the tenant.
func (c ExchangeConnectorClient) ListExchangeConnectors(ctx context.Context, options ListExchangeConnectorsOperationOptions) (result ListExchangeConnectorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListExchangeConnectorsCustomPager{},
		Path:          "/deviceManagement/exchangeConnectors",
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
		Values *[]beta.DeviceManagementExchangeConnector `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListExchangeConnectorsComplete retrieves all the results into a single object
func (c ExchangeConnectorClient) ListExchangeConnectorsComplete(ctx context.Context, options ListExchangeConnectorsOperationOptions) (ListExchangeConnectorsCompleteResult, error) {
	return c.ListExchangeConnectorsCompleteMatchingPredicate(ctx, options, DeviceManagementExchangeConnectorOperationPredicate{})
}

// ListExchangeConnectorsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ExchangeConnectorClient) ListExchangeConnectorsCompleteMatchingPredicate(ctx context.Context, options ListExchangeConnectorsOperationOptions, predicate DeviceManagementExchangeConnectorOperationPredicate) (result ListExchangeConnectorsCompleteResult, err error) {
	items := make([]beta.DeviceManagementExchangeConnector, 0)

	resp, err := c.ListExchangeConnectors(ctx, options)
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

	result = ListExchangeConnectorsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
