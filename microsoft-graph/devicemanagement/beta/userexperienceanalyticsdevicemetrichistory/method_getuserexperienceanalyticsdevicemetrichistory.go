package userexperienceanalyticsdevicemetrichistory

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserExperienceAnalyticsDeviceMetricHistoryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UserExperienceAnalyticsMetricHistory
}

type GetUserExperienceAnalyticsDeviceMetricHistoryOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetUserExperienceAnalyticsDeviceMetricHistoryOperationOptions() GetUserExperienceAnalyticsDeviceMetricHistoryOperationOptions {
	return GetUserExperienceAnalyticsDeviceMetricHistoryOperationOptions{}
}

func (o GetUserExperienceAnalyticsDeviceMetricHistoryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetUserExperienceAnalyticsDeviceMetricHistoryOperationOptions) ToOData() *odata.Query {
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

func (o GetUserExperienceAnalyticsDeviceMetricHistoryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetUserExperienceAnalyticsDeviceMetricHistory - Get userExperienceAnalyticsDeviceMetricHistory from deviceManagement.
// User experience analytics device metric history. The report will be retired on December 31, 2024. You can start using
// the Cloud PC connection quality report now via
// https://learn.microsoft.com/windows-365/enterprise/report-cloud-pc-connection-quality.
func (c UserExperienceAnalyticsDeviceMetricHistoryClient) GetUserExperienceAnalyticsDeviceMetricHistory(ctx context.Context, id beta.DeviceManagementUserExperienceAnalyticsDeviceMetricHistoryId, options GetUserExperienceAnalyticsDeviceMetricHistoryOperationOptions) (result GetUserExperienceAnalyticsDeviceMetricHistoryOperationResponse, err error) {
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

	var model beta.UserExperienceAnalyticsMetricHistory
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
