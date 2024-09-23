package microsofttunnelsitemicrosofttunnelserver

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

type GetMicrosoftTunnelSiteServerHealthMetricsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.KeyLongValuePair
}

type GetMicrosoftTunnelSiteServerHealthMetricsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.KeyLongValuePair
}

type GetMicrosoftTunnelSiteServerHealthMetricsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultGetMicrosoftTunnelSiteServerHealthMetricsOperationOptions() GetMicrosoftTunnelSiteServerHealthMetricsOperationOptions {
	return GetMicrosoftTunnelSiteServerHealthMetricsOperationOptions{}
}

func (o GetMicrosoftTunnelSiteServerHealthMetricsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetMicrosoftTunnelSiteServerHealthMetricsOperationOptions) ToOData() *odata.Query {
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

func (o GetMicrosoftTunnelSiteServerHealthMetricsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type GetMicrosoftTunnelSiteServerHealthMetricsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *GetMicrosoftTunnelSiteServerHealthMetricsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetMicrosoftTunnelSiteServerHealthMetrics - Invoke action getHealthMetrics
func (c MicrosoftTunnelSiteMicrosoftTunnelServerClient) GetMicrosoftTunnelSiteServerHealthMetrics(ctx context.Context, id beta.DeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId, input GetMicrosoftTunnelSiteServerHealthMetricsRequest, options GetMicrosoftTunnelSiteServerHealthMetricsOperationOptions) (result GetMicrosoftTunnelSiteServerHealthMetricsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &GetMicrosoftTunnelSiteServerHealthMetricsCustomPager{},
		Path:          fmt.Sprintf("%s/getHealthMetrics", id.ID()),
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
		Values *[]beta.KeyLongValuePair `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetMicrosoftTunnelSiteServerHealthMetricsComplete retrieves all the results into a single object
func (c MicrosoftTunnelSiteMicrosoftTunnelServerClient) GetMicrosoftTunnelSiteServerHealthMetricsComplete(ctx context.Context, id beta.DeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId, input GetMicrosoftTunnelSiteServerHealthMetricsRequest, options GetMicrosoftTunnelSiteServerHealthMetricsOperationOptions) (GetMicrosoftTunnelSiteServerHealthMetricsCompleteResult, error) {
	return c.GetMicrosoftTunnelSiteServerHealthMetricsCompleteMatchingPredicate(ctx, id, input, options, KeyLongValuePairOperationPredicate{})
}

// GetMicrosoftTunnelSiteServerHealthMetricsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MicrosoftTunnelSiteMicrosoftTunnelServerClient) GetMicrosoftTunnelSiteServerHealthMetricsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId, input GetMicrosoftTunnelSiteServerHealthMetricsRequest, options GetMicrosoftTunnelSiteServerHealthMetricsOperationOptions, predicate KeyLongValuePairOperationPredicate) (result GetMicrosoftTunnelSiteServerHealthMetricsCompleteResult, err error) {
	items := make([]beta.KeyLongValuePair, 0)

	resp, err := c.GetMicrosoftTunnelSiteServerHealthMetrics(ctx, id, input, options)
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

	result = GetMicrosoftTunnelSiteServerHealthMetricsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
