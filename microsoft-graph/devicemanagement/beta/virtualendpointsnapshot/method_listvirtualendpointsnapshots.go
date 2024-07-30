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

type ListVirtualEndpointSnapshotsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListVirtualEndpointSnapshotsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListVirtualEndpointSnapshots ...
func (c VirtualEndpointSnapshotClient) ListVirtualEndpointSnapshots(ctx context.Context) (result ListVirtualEndpointSnapshotsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListVirtualEndpointSnapshotsCustomPager{},
		Path:       "/deviceManagement/virtualEndpoint/snapshots",
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
func (c VirtualEndpointSnapshotClient) ListVirtualEndpointSnapshotsComplete(ctx context.Context) (ListVirtualEndpointSnapshotsCompleteResult, error) {
	return c.ListVirtualEndpointSnapshotsCompleteMatchingPredicate(ctx, CloudPCSnapshotOperationPredicate{})
}

// ListVirtualEndpointSnapshotsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VirtualEndpointSnapshotClient) ListVirtualEndpointSnapshotsCompleteMatchingPredicate(ctx context.Context, predicate CloudPCSnapshotOperationPredicate) (result ListVirtualEndpointSnapshotsCompleteResult, err error) {
	items := make([]beta.CloudPCSnapshot, 0)

	resp, err := c.ListVirtualEndpointSnapshots(ctx)
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
