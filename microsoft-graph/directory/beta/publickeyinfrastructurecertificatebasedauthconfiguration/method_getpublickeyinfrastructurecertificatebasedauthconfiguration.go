package publickeyinfrastructurecertificatebasedauthconfiguration

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetPublicKeyInfrastructureCertificateBasedAuthConfigurationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CertificateBasedAuthPki
}

type GetPublicKeyInfrastructureCertificateBasedAuthConfigurationOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetPublicKeyInfrastructureCertificateBasedAuthConfigurationOperationOptions() GetPublicKeyInfrastructureCertificateBasedAuthConfigurationOperationOptions {
	return GetPublicKeyInfrastructureCertificateBasedAuthConfigurationOperationOptions{}
}

func (o GetPublicKeyInfrastructureCertificateBasedAuthConfigurationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetPublicKeyInfrastructureCertificateBasedAuthConfigurationOperationOptions) ToOData() *odata.Query {
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

func (o GetPublicKeyInfrastructureCertificateBasedAuthConfigurationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetPublicKeyInfrastructureCertificateBasedAuthConfiguration - Get certificateBasedAuthPki. Read the properties and
// relationships of a certificateBasedAuthPki object.
func (c PublicKeyInfrastructureCertificateBasedAuthConfigurationClient) GetPublicKeyInfrastructureCertificateBasedAuthConfiguration(ctx context.Context, id beta.DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationId, options GetPublicKeyInfrastructureCertificateBasedAuthConfigurationOperationOptions) (result GetPublicKeyInfrastructureCertificateBasedAuthConfigurationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model beta.CertificateBasedAuthPki
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
