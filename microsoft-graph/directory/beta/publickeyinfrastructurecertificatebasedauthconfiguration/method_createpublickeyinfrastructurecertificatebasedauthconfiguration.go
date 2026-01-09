package publickeyinfrastructurecertificatebasedauthconfiguration

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreatePublicKeyInfrastructureCertificateBasedAuthConfigurationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CertificateBasedAuthPki
}

type CreatePublicKeyInfrastructureCertificateBasedAuthConfigurationOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreatePublicKeyInfrastructureCertificateBasedAuthConfigurationOperationOptions() CreatePublicKeyInfrastructureCertificateBasedAuthConfigurationOperationOptions {
	return CreatePublicKeyInfrastructureCertificateBasedAuthConfigurationOperationOptions{}
}

func (o CreatePublicKeyInfrastructureCertificateBasedAuthConfigurationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreatePublicKeyInfrastructureCertificateBasedAuthConfigurationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreatePublicKeyInfrastructureCertificateBasedAuthConfigurationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreatePublicKeyInfrastructureCertificateBasedAuthConfiguration - Create certificateBasedAuthPki. Create a new
// certificateBasedAuthPki object.
func (c PublicKeyInfrastructureCertificateBasedAuthConfigurationClient) CreatePublicKeyInfrastructureCertificateBasedAuthConfiguration(ctx context.Context, input beta.CertificateBasedAuthPki, options CreatePublicKeyInfrastructureCertificateBasedAuthConfigurationOperationOptions) (result CreatePublicKeyInfrastructureCertificateBasedAuthConfigurationOperationResponse, err error) {
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
		Path:          "/directory/publicKeyInfrastructure/certificateBasedAuthConfigurations",
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

	var model beta.CertificateBasedAuthPki
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
