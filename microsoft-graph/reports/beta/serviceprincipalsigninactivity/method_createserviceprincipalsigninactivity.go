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

type CreateServicePrincipalSignInActivityOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ServicePrincipalSignInActivity
}

type CreateServicePrincipalSignInActivityOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateServicePrincipalSignInActivityOperationOptions() CreateServicePrincipalSignInActivityOperationOptions {
	return CreateServicePrincipalSignInActivityOperationOptions{}
}

func (o CreateServicePrincipalSignInActivityOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateServicePrincipalSignInActivityOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateServicePrincipalSignInActivityOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateServicePrincipalSignInActivity - Create new navigation property to servicePrincipalSignInActivities for reports
func (c ServicePrincipalSignInActivityClient) CreateServicePrincipalSignInActivity(ctx context.Context, input beta.ServicePrincipalSignInActivity, options CreateServicePrincipalSignInActivityOperationOptions) (result CreateServicePrincipalSignInActivityOperationResponse, err error) {
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
		Path:          "/reports/servicePrincipalSignInActivities",
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

	var model beta.ServicePrincipalSignInActivity
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
