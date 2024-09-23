package slaazureadauthentication

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateSlaAzureADAuthenticationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateSlaAzureADAuthenticationOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateSlaAzureADAuthenticationOperationOptions() UpdateSlaAzureADAuthenticationOperationOptions {
	return UpdateSlaAzureADAuthenticationOperationOptions{}
}

func (o UpdateSlaAzureADAuthenticationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateSlaAzureADAuthenticationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateSlaAzureADAuthenticationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateSlaAzureADAuthentication - Update the navigation property azureADAuthentication in reports
func (c SlaAzureADAuthenticationClient) UpdateSlaAzureADAuthentication(ctx context.Context, input beta.AzureADAuthentication, options UpdateSlaAzureADAuthenticationOperationOptions) (result UpdateSlaAzureADAuthenticationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          "/reports/sla/azureADAuthentication",
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
