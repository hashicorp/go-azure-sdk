package androiddeviceownerenrollmentprofile

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateAndroidDeviceOwnerEnrollmentProfileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateAndroidDeviceOwnerEnrollmentProfileOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateAndroidDeviceOwnerEnrollmentProfileOperationOptions() UpdateAndroidDeviceOwnerEnrollmentProfileOperationOptions {
	return UpdateAndroidDeviceOwnerEnrollmentProfileOperationOptions{}
}

func (o UpdateAndroidDeviceOwnerEnrollmentProfileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateAndroidDeviceOwnerEnrollmentProfileOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateAndroidDeviceOwnerEnrollmentProfileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateAndroidDeviceOwnerEnrollmentProfile - Update the navigation property androidDeviceOwnerEnrollmentProfiles in
// deviceManagement
func (c AndroidDeviceOwnerEnrollmentProfileClient) UpdateAndroidDeviceOwnerEnrollmentProfile(ctx context.Context, id beta.DeviceManagementAndroidDeviceOwnerEnrollmentProfileId, input beta.AndroidDeviceOwnerEnrollmentProfile, options UpdateAndroidDeviceOwnerEnrollmentProfileOperationOptions) (result UpdateAndroidDeviceOwnerEnrollmentProfileOperationResponse, err error) {
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
