package virtualendpointserviceplan

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

type ListVirtualEndpointServicePlansOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CloudPCServicePlan
}

type ListVirtualEndpointServicePlansCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CloudPCServicePlan
}

type ListVirtualEndpointServicePlansCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListVirtualEndpointServicePlansCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListVirtualEndpointServicePlans ...
func (c VirtualEndpointServicePlanClient) ListVirtualEndpointServicePlans(ctx context.Context) (result ListVirtualEndpointServicePlansOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListVirtualEndpointServicePlansCustomPager{},
		Path:       "/deviceManagement/virtualEndpoint/servicePlans",
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
		Values *[]beta.CloudPCServicePlan `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListVirtualEndpointServicePlansComplete retrieves all the results into a single object
func (c VirtualEndpointServicePlanClient) ListVirtualEndpointServicePlansComplete(ctx context.Context) (ListVirtualEndpointServicePlansCompleteResult, error) {
	return c.ListVirtualEndpointServicePlansCompleteMatchingPredicate(ctx, CloudPCServicePlanOperationPredicate{})
}

// ListVirtualEndpointServicePlansCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VirtualEndpointServicePlanClient) ListVirtualEndpointServicePlansCompleteMatchingPredicate(ctx context.Context, predicate CloudPCServicePlanOperationPredicate) (result ListVirtualEndpointServicePlansCompleteResult, err error) {
	items := make([]beta.CloudPCServicePlan, 0)

	resp, err := c.ListVirtualEndpointServicePlans(ctx)
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

	result = ListVirtualEndpointServicePlansCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
