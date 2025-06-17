package certificateauthoritymutualtlsoauthconfiguration

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateCertificateAuthorityMutualTlsOauthConfigurationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.MutualTlsOauthConfiguration
}

type CreateCertificateAuthorityMutualTlsOauthConfigurationOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateCertificateAuthorityMutualTlsOauthConfigurationOperationOptions() CreateCertificateAuthorityMutualTlsOauthConfigurationOperationOptions {
	return CreateCertificateAuthorityMutualTlsOauthConfigurationOperationOptions{}
}

func (o CreateCertificateAuthorityMutualTlsOauthConfigurationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateCertificateAuthorityMutualTlsOauthConfigurationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateCertificateAuthorityMutualTlsOauthConfigurationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateCertificateAuthorityMutualTlsOauthConfiguration - Create mutualTlsOauthConfiguration. Create a
// mutualTlsOauthConfiguration resource that contains a specified certificate authority object.
func (c CertificateAuthorityMutualTlsOauthConfigurationClient) CreateCertificateAuthorityMutualTlsOauthConfiguration(ctx context.Context, input beta.MutualTlsOauthConfiguration, options CreateCertificateAuthorityMutualTlsOauthConfigurationOperationOptions) (result CreateCertificateAuthorityMutualTlsOauthConfigurationOperationResponse, err error) {
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
		Path:          "/directory/certificateAuthorities/mutualTlsOauthConfigurations",
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

	var model beta.MutualTlsOauthConfiguration
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
