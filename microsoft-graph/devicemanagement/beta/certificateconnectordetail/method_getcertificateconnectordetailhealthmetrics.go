package certificateconnectordetail

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

type GetCertificateConnectorDetailHealthMetricsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.KeyLongValuePair
}

type GetCertificateConnectorDetailHealthMetricsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.KeyLongValuePair
}

type GetCertificateConnectorDetailHealthMetricsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultGetCertificateConnectorDetailHealthMetricsOperationOptions() GetCertificateConnectorDetailHealthMetricsOperationOptions {
	return GetCertificateConnectorDetailHealthMetricsOperationOptions{}
}

func (o GetCertificateConnectorDetailHealthMetricsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetCertificateConnectorDetailHealthMetricsOperationOptions) ToOData() *odata.Query {
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

func (o GetCertificateConnectorDetailHealthMetricsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type GetCertificateConnectorDetailHealthMetricsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *GetCertificateConnectorDetailHealthMetricsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetCertificateConnectorDetailHealthMetrics - Invoke action getHealthMetrics
func (c CertificateConnectorDetailClient) GetCertificateConnectorDetailHealthMetrics(ctx context.Context, id beta.DeviceManagementCertificateConnectorDetailId, input GetCertificateConnectorDetailHealthMetricsRequest, options GetCertificateConnectorDetailHealthMetricsOperationOptions) (result GetCertificateConnectorDetailHealthMetricsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &GetCertificateConnectorDetailHealthMetricsCustomPager{},
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

// GetCertificateConnectorDetailHealthMetricsComplete retrieves all the results into a single object
func (c CertificateConnectorDetailClient) GetCertificateConnectorDetailHealthMetricsComplete(ctx context.Context, id beta.DeviceManagementCertificateConnectorDetailId, input GetCertificateConnectorDetailHealthMetricsRequest, options GetCertificateConnectorDetailHealthMetricsOperationOptions) (GetCertificateConnectorDetailHealthMetricsCompleteResult, error) {
	return c.GetCertificateConnectorDetailHealthMetricsCompleteMatchingPredicate(ctx, id, input, options, KeyLongValuePairOperationPredicate{})
}

// GetCertificateConnectorDetailHealthMetricsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CertificateConnectorDetailClient) GetCertificateConnectorDetailHealthMetricsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementCertificateConnectorDetailId, input GetCertificateConnectorDetailHealthMetricsRequest, options GetCertificateConnectorDetailHealthMetricsOperationOptions, predicate KeyLongValuePairOperationPredicate) (result GetCertificateConnectorDetailHealthMetricsCompleteResult, err error) {
	items := make([]beta.KeyLongValuePair, 0)

	resp, err := c.GetCertificateConnectorDetailHealthMetrics(ctx, id, input, options)
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

	result = GetCertificateConnectorDetailHealthMetricsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
