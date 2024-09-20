package userexperienceanalyticsdevicestartuphistory

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserExperienceAnalyticsDeviceStartupHistoryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UserExperienceAnalyticsDeviceStartupHistory
}

type GetUserExperienceAnalyticsDeviceStartupHistoryOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetUserExperienceAnalyticsDeviceStartupHistoryOperationOptions() GetUserExperienceAnalyticsDeviceStartupHistoryOperationOptions {
	return GetUserExperienceAnalyticsDeviceStartupHistoryOperationOptions{}
}

func (o GetUserExperienceAnalyticsDeviceStartupHistoryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetUserExperienceAnalyticsDeviceStartupHistoryOperationOptions) ToOData() *odata.Query {
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

func (o GetUserExperienceAnalyticsDeviceStartupHistoryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetUserExperienceAnalyticsDeviceStartupHistory - Get userExperienceAnalyticsDeviceStartupHistory from
// deviceManagement. User experience analytics device Startup History
func (c UserExperienceAnalyticsDeviceStartupHistoryClient) GetUserExperienceAnalyticsDeviceStartupHistory(ctx context.Context, id stable.DeviceManagementUserExperienceAnalyticsDeviceStartupHistoryId, options GetUserExperienceAnalyticsDeviceStartupHistoryOperationOptions) (result GetUserExperienceAnalyticsDeviceStartupHistoryOperationResponse, err error) {
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

	var model stable.UserExperienceAnalyticsDeviceStartupHistory
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
