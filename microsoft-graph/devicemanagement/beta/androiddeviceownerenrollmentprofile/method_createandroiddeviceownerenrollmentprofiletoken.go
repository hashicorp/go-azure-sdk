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

type CreateAndroidDeviceOwnerEnrollmentProfileTokenOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateAndroidDeviceOwnerEnrollmentProfileTokenOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateAndroidDeviceOwnerEnrollmentProfileTokenOperationOptions() CreateAndroidDeviceOwnerEnrollmentProfileTokenOperationOptions {
	return CreateAndroidDeviceOwnerEnrollmentProfileTokenOperationOptions{}
}

func (o CreateAndroidDeviceOwnerEnrollmentProfileTokenOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAndroidDeviceOwnerEnrollmentProfileTokenOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAndroidDeviceOwnerEnrollmentProfileTokenOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAndroidDeviceOwnerEnrollmentProfileToken - Invoke action createToken
func (c AndroidDeviceOwnerEnrollmentProfileClient) CreateAndroidDeviceOwnerEnrollmentProfileToken(ctx context.Context, id beta.DeviceManagementAndroidDeviceOwnerEnrollmentProfileId, input CreateAndroidDeviceOwnerEnrollmentProfileTokenRequest, options CreateAndroidDeviceOwnerEnrollmentProfileTokenOperationOptions) (result CreateAndroidDeviceOwnerEnrollmentProfileTokenOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/createToken", id.ID()),
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
