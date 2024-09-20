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

type ListVirtualEndpointCloudPCBulkResizesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CloudPCRemoteActionResult
}

type ListVirtualEndpointCloudPCBulkResizesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CloudPCRemoteActionResult
}

type ListVirtualEndpointCloudPCBulkResizesOperationOptions struct {
	Metadata *odata.Metadata
	Skip     *int64
	Top      *int64
}

func DefaultListVirtualEndpointCloudPCBulkResizesOperationOptions() ListVirtualEndpointCloudPCBulkResizesOperationOptions {
	return ListVirtualEndpointCloudPCBulkResizesOperationOptions{}
}

func (o ListVirtualEndpointCloudPCBulkResizesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListVirtualEndpointCloudPCBulkResizesOperationOptions) ToOData() *odata.Query {
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

func (o ListVirtualEndpointCloudPCBulkResizesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListVirtualEndpointCloudPCBulkResizesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListVirtualEndpointCloudPCBulkResizesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListVirtualEndpointCloudPCBulkResizes - Invoke action bulkResize. Perform a bulk resize action to resize a group of
// cloudPCs that successfully pass validation. If any devices can't be resized, those devices indicate 'resize failed'.
// The remaining devices are provisioned for the resize process.
func (c VirtualEndpointCloudPCClient) ListVirtualEndpointCloudPCBulkResizes(ctx context.Context, input ListVirtualEndpointCloudPCBulkResizesRequest, options ListVirtualEndpointCloudPCBulkResizesOperationOptions) (result ListVirtualEndpointCloudPCBulkResizesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListVirtualEndpointCloudPCBulkResizesCustomPager{},
		Path:          "/deviceManagement/virtualEndpoint/cloudPCs/bulkResize",
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
		Values *[]beta.CloudPCRemoteActionResult `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListVirtualEndpointCloudPCBulkResizesComplete retrieves all the results into a single object
func (c VirtualEndpointCloudPCClient) ListVirtualEndpointCloudPCBulkResizesComplete(ctx context.Context, input ListVirtualEndpointCloudPCBulkResizesRequest, options ListVirtualEndpointCloudPCBulkResizesOperationOptions) (ListVirtualEndpointCloudPCBulkResizesCompleteResult, error) {
	return c.ListVirtualEndpointCloudPCBulkResizesCompleteMatchingPredicate(ctx, input, options, CloudPCRemoteActionResultOperationPredicate{})
}

// ListVirtualEndpointCloudPCBulkResizesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VirtualEndpointCloudPCClient) ListVirtualEndpointCloudPCBulkResizesCompleteMatchingPredicate(ctx context.Context, input ListVirtualEndpointCloudPCBulkResizesRequest, options ListVirtualEndpointCloudPCBulkResizesOperationOptions, predicate CloudPCRemoteActionResultOperationPredicate) (result ListVirtualEndpointCloudPCBulkResizesCompleteResult, err error) {
	items := make([]beta.CloudPCRemoteActionResult, 0)

	resp, err := c.ListVirtualEndpointCloudPCBulkResizes(ctx, input, options)
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

	result = ListVirtualEndpointCloudPCBulkResizesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
