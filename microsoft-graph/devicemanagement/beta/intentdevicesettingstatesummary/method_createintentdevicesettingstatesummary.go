package intentdevicesettingstatesummary

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

type CreateIntentDeviceSettingStateSummaryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DeviceManagementIntentDeviceSettingStateSummary
}

// CreateIntentDeviceSettingStateSummary - Create new navigation property to deviceSettingStateSummaries for
// deviceManagement
func (c IntentDeviceSettingStateSummaryClient) CreateIntentDeviceSettingStateSummary(ctx context.Context, id beta.DeviceManagementIntentId, input beta.DeviceManagementIntentDeviceSettingStateSummary) (result CreateIntentDeviceSettingStateSummaryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/deviceSettingStateSummaries", id.ID()),
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

	var model beta.DeviceManagementIntentDeviceSettingStateSummary
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
