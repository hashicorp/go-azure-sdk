package virtualendpointreport

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateVirtualEndpointReportRetrieveConnectionQualityReportOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type CreateVirtualEndpointReportRetrieveConnectionQualityReportOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateVirtualEndpointReportRetrieveConnectionQualityReportOperationOptions() CreateVirtualEndpointReportRetrieveConnectionQualityReportOperationOptions {
	return CreateVirtualEndpointReportRetrieveConnectionQualityReportOperationOptions{}
}

func (o CreateVirtualEndpointReportRetrieveConnectionQualityReportOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateVirtualEndpointReportRetrieveConnectionQualityReportOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateVirtualEndpointReportRetrieveConnectionQualityReportOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateVirtualEndpointReportRetrieveConnectionQualityReport - Invoke action retrieveConnectionQualityReports
func (c VirtualEndpointReportClient) CreateVirtualEndpointReportRetrieveConnectionQualityReport(ctx context.Context, input CreateVirtualEndpointReportRetrieveConnectionQualityReportRequest, options CreateVirtualEndpointReportRetrieveConnectionQualityReportOperationOptions) (result CreateVirtualEndpointReportRetrieveConnectionQualityReportOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/virtualEndpoint/reports/retrieveConnectionQualityReports",
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
