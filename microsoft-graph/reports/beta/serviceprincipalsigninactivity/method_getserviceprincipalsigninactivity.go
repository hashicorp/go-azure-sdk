package serviceprincipalsigninactivity

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetServicePrincipalSignInActivityOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ServicePrincipalSignInActivity
}

type GetServicePrincipalSignInActivityOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetServicePrincipalSignInActivityOperationOptions() GetServicePrincipalSignInActivityOperationOptions {
	return GetServicePrincipalSignInActivityOperationOptions{}
}

func (o GetServicePrincipalSignInActivityOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetServicePrincipalSignInActivityOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetServicePrincipalSignInActivityOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetServicePrincipalSignInActivity - Get servicePrincipalSignInActivity. Get a servicePrincipalSignInActivity object
// that contains sign-in activity information for a service principal in a Microsoft Entra tenant. You can use a service
// principal as a client or resource. A service principal supports delegated or app-only authentication context.
func (c ServicePrincipalSignInActivityClient) GetServicePrincipalSignInActivity(ctx context.Context, id beta.ReportServicePrincipalSignInActivityId, options GetServicePrincipalSignInActivityOperationOptions) (result GetServicePrincipalSignInActivityOperationResponse, err error) {
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

	var model beta.ServicePrincipalSignInActivity
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
