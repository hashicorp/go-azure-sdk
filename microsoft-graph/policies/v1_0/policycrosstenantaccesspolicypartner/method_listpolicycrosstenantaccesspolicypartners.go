package policycrosstenantaccesspolicypartner

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

type ListPolicyCrossTenantAccessPolicyPartnersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.CrossTenantAccessPolicyConfigurationPartnerCollectionResponse
}

type ListPolicyCrossTenantAccessPolicyPartnersCompleteResult struct {
	Items []models.CrossTenantAccessPolicyConfigurationPartnerCollectionResponse
}

// ListPolicyCrossTenantAccessPolicyPartners ...
func (c PolicyCrossTenantAccessPolicyPartnerClient) ListPolicyCrossTenantAccessPolicyPartners(ctx context.Context) (result ListPolicyCrossTenantAccessPolicyPartnersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/policies/crossTenantAccessPolicy/partners",
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
		Values *[]models.CrossTenantAccessPolicyConfigurationPartnerCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPolicyCrossTenantAccessPolicyPartnersComplete retrieves all the results into a single object
func (c PolicyCrossTenantAccessPolicyPartnerClient) ListPolicyCrossTenantAccessPolicyPartnersComplete(ctx context.Context) (ListPolicyCrossTenantAccessPolicyPartnersCompleteResult, error) {
	return c.ListPolicyCrossTenantAccessPolicyPartnersCompleteMatchingPredicate(ctx, models.CrossTenantAccessPolicyConfigurationPartnerCollectionResponseOperationPredicate{})
}

// ListPolicyCrossTenantAccessPolicyPartnersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyCrossTenantAccessPolicyPartnerClient) ListPolicyCrossTenantAccessPolicyPartnersCompleteMatchingPredicate(ctx context.Context, predicate models.CrossTenantAccessPolicyConfigurationPartnerCollectionResponseOperationPredicate) (result ListPolicyCrossTenantAccessPolicyPartnersCompleteResult, err error) {
	items := make([]models.CrossTenantAccessPolicyConfigurationPartnerCollectionResponse, 0)

	resp, err := c.ListPolicyCrossTenantAccessPolicyPartners(ctx)
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

	result = ListPolicyCrossTenantAccessPolicyPartnersCompleteResult{
		Items: items,
	}
	return
}
