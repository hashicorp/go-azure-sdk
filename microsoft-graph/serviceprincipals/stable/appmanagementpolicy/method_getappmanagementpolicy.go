package appmanagementpolicy

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAppManagementPolicyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AppManagementPolicy
}

type GetAppManagementPolicyOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetAppManagementPolicyOperationOptions() GetAppManagementPolicyOperationOptions {
	return GetAppManagementPolicyOperationOptions{}
}

func (o GetAppManagementPolicyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAppManagementPolicyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetAppManagementPolicyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAppManagementPolicy - Get appManagementPolicies from servicePrincipals. The appManagementPolicy applied to this
// application.
func (c AppManagementPolicyClient) GetAppManagementPolicy(ctx context.Context, id stable.ServicePrincipalIdAppManagementPolicyId, options GetAppManagementPolicyOperationOptions) (result GetAppManagementPolicyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model stable.AppManagementPolicy
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
