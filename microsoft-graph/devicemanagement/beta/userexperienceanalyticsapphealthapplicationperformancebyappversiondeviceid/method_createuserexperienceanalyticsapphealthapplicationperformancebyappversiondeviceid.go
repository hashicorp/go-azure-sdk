package userexperienceanalyticsapphealthapplicationperformancebyappversiondeviceid

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId
}

// CreateUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId ...
func (c UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdClient) CreateUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId(ctx context.Context, input beta.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) (result CreateUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId",
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

	var model beta.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId
	result.Model = &model

	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
