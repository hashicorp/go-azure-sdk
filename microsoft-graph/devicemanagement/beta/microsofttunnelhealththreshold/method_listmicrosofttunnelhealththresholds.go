package microsofttunnelhealththreshold

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

type ListMicrosoftTunnelHealthThresholdsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MicrosoftTunnelHealthThreshold
}

type ListMicrosoftTunnelHealthThresholdsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MicrosoftTunnelHealthThreshold
}

type ListMicrosoftTunnelHealthThresholdsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListMicrosoftTunnelHealthThresholdsOperationOptions() ListMicrosoftTunnelHealthThresholdsOperationOptions {
	return ListMicrosoftTunnelHealthThresholdsOperationOptions{}
}

func (o ListMicrosoftTunnelHealthThresholdsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListMicrosoftTunnelHealthThresholdsOperationOptions) ToOData() *odata.Query {
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

func (o ListMicrosoftTunnelHealthThresholdsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListMicrosoftTunnelHealthThresholdsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMicrosoftTunnelHealthThresholdsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMicrosoftTunnelHealthThresholds - Get microsoftTunnelHealthThresholds from deviceManagement. Collection of
// MicrosoftTunnelHealthThreshold settings associated with account.
func (c MicrosoftTunnelHealthThresholdClient) ListMicrosoftTunnelHealthThresholds(ctx context.Context, options ListMicrosoftTunnelHealthThresholdsOperationOptions) (result ListMicrosoftTunnelHealthThresholdsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListMicrosoftTunnelHealthThresholdsCustomPager{},
		Path:          "/deviceManagement/microsoftTunnelHealthThresholds",
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
		Values *[]beta.MicrosoftTunnelHealthThreshold `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMicrosoftTunnelHealthThresholdsComplete retrieves all the results into a single object
func (c MicrosoftTunnelHealthThresholdClient) ListMicrosoftTunnelHealthThresholdsComplete(ctx context.Context, options ListMicrosoftTunnelHealthThresholdsOperationOptions) (ListMicrosoftTunnelHealthThresholdsCompleteResult, error) {
	return c.ListMicrosoftTunnelHealthThresholdsCompleteMatchingPredicate(ctx, options, MicrosoftTunnelHealthThresholdOperationPredicate{})
}

// ListMicrosoftTunnelHealthThresholdsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MicrosoftTunnelHealthThresholdClient) ListMicrosoftTunnelHealthThresholdsCompleteMatchingPredicate(ctx context.Context, options ListMicrosoftTunnelHealthThresholdsOperationOptions, predicate MicrosoftTunnelHealthThresholdOperationPredicate) (result ListMicrosoftTunnelHealthThresholdsCompleteResult, err error) {
	items := make([]beta.MicrosoftTunnelHealthThreshold, 0)

	resp, err := c.ListMicrosoftTunnelHealthThresholds(ctx, options)
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

	result = ListMicrosoftTunnelHealthThresholdsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
