package userexperienceanalyticsapphealthapplicationperformancebyappversiondetail

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails
}

type CreateUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailOperationOptions() CreateUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailOperationOptions {
	return CreateUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailOperationOptions{}
}

func (o CreateUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetail - Create new navigation property to
// userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails for deviceManagement
func (c UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailClient) CreateUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetail(ctx context.Context, input beta.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails, options CreateUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailOperationOptions) (result CreateUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails",
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

	var model beta.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
