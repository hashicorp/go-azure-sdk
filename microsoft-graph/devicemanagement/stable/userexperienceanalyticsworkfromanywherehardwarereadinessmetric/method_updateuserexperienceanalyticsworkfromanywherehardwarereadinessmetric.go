package userexperienceanalyticsworkfromanywherehardwarereadinessmetric

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricOperationOptions() UpdateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricOperationOptions {
	return UpdateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricOperationOptions{}
}

func (o UpdateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric - Update the navigation property
// userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric in deviceManagement
func (c UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricClient) UpdateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric(ctx context.Context, input stable.UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric, options UpdateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricOperationOptions) (result UpdateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          "/deviceManagement/userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric",
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
