package rolemanagementpolicy

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateRoleManagementPolicyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UnifiedRoleManagementPolicy
}

type CreateRoleManagementPolicyOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateRoleManagementPolicyOperationOptions() CreateRoleManagementPolicyOperationOptions {
	return CreateRoleManagementPolicyOperationOptions{}
}

func (o CreateRoleManagementPolicyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateRoleManagementPolicyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateRoleManagementPolicyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateRoleManagementPolicy - Create new navigation property to roleManagementPolicies for policies
func (c RoleManagementPolicyClient) CreateRoleManagementPolicy(ctx context.Context, input stable.UnifiedRoleManagementPolicy, options CreateRoleManagementPolicyOperationOptions) (result CreateRoleManagementPolicyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/policies/roleManagementPolicies",
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

	var model stable.UnifiedRoleManagementPolicy
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
