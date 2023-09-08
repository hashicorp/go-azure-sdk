package policyappmanagementpolicyappliesto

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

type ListPolicyAppManagementPolicyByIdAppliesTosOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObjectCollectionResponse
}

type ListPolicyAppManagementPolicyByIdAppliesTosCompleteResult struct {
	Items []models.DirectoryObjectCollectionResponse
}

// ListPolicyAppManagementPolicyByIdAppliesTos ...
func (c PolicyAppManagementPolicyAppliesToClient) ListPolicyAppManagementPolicyByIdAppliesTos(ctx context.Context, id PolicyAppManagementPolicyId) (result ListPolicyAppManagementPolicyByIdAppliesTosOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/appliesTo", id.ID()),
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
		Values *[]models.DirectoryObjectCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPolicyAppManagementPolicyByIdAppliesTosComplete retrieves all the results into a single object
func (c PolicyAppManagementPolicyAppliesToClient) ListPolicyAppManagementPolicyByIdAppliesTosComplete(ctx context.Context, id PolicyAppManagementPolicyId) (ListPolicyAppManagementPolicyByIdAppliesTosCompleteResult, error) {
	return c.ListPolicyAppManagementPolicyByIdAppliesTosCompleteMatchingPredicate(ctx, id, models.DirectoryObjectCollectionResponseOperationPredicate{})
}

// ListPolicyAppManagementPolicyByIdAppliesTosCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyAppManagementPolicyAppliesToClient) ListPolicyAppManagementPolicyByIdAppliesTosCompleteMatchingPredicate(ctx context.Context, id PolicyAppManagementPolicyId, predicate models.DirectoryObjectCollectionResponseOperationPredicate) (result ListPolicyAppManagementPolicyByIdAppliesTosCompleteResult, err error) {
	items := make([]models.DirectoryObjectCollectionResponse, 0)

	resp, err := c.ListPolicyAppManagementPolicyByIdAppliesTos(ctx, id)
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

	result = ListPolicyAppManagementPolicyByIdAppliesTosCompleteResult{
		Items: items,
	}
	return
}
