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

type GetUserExperienceAnalyticsAnomalyCorrelationGroupOverviewOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UserExperienceAnalyticsAnomalyCorrelationGroupOverview
}

type GetUserExperienceAnalyticsAnomalyCorrelationGroupOverviewOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetUserExperienceAnalyticsAnomalyCorrelationGroupOverviewOperationOptions() GetUserExperienceAnalyticsAnomalyCorrelationGroupOverviewOperationOptions {
	return GetUserExperienceAnalyticsAnomalyCorrelationGroupOverviewOperationOptions{}
}

func (o GetUserExperienceAnalyticsAnomalyCorrelationGroupOverviewOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetUserExperienceAnalyticsAnomalyCorrelationGroupOverviewOperationOptions) ToOData() *odata.Query {
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

func (o GetUserExperienceAnalyticsAnomalyCorrelationGroupOverviewOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetUserExperienceAnalyticsAnomalyCorrelationGroupOverview - Get
// userExperienceAnalyticsAnomalyCorrelationGroupOverview from deviceManagement. The user experience analytics anomaly
// correlation group overview entity contains the information for each correlation group of an anomaly.
func (c UserExperienceAnalyticsAnomalyCorrelationGroupOverviewClient) GetUserExperienceAnalyticsAnomalyCorrelationGroupOverview(ctx context.Context, id beta.DeviceManagementUserExperienceAnalyticsAnomalyCorrelationGroupOverviewId, options GetUserExperienceAnalyticsAnomalyCorrelationGroupOverviewOperationOptions) (result GetUserExperienceAnalyticsAnomalyCorrelationGroupOverviewOperationResponse, err error) {
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

	var model beta.UserExperienceAnalyticsAnomalyCorrelationGroupOverview
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
