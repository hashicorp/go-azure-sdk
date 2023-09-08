package policypermissiongrantpolicy

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListPolicyPermissionGrantPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PermissionGrantPolicyCollectionResponse
}

type ListPolicyPermissionGrantPoliciesCompleteResult struct {
	Items []models.PermissionGrantPolicyCollectionResponse
}

// ListPolicyPermissionGrantPolicies ...
func (c PolicyPermissionGrantPolicyClient) ListPolicyPermissionGrantPolicies(ctx context.Context) (result ListPolicyPermissionGrantPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/policies/permissionGrantPolicies",
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
		Values *[]models.PermissionGrantPolicyCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPolicyPermissionGrantPoliciesComplete retrieves all the results into a single object
func (c PolicyPermissionGrantPolicyClient) ListPolicyPermissionGrantPoliciesComplete(ctx context.Context) (ListPolicyPermissionGrantPoliciesCompleteResult, error) {
	return c.ListPolicyPermissionGrantPoliciesCompleteMatchingPredicate(ctx, models.PermissionGrantPolicyCollectionResponseOperationPredicate{})
}

// ListPolicyPermissionGrantPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyPermissionGrantPolicyClient) ListPolicyPermissionGrantPoliciesCompleteMatchingPredicate(ctx context.Context, predicate models.PermissionGrantPolicyCollectionResponseOperationPredicate) (result ListPolicyPermissionGrantPoliciesCompleteResult, err error) {
	items := make([]models.PermissionGrantPolicyCollectionResponse, 0)

	resp, err := c.ListPolicyPermissionGrantPolicies(ctx)
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

	result = ListPolicyPermissionGrantPoliciesCompleteResult{
		Items: items,
	}
	return
}
