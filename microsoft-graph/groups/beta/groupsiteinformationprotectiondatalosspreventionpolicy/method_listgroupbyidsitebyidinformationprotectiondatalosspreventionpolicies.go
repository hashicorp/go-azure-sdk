package groupsiteinformationprotectiondatalosspreventionpolicy

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

type ListGroupByIdSiteByIdInformationProtectionDataLossPreventionPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DataLossPreventionPolicyCollectionResponse
}

type ListGroupByIdSiteByIdInformationProtectionDataLossPreventionPoliciesCompleteResult struct {
	Items []models.DataLossPreventionPolicyCollectionResponse
}

// ListGroupByIdSiteByIdInformationProtectionDataLossPreventionPolicies ...
func (c GroupSiteInformationProtectionDataLossPreventionPolicyClient) ListGroupByIdSiteByIdInformationProtectionDataLossPreventionPolicies(ctx context.Context, id GroupSiteId) (result ListGroupByIdSiteByIdInformationProtectionDataLossPreventionPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/informationProtection/dataLossPreventionPolicies", id.ID()),
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

// ListGroupByIdSiteByIdInformationProtectionDataLossPreventionPoliciesComplete retrieves all the results into a single object
func (c GroupSiteInformationProtectionDataLossPreventionPolicyClient) ListGroupByIdSiteByIdInformationProtectionDataLossPreventionPoliciesComplete(ctx context.Context, id GroupSiteId) (ListGroupByIdSiteByIdInformationProtectionDataLossPreventionPoliciesCompleteResult, error) {
	return c.ListGroupByIdSiteByIdInformationProtectionDataLossPreventionPoliciesCompleteMatchingPredicate(ctx, id, models.DataLossPreventionPolicyCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdInformationProtectionDataLossPreventionPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteInformationProtectionDataLossPreventionPolicyClient) ListGroupByIdSiteByIdInformationProtectionDataLossPreventionPoliciesCompleteMatchingPredicate(ctx context.Context, id GroupSiteId, predicate models.DataLossPreventionPolicyCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdInformationProtectionDataLossPreventionPoliciesCompleteResult, err error) {
	items := make([]models.DataLossPreventionPolicyCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdInformationProtectionDataLossPreventionPolicies(ctx, id)
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

	result = ListGroupByIdSiteByIdInformationProtectionDataLossPreventionPoliciesCompleteResult{
		Items: items,
	}
	return
}
