package rolemanagementpolicyeffectiverule

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListRoleManagementPolicyEffectiveRulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleManagementPolicyRule
}

type ListRoleManagementPolicyEffectiveRulesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleManagementPolicyRule
}

type ListRoleManagementPolicyEffectiveRulesOperationOptions struct {
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

func DefaultListRoleManagementPolicyEffectiveRulesOperationOptions() ListRoleManagementPolicyEffectiveRulesOperationOptions {
	return ListRoleManagementPolicyEffectiveRulesOperationOptions{}
}

func (o ListRoleManagementPolicyEffectiveRulesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListRoleManagementPolicyEffectiveRulesOperationOptions) ToOData() *odata.Query {
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

func (o ListRoleManagementPolicyEffectiveRulesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListRoleManagementPolicyEffectiveRulesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListRoleManagementPolicyEffectiveRulesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListRoleManagementPolicyEffectiveRules - List effectiveRules. Get the unifiedRoleManagementPolicyRule resources from
// the effectiveRules navigation property. To retrieve rules for a policy that applies to Azure RBAC, use the Azure REST
// PIM API for role management policies.
func (c RoleManagementPolicyEffectiveRuleClient) ListRoleManagementPolicyEffectiveRules(ctx context.Context, id beta.PolicyRoleManagementPolicyId, options ListRoleManagementPolicyEffectiveRulesOperationOptions) (result ListRoleManagementPolicyEffectiveRulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListRoleManagementPolicyEffectiveRulesCustomPager{},
		Path:          fmt.Sprintf("%s/effectiveRules", id.ID()),
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]beta.UnifiedRoleManagementPolicyRule, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalUnifiedRoleManagementPolicyRuleImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.UnifiedRoleManagementPolicyRule (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListRoleManagementPolicyEffectiveRulesComplete retrieves all the results into a single object
func (c RoleManagementPolicyEffectiveRuleClient) ListRoleManagementPolicyEffectiveRulesComplete(ctx context.Context, id beta.PolicyRoleManagementPolicyId, options ListRoleManagementPolicyEffectiveRulesOperationOptions) (ListRoleManagementPolicyEffectiveRulesCompleteResult, error) {
	return c.ListRoleManagementPolicyEffectiveRulesCompleteMatchingPredicate(ctx, id, options, UnifiedRoleManagementPolicyRuleOperationPredicate{})
}

// ListRoleManagementPolicyEffectiveRulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementPolicyEffectiveRuleClient) ListRoleManagementPolicyEffectiveRulesCompleteMatchingPredicate(ctx context.Context, id beta.PolicyRoleManagementPolicyId, options ListRoleManagementPolicyEffectiveRulesOperationOptions, predicate UnifiedRoleManagementPolicyRuleOperationPredicate) (result ListRoleManagementPolicyEffectiveRulesCompleteResult, err error) {
	items := make([]beta.UnifiedRoleManagementPolicyRule, 0)

	resp, err := c.ListRoleManagementPolicyEffectiveRules(ctx, id, options)
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

	result = ListRoleManagementPolicyEffectiveRulesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
