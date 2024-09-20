package directoryrole

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateDirectoryRoleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DirectoryRole
}

type CreateDirectoryRoleOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateDirectoryRoleOperationOptions() CreateDirectoryRoleOperationOptions {
	return CreateDirectoryRoleOperationOptions{}
}

func (o CreateDirectoryRoleOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDirectoryRoleOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDirectoryRoleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDirectoryRole - Activate directoryRole. Activate a directory role. To read a directory role or update its
// members, it must first be activated in the tenant. The Company Administrators and the implicit user directory roles
// (User, Guest User, and Restricted Guest User roles) are activated by default. To access and assign members to other
// directory roles, you must first activate it with its corresponding directory role template ID.
func (c DirectoryRoleClient) CreateDirectoryRole(ctx context.Context, input beta.DirectoryRole, options CreateDirectoryRoleOperationOptions) (result CreateDirectoryRoleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/directoryRoles",
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

	var model beta.DirectoryRole
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
