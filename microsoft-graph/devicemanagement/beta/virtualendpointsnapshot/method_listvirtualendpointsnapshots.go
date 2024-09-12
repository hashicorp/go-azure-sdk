package virtualendpointsnapshot

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

type ListVirtualEndpointSnapshotsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CloudPCSnapshot
}

type ListVirtualEndpointSnapshotsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CloudPCSnapshot
}

type ListVirtualEndpointSnapshotsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListVirtualEndpointSnapshotsOperationOptions() ListVirtualEndpointSnapshotsOperationOptions {
	return ListVirtualEndpointSnapshotsOperationOptions{}
}

func (o ListVirtualEndpointSnapshotsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListVirtualEndpointSnapshotsOperationOptions) ToOData() *odata.Query {
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

func (o ListVirtualEndpointSnapshotsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListVirtualEndpointSnapshotsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListVirtualEndpointSnapshotsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListVirtualEndpointSnapshots - List snapshots. Get a list of cloudPcSnapshot objects and their properties.
func (c VirtualEndpointSnapshotClient) ListVirtualEndpointSnapshots(ctx context.Context, options ListVirtualEndpointSnapshotsOperationOptions) (result ListVirtualEndpointSnapshotsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListVirtualEndpointSnapshotsCustomPager{},
		Path:          "/deviceManagement/virtualEndpoint/snapshots",
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
		Values *[]beta.CloudPCSnapshot `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListVirtualEndpointSnapshotsComplete retrieves all the results into a single object
func (c VirtualEndpointSnapshotClient) ListVirtualEndpointSnapshotsComplete(ctx context.Context, options ListVirtualEndpointSnapshotsOperationOptions) (ListVirtualEndpointSnapshotsCompleteResult, error) {
	return c.ListVirtualEndpointSnapshotsCompleteMatchingPredicate(ctx, options, CloudPCSnapshotOperationPredicate{})
}

// ListVirtualEndpointSnapshotsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VirtualEndpointSnapshotClient) ListVirtualEndpointSnapshotsCompleteMatchingPredicate(ctx context.Context, options ListVirtualEndpointSnapshotsOperationOptions, predicate CloudPCSnapshotOperationPredicate) (result ListVirtualEndpointSnapshotsCompleteResult, err error) {
	items := make([]beta.CloudPCSnapshot, 0)

	resp, err := c.ListVirtualEndpointSnapshots(ctx, options)
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

	result = ListVirtualEndpointSnapshotsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
