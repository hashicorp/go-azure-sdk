package userexperienceanalyticsremoteconnection

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserExperienceAnalyticsRemoteConnectionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UserExperienceAnalyticsRemoteConnection
}

type GetUserExperienceAnalyticsRemoteConnectionOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetUserExperienceAnalyticsRemoteConnectionOperationOptions() GetUserExperienceAnalyticsRemoteConnectionOperationOptions {
	return GetUserExperienceAnalyticsRemoteConnectionOperationOptions{}
}

func (o GetUserExperienceAnalyticsRemoteConnectionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetUserExperienceAnalyticsRemoteConnectionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetUserExperienceAnalyticsRemoteConnectionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetUserExperienceAnalyticsRemoteConnection - Get userExperienceAnalyticsRemoteConnection from deviceManagement. User
// experience analytics remote connection. The report will be retired on December 31, 2024. You can start using the
// Cloud PC connection quality report now via
// https://learn.microsoft.com/windows-365/enterprise/report-cloud-pc-connection-quality.
func (c UserExperienceAnalyticsRemoteConnectionClient) GetUserExperienceAnalyticsRemoteConnection(ctx context.Context, id beta.DeviceManagementUserExperienceAnalyticsRemoteConnectionId, options GetUserExperienceAnalyticsRemoteConnectionOperationOptions) (result GetUserExperienceAnalyticsRemoteConnectionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
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

	var model beta.UserExperienceAnalyticsRemoteConnection
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
