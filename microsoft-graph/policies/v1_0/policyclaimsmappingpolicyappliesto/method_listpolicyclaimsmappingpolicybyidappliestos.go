package policyclaimsmappingpolicyappliesto

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

type ListPolicyClaimsMappingPolicyByIdAppliesTosOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObjectCollectionResponse
}

type ListPolicyClaimsMappingPolicyByIdAppliesTosCompleteResult struct {
	Items []models.DirectoryObjectCollectionResponse
}

// ListPolicyClaimsMappingPolicyByIdAppliesTos ...
func (c PolicyClaimsMappingPolicyAppliesToClient) ListPolicyClaimsMappingPolicyByIdAppliesTos(ctx context.Context, id PolicyClaimsMappingPolicyId) (result ListPolicyClaimsMappingPolicyByIdAppliesTosOperationResponse, err error) {
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

// ListPolicyClaimsMappingPolicyByIdAppliesTosComplete retrieves all the results into a single object
func (c PolicyClaimsMappingPolicyAppliesToClient) ListPolicyClaimsMappingPolicyByIdAppliesTosComplete(ctx context.Context, id PolicyClaimsMappingPolicyId) (ListPolicyClaimsMappingPolicyByIdAppliesTosCompleteResult, error) {
	return c.ListPolicyClaimsMappingPolicyByIdAppliesTosCompleteMatchingPredicate(ctx, id, models.DirectoryObjectCollectionResponseOperationPredicate{})
}

// ListPolicyClaimsMappingPolicyByIdAppliesTosCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyClaimsMappingPolicyAppliesToClient) ListPolicyClaimsMappingPolicyByIdAppliesTosCompleteMatchingPredicate(ctx context.Context, id PolicyClaimsMappingPolicyId, predicate models.DirectoryObjectCollectionResponseOperationPredicate) (result ListPolicyClaimsMappingPolicyByIdAppliesTosCompleteResult, err error) {
	items := make([]models.DirectoryObjectCollectionResponse, 0)

	resp, err := c.ListPolicyClaimsMappingPolicyByIdAppliesTos(ctx, id)
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

	result = ListPolicyClaimsMappingPolicyByIdAppliesTosCompleteResult{
		Items: items,
	}
	return
}
