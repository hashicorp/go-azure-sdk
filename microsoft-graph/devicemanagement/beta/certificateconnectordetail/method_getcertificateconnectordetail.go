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

type GetCertificateConnectorDetailOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CertificateConnectorDetails
}

type GetCertificateConnectorDetailOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetCertificateConnectorDetailOperationOptions() GetCertificateConnectorDetailOperationOptions {
	return GetCertificateConnectorDetailOperationOptions{}
}

func (o GetCertificateConnectorDetailOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetCertificateConnectorDetailOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetCertificateConnectorDetailOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetCertificateConnectorDetail - Get certificateConnectorDetails from deviceManagement. Collection of certificate
// connector details, each associated with a corresponding Intune Certificate Connector.
func (c CertificateConnectorDetailClient) GetCertificateConnectorDetail(ctx context.Context, id beta.DeviceManagementCertificateConnectorDetailId, options GetCertificateConnectorDetailOperationOptions) (result GetCertificateConnectorDetailOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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

	var model beta.CertificateConnectorDetails
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
