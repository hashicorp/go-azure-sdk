package computenodes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComputeNodeExtensionListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]NodeVMExtension
}

type ComputeNodeExtensionListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []NodeVMExtension
}

type ComputeNodeExtensionListOperationOptions struct {
	ClientRequestId       *string
	Maxresults            *int64
	OcpDate               *string
	ReturnClientRequestId *bool
	Select                *string
	Timeout               *int64
}

func DefaultComputeNodeExtensionListOperationOptions() ComputeNodeExtensionListOperationOptions {
	return ComputeNodeExtensionListOperationOptions{}
}

func (o ComputeNodeExtensionListOperationOptions) ToHeaders() *client.Headers {
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

func (o ComputeNodeExtensionListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ComputeNodeExtensionListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
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

type ComputeNodeExtensionListCustomPager struct {
	NextLink *odata.Link `json:"odata.nextLink"`
}

func (p *ComputeNodeExtensionListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ComputeNodeExtensionList ...
func (c ComputeNodesClient) ComputeNodeExtensionList(ctx context.Context, id NodeId, options ComputeNodeExtensionListOperationOptions) (result ComputeNodeExtensionListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ComputeNodeExtensionListCustomPager{},
		Path:          fmt.Sprintf("%s/extensions", id.Path()),
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
		Values *[]NodeVMExtension `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ComputeNodeExtensionListComplete retrieves all the results into a single object
func (c ComputeNodesClient) ComputeNodeExtensionListComplete(ctx context.Context, id NodeId, options ComputeNodeExtensionListOperationOptions) (ComputeNodeExtensionListCompleteResult, error) {
	return c.ComputeNodeExtensionListCompleteMatchingPredicate(ctx, id, options, NodeVMExtensionOperationPredicate{})
}

// ComputeNodeExtensionListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ComputeNodesClient) ComputeNodeExtensionListCompleteMatchingPredicate(ctx context.Context, id NodeId, options ComputeNodeExtensionListOperationOptions, predicate NodeVMExtensionOperationPredicate) (result ComputeNodeExtensionListCompleteResult, err error) {
	items := make([]NodeVMExtension, 0)

	resp, err := c.ComputeNodeExtensionList(ctx, id, options)
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

	result = ComputeNodeExtensionListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
