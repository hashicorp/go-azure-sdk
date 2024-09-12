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

type GetUserExperienceAnalyticsBatteryHealthRuntimeDetailOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UserExperienceAnalyticsBatteryHealthRuntimeDetails
}

type GetUserExperienceAnalyticsBatteryHealthRuntimeDetailOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetUserExperienceAnalyticsBatteryHealthRuntimeDetailOperationOptions() GetUserExperienceAnalyticsBatteryHealthRuntimeDetailOperationOptions {
	return GetUserExperienceAnalyticsBatteryHealthRuntimeDetailOperationOptions{}
}

func (o GetUserExperienceAnalyticsBatteryHealthRuntimeDetailOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetUserExperienceAnalyticsBatteryHealthRuntimeDetailOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetUserExperienceAnalyticsBatteryHealthRuntimeDetailOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetUserExperienceAnalyticsBatteryHealthRuntimeDetail - Get userExperienceAnalyticsBatteryHealthRuntimeDetails from
// deviceManagement. User Experience Analytics Battery Health Runtime Details
func (c UserExperienceAnalyticsBatteryHealthRuntimeDetailClient) GetUserExperienceAnalyticsBatteryHealthRuntimeDetail(ctx context.Context, options GetUserExperienceAnalyticsBatteryHealthRuntimeDetailOperationOptions) (result GetUserExperienceAnalyticsBatteryHealthRuntimeDetailOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/deviceManagement/userExperienceAnalyticsBatteryHealthRuntimeDetails",
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

	var model beta.UserExperienceAnalyticsBatteryHealthRuntimeDetails
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
