package androiddeviceownerenrollmentprofile

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteAndroidDeviceOwnerEnrollmentProfileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteAndroidDeviceOwnerEnrollmentProfileOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteAndroidDeviceOwnerEnrollmentProfileOperationOptions() DeleteAndroidDeviceOwnerEnrollmentProfileOperationOptions {
	return DeleteAndroidDeviceOwnerEnrollmentProfileOperationOptions{}
}

func (o DeleteAndroidDeviceOwnerEnrollmentProfileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteAndroidDeviceOwnerEnrollmentProfileOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteAndroidDeviceOwnerEnrollmentProfileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteAndroidDeviceOwnerEnrollmentProfile - Delete navigation property androidDeviceOwnerEnrollmentProfiles for
// deviceManagement
func (c AndroidDeviceOwnerEnrollmentProfileClient) DeleteAndroidDeviceOwnerEnrollmentProfile(ctx context.Context, id beta.DeviceManagementAndroidDeviceOwnerEnrollmentProfileId, options DeleteAndroidDeviceOwnerEnrollmentProfileOperationOptions) (result DeleteAndroidDeviceOwnerEnrollmentProfileOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
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

	return
}
