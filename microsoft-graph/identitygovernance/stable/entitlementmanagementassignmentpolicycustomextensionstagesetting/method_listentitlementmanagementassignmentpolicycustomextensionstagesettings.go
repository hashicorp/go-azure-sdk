package entitlementmanagementassignmentpolicycustomextensionstagesetting

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

type ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.CustomExtensionStageSetting
}

type ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.CustomExtensionStageSetting
}

type ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettings ...
func (c EntitlementManagementAssignmentPolicyCustomExtensionStageSettingClient) ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettings(ctx context.Context, id IdentityGovernanceEntitlementManagementAssignmentPolicyId) (result ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsCustomPager{},
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

// ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsComplete retrieves all the results into a single object
func (c EntitlementManagementAssignmentPolicyCustomExtensionStageSettingClient) ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsComplete(ctx context.Context, id IdentityGovernanceEntitlementManagementAssignmentPolicyId) (ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsCompleteResult, error) {
	return c.ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsCompleteMatchingPredicate(ctx, id, CustomExtensionStageSettingOperationPredicate{})
}

// ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAssignmentPolicyCustomExtensionStageSettingClient) ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceEntitlementManagementAssignmentPolicyId, predicate CustomExtensionStageSettingOperationPredicate) (result ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsCompleteResult, err error) {
	items := make([]stable.CustomExtensionStageSetting, 0)

	resp, err := c.ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettings(ctx, id)
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

	result = ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
