package virtualendpointfrontlineserviceplan

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

type ListVirtualEndpointFrontLineServicePlansOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CloudPCFrontLineServicePlan
}

type ListVirtualEndpointFrontLineServicePlansCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CloudPCFrontLineServicePlan
}

type ListVirtualEndpointFrontLineServicePlansCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListVirtualEndpointFrontLineServicePlansCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListVirtualEndpointFrontLineServicePlans ...
func (c VirtualEndpointFrontLineServicePlanClient) ListVirtualEndpointFrontLineServicePlans(ctx context.Context) (result ListVirtualEndpointFrontLineServicePlansOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListVirtualEndpointFrontLineServicePlansCustomPager{},
		Path:       "/deviceManagement/virtualEndpoint/frontLineServicePlans",
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
		Values *[]beta.CloudPCFrontLineServicePlan `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListVirtualEndpointFrontLineServicePlansComplete retrieves all the results into a single object
func (c VirtualEndpointFrontLineServicePlanClient) ListVirtualEndpointFrontLineServicePlansComplete(ctx context.Context) (ListVirtualEndpointFrontLineServicePlansCompleteResult, error) {
	return c.ListVirtualEndpointFrontLineServicePlansCompleteMatchingPredicate(ctx, CloudPCFrontLineServicePlanOperationPredicate{})
}

// ListVirtualEndpointFrontLineServicePlansCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VirtualEndpointFrontLineServicePlanClient) ListVirtualEndpointFrontLineServicePlansCompleteMatchingPredicate(ctx context.Context, predicate CloudPCFrontLineServicePlanOperationPredicate) (result ListVirtualEndpointFrontLineServicePlansCompleteResult, err error) {
	items := make([]beta.CloudPCFrontLineServicePlan, 0)

	resp, err := c.ListVirtualEndpointFrontLineServicePlans(ctx)
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

	result = ListVirtualEndpointFrontLineServicePlansCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
