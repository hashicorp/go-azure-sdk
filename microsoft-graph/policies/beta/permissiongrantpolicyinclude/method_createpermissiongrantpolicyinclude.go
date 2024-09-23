package permissiongrantpolicyinclude

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

type CreatePermissionGrantPolicyIncludeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.PermissionGrantConditionSet
}

type CreatePermissionGrantPolicyIncludeOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreatePermissionGrantPolicyIncludeOperationOptions() CreatePermissionGrantPolicyIncludeOperationOptions {
	return CreatePermissionGrantPolicyIncludeOperationOptions{}
}

func (o CreatePermissionGrantPolicyIncludeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreatePermissionGrantPolicyIncludeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreatePermissionGrantPolicyIncludeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreatePermissionGrantPolicyInclude - Create permissionGrantConditionSet in includes collection of
// permissionGrantPolicy. Add conditions under which a permission grant event is *included* in a permission grant
// policy. You do this by adding a permissionGrantConditionSet to the includes collection of a permissionGrantPolicy.
func (c PermissionGrantPolicyIncludeClient) CreatePermissionGrantPolicyInclude(ctx context.Context, id beta.PolicyPermissionGrantPolicyId, input beta.PermissionGrantConditionSet, options CreatePermissionGrantPolicyIncludeOperationOptions) (result CreatePermissionGrantPolicyIncludeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/includes", id.ID()),
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

	var model beta.PermissionGrantConditionSet
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
