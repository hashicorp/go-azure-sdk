package virtualendpointcloudpc

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

type ListVirtualEndpointCloudPCSOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CloudPC
}

type ListVirtualEndpointCloudPCSCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CloudPC
}

type ListVirtualEndpointCloudPCSOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListVirtualEndpointCloudPCSOperationOptions() ListVirtualEndpointCloudPCSOperationOptions {
	return ListVirtualEndpointCloudPCSOperationOptions{}
}

func (o ListVirtualEndpointCloudPCSOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListVirtualEndpointCloudPCSOperationOptions) ToOData() *odata.Query {
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

func (o ListVirtualEndpointCloudPCSOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListVirtualEndpointCloudPCSCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListVirtualEndpointCloudPCSCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListVirtualEndpointCloudPCS - Get cloudPC. Read the properties and relationships of a specific cloudPC object.
func (c VirtualEndpointCloudPCClient) ListVirtualEndpointCloudPCS(ctx context.Context, options ListVirtualEndpointCloudPCSOperationOptions) (result ListVirtualEndpointCloudPCSOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListVirtualEndpointCloudPCSCustomPager{},
		Path:          "/deviceManagement/virtualEndpoint/cloudPCs",
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
		Values *[]beta.CloudPC `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListVirtualEndpointCloudPCSComplete retrieves all the results into a single object
func (c VirtualEndpointCloudPCClient) ListVirtualEndpointCloudPCSComplete(ctx context.Context, options ListVirtualEndpointCloudPCSOperationOptions) (ListVirtualEndpointCloudPCSCompleteResult, error) {
	return c.ListVirtualEndpointCloudPCSCompleteMatchingPredicate(ctx, options, CloudPCOperationPredicate{})
}

// ListVirtualEndpointCloudPCSCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VirtualEndpointCloudPCClient) ListVirtualEndpointCloudPCSCompleteMatchingPredicate(ctx context.Context, options ListVirtualEndpointCloudPCSOperationOptions, predicate CloudPCOperationPredicate) (result ListVirtualEndpointCloudPCSCompleteResult, err error) {
	items := make([]beta.CloudPC, 0)

	resp, err := c.ListVirtualEndpointCloudPCS(ctx, options)
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

	result = ListVirtualEndpointCloudPCSCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
