package adminconsentrequestpolicy

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateAdminConsentRequestPolicyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateAdminConsentRequestPolicyOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateAdminConsentRequestPolicyOperationOptions() UpdateAdminConsentRequestPolicyOperationOptions {
	return UpdateAdminConsentRequestPolicyOperationOptions{}
}

func (o UpdateAdminConsentRequestPolicyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateAdminConsentRequestPolicyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateAdminConsentRequestPolicyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateAdminConsentRequestPolicy - Update adminConsentRequestPolicy. Update the properties of an
// adminConsentRequestPolicy object.
func (c AdminConsentRequestPolicyClient) UpdateAdminConsentRequestPolicy(ctx context.Context, input beta.AdminConsentRequestPolicy, options UpdateAdminConsentRequestPolicyOperationOptions) (result UpdateAdminConsentRequestPolicyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          "/policies/adminConsentRequestPolicy",
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
