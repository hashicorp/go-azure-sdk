package remoteprivateendpointconnectionarmresourceoperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsGetPrivateEndpointConnectionListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]RemotePrivateEndpointConnectionARMResource
}

type WebAppsGetPrivateEndpointConnectionListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []RemotePrivateEndpointConnectionARMResource
}

type WebAppsGetPrivateEndpointConnectionListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsGetPrivateEndpointConnectionListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsGetPrivateEndpointConnectionList ...
func (c RemotePrivateEndpointConnectionARMResourceOperationGroupClient) WebAppsGetPrivateEndpointConnectionList(ctx context.Context, id commonids.AppServiceId) (result WebAppsGetPrivateEndpointConnectionListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsGetPrivateEndpointConnectionListCustomPager{},
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

// WebAppsGetPrivateEndpointConnectionListComplete retrieves all the results into a single object
func (c RemotePrivateEndpointConnectionARMResourceOperationGroupClient) WebAppsGetPrivateEndpointConnectionListComplete(ctx context.Context, id commonids.AppServiceId) (WebAppsGetPrivateEndpointConnectionListCompleteResult, error) {
	return c.WebAppsGetPrivateEndpointConnectionListCompleteMatchingPredicate(ctx, id, RemotePrivateEndpointConnectionARMResourceOperationPredicate{})
}

// WebAppsGetPrivateEndpointConnectionListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RemotePrivateEndpointConnectionARMResourceOperationGroupClient) WebAppsGetPrivateEndpointConnectionListCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceId, predicate RemotePrivateEndpointConnectionARMResourceOperationPredicate) (result WebAppsGetPrivateEndpointConnectionListCompleteResult, err error) {
	items := make([]RemotePrivateEndpointConnectionARMResource, 0)

	resp, err := c.WebAppsGetPrivateEndpointConnectionList(ctx, id)
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

	result = WebAppsGetPrivateEndpointConnectionListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
