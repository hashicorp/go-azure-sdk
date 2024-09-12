package rolemanagementpolicyrule

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateRoleManagementPolicyRuleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// UpdateRoleManagementPolicyRule - Update unifiedRoleManagementPolicyRule. Update a rule defined for a role management
// policy. The rule can be one of the following types that are derived from the unifiedRoleManagementPolicyRule object:
// For more information about rules for Microsoft Entra roles and examples of updating rules, see the following
// articles:
func (c RoleManagementPolicyRuleClient) UpdateRoleManagementPolicyRule(ctx context.Context, id stable.PolicyRoleManagementPolicyIdRuleId, input stable.UnifiedRoleManagementPolicyRule) (result UpdateRoleManagementPolicyRuleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPatch,
		Path:       id.ID(),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
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

	return
}
