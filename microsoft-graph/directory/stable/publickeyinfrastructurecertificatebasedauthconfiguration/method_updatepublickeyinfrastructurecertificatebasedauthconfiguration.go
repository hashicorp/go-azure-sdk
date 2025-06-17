package publickeyinfrastructurecertificatebasedauthconfiguration

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdatePublicKeyInfrastructureCertificateBasedAuthConfigurationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdatePublicKeyInfrastructureCertificateBasedAuthConfigurationOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdatePublicKeyInfrastructureCertificateBasedAuthConfigurationOperationOptions() UpdatePublicKeyInfrastructureCertificateBasedAuthConfigurationOperationOptions {
	return UpdatePublicKeyInfrastructureCertificateBasedAuthConfigurationOperationOptions{}
}

func (o UpdatePublicKeyInfrastructureCertificateBasedAuthConfigurationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdatePublicKeyInfrastructureCertificateBasedAuthConfigurationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdatePublicKeyInfrastructureCertificateBasedAuthConfigurationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdatePublicKeyInfrastructureCertificateBasedAuthConfiguration - Update the navigation property
// certificateBasedAuthConfigurations in directory
func (c PublicKeyInfrastructureCertificateBasedAuthConfigurationClient) UpdatePublicKeyInfrastructureCertificateBasedAuthConfiguration(ctx context.Context, id stable.DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationId, input stable.CertificateBasedAuthPki, options UpdatePublicKeyInfrastructureCertificateBasedAuthConfigurationOperationOptions) (result UpdatePublicKeyInfrastructureCertificateBasedAuthConfigurationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
