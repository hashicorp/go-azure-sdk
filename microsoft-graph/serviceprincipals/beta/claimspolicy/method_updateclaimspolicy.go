package claimspolicy

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

type UpdateClaimsPolicyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateClaimsPolicyOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateClaimsPolicyOperationOptions() UpdateClaimsPolicyOperationOptions {
	return UpdateClaimsPolicyOperationOptions{}
}

func (o UpdateClaimsPolicyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateClaimsPolicyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateClaimsPolicyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateClaimsPolicy - Create or replace claimsPolicy. Create a new customClaimsPolicy object if it doesn't exist, or
// replace an existing one.
func (c ClaimsPolicyClient) UpdateClaimsPolicy(ctx context.Context, id beta.ServicePrincipalId, input beta.CustomClaimsPolicy, options UpdateClaimsPolicyOperationOptions) (result UpdateClaimsPolicyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/claimsPolicy", id.ID()),
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
