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

type KubernetesClustersListBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]KubernetesCluster
}

type KubernetesClustersListBySubscriptionCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []KubernetesCluster
}

// KubernetesClustersListBySubscription ...
func (c NetworkcloudsClient) KubernetesClustersListBySubscription(ctx context.Context, id commonids.SubscriptionId) (result KubernetesClustersListBySubscriptionOperationResponse, err error) {
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

// KubernetesClustersListBySubscriptionComplete retrieves all the results into a single object
func (c NetworkcloudsClient) KubernetesClustersListBySubscriptionComplete(ctx context.Context, id commonids.SubscriptionId) (KubernetesClustersListBySubscriptionCompleteResult, error) {
	return c.KubernetesClustersListBySubscriptionCompleteMatchingPredicate(ctx, id, KubernetesClusterOperationPredicate{})
}

// KubernetesClustersListBySubscriptionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) KubernetesClustersListBySubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate KubernetesClusterOperationPredicate) (result KubernetesClustersListBySubscriptionCompleteResult, err error) {
	items := make([]KubernetesCluster, 0)

	resp, err := c.KubernetesClustersListBySubscription(ctx, id)
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

	result = KubernetesClustersListBySubscriptionCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
