package applications

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ApplicationSummary
}

type ApplicationListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ApplicationSummary
}

type ApplicationListOperationOptions struct {
	ClientRequestId       *string
	Maxresults            *int64
	OcpDate               *string
	ReturnClientRequestId *bool
	Timeout               *int64
}

func DefaultApplicationListOperationOptions() ApplicationListOperationOptions {
	return ApplicationListOperationOptions{}
}

func (o ApplicationListOperationOptions) ToHeaders() *client.Headers {
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

func (o ApplicationListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ApplicationListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Maxresults != nil {
		out.Append("maxresults", fmt.Sprintf("%v", *o.Maxresults))
	}
	if o.Timeout != nil {
		out.Append("timeout", fmt.Sprintf("%v", *o.Timeout))
	}
	return &out
}

type ApplicationListCustomPager struct {
	NextLink *odata.Link `json:"odata.nextLink"`
}

func (p *ApplicationListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ApplicationList ...
func (c ApplicationsClient) ApplicationList(ctx context.Context, options ApplicationListOperationOptions) (result ApplicationListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ApplicationListCustomPager{},
		Path:          "/applications",
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
		Values *[]ApplicationSummary `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ApplicationListComplete retrieves all the results into a single object
func (c ApplicationsClient) ApplicationListComplete(ctx context.Context, options ApplicationListOperationOptions) (ApplicationListCompleteResult, error) {
	return c.ApplicationListCompleteMatchingPredicate(ctx, options, ApplicationSummaryOperationPredicate{})
}

// ApplicationListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ApplicationsClient) ApplicationListCompleteMatchingPredicate(ctx context.Context, options ApplicationListOperationOptions, predicate ApplicationSummaryOperationPredicate) (result ApplicationListCompleteResult, err error) {
	items := make([]ApplicationSummary, 0)

	resp, err := c.ApplicationList(ctx, options)
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

	result = ApplicationListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
