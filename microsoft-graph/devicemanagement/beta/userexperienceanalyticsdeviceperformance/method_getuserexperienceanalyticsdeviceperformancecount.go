package userexperienceanalyticsdeviceperformance

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserExperienceAnalyticsDevicePerformanceCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetUserExperienceAnalyticsDevicePerformanceCountOperationOptions struct {
	Filter   *string
	Metadata *odata.Metadata
	Search   *string
}

func DefaultGetUserExperienceAnalyticsDevicePerformanceCountOperationOptions() GetUserExperienceAnalyticsDevicePerformanceCountOperationOptions {
	return GetUserExperienceAnalyticsDevicePerformanceCountOperationOptions{}
}

func (o GetUserExperienceAnalyticsDevicePerformanceCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetUserExperienceAnalyticsDevicePerformanceCountOperationOptions) ToOData() *odata.Query {
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

func (o GetUserExperienceAnalyticsDevicePerformanceCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetUserExperienceAnalyticsDevicePerformanceCount - Get the number of the resource
func (c UserExperienceAnalyticsDevicePerformanceClient) GetUserExperienceAnalyticsDevicePerformanceCount(ctx context.Context, options GetUserExperienceAnalyticsDevicePerformanceCountOperationOptions) (result GetUserExperienceAnalyticsDevicePerformanceCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/deviceManagement/userExperienceAnalyticsDevicePerformance/$count",
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
