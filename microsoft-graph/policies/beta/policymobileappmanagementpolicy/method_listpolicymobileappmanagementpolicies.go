package policymobileappmanagementpolicy

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

type ListPolicyMobileAppManagementPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.MobilityManagementPolicyCollectionResponse
}

type ListPolicyMobileAppManagementPoliciesCompleteResult struct {
	Items []models.MobilityManagementPolicyCollectionResponse
}

// ListPolicyMobileAppManagementPolicies ...
func (c PolicyMobileAppManagementPolicyClient) ListPolicyMobileAppManagementPolicies(ctx context.Context) (result ListPolicyMobileAppManagementPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/policies/mobileAppManagementPolicies",
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
		Values *[]models.MobilityManagementPolicyCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPolicyMobileAppManagementPoliciesComplete retrieves all the results into a single object
func (c PolicyMobileAppManagementPolicyClient) ListPolicyMobileAppManagementPoliciesComplete(ctx context.Context) (ListPolicyMobileAppManagementPoliciesCompleteResult, error) {
	return c.ListPolicyMobileAppManagementPoliciesCompleteMatchingPredicate(ctx, models.MobilityManagementPolicyCollectionResponseOperationPredicate{})
}

// ListPolicyMobileAppManagementPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyMobileAppManagementPolicyClient) ListPolicyMobileAppManagementPoliciesCompleteMatchingPredicate(ctx context.Context, predicate models.MobilityManagementPolicyCollectionResponseOperationPredicate) (result ListPolicyMobileAppManagementPoliciesCompleteResult, err error) {
	items := make([]models.MobilityManagementPolicyCollectionResponse, 0)

	resp, err := c.ListPolicyMobileAppManagementPolicies(ctx)
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

	result = ListPolicyMobileAppManagementPoliciesCompleteResult{
		Items: items,
	}
	return
}
