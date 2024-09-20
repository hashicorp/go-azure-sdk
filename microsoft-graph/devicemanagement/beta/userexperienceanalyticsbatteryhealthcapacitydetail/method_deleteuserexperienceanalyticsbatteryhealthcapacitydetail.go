package userexperienceanalyticsbatteryhealthcapacitydetail

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteUserExperienceAnalyticsBatteryHealthCapacityDetailOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteUserExperienceAnalyticsBatteryHealthCapacityDetailOperationOptions struct {
	IfMatch  *string
	Metadata *odata.Metadata
}

func DefaultDeleteUserExperienceAnalyticsBatteryHealthCapacityDetailOperationOptions() DeleteUserExperienceAnalyticsBatteryHealthCapacityDetailOperationOptions {
	return DeleteUserExperienceAnalyticsBatteryHealthCapacityDetailOperationOptions{}
}

func (o DeleteUserExperienceAnalyticsBatteryHealthCapacityDetailOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteUserExperienceAnalyticsBatteryHealthCapacityDetailOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteUserExperienceAnalyticsBatteryHealthCapacityDetailOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteUserExperienceAnalyticsBatteryHealthCapacityDetail - Delete navigation property
// userExperienceAnalyticsBatteryHealthCapacityDetails for deviceManagement
func (c UserExperienceAnalyticsBatteryHealthCapacityDetailClient) DeleteUserExperienceAnalyticsBatteryHealthCapacityDetail(ctx context.Context, options DeleteUserExperienceAnalyticsBatteryHealthCapacityDetailOperationOptions) (result DeleteUserExperienceAnalyticsBatteryHealthCapacityDetailOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
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

	return
}
