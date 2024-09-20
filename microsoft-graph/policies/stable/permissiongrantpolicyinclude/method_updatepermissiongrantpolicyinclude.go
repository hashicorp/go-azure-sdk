package permissiongrantpolicyinclude

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdatePermissionGrantPolicyIncludeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdatePermissionGrantPolicyIncludeOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdatePermissionGrantPolicyIncludeOperationOptions() UpdatePermissionGrantPolicyIncludeOperationOptions {
	return UpdatePermissionGrantPolicyIncludeOperationOptions{}
}

func (o UpdatePermissionGrantPolicyIncludeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdatePermissionGrantPolicyIncludeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdatePermissionGrantPolicyIncludeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdatePermissionGrantPolicyInclude - Update the navigation property includes in policies
func (c PermissionGrantPolicyIncludeClient) UpdatePermissionGrantPolicyInclude(ctx context.Context, id stable.PolicyPermissionGrantPolicyIdIncludeId, input stable.PermissionGrantConditionSet, options UpdatePermissionGrantPolicyIncludeOperationOptions) (result UpdatePermissionGrantPolicyIncludeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
