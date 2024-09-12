package userexperienceanalyticsmetrichistory

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserExperienceAnalyticsMetricHistoryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UserExperienceAnalyticsMetricHistory
}

type GetUserExperienceAnalyticsMetricHistoryOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetUserExperienceAnalyticsMetricHistoryOperationOptions() GetUserExperienceAnalyticsMetricHistoryOperationOptions {
	return GetUserExperienceAnalyticsMetricHistoryOperationOptions{}
}

func (o GetUserExperienceAnalyticsMetricHistoryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetUserExperienceAnalyticsMetricHistoryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetUserExperienceAnalyticsMetricHistoryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetUserExperienceAnalyticsMetricHistory - Get userExperienceAnalyticsMetricHistory from deviceManagement. User
// experience analytics metric history
func (c UserExperienceAnalyticsMetricHistoryClient) GetUserExperienceAnalyticsMetricHistory(ctx context.Context, id stable.DeviceManagementUserExperienceAnalyticsMetricHistoryId, options GetUserExperienceAnalyticsMetricHistoryOperationOptions) (result GetUserExperienceAnalyticsMetricHistoryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model stable.UserExperienceAnalyticsMetricHistory
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
