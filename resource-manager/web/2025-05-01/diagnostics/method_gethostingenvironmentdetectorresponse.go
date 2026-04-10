package diagnostics

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetHostingEnvironmentDetectorResponseOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *DetectorResponse
}

type GetHostingEnvironmentDetectorResponseOperationOptions struct {
	EndTime   *string
	StartTime *string
	TimeGrain *string
}

func DefaultGetHostingEnvironmentDetectorResponseOperationOptions() GetHostingEnvironmentDetectorResponseOperationOptions {
	return GetHostingEnvironmentDetectorResponseOperationOptions{}
}

func (o GetHostingEnvironmentDetectorResponseOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetHostingEnvironmentDetectorResponseOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o GetHostingEnvironmentDetectorResponseOperationOptions) ToQuery() *client.QueryParams {
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

// GetHostingEnvironmentDetectorResponse ...
func (c DiagnosticsClient) GetHostingEnvironmentDetectorResponse(ctx context.Context, id HostingEnvironmentDetectorId, options GetHostingEnvironmentDetectorResponseOperationOptions) (result GetHostingEnvironmentDetectorResponseOperationResponse, err error) {
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

	var model DetectorResponse
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
