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

type ListReusablePolicySettingReferencingConfigurationPoliciesOperationOptions struct {
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

func DefaultListReusablePolicySettingReferencingConfigurationPoliciesOperationOptions() ListReusablePolicySettingReferencingConfigurationPoliciesOperationOptions {
	return ListReusablePolicySettingReferencingConfigurationPoliciesOperationOptions{}
}

func (o ListReusablePolicySettingReferencingConfigurationPoliciesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListReusablePolicySettingReferencingConfigurationPoliciesOperationOptions) ToOData() *odata.Query {
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

func (o ListReusablePolicySettingReferencingConfigurationPoliciesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
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

// ListReusablePolicySettingReferencingConfigurationPolicies - Get referencingConfigurationPolicies from
// deviceManagement. configuration policies referencing the current reusable setting. This property is read-only.
func (c ReusablePolicySettingReferencingConfigurationPolicyClient) ListReusablePolicySettingReferencingConfigurationPolicies(ctx context.Context, id beta.DeviceManagementReusablePolicySettingId, options ListReusablePolicySettingReferencingConfigurationPoliciesOperationOptions) (result ListReusablePolicySettingReferencingConfigurationPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListReusablePolicySettingReferencingConfigurationPoliciesCustomPager{},
		Path:          fmt.Sprintf("%s/referencingConfigurationPolicies", id.ID()),
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
		Values *[]beta.DeviceManagementConfigurationPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListReusablePolicySettingReferencingConfigurationPoliciesComplete retrieves all the results into a single object
func (c ReusablePolicySettingReferencingConfigurationPolicyClient) ListReusablePolicySettingReferencingConfigurationPoliciesComplete(ctx context.Context, id beta.DeviceManagementReusablePolicySettingId, options ListReusablePolicySettingReferencingConfigurationPoliciesOperationOptions) (ListReusablePolicySettingReferencingConfigurationPoliciesCompleteResult, error) {
	return c.ListReusablePolicySettingReferencingConfigurationPoliciesCompleteMatchingPredicate(ctx, id, options, DeviceManagementConfigurationPolicyOperationPredicate{})
}

// ListReusablePolicySettingReferencingConfigurationPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ReusablePolicySettingReferencingConfigurationPolicyClient) ListReusablePolicySettingReferencingConfigurationPoliciesCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementReusablePolicySettingId, options ListReusablePolicySettingReferencingConfigurationPoliciesOperationOptions, predicate DeviceManagementConfigurationPolicyOperationPredicate) (result ListReusablePolicySettingReferencingConfigurationPoliciesCompleteResult, err error) {
	items := make([]beta.DeviceManagementConfigurationPolicy, 0)

	resp, err := c.ListReusablePolicySettingReferencingConfigurationPolicies(ctx, id, options)
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
