package serviceprincipalcreationpolicyexclude

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateServicePrincipalCreationPolicyExcludeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateServicePrincipalCreationPolicyExcludeOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateServicePrincipalCreationPolicyExcludeOperationOptions() UpdateServicePrincipalCreationPolicyExcludeOperationOptions {
	return UpdateServicePrincipalCreationPolicyExcludeOperationOptions{}
}

func (o UpdateServicePrincipalCreationPolicyExcludeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateServicePrincipalCreationPolicyExcludeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateServicePrincipalCreationPolicyExcludeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateServicePrincipalCreationPolicyExclude - Update the navigation property excludes in policies
func (c ServicePrincipalCreationPolicyExcludeClient) UpdateServicePrincipalCreationPolicyExclude(ctx context.Context, id beta.PolicyServicePrincipalCreationPolicyIdExcludeId, input beta.ServicePrincipalCreationConditionSet, options UpdateServicePrincipalCreationPolicyExcludeOperationOptions) (result UpdateServicePrincipalCreationPolicyExcludeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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

	return
}
