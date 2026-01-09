package userexperienceanalyticsbatteryhealthdeviceruntimehistory

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory
}

type CreateUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryOperationOptions() CreateUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryOperationOptions {
	return CreateUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryOperationOptions{}
}

func (o CreateUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory - Create new navigation property to
// userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory for deviceManagement
func (c UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryClient) CreateUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory(ctx context.Context, input beta.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory, options CreateUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryOperationOptions) (result CreateUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryOperationResponse, err error) {
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
		Path:          "/deviceManagement/userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory",
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

	var model beta.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
