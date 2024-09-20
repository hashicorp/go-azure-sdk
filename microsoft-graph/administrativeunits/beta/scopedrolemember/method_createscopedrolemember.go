package scopedrolemember

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

type CreateScopedRoleMemberOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ScopedRoleMembership
}

type CreateScopedRoleMemberOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateScopedRoleMemberOperationOptions() CreateScopedRoleMemberOperationOptions {
	return CreateScopedRoleMemberOperationOptions{}
}

func (o CreateScopedRoleMemberOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateScopedRoleMemberOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateScopedRoleMemberOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateScopedRoleMember - Add a scopedRoleMember. Assign a Microsoft Entra role with administrative unit scope. For a
// list of roles that can be assigned with administrative unit scope, see Assign Microsoft Entra roles with
// administrative unit scope.
func (c ScopedRoleMemberClient) CreateScopedRoleMember(ctx context.Context, id beta.AdministrativeUnitId, input beta.ScopedRoleMembership, options CreateScopedRoleMemberOperationOptions) (result CreateScopedRoleMemberOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/scopedRoleMembers", id.ID()),
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

	var model beta.ScopedRoleMembership
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
