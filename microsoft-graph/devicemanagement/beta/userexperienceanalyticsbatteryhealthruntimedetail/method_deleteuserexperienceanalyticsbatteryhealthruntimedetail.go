package userexperienceanalyticsbatteryhealthruntimedetail

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteUserExperienceAnalyticsBatteryHealthRuntimeDetailOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteUserExperienceAnalyticsBatteryHealthRuntimeDetailOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteUserExperienceAnalyticsBatteryHealthRuntimeDetailOperationOptions() DeleteUserExperienceAnalyticsBatteryHealthRuntimeDetailOperationOptions {
	return DeleteUserExperienceAnalyticsBatteryHealthRuntimeDetailOperationOptions{}
}

func (o DeleteUserExperienceAnalyticsBatteryHealthRuntimeDetailOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteUserExperienceAnalyticsBatteryHealthRuntimeDetailOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteUserExperienceAnalyticsBatteryHealthRuntimeDetailOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteUserExperienceAnalyticsBatteryHealthRuntimeDetail - Delete navigation property
// userExperienceAnalyticsBatteryHealthRuntimeDetails for deviceManagement
func (c UserExperienceAnalyticsBatteryHealthRuntimeDetailClient) DeleteUserExperienceAnalyticsBatteryHealthRuntimeDetail(ctx context.Context, options DeleteUserExperienceAnalyticsBatteryHealthRuntimeDetailOperationOptions) (result DeleteUserExperienceAnalyticsBatteryHealthRuntimeDetailOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
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

	return
}
