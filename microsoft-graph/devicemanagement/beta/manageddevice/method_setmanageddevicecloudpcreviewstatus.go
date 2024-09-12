package manageddevice

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

type SetManagedDeviceCloudPCReviewStatusOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// SetManagedDeviceCloudPCReviewStatus - Invoke action setCloudPcReviewStatus. Set the review status of a specific Cloud
// PC device. Use this API to set the review status of a Cloud PC to in review if you consider a Cloud PC as suspicious.
// After the review is completed, use this API again to set the Cloud PC back to a normal state.
func (c ManagedDeviceClient) SetManagedDeviceCloudPCReviewStatus(ctx context.Context, id beta.DeviceManagementManagedDeviceId, input SetManagedDeviceCloudPCReviewStatusRequest) (result SetManagedDeviceCloudPCReviewStatusOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/setCloudPcReviewStatus", id.ID()),
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
