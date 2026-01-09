package publickeyinfrastructurecertificatebasedauthconfiguration

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreatePublicKeyInfrastructureCertificateBasedAuthConfigurationUploadOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreatePublicKeyInfrastructureCertificateBasedAuthConfigurationUploadOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreatePublicKeyInfrastructureCertificateBasedAuthConfigurationUploadOperationOptions() CreatePublicKeyInfrastructureCertificateBasedAuthConfigurationUploadOperationOptions {
	return CreatePublicKeyInfrastructureCertificateBasedAuthConfigurationUploadOperationOptions{}
}

func (o CreatePublicKeyInfrastructureCertificateBasedAuthConfigurationUploadOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreatePublicKeyInfrastructureCertificateBasedAuthConfigurationUploadOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreatePublicKeyInfrastructureCertificateBasedAuthConfigurationUploadOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreatePublicKeyInfrastructureCertificateBasedAuthConfigurationUpload - Invoke action upload. Append additional
// certificate authority details to a certificateBasedAuthPki resource. Only one operation can run at a time and this
// operation can take up to 30 minutes to complete. To know whether another upload is in progress, call the Get
// certificateBasedAuthPki. The status property will have the value running.
func (c PublicKeyInfrastructureCertificateBasedAuthConfigurationClient) CreatePublicKeyInfrastructureCertificateBasedAuthConfigurationUpload(ctx context.Context, id beta.DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationId, input CreatePublicKeyInfrastructureCertificateBasedAuthConfigurationUploadRequest, options CreatePublicKeyInfrastructureCertificateBasedAuthConfigurationUploadOperationOptions) (result CreatePublicKeyInfrastructureCertificateBasedAuthConfigurationUploadOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/upload", id.ID()),
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
