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

type ListReusablePolicySettingReferencingConfigurationPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementConfigurationPolicy
}

type ListReusablePolicySettingReferencingConfigurationPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementConfigurationPolicy
}

type ListReusablePolicySettingReferencingConfigurationPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListReusablePolicySettingReferencingConfigurationPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListReusablePolicySettingReferencingConfigurationPolicies ...
func (c ReusablePolicySettingReferencingConfigurationPolicyClient) ListReusablePolicySettingReferencingConfigurationPolicies(ctx context.Context, id DeviceManagementReusablePolicySettingId) (result ListReusablePolicySettingReferencingConfigurationPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListReusablePolicySettingReferencingConfigurationPoliciesCustomPager{},
		Path:       fmt.Sprintf("%s/referencingConfigurationPolicies", id.ID()),
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
		Values *[]beta.DeviceManagementConfigurationPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListReusablePolicySettingReferencingConfigurationPoliciesComplete retrieves all the results into a single object
func (c ReusablePolicySettingReferencingConfigurationPolicyClient) ListReusablePolicySettingReferencingConfigurationPoliciesComplete(ctx context.Context, id DeviceManagementReusablePolicySettingId) (ListReusablePolicySettingReferencingConfigurationPoliciesCompleteResult, error) {
	return c.ListReusablePolicySettingReferencingConfigurationPoliciesCompleteMatchingPredicate(ctx, id, DeviceManagementConfigurationPolicyOperationPredicate{})
}

// ListReusablePolicySettingReferencingConfigurationPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ReusablePolicySettingReferencingConfigurationPolicyClient) ListReusablePolicySettingReferencingConfigurationPoliciesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementReusablePolicySettingId, predicate DeviceManagementConfigurationPolicyOperationPredicate) (result ListReusablePolicySettingReferencingConfigurationPoliciesCompleteResult, err error) {
	items := make([]beta.DeviceManagementConfigurationPolicy, 0)

	resp, err := c.ListReusablePolicySettingReferencingConfigurationPolicies(ctx, id)
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

	result = ListReusablePolicySettingReferencingConfigurationPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
