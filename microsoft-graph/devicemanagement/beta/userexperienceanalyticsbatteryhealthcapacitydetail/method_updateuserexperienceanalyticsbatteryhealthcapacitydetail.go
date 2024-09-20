package userexperienceanalyticsbatteryhealthcapacitydetail

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateUserExperienceAnalyticsBatteryHealthCapacityDetailOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateUserExperienceAnalyticsBatteryHealthCapacityDetailOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateUserExperienceAnalyticsBatteryHealthCapacityDetailOperationOptions() UpdateUserExperienceAnalyticsBatteryHealthCapacityDetailOperationOptions {
	return UpdateUserExperienceAnalyticsBatteryHealthCapacityDetailOperationOptions{}
}

func (o UpdateUserExperienceAnalyticsBatteryHealthCapacityDetailOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateUserExperienceAnalyticsBatteryHealthCapacityDetailOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateUserExperienceAnalyticsBatteryHealthCapacityDetailOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateUserExperienceAnalyticsBatteryHealthCapacityDetail - Update the navigation property
// userExperienceAnalyticsBatteryHealthCapacityDetails in deviceManagement
func (c UserExperienceAnalyticsBatteryHealthCapacityDetailClient) UpdateUserExperienceAnalyticsBatteryHealthCapacityDetail(ctx context.Context, input beta.UserExperienceAnalyticsBatteryHealthCapacityDetails, options UpdateUserExperienceAnalyticsBatteryHealthCapacityDetailOperationOptions) (result UpdateUserExperienceAnalyticsBatteryHealthCapacityDetailOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          "/deviceManagement/userExperienceAnalyticsBatteryHealthCapacityDetails",
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
