package detectordefinitionresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiagnosticsExecuteSiteDetectorOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *DiagnosticDetectorResponse
}

type DiagnosticsExecuteSiteDetectorOperationOptions struct {
	EndTime   *string
	StartTime *string
	TimeGrain *string
}

func DefaultDiagnosticsExecuteSiteDetectorOperationOptions() DiagnosticsExecuteSiteDetectorOperationOptions {
	return DiagnosticsExecuteSiteDetectorOperationOptions{}
}

func (o DiagnosticsExecuteSiteDetectorOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DiagnosticsExecuteSiteDetectorOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DiagnosticsExecuteSiteDetectorOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.EndTime != nil {
		out.Append("endTime", fmt.Sprintf("%v", *o.EndTime))
	}
	if o.StartTime != nil {
		out.Append("startTime", fmt.Sprintf("%v", *o.StartTime))
	}
	if o.TimeGrain != nil {
		out.Append("timeGrain", fmt.Sprintf("%v", *o.TimeGrain))
	}
	return &out
}

// DiagnosticsExecuteSiteDetector ...
func (c DetectorDefinitionResourcesClient) DiagnosticsExecuteSiteDetector(ctx context.Context, id DiagnosticDetectorId, options DiagnosticsExecuteSiteDetectorOperationOptions) (result DiagnosticsExecuteSiteDetectorOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/execute", id.ID()),
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

	var model DiagnosticDetectorResponse
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
