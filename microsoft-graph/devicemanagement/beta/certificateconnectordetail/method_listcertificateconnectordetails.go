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

type ListCertificateConnectorDetailsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CertificateConnectorDetails
}

type ListCertificateConnectorDetailsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CertificateConnectorDetails
}

type ListCertificateConnectorDetailsOperationOptions struct {
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

func DefaultListCertificateConnectorDetailsOperationOptions() ListCertificateConnectorDetailsOperationOptions {
	return ListCertificateConnectorDetailsOperationOptions{}
}

func (o ListCertificateConnectorDetailsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCertificateConnectorDetailsOperationOptions) ToOData() *odata.Query {
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

func (o ListCertificateConnectorDetailsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListCertificateConnectorDetailsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCertificateConnectorDetailsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCertificateConnectorDetails - Get certificateConnectorDetails from deviceManagement. Collection of certificate
// connector details, each associated with a corresponding Intune Certificate Connector.
func (c CertificateConnectorDetailClient) ListCertificateConnectorDetails(ctx context.Context, options ListCertificateConnectorDetailsOperationOptions) (result ListCertificateConnectorDetailsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCertificateConnectorDetailsCustomPager{},
		Path:          "/deviceManagement/certificateConnectorDetails",
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
		Values *[]beta.CertificateConnectorDetails `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCertificateConnectorDetailsComplete retrieves all the results into a single object
func (c CertificateConnectorDetailClient) ListCertificateConnectorDetailsComplete(ctx context.Context, options ListCertificateConnectorDetailsOperationOptions) (ListCertificateConnectorDetailsCompleteResult, error) {
	return c.ListCertificateConnectorDetailsCompleteMatchingPredicate(ctx, options, CertificateConnectorDetailsOperationPredicate{})
}

// ListCertificateConnectorDetailsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CertificateConnectorDetailClient) ListCertificateConnectorDetailsCompleteMatchingPredicate(ctx context.Context, options ListCertificateConnectorDetailsOperationOptions, predicate CertificateConnectorDetailsOperationPredicate) (result ListCertificateConnectorDetailsCompleteResult, err error) {
	items := make([]beta.CertificateConnectorDetails, 0)

	resp, err := c.ListCertificateConnectorDetails(ctx, options)
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

	result = ListCertificateConnectorDetailsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
