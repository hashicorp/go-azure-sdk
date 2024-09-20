package ndesconnector

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

type ListNdesConnectorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.NdesConnector
}

type ListNdesConnectorsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.NdesConnector
}

type ListNdesConnectorsOperationOptions struct {
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

func DefaultListNdesConnectorsOperationOptions() ListNdesConnectorsOperationOptions {
	return ListNdesConnectorsOperationOptions{}
}

func (o ListNdesConnectorsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListNdesConnectorsOperationOptions) ToOData() *odata.Query {
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

func (o ListNdesConnectorsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListNdesConnectorsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListNdesConnectorsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListNdesConnectors - Get ndesConnectors from deviceManagement. The collection of Ndes connectors for this account.
func (c NdesConnectorClient) ListNdesConnectors(ctx context.Context, options ListNdesConnectorsOperationOptions) (result ListNdesConnectorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListNdesConnectorsCustomPager{},
		Path:          "/deviceManagement/ndesConnectors",
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
		Values *[]beta.NdesConnector `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListNdesConnectorsComplete retrieves all the results into a single object
func (c NdesConnectorClient) ListNdesConnectorsComplete(ctx context.Context, options ListNdesConnectorsOperationOptions) (ListNdesConnectorsCompleteResult, error) {
	return c.ListNdesConnectorsCompleteMatchingPredicate(ctx, options, NdesConnectorOperationPredicate{})
}

// ListNdesConnectorsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NdesConnectorClient) ListNdesConnectorsCompleteMatchingPredicate(ctx context.Context, options ListNdesConnectorsOperationOptions, predicate NdesConnectorOperationPredicate) (result ListNdesConnectorsCompleteResult, err error) {
	items := make([]beta.NdesConnector, 0)

	resp, err := c.ListNdesConnectors(ctx, options)
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

	result = ListNdesConnectorsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
