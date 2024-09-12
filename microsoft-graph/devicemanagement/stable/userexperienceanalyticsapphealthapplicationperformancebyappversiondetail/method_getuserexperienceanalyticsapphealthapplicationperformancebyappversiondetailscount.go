package userexperienceanalyticsapphealthapplicationperformancebyappversiondetail

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCountOperationOptions struct {
	Filter *string
	Search *string
}

func DefaultGetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCountOperationOptions() GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCountOperationOptions {
	return GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCountOperationOptions{}
}

func (o GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	return &out
}

func (o GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCount - Get the number of the resource
func (c UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailClient) GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCount(ctx context.Context, options GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCountOperationOptions) (result GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails/$count",
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
