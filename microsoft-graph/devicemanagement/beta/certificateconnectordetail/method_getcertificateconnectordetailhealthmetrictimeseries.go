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

type GetCertificateConnectorDetailHealthMetricTimeSeriesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CertificateConnectorHealthMetricValue
}

type GetCertificateConnectorDetailHealthMetricTimeSeriesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CertificateConnectorHealthMetricValue
}

type GetCertificateConnectorDetailHealthMetricTimeSeriesOperationOptions struct {
	Metadata *odata.Metadata
	Skip     *int64
	Top      *int64
}

func DefaultGetCertificateConnectorDetailHealthMetricTimeSeriesOperationOptions() GetCertificateConnectorDetailHealthMetricTimeSeriesOperationOptions {
	return GetCertificateConnectorDetailHealthMetricTimeSeriesOperationOptions{}
}

func (o GetCertificateConnectorDetailHealthMetricTimeSeriesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetCertificateConnectorDetailHealthMetricTimeSeriesOperationOptions) ToOData() *odata.Query {
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

func (o GetCertificateConnectorDetailHealthMetricTimeSeriesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type GetCertificateConnectorDetailHealthMetricTimeSeriesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *GetCertificateConnectorDetailHealthMetricTimeSeriesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetCertificateConnectorDetailHealthMetricTimeSeries - Invoke action getHealthMetricTimeSeries
func (c CertificateConnectorDetailClient) GetCertificateConnectorDetailHealthMetricTimeSeries(ctx context.Context, id beta.DeviceManagementCertificateConnectorDetailId, input GetCertificateConnectorDetailHealthMetricTimeSeriesRequest, options GetCertificateConnectorDetailHealthMetricTimeSeriesOperationOptions) (result GetCertificateConnectorDetailHealthMetricTimeSeriesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &GetCertificateConnectorDetailHealthMetricTimeSeriesCustomPager{},
		Path:          fmt.Sprintf("%s/getHealthMetricTimeSeries", id.ID()),
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
		Values *[]beta.CertificateConnectorHealthMetricValue `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetCertificateConnectorDetailHealthMetricTimeSeriesComplete retrieves all the results into a single object
func (c CertificateConnectorDetailClient) GetCertificateConnectorDetailHealthMetricTimeSeriesComplete(ctx context.Context, id beta.DeviceManagementCertificateConnectorDetailId, input GetCertificateConnectorDetailHealthMetricTimeSeriesRequest, options GetCertificateConnectorDetailHealthMetricTimeSeriesOperationOptions) (GetCertificateConnectorDetailHealthMetricTimeSeriesCompleteResult, error) {
	return c.GetCertificateConnectorDetailHealthMetricTimeSeriesCompleteMatchingPredicate(ctx, id, input, options, CertificateConnectorHealthMetricValueOperationPredicate{})
}

// GetCertificateConnectorDetailHealthMetricTimeSeriesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CertificateConnectorDetailClient) GetCertificateConnectorDetailHealthMetricTimeSeriesCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementCertificateConnectorDetailId, input GetCertificateConnectorDetailHealthMetricTimeSeriesRequest, options GetCertificateConnectorDetailHealthMetricTimeSeriesOperationOptions, predicate CertificateConnectorHealthMetricValueOperationPredicate) (result GetCertificateConnectorDetailHealthMetricTimeSeriesCompleteResult, err error) {
	items := make([]beta.CertificateConnectorHealthMetricValue, 0)

	resp, err := c.GetCertificateConnectorDetailHealthMetricTimeSeries(ctx, id, input, options)
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

	result = GetCertificateConnectorDetailHealthMetricTimeSeriesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
