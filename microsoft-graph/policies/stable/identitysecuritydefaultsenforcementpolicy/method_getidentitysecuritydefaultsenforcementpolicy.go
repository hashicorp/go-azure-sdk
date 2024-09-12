package identitysecuritydefaultsenforcementpolicy

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetIdentitySecurityDefaultsEnforcementPolicyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.IdentitySecurityDefaultsEnforcementPolicy
}

type GetIdentitySecurityDefaultsEnforcementPolicyOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetIdentitySecurityDefaultsEnforcementPolicyOperationOptions() GetIdentitySecurityDefaultsEnforcementPolicyOperationOptions {
	return GetIdentitySecurityDefaultsEnforcementPolicyOperationOptions{}
}

func (o GetIdentitySecurityDefaultsEnforcementPolicyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetIdentitySecurityDefaultsEnforcementPolicyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetIdentitySecurityDefaultsEnforcementPolicyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetIdentitySecurityDefaultsEnforcementPolicy - Get identitySecurityDefaultsEnforcementPolicy. Retrieve the properties
// of an identitySecurityDefaultsEnforcementPolicy object.
func (c IdentitySecurityDefaultsEnforcementPolicyClient) GetIdentitySecurityDefaultsEnforcementPolicy(ctx context.Context, options GetIdentitySecurityDefaultsEnforcementPolicyOperationOptions) (result GetIdentitySecurityDefaultsEnforcementPolicyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/policies/identitySecurityDefaultsEnforcementPolicy",
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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

	var model stable.IdentitySecurityDefaultsEnforcementPolicy
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
