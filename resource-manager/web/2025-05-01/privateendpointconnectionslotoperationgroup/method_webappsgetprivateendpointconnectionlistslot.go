package privateendpointconnectionslotoperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsGetPrivateEndpointConnectionListSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]RemotePrivateEndpointConnectionARMResource
}

type WebAppsGetPrivateEndpointConnectionListSlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []RemotePrivateEndpointConnectionARMResource
}

type WebAppsGetPrivateEndpointConnectionListSlotCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsGetPrivateEndpointConnectionListSlotCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsGetPrivateEndpointConnectionListSlot ...
func (c PrivateEndpointConnectionSlotOperationGroupClient) WebAppsGetPrivateEndpointConnectionListSlot(ctx context.Context, id SlotId) (result WebAppsGetPrivateEndpointConnectionListSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsGetPrivateEndpointConnectionListSlotCustomPager{},
		Path:       fmt.Sprintf("%s/privateEndpointConnections", id.ID()),
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
		Values *[]RemotePrivateEndpointConnectionARMResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WebAppsGetPrivateEndpointConnectionListSlotComplete retrieves all the results into a single object
func (c PrivateEndpointConnectionSlotOperationGroupClient) WebAppsGetPrivateEndpointConnectionListSlotComplete(ctx context.Context, id SlotId) (WebAppsGetPrivateEndpointConnectionListSlotCompleteResult, error) {
	return c.WebAppsGetPrivateEndpointConnectionListSlotCompleteMatchingPredicate(ctx, id, RemotePrivateEndpointConnectionARMResourceOperationPredicate{})
}

// WebAppsGetPrivateEndpointConnectionListSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivateEndpointConnectionSlotOperationGroupClient) WebAppsGetPrivateEndpointConnectionListSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate RemotePrivateEndpointConnectionARMResourceOperationPredicate) (result WebAppsGetPrivateEndpointConnectionListSlotCompleteResult, err error) {
	items := make([]RemotePrivateEndpointConnectionARMResource, 0)

	resp, err := c.WebAppsGetPrivateEndpointConnectionListSlot(ctx, id)
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

	result = WebAppsGetPrivateEndpointConnectionListSlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
