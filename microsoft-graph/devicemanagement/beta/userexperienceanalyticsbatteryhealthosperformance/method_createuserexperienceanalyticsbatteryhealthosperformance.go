package userexperienceanalyticsbatteryhealthosperformance

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserExperienceAnalyticsBatteryHealthOsPerformanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UserExperienceAnalyticsBatteryHealthOsPerformance
}

type CreateUserExperienceAnalyticsBatteryHealthOsPerformanceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateUserExperienceAnalyticsBatteryHealthOsPerformanceOperationOptions() CreateUserExperienceAnalyticsBatteryHealthOsPerformanceOperationOptions {
	return CreateUserExperienceAnalyticsBatteryHealthOsPerformanceOperationOptions{}
}

func (o CreateUserExperienceAnalyticsBatteryHealthOsPerformanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateUserExperienceAnalyticsBatteryHealthOsPerformanceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateUserExperienceAnalyticsBatteryHealthOsPerformanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateUserExperienceAnalyticsBatteryHealthOsPerformance - Create new navigation property to
// userExperienceAnalyticsBatteryHealthOsPerformance for deviceManagement
func (c UserExperienceAnalyticsBatteryHealthOsPerformanceClient) CreateUserExperienceAnalyticsBatteryHealthOsPerformance(ctx context.Context, input beta.UserExperienceAnalyticsBatteryHealthOsPerformance, options CreateUserExperienceAnalyticsBatteryHealthOsPerformanceOperationOptions) (result CreateUserExperienceAnalyticsBatteryHealthOsPerformanceOperationResponse, err error) {
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
		Path:          "/deviceManagement/userExperienceAnalyticsBatteryHealthOsPerformance",
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

	var model beta.UserExperienceAnalyticsBatteryHealthOsPerformance
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
