package userexperienceanalyticsdevicestartupprocess

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateUserExperienceAnalyticsDeviceStartupProcessOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateUserExperienceAnalyticsDeviceStartupProcessOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateUserExperienceAnalyticsDeviceStartupProcessOperationOptions() UpdateUserExperienceAnalyticsDeviceStartupProcessOperationOptions {
	return UpdateUserExperienceAnalyticsDeviceStartupProcessOperationOptions{}
}

func (o UpdateUserExperienceAnalyticsDeviceStartupProcessOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateUserExperienceAnalyticsDeviceStartupProcessOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateUserExperienceAnalyticsDeviceStartupProcessOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateUserExperienceAnalyticsDeviceStartupProcess - Update the navigation property
// userExperienceAnalyticsDeviceStartupProcesses in deviceManagement
func (c UserExperienceAnalyticsDeviceStartupProcessClient) UpdateUserExperienceAnalyticsDeviceStartupProcess(ctx context.Context, id stable.DeviceManagementUserExperienceAnalyticsDeviceStartupProcessId, input stable.UserExperienceAnalyticsDeviceStartupProcess, options UpdateUserExperienceAnalyticsDeviceStartupProcessOperationOptions) (result UpdateUserExperienceAnalyticsDeviceStartupProcessOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
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
