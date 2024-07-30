package userexperienceanalyticsapphealthapplicationperformancebyappversion

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *int64
}

// GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionCount ...
func (c UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionClient) GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionCount(ctx context.Context) (result GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion/$count",
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

	var model int64
	result.Model = &model

	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
