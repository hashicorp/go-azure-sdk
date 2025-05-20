package cloudhsmclustergetbysubscription

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

type CloudHsmClustersListBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]CloudHsmCluster
}

type CloudHsmClustersListBySubscriptionCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []CloudHsmCluster
}

type CloudHsmClustersListBySubscriptionCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *CloudHsmClustersListBySubscriptionCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// CloudHsmClustersListBySubscription ...
func (c CloudHSMClusterGetBySubscriptionClient) CloudHsmClustersListBySubscription(ctx context.Context, id commonids.SubscriptionId) (result CloudHsmClustersListBySubscriptionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &CloudHsmClustersListBySubscriptionCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.HardwareSecurityModules/cloudHsmClusters", id.ID()),
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
		Values *[]CloudHsmCluster `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// CloudHsmClustersListBySubscriptionComplete retrieves all the results into a single object
func (c CloudHSMClusterGetBySubscriptionClient) CloudHsmClustersListBySubscriptionComplete(ctx context.Context, id commonids.SubscriptionId) (CloudHsmClustersListBySubscriptionCompleteResult, error) {
	return c.CloudHsmClustersListBySubscriptionCompleteMatchingPredicate(ctx, id, CloudHsmClusterOperationPredicate{})
}

// CloudHsmClustersListBySubscriptionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CloudHSMClusterGetBySubscriptionClient) CloudHsmClustersListBySubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate CloudHsmClusterOperationPredicate) (result CloudHsmClustersListBySubscriptionCompleteResult, err error) {
	items := make([]CloudHsmCluster, 0)

	resp, err := c.CloudHsmClustersListBySubscription(ctx, id)
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

	result = CloudHsmClustersListBySubscriptionCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
