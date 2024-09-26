package userexperienceanalyticsapphealthosversionperformance

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateUserExperienceAnalyticsAppHealthOSVersionPerformanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateUserExperienceAnalyticsAppHealthOSVersionPerformanceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateUserExperienceAnalyticsAppHealthOSVersionPerformanceOperationOptions() UpdateUserExperienceAnalyticsAppHealthOSVersionPerformanceOperationOptions {
	return UpdateUserExperienceAnalyticsAppHealthOSVersionPerformanceOperationOptions{}
}

func (o UpdateUserExperienceAnalyticsAppHealthOSVersionPerformanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateUserExperienceAnalyticsAppHealthOSVersionPerformanceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateUserExperienceAnalyticsAppHealthOSVersionPerformanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateUserExperienceAnalyticsAppHealthOSVersionPerformance - Update the navigation property
// userExperienceAnalyticsAppHealthOSVersionPerformance in deviceManagement
func (c UserExperienceAnalyticsAppHealthOSVersionPerformanceClient) UpdateUserExperienceAnalyticsAppHealthOSVersionPerformance(ctx context.Context, id stable.DeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceId, input stable.UserExperienceAnalyticsAppHealthOSVersionPerformance, options UpdateUserExperienceAnalyticsAppHealthOSVersionPerformanceOperationOptions) (result UpdateUserExperienceAnalyticsAppHealthOSVersionPerformanceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
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
