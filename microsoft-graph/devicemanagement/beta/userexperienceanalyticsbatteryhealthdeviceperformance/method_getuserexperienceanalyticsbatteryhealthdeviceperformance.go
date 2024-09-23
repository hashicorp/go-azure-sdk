package userexperienceanalyticsbatteryhealthdeviceperformance

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserExperienceAnalyticsBatteryHealthDevicePerformanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UserExperienceAnalyticsBatteryHealthDevicePerformance
}

type GetUserExperienceAnalyticsBatteryHealthDevicePerformanceOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetUserExperienceAnalyticsBatteryHealthDevicePerformanceOperationOptions() GetUserExperienceAnalyticsBatteryHealthDevicePerformanceOperationOptions {
	return GetUserExperienceAnalyticsBatteryHealthDevicePerformanceOperationOptions{}
}

func (o GetUserExperienceAnalyticsBatteryHealthDevicePerformanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetUserExperienceAnalyticsBatteryHealthDevicePerformanceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetUserExperienceAnalyticsBatteryHealthDevicePerformanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetUserExperienceAnalyticsBatteryHealthDevicePerformance - Get userExperienceAnalyticsBatteryHealthDevicePerformance
// from deviceManagement. User Experience Analytics Battery Health Device Performance
func (c UserExperienceAnalyticsBatteryHealthDevicePerformanceClient) GetUserExperienceAnalyticsBatteryHealthDevicePerformance(ctx context.Context, id beta.DeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceId, options GetUserExperienceAnalyticsBatteryHealthDevicePerformanceOperationOptions) (result GetUserExperienceAnalyticsBatteryHealthDevicePerformanceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
		RetryFunc:     options.RetryFunc,
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

	var model beta.UserExperienceAnalyticsBatteryHealthDevicePerformance
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
