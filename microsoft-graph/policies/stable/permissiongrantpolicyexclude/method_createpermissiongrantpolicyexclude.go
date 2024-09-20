package permissiongrantpolicyexclude

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreatePermissionGrantPolicyExcludeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.PermissionGrantConditionSet
}

type CreatePermissionGrantPolicyExcludeOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreatePermissionGrantPolicyExcludeOperationOptions() CreatePermissionGrantPolicyExcludeOperationOptions {
	return CreatePermissionGrantPolicyExcludeOperationOptions{}
}

func (o CreatePermissionGrantPolicyExcludeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreatePermissionGrantPolicyExcludeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreatePermissionGrantPolicyExcludeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreatePermissionGrantPolicyExclude - Create permissionGrantConditionSet in excludes collection of
// permissionGrantPolicy. Add conditions under which a permission grant event is *excluded* in a permission grant
// policy. You do this by adding a permissionGrantConditionSet to the excludes collection of a permissionGrantPolicy.
func (c PermissionGrantPolicyExcludeClient) CreatePermissionGrantPolicyExclude(ctx context.Context, id stable.PolicyPermissionGrantPolicyId, input stable.PermissionGrantConditionSet, options CreatePermissionGrantPolicyExcludeOperationOptions) (result CreatePermissionGrantPolicyExcludeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/excludes", id.ID()),
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

	var model stable.PermissionGrantConditionSet
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
