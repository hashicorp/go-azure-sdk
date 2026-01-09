package crosstenantaccesspolicypartner

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateCrossTenantAccessPolicyPartnerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CrossTenantAccessPolicyConfigurationPartner
}

type CreateCrossTenantAccessPolicyPartnerOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateCrossTenantAccessPolicyPartnerOperationOptions() CreateCrossTenantAccessPolicyPartnerOperationOptions {
	return CreateCrossTenantAccessPolicyPartnerOperationOptions{}
}

func (o CreateCrossTenantAccessPolicyPartnerOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateCrossTenantAccessPolicyPartnerOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateCrossTenantAccessPolicyPartnerOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateCrossTenantAccessPolicyPartner - Create crossTenantAccessPolicyConfigurationPartner. Create a new partner
// configuration in a cross-tenant access policy.
func (c CrossTenantAccessPolicyPartnerClient) CreateCrossTenantAccessPolicyPartner(ctx context.Context, input beta.CrossTenantAccessPolicyConfigurationPartner, options CreateCrossTenantAccessPolicyPartnerOperationOptions) (result CreateCrossTenantAccessPolicyPartnerOperationResponse, err error) {
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
		Path:          "/policies/crossTenantAccessPolicy/partners",
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

	var model beta.CrossTenantAccessPolicyConfigurationPartner
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
