package rolemanagementpolicyrule

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListRoleManagementPolicyRulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleManagementPolicyRule
}

type ListRoleManagementPolicyRulesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleManagementPolicyRule
}

type ListRoleManagementPolicyRulesOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListRoleManagementPolicyRulesOperationOptions() ListRoleManagementPolicyRulesOperationOptions {
	return ListRoleManagementPolicyRulesOperationOptions{}
}

func (o ListRoleManagementPolicyRulesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListRoleManagementPolicyRulesOperationOptions) ToOData() *odata.Query {
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

func (o ListRoleManagementPolicyRulesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListRoleManagementPolicyRulesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListRoleManagementPolicyRulesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListRoleManagementPolicyRules - List rules (for a role management policy). Get the rules defined for a role
// management policy. The rules are a collection of following types that are derived from the
// unifiedRoleManagementPolicyRule object: + unifiedRoleManagementPolicyApprovalRule +
// unifiedRoleManagementPolicyAuthenticationContextRule + unifiedRoleManagementPolicyEnablementRule +
// unifiedRoleManagementPolicyExpirationRule + unifiedRoleManagementPolicyNotificationRule To retrieve rules for a
// policy that applies to Azure RBAC, use the Azure REST PIM API for role management policies.
func (c RoleManagementPolicyRuleClient) ListRoleManagementPolicyRules(ctx context.Context, id beta.PolicyRoleManagementPolicyId, options ListRoleManagementPolicyRulesOperationOptions) (result ListRoleManagementPolicyRulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListRoleManagementPolicyRulesCustomPager{},
		Path:          fmt.Sprintf("%s/rules", id.ID()),
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

// ListRoleManagementPolicyRulesComplete retrieves all the results into a single object
func (c RoleManagementPolicyRuleClient) ListRoleManagementPolicyRulesComplete(ctx context.Context, id beta.PolicyRoleManagementPolicyId, options ListRoleManagementPolicyRulesOperationOptions) (ListRoleManagementPolicyRulesCompleteResult, error) {
	return c.ListRoleManagementPolicyRulesCompleteMatchingPredicate(ctx, id, options, UnifiedRoleManagementPolicyRuleOperationPredicate{})
}

// ListRoleManagementPolicyRulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementPolicyRuleClient) ListRoleManagementPolicyRulesCompleteMatchingPredicate(ctx context.Context, id beta.PolicyRoleManagementPolicyId, options ListRoleManagementPolicyRulesOperationOptions, predicate UnifiedRoleManagementPolicyRuleOperationPredicate) (result ListRoleManagementPolicyRulesCompleteResult, err error) {
	items := make([]beta.UnifiedRoleManagementPolicyRule, 0)

	resp, err := c.ListRoleManagementPolicyRules(ctx, id, options)
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

	result = ListRoleManagementPolicyRulesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
