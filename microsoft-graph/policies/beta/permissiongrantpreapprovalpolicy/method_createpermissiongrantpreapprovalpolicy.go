package permissiongrantpreapprovalpolicy

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreatePermissionGrantPreApprovalPolicyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.PermissionGrantPreApprovalPolicy
}

type CreatePermissionGrantPreApprovalPolicyOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreatePermissionGrantPreApprovalPolicyOperationOptions() CreatePermissionGrantPreApprovalPolicyOperationOptions {
	return CreatePermissionGrantPreApprovalPolicyOperationOptions{}
}

func (o CreatePermissionGrantPreApprovalPolicyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreatePermissionGrantPreApprovalPolicyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreatePermissionGrantPreApprovalPolicyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreatePermissionGrantPreApprovalPolicy - Create permissionGrantPreApprovalPolicy. Create a new
// permissionGrantPreApprovalPolicy object.
func (c PermissionGrantPreApprovalPolicyClient) CreatePermissionGrantPreApprovalPolicy(ctx context.Context, input beta.PermissionGrantPreApprovalPolicy, options CreatePermissionGrantPreApprovalPolicyOperationOptions) (result CreatePermissionGrantPreApprovalPolicyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/policies/permissionGrantPreApprovalPolicies",
		RetryFunc:     options.RetryFunc,
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

	var model beta.PermissionGrantPreApprovalPolicy
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
