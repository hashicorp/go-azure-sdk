package userexperienceanalyticsbatteryhealthosperformance

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserExperienceAnalyticsBatteryHealthOsPerformanceCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetUserExperienceAnalyticsBatteryHealthOsPerformanceCountOperationOptions struct {
	Filter   *string
	Metadata *odata.Metadata
	Search   *string
}

func DefaultGetUserExperienceAnalyticsBatteryHealthOsPerformanceCountOperationOptions() GetUserExperienceAnalyticsBatteryHealthOsPerformanceCountOperationOptions {
	return GetUserExperienceAnalyticsBatteryHealthOsPerformanceCountOperationOptions{}
}

func (o GetUserExperienceAnalyticsBatteryHealthOsPerformanceCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetUserExperienceAnalyticsBatteryHealthOsPerformanceCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	return &out
}

func (o GetUserExperienceAnalyticsBatteryHealthOsPerformanceCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetUserExperienceAnalyticsBatteryHealthOsPerformanceCount - Get the number of the resource
func (c UserExperienceAnalyticsBatteryHealthOsPerformanceClient) GetUserExperienceAnalyticsBatteryHealthOsPerformanceCount(ctx context.Context, options GetUserExperienceAnalyticsBatteryHealthOsPerformanceCountOperationOptions) (result GetUserExperienceAnalyticsBatteryHealthOsPerformanceCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/deviceManagement/userExperienceAnalyticsBatteryHealthOsPerformance/$count",
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
