package virtualendpointdeviceimage

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

type ListVirtualEndpointDeviceImagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CloudPCDeviceImage
}

type ListVirtualEndpointDeviceImagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CloudPCDeviceImage
}

type ListVirtualEndpointDeviceImagesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListVirtualEndpointDeviceImagesOperationOptions() ListVirtualEndpointDeviceImagesOperationOptions {
	return ListVirtualEndpointDeviceImagesOperationOptions{}
}

func (o ListVirtualEndpointDeviceImagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListVirtualEndpointDeviceImagesOperationOptions) ToOData() *odata.Query {
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

func (o ListVirtualEndpointDeviceImagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListVirtualEndpointDeviceImagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListVirtualEndpointDeviceImagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListVirtualEndpointDeviceImages - List deviceImages. List the properties and relationships of the cloudPcDeviceImage
// objects (OS images) uploaded to Cloud PC.
func (c VirtualEndpointDeviceImageClient) ListVirtualEndpointDeviceImages(ctx context.Context, options ListVirtualEndpointDeviceImagesOperationOptions) (result ListVirtualEndpointDeviceImagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListVirtualEndpointDeviceImagesCustomPager{},
		Path:          "/deviceManagement/virtualEndpoint/deviceImages",
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
		Values *[]beta.CloudPCDeviceImage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListVirtualEndpointDeviceImagesComplete retrieves all the results into a single object
func (c VirtualEndpointDeviceImageClient) ListVirtualEndpointDeviceImagesComplete(ctx context.Context, options ListVirtualEndpointDeviceImagesOperationOptions) (ListVirtualEndpointDeviceImagesCompleteResult, error) {
	return c.ListVirtualEndpointDeviceImagesCompleteMatchingPredicate(ctx, options, CloudPCDeviceImageOperationPredicate{})
}

// ListVirtualEndpointDeviceImagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VirtualEndpointDeviceImageClient) ListVirtualEndpointDeviceImagesCompleteMatchingPredicate(ctx context.Context, options ListVirtualEndpointDeviceImagesOperationOptions, predicate CloudPCDeviceImageOperationPredicate) (result ListVirtualEndpointDeviceImagesCompleteResult, err error) {
	items := make([]beta.CloudPCDeviceImage, 0)

	resp, err := c.ListVirtualEndpointDeviceImages(ctx, options)
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

	result = ListVirtualEndpointDeviceImagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
