package userexperienceanalyticscategorymetricvalue

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserExperienceAnalyticsCategoryMetricValueOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UserExperienceAnalyticsMetric
}

type CreateUserExperienceAnalyticsCategoryMetricValueOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateUserExperienceAnalyticsCategoryMetricValueOperationOptions() CreateUserExperienceAnalyticsCategoryMetricValueOperationOptions {
	return CreateUserExperienceAnalyticsCategoryMetricValueOperationOptions{}
}

func (o CreateUserExperienceAnalyticsCategoryMetricValueOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateUserExperienceAnalyticsCategoryMetricValueOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateUserExperienceAnalyticsCategoryMetricValueOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateUserExperienceAnalyticsCategoryMetricValue - Create new navigation property to metricValues for
// deviceManagement
func (c UserExperienceAnalyticsCategoryMetricValueClient) CreateUserExperienceAnalyticsCategoryMetricValue(ctx context.Context, id beta.DeviceManagementUserExperienceAnalyticsCategoryId, input beta.UserExperienceAnalyticsMetric, options CreateUserExperienceAnalyticsCategoryMetricValueOperationOptions) (result CreateUserExperienceAnalyticsCategoryMetricValueOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/metricValues", id.ID()),
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

	var model beta.UserExperienceAnalyticsMetric
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
