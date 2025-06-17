package publickeyinfrastructurecertificatebasedauthconfigurationcertificateauthority

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdatePublicKeyInfrastructureCertificateBasedAuthConfigurationCertificateAuthorityOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdatePublicKeyInfrastructureCertificateBasedAuthConfigurationCertificateAuthorityOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdatePublicKeyInfrastructureCertificateBasedAuthConfigurationCertificateAuthorityOperationOptions() UpdatePublicKeyInfrastructureCertificateBasedAuthConfigurationCertificateAuthorityOperationOptions {
	return UpdatePublicKeyInfrastructureCertificateBasedAuthConfigurationCertificateAuthorityOperationOptions{}
}

func (o UpdatePublicKeyInfrastructureCertificateBasedAuthConfigurationCertificateAuthorityOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdatePublicKeyInfrastructureCertificateBasedAuthConfigurationCertificateAuthorityOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdatePublicKeyInfrastructureCertificateBasedAuthConfigurationCertificateAuthorityOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdatePublicKeyInfrastructureCertificateBasedAuthConfigurationCertificateAuthority - Update the navigation property
// certificateAuthorities in directory
func (c PublicKeyInfrastructureCertificateBasedAuthConfigurationCertificateAuthorityClient) UpdatePublicKeyInfrastructureCertificateBasedAuthConfigurationCertificateAuthority(ctx context.Context, id stable.DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityId, input stable.CertificateAuthorityDetail, options UpdatePublicKeyInfrastructureCertificateBasedAuthConfigurationCertificateAuthorityOperationOptions) (result UpdatePublicKeyInfrastructureCertificateBasedAuthConfigurationCertificateAuthorityOperationResponse, err error) {
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
