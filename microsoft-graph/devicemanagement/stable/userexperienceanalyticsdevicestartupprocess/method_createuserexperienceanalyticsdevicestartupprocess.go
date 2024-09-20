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

type CreateUserExperienceAnalyticsDeviceStartupProcessOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UserExperienceAnalyticsDeviceStartupProcess
}

type CreateUserExperienceAnalyticsDeviceStartupProcessOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateUserExperienceAnalyticsDeviceStartupProcessOperationOptions() CreateUserExperienceAnalyticsDeviceStartupProcessOperationOptions {
	return CreateUserExperienceAnalyticsDeviceStartupProcessOperationOptions{}
}

func (o CreateUserExperienceAnalyticsDeviceStartupProcessOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateUserExperienceAnalyticsDeviceStartupProcessOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateUserExperienceAnalyticsDeviceStartupProcessOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateUserExperienceAnalyticsDeviceStartupProcess - Create new navigation property to
// userExperienceAnalyticsDeviceStartupProcesses for deviceManagement
func (c UserExperienceAnalyticsDeviceStartupProcessClient) CreateUserExperienceAnalyticsDeviceStartupProcess(ctx context.Context, input stable.UserExperienceAnalyticsDeviceStartupProcess, options CreateUserExperienceAnalyticsDeviceStartupProcessOperationOptions) (result CreateUserExperienceAnalyticsDeviceStartupProcessOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/userExperienceAnalyticsDeviceStartupProcesses",
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

	var model stable.UserExperienceAnalyticsDeviceStartupProcess
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
