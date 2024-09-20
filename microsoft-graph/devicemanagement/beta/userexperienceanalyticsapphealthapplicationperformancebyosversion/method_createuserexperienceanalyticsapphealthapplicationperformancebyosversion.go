package userexperienceanalyticsapphealthapplicationperformancebyosversion

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion
}

type CreateUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionOperationOptions() CreateUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionOperationOptions {
	return CreateUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionOperationOptions{}
}

func (o CreateUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion - Create new navigation property to
// userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion for deviceManagement
func (c UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionClient) CreateUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion(ctx context.Context, input beta.UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion, options CreateUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionOperationOptions) (result CreateUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion",
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

	var model beta.UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
