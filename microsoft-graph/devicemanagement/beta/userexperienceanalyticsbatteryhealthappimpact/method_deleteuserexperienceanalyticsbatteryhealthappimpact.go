package userexperienceanalyticsbatteryhealthappimpact

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteUserExperienceAnalyticsBatteryHealthAppImpactOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteUserExperienceAnalyticsBatteryHealthAppImpactOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteUserExperienceAnalyticsBatteryHealthAppImpactOperationOptions() DeleteUserExperienceAnalyticsBatteryHealthAppImpactOperationOptions {
	return DeleteUserExperienceAnalyticsBatteryHealthAppImpactOperationOptions{}
}

func (o DeleteUserExperienceAnalyticsBatteryHealthAppImpactOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteUserExperienceAnalyticsBatteryHealthAppImpactOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteUserExperienceAnalyticsBatteryHealthAppImpactOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteUserExperienceAnalyticsBatteryHealthAppImpact - Delete navigation property
// userExperienceAnalyticsBatteryHealthAppImpact for deviceManagement
func (c UserExperienceAnalyticsBatteryHealthAppImpactClient) DeleteUserExperienceAnalyticsBatteryHealthAppImpact(ctx context.Context, id beta.DeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactId, options DeleteUserExperienceAnalyticsBatteryHealthAppImpactOperationOptions) (result DeleteUserExperienceAnalyticsBatteryHealthAppImpactOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
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

	return
}
