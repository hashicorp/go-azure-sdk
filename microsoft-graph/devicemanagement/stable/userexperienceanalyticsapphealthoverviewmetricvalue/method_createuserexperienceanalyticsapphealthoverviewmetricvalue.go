package userexperienceanalyticsapphealthoverviewmetricvalue

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserExperienceAnalyticsAppHealthOverviewMetricValueOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UserExperienceAnalyticsMetric
}

type CreateUserExperienceAnalyticsAppHealthOverviewMetricValueOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateUserExperienceAnalyticsAppHealthOverviewMetricValueOperationOptions() CreateUserExperienceAnalyticsAppHealthOverviewMetricValueOperationOptions {
	return CreateUserExperienceAnalyticsAppHealthOverviewMetricValueOperationOptions{}
}

func (o CreateUserExperienceAnalyticsAppHealthOverviewMetricValueOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateUserExperienceAnalyticsAppHealthOverviewMetricValueOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateUserExperienceAnalyticsAppHealthOverviewMetricValueOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateUserExperienceAnalyticsAppHealthOverviewMetricValue - Create new navigation property to metricValues for
// deviceManagement
func (c UserExperienceAnalyticsAppHealthOverviewMetricValueClient) CreateUserExperienceAnalyticsAppHealthOverviewMetricValue(ctx context.Context, input stable.UserExperienceAnalyticsMetric, options CreateUserExperienceAnalyticsAppHealthOverviewMetricValueOperationOptions) (result CreateUserExperienceAnalyticsAppHealthOverviewMetricValueOperationResponse, err error) {
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
		Path:          "/deviceManagement/userExperienceAnalyticsAppHealthOverview/metricValues",
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

	var model stable.UserExperienceAnalyticsMetric
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
