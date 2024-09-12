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

type ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsOperationOptions() ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsOperationOptions {
	return ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsOperationOptions{}
}

func (o ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
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

// ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettings - Get customExtensionStageSettings
// from identityGovernance. The collection of stages when to execute one or more custom access package workflow
// extensions. Supports $expand.
func (c EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingClient) ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettings(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyId, options ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsOperationOptions) (result ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsCustomPager{},
		Path:          fmt.Sprintf("%s/customExtensionStageSettings", id.ID()),
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
func (c EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingClient) ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsComplete(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyId, options ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsOperationOptions) (ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsCompleteMatchingPredicate(ctx, id, options, CustomExtensionStageSettingOperationPredicate{})
}

// ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingClient) ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsCompleteMatchingPredicate(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyId, options ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsOperationOptions, predicate CustomExtensionStageSettingOperationPredicate) (result ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingsCompleteResult, err error) {
	items := make([]stable.CustomExtensionStageSetting, 0)

	resp, err := c.ListEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettings(ctx, id, options)
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
