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

type CreateCertificateConnectorDetailOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CertificateConnectorDetails
}

type CreateCertificateConnectorDetailOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateCertificateConnectorDetailOperationOptions() CreateCertificateConnectorDetailOperationOptions {
	return CreateCertificateConnectorDetailOperationOptions{}
}

func (o CreateCertificateConnectorDetailOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateCertificateConnectorDetailOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateCertificateConnectorDetailOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateCertificateConnectorDetail - Create new navigation property to certificateConnectorDetails for deviceManagement
func (c CertificateConnectorDetailClient) CreateCertificateConnectorDetail(ctx context.Context, input beta.CertificateConnectorDetails, options CreateCertificateConnectorDetailOperationOptions) (result CreateCertificateConnectorDetailOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/certificateConnectorDetails",
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

	var model beta.CertificateConnectorDetails
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
