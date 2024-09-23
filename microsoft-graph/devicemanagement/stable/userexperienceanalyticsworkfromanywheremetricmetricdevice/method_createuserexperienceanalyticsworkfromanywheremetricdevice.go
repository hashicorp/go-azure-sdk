package userexperienceanalyticsworkfromanywheremetricmetricdevice

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserExperienceAnalyticsWorkFromAnywhereMetricDeviceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UserExperienceAnalyticsWorkFromAnywhereDevice
}

type CreateUserExperienceAnalyticsWorkFromAnywhereMetricDeviceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateUserExperienceAnalyticsWorkFromAnywhereMetricDeviceOperationOptions() CreateUserExperienceAnalyticsWorkFromAnywhereMetricDeviceOperationOptions {
	return CreateUserExperienceAnalyticsWorkFromAnywhereMetricDeviceOperationOptions{}
}

func (o CreateUserExperienceAnalyticsWorkFromAnywhereMetricDeviceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateUserExperienceAnalyticsWorkFromAnywhereMetricDeviceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateUserExperienceAnalyticsWorkFromAnywhereMetricDeviceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateUserExperienceAnalyticsWorkFromAnywhereMetricDevice - Create new navigation property to metricDevices for
// deviceManagement
func (c UserExperienceAnalyticsWorkFromAnywhereMetricMetricDeviceClient) CreateUserExperienceAnalyticsWorkFromAnywhereMetricDevice(ctx context.Context, id stable.DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId, input stable.UserExperienceAnalyticsWorkFromAnywhereDevice, options CreateUserExperienceAnalyticsWorkFromAnywhereMetricDeviceOperationOptions) (result CreateUserExperienceAnalyticsWorkFromAnywhereMetricDeviceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/metricDevices", id.ID()),
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

	var model stable.UserExperienceAnalyticsWorkFromAnywhereDevice
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
