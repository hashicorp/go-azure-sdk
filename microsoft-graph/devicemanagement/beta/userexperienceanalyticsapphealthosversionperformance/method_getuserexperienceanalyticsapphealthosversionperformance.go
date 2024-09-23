package userexperienceanalyticsapphealthosversionperformance

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserExperienceAnalyticsAppHealthOSVersionPerformanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UserExperienceAnalyticsAppHealthOSVersionPerformance
}

type GetUserExperienceAnalyticsAppHealthOSVersionPerformanceOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetUserExperienceAnalyticsAppHealthOSVersionPerformanceOperationOptions() GetUserExperienceAnalyticsAppHealthOSVersionPerformanceOperationOptions {
	return GetUserExperienceAnalyticsAppHealthOSVersionPerformanceOperationOptions{}
}

func (o GetUserExperienceAnalyticsAppHealthOSVersionPerformanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetUserExperienceAnalyticsAppHealthOSVersionPerformanceOperationOptions) ToOData() *odata.Query {
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

func (o GetUserExperienceAnalyticsAppHealthOSVersionPerformanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetUserExperienceAnalyticsAppHealthOSVersionPerformance - Get userExperienceAnalyticsAppHealthOSVersionPerformance
// from deviceManagement. User experience analytics appHealth OS version Performance
func (c UserExperienceAnalyticsAppHealthOSVersionPerformanceClient) GetUserExperienceAnalyticsAppHealthOSVersionPerformance(ctx context.Context, id beta.DeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceId, options GetUserExperienceAnalyticsAppHealthOSVersionPerformanceOperationOptions) (result GetUserExperienceAnalyticsAppHealthOSVersionPerformanceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
		RetryFunc:     options.RetryFunc,
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

	var model beta.UserExperienceAnalyticsAppHealthOSVersionPerformance
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
