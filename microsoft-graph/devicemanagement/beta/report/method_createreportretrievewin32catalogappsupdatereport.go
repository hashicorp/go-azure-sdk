package report

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateReportRetrieveWin32CatalogAppsUpdateReportOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type CreateReportRetrieveWin32CatalogAppsUpdateReportOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateReportRetrieveWin32CatalogAppsUpdateReportOperationOptions() CreateReportRetrieveWin32CatalogAppsUpdateReportOperationOptions {
	return CreateReportRetrieveWin32CatalogAppsUpdateReportOperationOptions{}
}

func (o CreateReportRetrieveWin32CatalogAppsUpdateReportOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateReportRetrieveWin32CatalogAppsUpdateReportOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateReportRetrieveWin32CatalogAppsUpdateReportOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateReportRetrieveWin32CatalogAppsUpdateReport - Invoke action retrieveWin32CatalogAppsUpdateReport
func (c ReportClient) CreateReportRetrieveWin32CatalogAppsUpdateReport(ctx context.Context, input CreateReportRetrieveWin32CatalogAppsUpdateReportRequest, options CreateReportRetrieveWin32CatalogAppsUpdateReportOperationOptions) (result CreateReportRetrieveWin32CatalogAppsUpdateReportOperationResponse, err error) {
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
		Path:          "/deviceManagement/reports/retrieveWin32CatalogAppsUpdateReport",
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
