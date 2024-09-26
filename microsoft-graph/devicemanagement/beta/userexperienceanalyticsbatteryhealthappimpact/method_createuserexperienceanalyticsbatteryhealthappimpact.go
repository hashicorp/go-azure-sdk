package userexperienceanalyticsbatteryhealthappimpact

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserExperienceAnalyticsBatteryHealthAppImpactOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UserExperienceAnalyticsBatteryHealthAppImpact
}

type CreateUserExperienceAnalyticsBatteryHealthAppImpactOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateUserExperienceAnalyticsBatteryHealthAppImpactOperationOptions() CreateUserExperienceAnalyticsBatteryHealthAppImpactOperationOptions {
	return CreateUserExperienceAnalyticsBatteryHealthAppImpactOperationOptions{}
}

func (o CreateUserExperienceAnalyticsBatteryHealthAppImpactOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateUserExperienceAnalyticsBatteryHealthAppImpactOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateUserExperienceAnalyticsBatteryHealthAppImpactOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateUserExperienceAnalyticsBatteryHealthAppImpact - Create new navigation property to
// userExperienceAnalyticsBatteryHealthAppImpact for deviceManagement
func (c UserExperienceAnalyticsBatteryHealthAppImpactClient) CreateUserExperienceAnalyticsBatteryHealthAppImpact(ctx context.Context, input beta.UserExperienceAnalyticsBatteryHealthAppImpact, options CreateUserExperienceAnalyticsBatteryHealthAppImpactOperationOptions) (result CreateUserExperienceAnalyticsBatteryHealthAppImpactOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/userExperienceAnalyticsBatteryHealthAppImpact",
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

	var model beta.UserExperienceAnalyticsBatteryHealthAppImpact
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
