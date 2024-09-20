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

type UpdateUserExperienceAnalyticsRemoteConnectionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateUserExperienceAnalyticsRemoteConnectionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateUserExperienceAnalyticsRemoteConnectionOperationOptions() UpdateUserExperienceAnalyticsRemoteConnectionOperationOptions {
	return UpdateUserExperienceAnalyticsRemoteConnectionOperationOptions{}
}

func (o UpdateUserExperienceAnalyticsRemoteConnectionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateUserExperienceAnalyticsRemoteConnectionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateUserExperienceAnalyticsRemoteConnectionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateUserExperienceAnalyticsRemoteConnection - Update the navigation property
// userExperienceAnalyticsRemoteConnection in deviceManagement
func (c UserExperienceAnalyticsRemoteConnectionClient) UpdateUserExperienceAnalyticsRemoteConnection(ctx context.Context, id beta.DeviceManagementUserExperienceAnalyticsRemoteConnectionId, input beta.UserExperienceAnalyticsRemoteConnection, options UpdateUserExperienceAnalyticsRemoteConnectionOperationOptions) (result UpdateUserExperienceAnalyticsRemoteConnectionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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

	return
}
