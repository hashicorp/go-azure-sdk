package userexperienceanalyticsbatteryhealthruntimedetail

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateUserExperienceAnalyticsBatteryHealthRuntimeDetailOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateUserExperienceAnalyticsBatteryHealthRuntimeDetailOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateUserExperienceAnalyticsBatteryHealthRuntimeDetailOperationOptions() UpdateUserExperienceAnalyticsBatteryHealthRuntimeDetailOperationOptions {
	return UpdateUserExperienceAnalyticsBatteryHealthRuntimeDetailOperationOptions{}
}

func (o UpdateUserExperienceAnalyticsBatteryHealthRuntimeDetailOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateUserExperienceAnalyticsBatteryHealthRuntimeDetailOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateUserExperienceAnalyticsBatteryHealthRuntimeDetailOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateUserExperienceAnalyticsBatteryHealthRuntimeDetail - Update the navigation property
// userExperienceAnalyticsBatteryHealthRuntimeDetails in deviceManagement
func (c UserExperienceAnalyticsBatteryHealthRuntimeDetailClient) UpdateUserExperienceAnalyticsBatteryHealthRuntimeDetail(ctx context.Context, input beta.UserExperienceAnalyticsBatteryHealthRuntimeDetails, options UpdateUserExperienceAnalyticsBatteryHealthRuntimeDetailOperationOptions) (result UpdateUserExperienceAnalyticsBatteryHealthRuntimeDetailOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          "/deviceManagement/userExperienceAnalyticsBatteryHealthRuntimeDetails",
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
