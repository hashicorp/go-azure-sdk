package publickeyinfrastructure

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetPublicKeyInfrastructureOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.PublicKeyInfrastructureRoot
}

type GetPublicKeyInfrastructureOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetPublicKeyInfrastructureOperationOptions() GetPublicKeyInfrastructureOperationOptions {
	return GetPublicKeyInfrastructureOperationOptions{}
}

func (o GetPublicKeyInfrastructureOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetPublicKeyInfrastructureOperationOptions) ToOData() *odata.Query {
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

func (o GetPublicKeyInfrastructureOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetPublicKeyInfrastructure - Get publicKeyInfrastructure from directory. The collection of public key infrastructure
// instances for the certificate-based authentication feature for users in a Microsoft Entra tenant.
func (c PublicKeyInfrastructureClient) GetPublicKeyInfrastructure(ctx context.Context, options GetPublicKeyInfrastructureOperationOptions) (result GetPublicKeyInfrastructureOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/directory/publicKeyInfrastructure",
		RetryFunc:     options.RetryFunc,
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

	var model beta.PublicKeyInfrastructureRoot
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
