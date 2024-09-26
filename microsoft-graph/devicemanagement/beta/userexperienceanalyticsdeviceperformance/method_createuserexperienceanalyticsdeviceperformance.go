package userexperienceanalyticsdeviceperformance

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserExperienceAnalyticsDevicePerformanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UserExperienceAnalyticsDevicePerformance
}

type CreateUserExperienceAnalyticsDevicePerformanceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateUserExperienceAnalyticsDevicePerformanceOperationOptions() CreateUserExperienceAnalyticsDevicePerformanceOperationOptions {
	return CreateUserExperienceAnalyticsDevicePerformanceOperationOptions{}
}

func (o CreateUserExperienceAnalyticsDevicePerformanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateUserExperienceAnalyticsDevicePerformanceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateUserExperienceAnalyticsDevicePerformanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateUserExperienceAnalyticsDevicePerformance - Create new navigation property to
// userExperienceAnalyticsDevicePerformance for deviceManagement
func (c UserExperienceAnalyticsDevicePerformanceClient) CreateUserExperienceAnalyticsDevicePerformance(ctx context.Context, input beta.UserExperienceAnalyticsDevicePerformance, options CreateUserExperienceAnalyticsDevicePerformanceOperationOptions) (result CreateUserExperienceAnalyticsDevicePerformanceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/userExperienceAnalyticsDevicePerformance",
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

	var model beta.UserExperienceAnalyticsDevicePerformance
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
