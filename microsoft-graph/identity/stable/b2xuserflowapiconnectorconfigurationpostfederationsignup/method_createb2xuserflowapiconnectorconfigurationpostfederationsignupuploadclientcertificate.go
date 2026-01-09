package b2xuserflowapiconnectorconfigurationpostfederationsignup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateB2xUserFlowApiConnectorConfigurationPostFederationSignupUploadClientCertificateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.IdentityApiConnector
}

type CreateB2xUserFlowApiConnectorConfigurationPostFederationSignupUploadClientCertificateOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateB2xUserFlowApiConnectorConfigurationPostFederationSignupUploadClientCertificateOperationOptions() CreateB2xUserFlowApiConnectorConfigurationPostFederationSignupUploadClientCertificateOperationOptions {
	return CreateB2xUserFlowApiConnectorConfigurationPostFederationSignupUploadClientCertificateOperationOptions{}
}

func (o CreateB2xUserFlowApiConnectorConfigurationPostFederationSignupUploadClientCertificateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateB2xUserFlowApiConnectorConfigurationPostFederationSignupUploadClientCertificateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateB2xUserFlowApiConnectorConfigurationPostFederationSignupUploadClientCertificateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateB2xUserFlowApiConnectorConfigurationPostFederationSignupUploadClientCertificate - Invoke action
// uploadClientCertificate. Upload a PKCS 12 format key (.pfx) to an API connector's authentication configuration. The
// input is a base-64 encoded value of the PKCS 12 certificate contents. This method returns an apiConnector.
func (c B2xUserFlowApiConnectorConfigurationPostFederationSignupClient) CreateB2xUserFlowApiConnectorConfigurationPostFederationSignupUploadClientCertificate(ctx context.Context, id stable.IdentityB2xUserFlowId, input CreateB2xUserFlowApiConnectorConfigurationPostFederationSignupUploadClientCertificateRequest, options CreateB2xUserFlowApiConnectorConfigurationPostFederationSignupUploadClientCertificateOperationOptions) (result CreateB2xUserFlowApiConnectorConfigurationPostFederationSignupUploadClientCertificateOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/apiConnectorConfiguration/postFederationSignup/uploadClientCertificate", id.ID()),
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

	var model stable.IdentityApiConnector
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
