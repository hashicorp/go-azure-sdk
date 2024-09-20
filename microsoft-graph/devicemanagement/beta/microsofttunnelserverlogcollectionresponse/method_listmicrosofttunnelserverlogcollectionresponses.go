package microsofttunnelserverlogcollectionresponse

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

type ListMicrosoftTunnelServerLogCollectionResponsesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MicrosoftTunnelServerLogCollectionResponse
}

type ListMicrosoftTunnelServerLogCollectionResponsesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MicrosoftTunnelServerLogCollectionResponse
}

type ListMicrosoftTunnelServerLogCollectionResponsesOperationOptions struct {
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

func DefaultListMicrosoftTunnelServerLogCollectionResponsesOperationOptions() ListMicrosoftTunnelServerLogCollectionResponsesOperationOptions {
	return ListMicrosoftTunnelServerLogCollectionResponsesOperationOptions{}
}

func (o ListMicrosoftTunnelServerLogCollectionResponsesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListMicrosoftTunnelServerLogCollectionResponsesOperationOptions) ToOData() *odata.Query {
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

func (o ListMicrosoftTunnelServerLogCollectionResponsesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListMicrosoftTunnelServerLogCollectionResponsesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMicrosoftTunnelServerLogCollectionResponsesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMicrosoftTunnelServerLogCollectionResponses - Get microsoftTunnelServerLogCollectionResponses from
// deviceManagement. Collection of MicrosoftTunnelServerLogCollectionResponse settings associated with account.
func (c MicrosoftTunnelServerLogCollectionResponseClient) ListMicrosoftTunnelServerLogCollectionResponses(ctx context.Context, options ListMicrosoftTunnelServerLogCollectionResponsesOperationOptions) (result ListMicrosoftTunnelServerLogCollectionResponsesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListMicrosoftTunnelServerLogCollectionResponsesCustomPager{},
		Path:          "/deviceManagement/microsoftTunnelServerLogCollectionResponses",
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
		Values *[]beta.MicrosoftTunnelServerLogCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMicrosoftTunnelServerLogCollectionResponsesComplete retrieves all the results into a single object
func (c MicrosoftTunnelServerLogCollectionResponseClient) ListMicrosoftTunnelServerLogCollectionResponsesComplete(ctx context.Context, options ListMicrosoftTunnelServerLogCollectionResponsesOperationOptions) (ListMicrosoftTunnelServerLogCollectionResponsesCompleteResult, error) {
	return c.ListMicrosoftTunnelServerLogCollectionResponsesCompleteMatchingPredicate(ctx, options, MicrosoftTunnelServerLogCollectionResponseOperationPredicate{})
}

// ListMicrosoftTunnelServerLogCollectionResponsesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MicrosoftTunnelServerLogCollectionResponseClient) ListMicrosoftTunnelServerLogCollectionResponsesCompleteMatchingPredicate(ctx context.Context, options ListMicrosoftTunnelServerLogCollectionResponsesOperationOptions, predicate MicrosoftTunnelServerLogCollectionResponseOperationPredicate) (result ListMicrosoftTunnelServerLogCollectionResponsesCompleteResult, err error) {
	items := make([]beta.MicrosoftTunnelServerLogCollectionResponse, 0)

	resp, err := c.ListMicrosoftTunnelServerLogCollectionResponses(ctx, options)
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

	result = ListMicrosoftTunnelServerLogCollectionResponsesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
