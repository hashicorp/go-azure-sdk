package b2xuserflowapiconnectorconfigurationpostattributecollection

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateB2xUserFlowApiConnectorConfigurationPostAttributeCollectionUploadClientCertificateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.IdentityApiConnector
}

type CreateB2xUserFlowApiConnectorConfigurationPostAttributeCollectionUploadClientCertificateOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateB2xUserFlowApiConnectorConfigurationPostAttributeCollectionUploadClientCertificateOperationOptions() CreateB2xUserFlowApiConnectorConfigurationPostAttributeCollectionUploadClientCertificateOperationOptions {
	return CreateB2xUserFlowApiConnectorConfigurationPostAttributeCollectionUploadClientCertificateOperationOptions{}
}

func (o CreateB2xUserFlowApiConnectorConfigurationPostAttributeCollectionUploadClientCertificateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateB2xUserFlowApiConnectorConfigurationPostAttributeCollectionUploadClientCertificateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateB2xUserFlowApiConnectorConfigurationPostAttributeCollectionUploadClientCertificateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateB2xUserFlowApiConnectorConfigurationPostAttributeCollectionUploadClientCertificate - Invoke action
// uploadClientCertificate. Upload a PKCS 12 format key (.pfx) to an API connector's authentication configuration. The
// input is a base-64 encoded value of the PKCS 12 certificate contents. This method returns an apiConnector.
func (c B2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient) CreateB2xUserFlowApiConnectorConfigurationPostAttributeCollectionUploadClientCertificate(ctx context.Context, id stable.IdentityB2xUserFlowId, input CreateB2xUserFlowApiConnectorConfigurationPostAttributeCollectionUploadClientCertificateRequest, options CreateB2xUserFlowApiConnectorConfigurationPostAttributeCollectionUploadClientCertificateOperationOptions) (result CreateB2xUserFlowApiConnectorConfigurationPostAttributeCollectionUploadClientCertificateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/apiConnectorConfiguration/postAttributeCollection/uploadClientCertificate", id.ID()),
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
