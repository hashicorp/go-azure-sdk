package userexperienceanalyticsbaselinerebootanalyticsmetric

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

type GetUserExperienceAnalyticsBaselineRebootAnalyticsMetricOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UserExperienceAnalyticsCategory
}

type GetUserExperienceAnalyticsBaselineRebootAnalyticsMetricOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetUserExperienceAnalyticsBaselineRebootAnalyticsMetricOperationOptions() GetUserExperienceAnalyticsBaselineRebootAnalyticsMetricOperationOptions {
	return GetUserExperienceAnalyticsBaselineRebootAnalyticsMetricOperationOptions{}
}

func (o GetUserExperienceAnalyticsBaselineRebootAnalyticsMetricOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetUserExperienceAnalyticsBaselineRebootAnalyticsMetricOperationOptions) ToOData() *odata.Query {
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

func (o GetUserExperienceAnalyticsBaselineRebootAnalyticsMetricOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetUserExperienceAnalyticsBaselineRebootAnalyticsMetric - Get rebootAnalyticsMetrics from deviceManagement. The
// scores and insights for the reboot analytics metrics.
func (c UserExperienceAnalyticsBaselineRebootAnalyticsMetricClient) GetUserExperienceAnalyticsBaselineRebootAnalyticsMetric(ctx context.Context, id beta.DeviceManagementUserExperienceAnalyticsBaselineId, options GetUserExperienceAnalyticsBaselineRebootAnalyticsMetricOperationOptions) (result GetUserExperienceAnalyticsBaselineRebootAnalyticsMetricOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/rebootAnalyticsMetrics", id.ID()),
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

	var model beta.UserExperienceAnalyticsCategory
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
