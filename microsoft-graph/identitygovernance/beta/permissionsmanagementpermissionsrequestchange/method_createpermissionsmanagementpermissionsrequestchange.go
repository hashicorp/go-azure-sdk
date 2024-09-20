package permissionsmanagementpermissionsrequestchange

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreatePermissionsManagementPermissionsRequestChangeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.PermissionsRequestChange
}

type CreatePermissionsManagementPermissionsRequestChangeOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreatePermissionsManagementPermissionsRequestChangeOperationOptions() CreatePermissionsManagementPermissionsRequestChangeOperationOptions {
	return CreatePermissionsManagementPermissionsRequestChangeOperationOptions{}
}

func (o CreatePermissionsManagementPermissionsRequestChangeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreatePermissionsManagementPermissionsRequestChangeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreatePermissionsManagementPermissionsRequestChangeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreatePermissionsManagementPermissionsRequestChange - Create new navigation property to permissionsRequestChanges for
// identityGovernance
func (c PermissionsManagementPermissionsRequestChangeClient) CreatePermissionsManagementPermissionsRequestChange(ctx context.Context, input beta.PermissionsRequestChange, options CreatePermissionsManagementPermissionsRequestChangeOperationOptions) (result CreatePermissionsManagementPermissionsRequestChangeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/identityGovernance/permissionsManagement/permissionsRequestChanges",
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

	var model beta.PermissionsRequestChange
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
