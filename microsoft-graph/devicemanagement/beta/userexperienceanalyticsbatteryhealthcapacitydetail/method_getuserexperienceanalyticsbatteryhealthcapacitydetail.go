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

type GetUserExperienceAnalyticsBatteryHealthCapacityDetailOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UserExperienceAnalyticsBatteryHealthCapacityDetails
}

type GetUserExperienceAnalyticsBatteryHealthCapacityDetailOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetUserExperienceAnalyticsBatteryHealthCapacityDetailOperationOptions() GetUserExperienceAnalyticsBatteryHealthCapacityDetailOperationOptions {
	return GetUserExperienceAnalyticsBatteryHealthCapacityDetailOperationOptions{}
}

func (o GetUserExperienceAnalyticsBatteryHealthCapacityDetailOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetUserExperienceAnalyticsBatteryHealthCapacityDetailOperationOptions) ToOData() *odata.Query {
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

func (o GetUserExperienceAnalyticsBatteryHealthCapacityDetailOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetUserExperienceAnalyticsBatteryHealthCapacityDetail - Get userExperienceAnalyticsBatteryHealthCapacityDetails from
// deviceManagement. User Experience Analytics Battery Health Capacity Details
func (c UserExperienceAnalyticsBatteryHealthCapacityDetailClient) GetUserExperienceAnalyticsBatteryHealthCapacityDetail(ctx context.Context, options GetUserExperienceAnalyticsBatteryHealthCapacityDetailOperationOptions) (result GetUserExperienceAnalyticsBatteryHealthCapacityDetailOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/deviceManagement/userExperienceAnalyticsBatteryHealthCapacityDetails",
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

	var model beta.UserExperienceAnalyticsBatteryHealthCapacityDetails
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
