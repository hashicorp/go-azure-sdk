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

type GetAdminConsentRequestPolicyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AdminConsentRequestPolicy
}

type GetAdminConsentRequestPolicyOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetAdminConsentRequestPolicyOperationOptions() GetAdminConsentRequestPolicyOperationOptions {
	return GetAdminConsentRequestPolicyOperationOptions{}
}

func (o GetAdminConsentRequestPolicyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAdminConsentRequestPolicyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetAdminConsentRequestPolicyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAdminConsentRequestPolicy - Get adminConsentRequestPolicy. Read the properties and relationships of an
// adminConsentRequestPolicy object.
func (c AdminConsentRequestPolicyClient) GetAdminConsentRequestPolicy(ctx context.Context, options GetAdminConsentRequestPolicyOperationOptions) (result GetAdminConsentRequestPolicyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/policies/adminConsentRequestPolicy",
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

	var model beta.AdminConsentRequestPolicy
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
