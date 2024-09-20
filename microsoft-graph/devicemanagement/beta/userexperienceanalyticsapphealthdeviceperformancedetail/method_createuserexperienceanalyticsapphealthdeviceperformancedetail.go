package userexperienceanalyticsapphealthdeviceperformancedetail

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserExperienceAnalyticsAppHealthDevicePerformanceDetailOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UserExperienceAnalyticsAppHealthDevicePerformanceDetails
}

type CreateUserExperienceAnalyticsAppHealthDevicePerformanceDetailOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateUserExperienceAnalyticsAppHealthDevicePerformanceDetailOperationOptions() CreateUserExperienceAnalyticsAppHealthDevicePerformanceDetailOperationOptions {
	return CreateUserExperienceAnalyticsAppHealthDevicePerformanceDetailOperationOptions{}
}

func (o CreateUserExperienceAnalyticsAppHealthDevicePerformanceDetailOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateUserExperienceAnalyticsAppHealthDevicePerformanceDetailOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateUserExperienceAnalyticsAppHealthDevicePerformanceDetailOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateUserExperienceAnalyticsAppHealthDevicePerformanceDetail - Create new navigation property to
// userExperienceAnalyticsAppHealthDevicePerformanceDetails for deviceManagement
func (c UserExperienceAnalyticsAppHealthDevicePerformanceDetailClient) CreateUserExperienceAnalyticsAppHealthDevicePerformanceDetail(ctx context.Context, input beta.UserExperienceAnalyticsAppHealthDevicePerformanceDetails, options CreateUserExperienceAnalyticsAppHealthDevicePerformanceDetailOperationOptions) (result CreateUserExperienceAnalyticsAppHealthDevicePerformanceDetailOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/userExperienceAnalyticsAppHealthDevicePerformanceDetails",
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

	var model beta.UserExperienceAnalyticsAppHealthDevicePerformanceDetails
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
