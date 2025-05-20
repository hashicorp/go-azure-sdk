package cloudhsmclustergetbyresourcegroup

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

type CloudHsmClustersListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]CloudHsmCluster
}

type CloudHsmClustersListByResourceGroupCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []CloudHsmCluster
}

type CloudHsmClustersListByResourceGroupCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *CloudHsmClustersListByResourceGroupCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// CloudHsmClustersListByResourceGroup ...
func (c CloudHSMClusterGetByResourceGroupClient) CloudHsmClustersListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (result CloudHsmClustersListByResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &CloudHsmClustersListByResourceGroupCustomPager{},
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

// CloudHsmClustersListByResourceGroupComplete retrieves all the results into a single object
func (c CloudHSMClusterGetByResourceGroupClient) CloudHsmClustersListByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId) (CloudHsmClustersListByResourceGroupCompleteResult, error) {
	return c.CloudHsmClustersListByResourceGroupCompleteMatchingPredicate(ctx, id, CloudHsmClusterOperationPredicate{})
}

// CloudHsmClustersListByResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CloudHSMClusterGetByResourceGroupClient) CloudHsmClustersListByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate CloudHsmClusterOperationPredicate) (result CloudHsmClustersListByResourceGroupCompleteResult, err error) {
	items := make([]CloudHsmCluster, 0)

	resp, err := c.CloudHsmClustersListByResourceGroup(ctx, id)
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

	result = CloudHsmClustersListByResourceGroupCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
