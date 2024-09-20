package certificateauthoritycertificatebasedapplicationconfigurationtrustedcertificateauthority

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityOperationOptions() UpdateCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityOperationOptions {
	return UpdateCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityOperationOptions{}
}

func (o UpdateCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthority - Update
// certificateAuthorityAsEntity. Update the properties of a certificateAuthorityAsEntity object.
func (c CertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityClient) UpdateCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthority(ctx context.Context, id beta.DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIdTrustedCertificateAuthorityId, input beta.CertificateAuthorityAsEntity, options UpdateCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityOperationOptions) (result UpdateCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
