package userexperienceanalyticsapphealthdevicemodelperformance

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserExperienceAnalyticsAppHealthDeviceModelPerformanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UserExperienceAnalyticsAppHealthDeviceModelPerformance
}

type CreateUserExperienceAnalyticsAppHealthDeviceModelPerformanceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateUserExperienceAnalyticsAppHealthDeviceModelPerformanceOperationOptions() CreateUserExperienceAnalyticsAppHealthDeviceModelPerformanceOperationOptions {
	return CreateUserExperienceAnalyticsAppHealthDeviceModelPerformanceOperationOptions{}
}

func (o CreateUserExperienceAnalyticsAppHealthDeviceModelPerformanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateUserExperienceAnalyticsAppHealthDeviceModelPerformanceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateUserExperienceAnalyticsAppHealthDeviceModelPerformanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateUserExperienceAnalyticsAppHealthDeviceModelPerformance - Create new navigation property to
// userExperienceAnalyticsAppHealthDeviceModelPerformance for deviceManagement
func (c UserExperienceAnalyticsAppHealthDeviceModelPerformanceClient) CreateUserExperienceAnalyticsAppHealthDeviceModelPerformance(ctx context.Context, input stable.UserExperienceAnalyticsAppHealthDeviceModelPerformance, options CreateUserExperienceAnalyticsAppHealthDeviceModelPerformanceOperationOptions) (result CreateUserExperienceAnalyticsAppHealthDeviceModelPerformanceOperationResponse, err error) {
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
		Path:          "/deviceManagement/userExperienceAnalyticsAppHealthDeviceModelPerformance",
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

	var model stable.UserExperienceAnalyticsAppHealthDeviceModelPerformance
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
