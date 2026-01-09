package rolemanagementpolicyrule

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetRoleManagementPolicyRuleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.UnifiedRoleManagementPolicyRule
}

type GetRoleManagementPolicyRuleOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetRoleManagementPolicyRuleOperationOptions() GetRoleManagementPolicyRuleOperationOptions {
	return GetRoleManagementPolicyRuleOperationOptions{}
}

func (o GetRoleManagementPolicyRuleOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetRoleManagementPolicyRuleOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetRoleManagementPolicyRuleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetRoleManagementPolicyRule - Get unifiedRoleManagementPolicyRule. Retrieve a rule or settings defined for a role
// management policy. The rule can be one of the following types that are derived from the
// unifiedRoleManagementPolicyRule object
func (c RoleManagementPolicyRuleClient) GetRoleManagementPolicyRule(ctx context.Context, id stable.PolicyRoleManagementPolicyIdRuleId, options GetRoleManagementPolicyRuleOperationOptions) (result GetRoleManagementPolicyRuleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := stable.UnmarshalUnifiedRoleManagementPolicyRuleImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
