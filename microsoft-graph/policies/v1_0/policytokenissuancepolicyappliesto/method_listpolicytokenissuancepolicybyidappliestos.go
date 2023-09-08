package policytokenissuancepolicyappliesto

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListPolicyTokenIssuancePolicyByIdAppliesTosOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObjectCollectionResponse
}

type ListPolicyTokenIssuancePolicyByIdAppliesTosCompleteResult struct {
	Items []models.DirectoryObjectCollectionResponse
}

// ListPolicyTokenIssuancePolicyByIdAppliesTos ...
func (c PolicyTokenIssuancePolicyAppliesToClient) ListPolicyTokenIssuancePolicyByIdAppliesTos(ctx context.Context, id PolicyTokenIssuancePolicyId) (result ListPolicyTokenIssuancePolicyByIdAppliesTosOperationResponse, err error) {
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

// ListPolicyTokenIssuancePolicyByIdAppliesTosComplete retrieves all the results into a single object
func (c PolicyTokenIssuancePolicyAppliesToClient) ListPolicyTokenIssuancePolicyByIdAppliesTosComplete(ctx context.Context, id PolicyTokenIssuancePolicyId) (ListPolicyTokenIssuancePolicyByIdAppliesTosCompleteResult, error) {
	return c.ListPolicyTokenIssuancePolicyByIdAppliesTosCompleteMatchingPredicate(ctx, id, models.DirectoryObjectCollectionResponseOperationPredicate{})
}

// ListPolicyTokenIssuancePolicyByIdAppliesTosCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyTokenIssuancePolicyAppliesToClient) ListPolicyTokenIssuancePolicyByIdAppliesTosCompleteMatchingPredicate(ctx context.Context, id PolicyTokenIssuancePolicyId, predicate models.DirectoryObjectCollectionResponseOperationPredicate) (result ListPolicyTokenIssuancePolicyByIdAppliesTosCompleteResult, err error) {
	items := make([]models.DirectoryObjectCollectionResponse, 0)

	resp, err := c.ListPolicyTokenIssuancePolicyByIdAppliesTos(ctx, id)
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

	result = ListPolicyTokenIssuancePolicyByIdAppliesTosCompleteResult{
		Items: items,
	}
	return
}
