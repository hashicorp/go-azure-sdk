package userexperienceanalyticsdevicestartuphistory

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserExperienceAnalyticsDeviceStartupHistoryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UserExperienceAnalyticsDeviceStartupHistory
}

type CreateUserExperienceAnalyticsDeviceStartupHistoryOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateUserExperienceAnalyticsDeviceStartupHistoryOperationOptions() CreateUserExperienceAnalyticsDeviceStartupHistoryOperationOptions {
	return CreateUserExperienceAnalyticsDeviceStartupHistoryOperationOptions{}
}

func (o CreateUserExperienceAnalyticsDeviceStartupHistoryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateUserExperienceAnalyticsDeviceStartupHistoryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateUserExperienceAnalyticsDeviceStartupHistoryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateUserExperienceAnalyticsDeviceStartupHistory - Create new navigation property to
// userExperienceAnalyticsDeviceStartupHistory for deviceManagement
func (c UserExperienceAnalyticsDeviceStartupHistoryClient) CreateUserExperienceAnalyticsDeviceStartupHistory(ctx context.Context, input stable.UserExperienceAnalyticsDeviceStartupHistory, options CreateUserExperienceAnalyticsDeviceStartupHistoryOperationOptions) (result CreateUserExperienceAnalyticsDeviceStartupHistoryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/userExperienceAnalyticsDeviceStartupHistory",
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

	var model stable.UserExperienceAnalyticsDeviceStartupHistory
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
