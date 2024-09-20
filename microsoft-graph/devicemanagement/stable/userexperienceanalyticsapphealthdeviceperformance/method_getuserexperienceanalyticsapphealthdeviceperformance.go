package userexperienceanalyticsapphealthdeviceperformance

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserExperienceAnalyticsAppHealthDevicePerformanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UserExperienceAnalyticsAppHealthDevicePerformance
}

type GetUserExperienceAnalyticsAppHealthDevicePerformanceOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetUserExperienceAnalyticsAppHealthDevicePerformanceOperationOptions() GetUserExperienceAnalyticsAppHealthDevicePerformanceOperationOptions {
	return GetUserExperienceAnalyticsAppHealthDevicePerformanceOperationOptions{}
}

func (o GetUserExperienceAnalyticsAppHealthDevicePerformanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetUserExperienceAnalyticsAppHealthDevicePerformanceOperationOptions) ToOData() *odata.Query {
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

func (o GetUserExperienceAnalyticsAppHealthDevicePerformanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetUserExperienceAnalyticsAppHealthDevicePerformance - Get userExperienceAnalyticsAppHealthDevicePerformance from
// deviceManagement. User experience analytics appHealth Device Performance
func (c UserExperienceAnalyticsAppHealthDevicePerformanceClient) GetUserExperienceAnalyticsAppHealthDevicePerformance(ctx context.Context, id stable.DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceId, options GetUserExperienceAnalyticsAppHealthDevicePerformanceOperationOptions) (result GetUserExperienceAnalyticsAppHealthDevicePerformanceOperationResponse, err error) {
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

	var model stable.UserExperienceAnalyticsAppHealthDevicePerformance
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
