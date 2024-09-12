package userexperienceanalyticsapphealthdevicemodelperformance

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserExperienceAnalyticsAppHealthDeviceModelPerformanceCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetUserExperienceAnalyticsAppHealthDeviceModelPerformanceCountOperationOptions struct {
	Filter *string
	Search *string
}

func DefaultGetUserExperienceAnalyticsAppHealthDeviceModelPerformanceCountOperationOptions() GetUserExperienceAnalyticsAppHealthDeviceModelPerformanceCountOperationOptions {
	return GetUserExperienceAnalyticsAppHealthDeviceModelPerformanceCountOperationOptions{}
}

func (o GetUserExperienceAnalyticsAppHealthDeviceModelPerformanceCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetUserExperienceAnalyticsAppHealthDeviceModelPerformanceCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	return &out
}

func (o GetUserExperienceAnalyticsAppHealthDeviceModelPerformanceCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetUserExperienceAnalyticsAppHealthDeviceModelPerformanceCount - Get the number of the resource
func (c UserExperienceAnalyticsAppHealthDeviceModelPerformanceClient) GetUserExperienceAnalyticsAppHealthDeviceModelPerformanceCount(ctx context.Context, options GetUserExperienceAnalyticsAppHealthDeviceModelPerformanceCountOperationOptions) (result GetUserExperienceAnalyticsAppHealthDeviceModelPerformanceCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/deviceManagement/userExperienceAnalyticsAppHealthDeviceModelPerformance/$count",
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
