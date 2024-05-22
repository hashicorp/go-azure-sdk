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

type ListAtManagementGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DeploymentStack
}

type ListAtManagementGroupCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DeploymentStack
}

// ListAtManagementGroup ...
func (c DeploymentStacksClient) ListAtManagementGroup(ctx context.Context, id commonids.ManagementGroupId) (result ListAtManagementGroupOperationResponse, err error) {
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

// ListAtManagementGroupComplete retrieves all the results into a single object
func (c DeploymentStacksClient) ListAtManagementGroupComplete(ctx context.Context, id commonids.ManagementGroupId) (ListAtManagementGroupCompleteResult, error) {
	return c.ListAtManagementGroupCompleteMatchingPredicate(ctx, id, DeploymentStackOperationPredicate{})
}

// ListAtManagementGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeploymentStacksClient) ListAtManagementGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ManagementGroupId, predicate DeploymentStackOperationPredicate) (result ListAtManagementGroupCompleteResult, err error) {
	items := make([]DeploymentStack, 0)

	resp, err := c.ListAtManagementGroup(ctx, id)
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

	result = ListAtManagementGroupCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
