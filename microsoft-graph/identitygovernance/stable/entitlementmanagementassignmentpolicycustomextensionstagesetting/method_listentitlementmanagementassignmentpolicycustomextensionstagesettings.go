package entitlementmanagementassignmentpolicycustomextensionstagesetting

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
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

type ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsOperationOptions() ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsOperationOptions {
	return ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsOperationOptions{}
}

func (o ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsOperationOptions) ToOData() *odata.Query {
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
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
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

func (o ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
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

// ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettings - Get customExtensionStageSettings from
// identityGovernance. The collection of stages when to execute one or more custom access package workflow extensions.
// Supports $expand.
func (c EntitlementManagementAssignmentPolicyCustomExtensionStageSettingClient) ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettings(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAssignmentPolicyId, options ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsOperationOptions) (result ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsCustomPager{},
		Path:          fmt.Sprintf("%s/customExtensionStageSettings", id.ID()),
		RetryFunc:     options.RetryFunc,
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
func (c EntitlementManagementAssignmentPolicyCustomExtensionStageSettingClient) ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsComplete(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAssignmentPolicyId, options ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsOperationOptions) (ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsCompleteResult, error) {
	return c.ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsCompleteMatchingPredicate(ctx, id, options, CustomExtensionStageSettingOperationPredicate{})
}

// ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAssignmentPolicyCustomExtensionStageSettingClient) ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsCompleteMatchingPredicate(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAssignmentPolicyId, options ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsOperationOptions, predicate CustomExtensionStageSettingOperationPredicate) (result ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettingsCompleteResult, err error) {
	items := make([]stable.CustomExtensionStageSetting, 0)

	resp, err := c.ListEntitlementManagementAssignmentPolicyCustomExtensionStageSettings(ctx, id, options)
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
