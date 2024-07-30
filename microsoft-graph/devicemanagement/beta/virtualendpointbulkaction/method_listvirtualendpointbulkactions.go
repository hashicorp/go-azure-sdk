package virtualendpointbulkaction

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

type ListVirtualEndpointBulkActionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CloudPCBulkAction
}

type ListVirtualEndpointBulkActionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CloudPCBulkAction
}

type ListVirtualEndpointBulkActionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListVirtualEndpointBulkActionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListVirtualEndpointBulkActions ...
func (c VirtualEndpointBulkActionClient) ListVirtualEndpointBulkActions(ctx context.Context) (result ListVirtualEndpointBulkActionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListVirtualEndpointBulkActionsCustomPager{},
		Path:       "/deviceManagement/virtualEndpoint/bulkActions",
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
		Values *[]beta.CloudPCBulkAction `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListVirtualEndpointBulkActionsComplete retrieves all the results into a single object
func (c VirtualEndpointBulkActionClient) ListVirtualEndpointBulkActionsComplete(ctx context.Context) (ListVirtualEndpointBulkActionsCompleteResult, error) {
	return c.ListVirtualEndpointBulkActionsCompleteMatchingPredicate(ctx, CloudPCBulkActionOperationPredicate{})
}

// ListVirtualEndpointBulkActionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VirtualEndpointBulkActionClient) ListVirtualEndpointBulkActionsCompleteMatchingPredicate(ctx context.Context, predicate CloudPCBulkActionOperationPredicate) (result ListVirtualEndpointBulkActionsCompleteResult, err error) {
	items := make([]beta.CloudPCBulkAction, 0)

	resp, err := c.ListVirtualEndpointBulkActions(ctx)
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

	result = ListVirtualEndpointBulkActionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
