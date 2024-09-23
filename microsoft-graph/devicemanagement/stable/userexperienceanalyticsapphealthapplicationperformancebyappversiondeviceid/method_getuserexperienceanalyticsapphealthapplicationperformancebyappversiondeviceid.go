package userexperienceanalyticsapphealthapplicationperformancebyappversiondeviceid

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId
}

type GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdOperationOptions() GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdOperationOptions {
	return GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdOperationOptions{}
}

func (o GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdOperationOptions) ToOData() *odata.Query {
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

func (o GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId - Get
// userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId from deviceManagement. User experience
// analytics appHealth Application Performance by App Version Device Id
func (c UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdClient) GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId(ctx context.Context, id stable.DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdId, options GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdOperationOptions) (result GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdOperationResponse, err error) {
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

	var model stable.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
