package entitlementmanagementaccesspackageassignmentpolicycustomextensionstagesetting

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.CustomExtensionStageSetting
}

type ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.CustomExtensionStageSetting
}

type ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettings ...
func (c EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingClient) ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettings(ctx context.Context, id IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyId) (result ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsCustomPager{},
		Path:       fmt.Sprintf("%s/customExtensionStageSettings", id.ID()),
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
		Values *[]stable.CustomExtensionStageSetting `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingClient) ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsComplete(ctx context.Context, id IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyId) (ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsCompleteMatchingPredicate(ctx, id, CustomExtensionStageSettingOperationPredicate{})
}

// ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingClient) ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyId, predicate CustomExtensionStageSettingOperationPredicate) (result ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsCompleteResult, err error) {
	items := make([]stable.CustomExtensionStageSetting, 0)

	resp, err := c.ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettings(ctx, id)
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

	result = ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
