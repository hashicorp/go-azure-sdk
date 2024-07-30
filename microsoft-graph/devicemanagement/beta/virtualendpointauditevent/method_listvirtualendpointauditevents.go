package virtualendpointauditevent

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

type ListVirtualEndpointAuditEventsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CloudPCAuditEvent
}

type ListVirtualEndpointAuditEventsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CloudPCAuditEvent
}

type ListVirtualEndpointAuditEventsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListVirtualEndpointAuditEventsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListVirtualEndpointAuditEvents ...
func (c VirtualEndpointAuditEventClient) ListVirtualEndpointAuditEvents(ctx context.Context) (result ListVirtualEndpointAuditEventsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListVirtualEndpointAuditEventsCustomPager{},
		Path:       "/deviceManagement/virtualEndpoint/auditEvents",
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
		Values *[]beta.CloudPCAuditEvent `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListVirtualEndpointAuditEventsComplete retrieves all the results into a single object
func (c VirtualEndpointAuditEventClient) ListVirtualEndpointAuditEventsComplete(ctx context.Context) (ListVirtualEndpointAuditEventsCompleteResult, error) {
	return c.ListVirtualEndpointAuditEventsCompleteMatchingPredicate(ctx, CloudPCAuditEventOperationPredicate{})
}

// ListVirtualEndpointAuditEventsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VirtualEndpointAuditEventClient) ListVirtualEndpointAuditEventsCompleteMatchingPredicate(ctx context.Context, predicate CloudPCAuditEventOperationPredicate) (result ListVirtualEndpointAuditEventsCompleteResult, err error) {
	items := make([]beta.CloudPCAuditEvent, 0)

	resp, err := c.ListVirtualEndpointAuditEvents(ctx)
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

	result = ListVirtualEndpointAuditEventsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
