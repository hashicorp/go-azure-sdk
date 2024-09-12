package userexperienceanalyticsapphealthdevicemodelperformance

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserExperienceAnalyticsAppHealthDeviceModelPerformanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UserExperienceAnalyticsAppHealthDeviceModelPerformance
}

type GetUserExperienceAnalyticsAppHealthDeviceModelPerformanceOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetUserExperienceAnalyticsAppHealthDeviceModelPerformanceOperationOptions() GetUserExperienceAnalyticsAppHealthDeviceModelPerformanceOperationOptions {
	return GetUserExperienceAnalyticsAppHealthDeviceModelPerformanceOperationOptions{}
}

func (o GetUserExperienceAnalyticsAppHealthDeviceModelPerformanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetUserExperienceAnalyticsAppHealthDeviceModelPerformanceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetUserExperienceAnalyticsAppHealthDeviceModelPerformanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetUserExperienceAnalyticsAppHealthDeviceModelPerformance - Get
// userExperienceAnalyticsAppHealthDeviceModelPerformance from deviceManagement. User experience analytics appHealth
// Model Performance
func (c UserExperienceAnalyticsAppHealthDeviceModelPerformanceClient) GetUserExperienceAnalyticsAppHealthDeviceModelPerformance(ctx context.Context, id stable.DeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceId, options GetUserExperienceAnalyticsAppHealthDeviceModelPerformanceOperationOptions) (result GetUserExperienceAnalyticsAppHealthDeviceModelPerformanceOperationResponse, err error) {
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

	var model stable.UserExperienceAnalyticsAppHealthDeviceModelPerformance
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
