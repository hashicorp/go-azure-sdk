package deploymentstacks

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

type ListAtResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DeploymentStack
}

type ListAtResourceGroupCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DeploymentStack
}

// ListAtResourceGroup ...
func (c DeploymentStacksClient) ListAtResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (result ListAtResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.Resources/deploymentStacks", id.ID()),
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
		Values *[]DeploymentStack `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAtResourceGroupComplete retrieves all the results into a single object
func (c DeploymentStacksClient) ListAtResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId) (ListAtResourceGroupCompleteResult, error) {
	return c.ListAtResourceGroupCompleteMatchingPredicate(ctx, id, DeploymentStackOperationPredicate{})
}

// ListAtResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeploymentStacksClient) ListAtResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate DeploymentStackOperationPredicate) (result ListAtResourceGroupCompleteResult, err error) {
	items := make([]DeploymentStack, 0)

	resp, err := c.ListAtResourceGroup(ctx, id)
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

	result = ListAtResourceGroupCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
