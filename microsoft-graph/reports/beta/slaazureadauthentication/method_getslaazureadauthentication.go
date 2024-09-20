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

type GetSlaAzureADAuthenticationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AzureADAuthentication
}

type GetSlaAzureADAuthenticationOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetSlaAzureADAuthenticationOperationOptions() GetSlaAzureADAuthenticationOperationOptions {
	return GetSlaAzureADAuthenticationOperationOptions{}
}

func (o GetSlaAzureADAuthenticationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetSlaAzureADAuthenticationOperationOptions) ToOData() *odata.Query {
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

func (o GetSlaAzureADAuthenticationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetSlaAzureADAuthentication - Get azureADAuthentication. Read the properties and relationships of an
// azureADAuthentication object to find the level of Microsoft Entra authentication availability for your tenant. The
// Microsoft Entra service Level Agreement (SLA) commits to at least 99.99% authentication availability, as described in
// Microsoft Entra SLA performance. This object provides you with your tenant's actual performance against this
// commitment.
func (c SlaAzureADAuthenticationClient) GetSlaAzureADAuthentication(ctx context.Context, options GetSlaAzureADAuthenticationOperationOptions) (result GetSlaAzureADAuthenticationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/reports/sla/azureADAuthentication",
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

	var model beta.AzureADAuthentication
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
