package certificates

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Certificate
}

type CertificateListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Certificate
}

type CertificateListOperationOptions struct {
	ClientRequestId       *string
	Filter                *string
	Maxresults            *int64
	OcpDate               *string
	ReturnClientRequestId *bool
	Select                *string
	Timeout               *int64
}

func DefaultCertificateListOperationOptions() CertificateListOperationOptions {
	return CertificateListOperationOptions{}
}

func (o CertificateListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.ClientRequestId != nil {
		out.Append("client-request-id", fmt.Sprintf("%v", *o.ClientRequestId))
	}
	if o.OcpDate != nil {
		out.Append("ocp-date", fmt.Sprintf("%v", *o.OcpDate))
	}
	if o.ReturnClientRequestId != nil {
		out.Append("return-client-request-id", fmt.Sprintf("%v", *o.ReturnClientRequestId))
	}
	return &out
}

func (o CertificateListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o CertificateListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Maxresults != nil {
		out.Append("maxresults", fmt.Sprintf("%v", *o.Maxresults))
	}
	if o.Select != nil {
		out.Append("$select", fmt.Sprintf("%v", *o.Select))
	}
	if o.Timeout != nil {
		out.Append("timeout", fmt.Sprintf("%v", *o.Timeout))
	}
	return &out
}

type CertificateListCustomPager struct {
	NextLink *odata.Link `json:"odata.nextLink"`
}

func (p *CertificateListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// CertificateList ...
func (c CertificatesClient) CertificateList(ctx context.Context, options CertificateListOperationOptions) (result CertificateListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &CertificateListCustomPager{},
		Path:          "/certificates",
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
		Values *[]Certificate `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// CertificateListComplete retrieves all the results into a single object
func (c CertificatesClient) CertificateListComplete(ctx context.Context, options CertificateListOperationOptions) (CertificateListCompleteResult, error) {
	return c.CertificateListCompleteMatchingPredicate(ctx, options, CertificateOperationPredicate{})
}

// CertificateListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CertificatesClient) CertificateListCompleteMatchingPredicate(ctx context.Context, options CertificateListOperationOptions, predicate CertificateOperationPredicate) (result CertificateListCompleteResult, err error) {
	items := make([]Certificate, 0)

	resp, err := c.CertificateList(ctx, options)
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

	result = CertificateListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
