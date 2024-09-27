package userexperienceanalyticsworkfromanywheremetricmetricdevice

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateUserExperienceAnalyticsWorkFromAnywhereMetricDeviceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateUserExperienceAnalyticsWorkFromAnywhereMetricDeviceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateUserExperienceAnalyticsWorkFromAnywhereMetricDeviceOperationOptions() UpdateUserExperienceAnalyticsWorkFromAnywhereMetricDeviceOperationOptions {
	return UpdateUserExperienceAnalyticsWorkFromAnywhereMetricDeviceOperationOptions{}
}

func (o UpdateUserExperienceAnalyticsWorkFromAnywhereMetricDeviceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateUserExperienceAnalyticsWorkFromAnywhereMetricDeviceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateUserExperienceAnalyticsWorkFromAnywhereMetricDeviceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateUserExperienceAnalyticsWorkFromAnywhereMetricDevice - Update the navigation property metricDevices in
// deviceManagement
func (c UserExperienceAnalyticsWorkFromAnywhereMetricMetricDeviceClient) UpdateUserExperienceAnalyticsWorkFromAnywhereMetricDevice(ctx context.Context, id stable.DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricIdMetricDeviceId, input stable.UserExperienceAnalyticsWorkFromAnywhereDevice, options UpdateUserExperienceAnalyticsWorkFromAnywhereMetricDeviceOperationOptions) (result UpdateUserExperienceAnalyticsWorkFromAnywhereMetricDeviceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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

	return
}
