package userexperienceanalyticsanomalycorrelationgroupoverview

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserExperienceAnalyticsAnomalyCorrelationGroupOverviewOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UserExperienceAnalyticsAnomalyCorrelationGroupOverview
}

type CreateUserExperienceAnalyticsAnomalyCorrelationGroupOverviewOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateUserExperienceAnalyticsAnomalyCorrelationGroupOverviewOperationOptions() CreateUserExperienceAnalyticsAnomalyCorrelationGroupOverviewOperationOptions {
	return CreateUserExperienceAnalyticsAnomalyCorrelationGroupOverviewOperationOptions{}
}

func (o CreateUserExperienceAnalyticsAnomalyCorrelationGroupOverviewOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateUserExperienceAnalyticsAnomalyCorrelationGroupOverviewOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateUserExperienceAnalyticsAnomalyCorrelationGroupOverviewOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateUserExperienceAnalyticsAnomalyCorrelationGroupOverview - Create new navigation property to
// userExperienceAnalyticsAnomalyCorrelationGroupOverview for deviceManagement
func (c UserExperienceAnalyticsAnomalyCorrelationGroupOverviewClient) CreateUserExperienceAnalyticsAnomalyCorrelationGroupOverview(ctx context.Context, input beta.UserExperienceAnalyticsAnomalyCorrelationGroupOverview, options CreateUserExperienceAnalyticsAnomalyCorrelationGroupOverviewOperationOptions) (result CreateUserExperienceAnalyticsAnomalyCorrelationGroupOverviewOperationResponse, err error) {
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
		Path:          "/deviceManagement/userExperienceAnalyticsAnomalyCorrelationGroupOverview",
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

	var model beta.UserExperienceAnalyticsAnomalyCorrelationGroupOverview
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
