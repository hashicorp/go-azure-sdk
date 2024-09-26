package certificateconnectordetail

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateCertificateConnectorDetailOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateCertificateConnectorDetailOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateCertificateConnectorDetailOperationOptions() UpdateCertificateConnectorDetailOperationOptions {
	return UpdateCertificateConnectorDetailOperationOptions{}
}

func (o UpdateCertificateConnectorDetailOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateCertificateConnectorDetailOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateCertificateConnectorDetailOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateCertificateConnectorDetail - Update the navigation property certificateConnectorDetails in deviceManagement
func (c CertificateConnectorDetailClient) UpdateCertificateConnectorDetail(ctx context.Context, id beta.DeviceManagementCertificateConnectorDetailId, input beta.CertificateConnectorDetails, options UpdateCertificateConnectorDetailOperationOptions) (result UpdateCertificateConnectorDetailOperationResponse, err error) {
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
