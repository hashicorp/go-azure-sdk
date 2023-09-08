package meinformationprotectiondatalosspreventionpolicy

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

type ListMeInformationProtectionDataLossPreventionPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DataLossPreventionPolicyCollectionResponse
}

type ListMeInformationProtectionDataLossPreventionPoliciesCompleteResult struct {
	Items []models.DataLossPreventionPolicyCollectionResponse
}

// ListMeInformationProtectionDataLossPreventionPolicies ...
func (c MeInformationProtectionDataLossPreventionPolicyClient) ListMeInformationProtectionDataLossPreventionPolicies(ctx context.Context) (result ListMeInformationProtectionDataLossPreventionPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/informationProtection/dataLossPreventionPolicies",
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
		Values *[]models.DataLossPreventionPolicyCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeInformationProtectionDataLossPreventionPoliciesComplete retrieves all the results into a single object
func (c MeInformationProtectionDataLossPreventionPolicyClient) ListMeInformationProtectionDataLossPreventionPoliciesComplete(ctx context.Context) (ListMeInformationProtectionDataLossPreventionPoliciesCompleteResult, error) {
	return c.ListMeInformationProtectionDataLossPreventionPoliciesCompleteMatchingPredicate(ctx, models.DataLossPreventionPolicyCollectionResponseOperationPredicate{})
}

// ListMeInformationProtectionDataLossPreventionPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeInformationProtectionDataLossPreventionPolicyClient) ListMeInformationProtectionDataLossPreventionPoliciesCompleteMatchingPredicate(ctx context.Context, predicate models.DataLossPreventionPolicyCollectionResponseOperationPredicate) (result ListMeInformationProtectionDataLossPreventionPoliciesCompleteResult, err error) {
	items := make([]models.DataLossPreventionPolicyCollectionResponse, 0)

	resp, err := c.ListMeInformationProtectionDataLossPreventionPolicies(ctx)
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

	result = ListMeInformationProtectionDataLossPreventionPoliciesCompleteResult{
		Items: items,
	}
	return
}
