package reusablepolicysettingreferencingconfigurationpolicy

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignReusablePolicySettingReferencingConfigurationPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementConfigurationPolicyAssignment
}

type AssignReusablePolicySettingReferencingConfigurationPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementConfigurationPolicyAssignment
}

type AssignReusablePolicySettingReferencingConfigurationPoliciesOperationOptions struct {
	Metadata *odata.Metadata
	Skip     *int64
	Top      *int64
}

func DefaultAssignReusablePolicySettingReferencingConfigurationPoliciesOperationOptions() AssignReusablePolicySettingReferencingConfigurationPoliciesOperationOptions {
	return AssignReusablePolicySettingReferencingConfigurationPoliciesOperationOptions{}
}

func (o AssignReusablePolicySettingReferencingConfigurationPoliciesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AssignReusablePolicySettingReferencingConfigurationPoliciesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o AssignReusablePolicySettingReferencingConfigurationPoliciesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type AssignReusablePolicySettingReferencingConfigurationPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *AssignReusablePolicySettingReferencingConfigurationPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AssignReusablePolicySettingReferencingConfigurationPolicies - Invoke action assign
func (c ReusablePolicySettingReferencingConfigurationPolicyClient) AssignReusablePolicySettingReferencingConfigurationPolicies(ctx context.Context, id beta.DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId, input AssignReusablePolicySettingReferencingConfigurationPoliciesRequest, options AssignReusablePolicySettingReferencingConfigurationPoliciesOperationOptions) (result AssignReusablePolicySettingReferencingConfigurationPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &AssignReusablePolicySettingReferencingConfigurationPoliciesCustomPager{},
		Path:          fmt.Sprintf("%s/assign", id.ID()),
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
		Values *[]beta.DeviceManagementConfigurationPolicyAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AssignReusablePolicySettingReferencingConfigurationPoliciesComplete retrieves all the results into a single object
func (c ReusablePolicySettingReferencingConfigurationPolicyClient) AssignReusablePolicySettingReferencingConfigurationPoliciesComplete(ctx context.Context, id beta.DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId, input AssignReusablePolicySettingReferencingConfigurationPoliciesRequest, options AssignReusablePolicySettingReferencingConfigurationPoliciesOperationOptions) (AssignReusablePolicySettingReferencingConfigurationPoliciesCompleteResult, error) {
	return c.AssignReusablePolicySettingReferencingConfigurationPoliciesCompleteMatchingPredicate(ctx, id, input, options, DeviceManagementConfigurationPolicyAssignmentOperationPredicate{})
}

// AssignReusablePolicySettingReferencingConfigurationPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ReusablePolicySettingReferencingConfigurationPolicyClient) AssignReusablePolicySettingReferencingConfigurationPoliciesCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId, input AssignReusablePolicySettingReferencingConfigurationPoliciesRequest, options AssignReusablePolicySettingReferencingConfigurationPoliciesOperationOptions, predicate DeviceManagementConfigurationPolicyAssignmentOperationPredicate) (result AssignReusablePolicySettingReferencingConfigurationPoliciesCompleteResult, err error) {
	items := make([]beta.DeviceManagementConfigurationPolicyAssignment, 0)

	resp, err := c.AssignReusablePolicySettingReferencingConfigurationPolicies(ctx, id, input, options)
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

	result = AssignReusablePolicySettingReferencingConfigurationPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
