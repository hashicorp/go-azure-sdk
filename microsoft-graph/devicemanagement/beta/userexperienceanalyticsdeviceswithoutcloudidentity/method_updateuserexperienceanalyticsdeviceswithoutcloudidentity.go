package userexperienceanalyticsdeviceswithoutcloudidentity

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateUserExperienceAnalyticsDevicesWithoutCloudIdentityOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateUserExperienceAnalyticsDevicesWithoutCloudIdentityOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateUserExperienceAnalyticsDevicesWithoutCloudIdentityOperationOptions() UpdateUserExperienceAnalyticsDevicesWithoutCloudIdentityOperationOptions {
	return UpdateUserExperienceAnalyticsDevicesWithoutCloudIdentityOperationOptions{}
}

func (o UpdateUserExperienceAnalyticsDevicesWithoutCloudIdentityOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateUserExperienceAnalyticsDevicesWithoutCloudIdentityOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateUserExperienceAnalyticsDevicesWithoutCloudIdentityOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateUserExperienceAnalyticsDevicesWithoutCloudIdentity - Update the navigation property
// userExperienceAnalyticsDevicesWithoutCloudIdentity in deviceManagement
func (c UserExperienceAnalyticsDevicesWithoutCloudIdentityClient) UpdateUserExperienceAnalyticsDevicesWithoutCloudIdentity(ctx context.Context, id beta.DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityId, input beta.UserExperienceAnalyticsDeviceWithoutCloudIdentity, options UpdateUserExperienceAnalyticsDevicesWithoutCloudIdentityOperationOptions) (result UpdateUserExperienceAnalyticsDevicesWithoutCloudIdentityOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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

	return
}
