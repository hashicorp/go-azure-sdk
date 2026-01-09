package certificateauthoritymutualtlsoauthconfiguration

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateCertificateAuthorityMutualTlsOauthConfigurationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateCertificateAuthorityMutualTlsOauthConfigurationOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateCertificateAuthorityMutualTlsOauthConfigurationOperationOptions() UpdateCertificateAuthorityMutualTlsOauthConfigurationOperationOptions {
	return UpdateCertificateAuthorityMutualTlsOauthConfigurationOperationOptions{}
}

func (o UpdateCertificateAuthorityMutualTlsOauthConfigurationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateCertificateAuthorityMutualTlsOauthConfigurationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateCertificateAuthorityMutualTlsOauthConfigurationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateCertificateAuthorityMutualTlsOauthConfiguration - Update mutualTlsOauthConfiguration. Update the specified
// mutualTlsOauthConfiguration resource. You can only update the following two properties: displayName,
// certificateAuthority. To update a subset of objects in the certificateAuthorities collection, first get the complete
// list, make your modifications, and then repost the entire contents of the certificateAuthorities attribute list in
// the request body. Excluding a subset of objects removes them from the collection.
func (c CertificateAuthorityMutualTlsOauthConfigurationClient) UpdateCertificateAuthorityMutualTlsOauthConfiguration(ctx context.Context, id beta.DirectoryCertificateAuthorityMutualTlsOauthConfigurationId, input beta.MutualTlsOauthConfiguration, options UpdateCertificateAuthorityMutualTlsOauthConfigurationOperationOptions) (result UpdateCertificateAuthorityMutualTlsOauthConfigurationOperationResponse, err error) {
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
