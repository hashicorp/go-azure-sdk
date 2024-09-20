package userexperienceanalyticsbatteryhealthdeviceappimpact

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateUserExperienceAnalyticsBatteryHealthDeviceAppImpactOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateUserExperienceAnalyticsBatteryHealthDeviceAppImpactOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateUserExperienceAnalyticsBatteryHealthDeviceAppImpactOperationOptions() UpdateUserExperienceAnalyticsBatteryHealthDeviceAppImpactOperationOptions {
	return UpdateUserExperienceAnalyticsBatteryHealthDeviceAppImpactOperationOptions{}
}

func (o UpdateUserExperienceAnalyticsBatteryHealthDeviceAppImpactOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateUserExperienceAnalyticsBatteryHealthDeviceAppImpactOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateUserExperienceAnalyticsBatteryHealthDeviceAppImpactOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateUserExperienceAnalyticsBatteryHealthDeviceAppImpact - Update the navigation property
// userExperienceAnalyticsBatteryHealthDeviceAppImpact in deviceManagement
func (c UserExperienceAnalyticsBatteryHealthDeviceAppImpactClient) UpdateUserExperienceAnalyticsBatteryHealthDeviceAppImpact(ctx context.Context, id beta.DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactId, input beta.UserExperienceAnalyticsBatteryHealthDeviceAppImpact, options UpdateUserExperienceAnalyticsBatteryHealthDeviceAppImpactOperationOptions) (result UpdateUserExperienceAnalyticsBatteryHealthDeviceAppImpactOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
