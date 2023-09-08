package applicationappmanagementpolicy

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

type ListApplicationByIdAppManagementPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AppManagementPolicyCollectionResponse
}

type ListApplicationByIdAppManagementPoliciesCompleteResult struct {
	Items []models.AppManagementPolicyCollectionResponse
}

// ListApplicationByIdAppManagementPolicies ...
func (c ApplicationAppManagementPolicyClient) ListApplicationByIdAppManagementPolicies(ctx context.Context, id ApplicationId) (result ListApplicationByIdAppManagementPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/appManagementPolicies", id.ID()),
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
		Values *[]models.AppManagementPolicyCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListApplicationByIdAppManagementPoliciesComplete retrieves all the results into a single object
func (c ApplicationAppManagementPolicyClient) ListApplicationByIdAppManagementPoliciesComplete(ctx context.Context, id ApplicationId) (ListApplicationByIdAppManagementPoliciesCompleteResult, error) {
	return c.ListApplicationByIdAppManagementPoliciesCompleteMatchingPredicate(ctx, id, models.AppManagementPolicyCollectionResponseOperationPredicate{})
}

// ListApplicationByIdAppManagementPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ApplicationAppManagementPolicyClient) ListApplicationByIdAppManagementPoliciesCompleteMatchingPredicate(ctx context.Context, id ApplicationId, predicate models.AppManagementPolicyCollectionResponseOperationPredicate) (result ListApplicationByIdAppManagementPoliciesCompleteResult, err error) {
	items := make([]models.AppManagementPolicyCollectionResponse, 0)

	resp, err := c.ListApplicationByIdAppManagementPolicies(ctx, id)
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

	result = ListApplicationByIdAppManagementPoliciesCompleteResult{
		Items: items,
	}
	return
}
