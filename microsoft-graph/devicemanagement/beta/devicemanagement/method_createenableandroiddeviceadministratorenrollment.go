package devicemanagement

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateEnableAndroidDeviceAdministratorEnrollmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// CreateEnableAndroidDeviceAdministratorEnrollment ...
func (c DeviceManagementClient) CreateEnableAndroidDeviceAdministratorEnrollment(ctx context.Context) (result CreateEnableAndroidDeviceAdministratorEnrollmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       "/deviceManagement/enableAndroidDeviceAdministratorEnrollment",
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
