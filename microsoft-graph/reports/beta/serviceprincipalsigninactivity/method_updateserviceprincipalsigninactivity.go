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

type UpdateServicePrincipalSignInActivityOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateServicePrincipalSignInActivityOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateServicePrincipalSignInActivityOperationOptions() UpdateServicePrincipalSignInActivityOperationOptions {
	return UpdateServicePrincipalSignInActivityOperationOptions{}
}

func (o UpdateServicePrincipalSignInActivityOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateServicePrincipalSignInActivityOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateServicePrincipalSignInActivityOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateServicePrincipalSignInActivity - Update the navigation property servicePrincipalSignInActivities in reports
func (c ServicePrincipalSignInActivityClient) UpdateServicePrincipalSignInActivity(ctx context.Context, id beta.ReportServicePrincipalSignInActivityId, input beta.ServicePrincipalSignInActivity, options UpdateServicePrincipalSignInActivityOperationOptions) (result UpdateServicePrincipalSignInActivityOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
