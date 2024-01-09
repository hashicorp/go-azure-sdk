package networkclouds

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

type KubernetesClustersListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]KubernetesCluster
}

type KubernetesClustersListByResourceGroupCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []KubernetesCluster
}

// KubernetesClustersListByResourceGroup ...
func (c NetworkcloudsClient) KubernetesClustersListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (result KubernetesClustersListByResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.NetworkCloud/kubernetesClusters", id.ID()),
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
		Values *[]KubernetesCluster `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// KubernetesClustersListByResourceGroupComplete retrieves all the results into a single object
func (c NetworkcloudsClient) KubernetesClustersListByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId) (KubernetesClustersListByResourceGroupCompleteResult, error) {
	return c.KubernetesClustersListByResourceGroupCompleteMatchingPredicate(ctx, id, KubernetesClusterOperationPredicate{})
}

// KubernetesClustersListByResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) KubernetesClustersListByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate KubernetesClusterOperationPredicate) (result KubernetesClustersListByResourceGroupCompleteResult, err error) {
	items := make([]KubernetesCluster, 0)

	resp, err := c.KubernetesClustersListByResourceGroup(ctx, id)
	if err != nil {
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

	result = KubernetesClustersListByResourceGroupCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
