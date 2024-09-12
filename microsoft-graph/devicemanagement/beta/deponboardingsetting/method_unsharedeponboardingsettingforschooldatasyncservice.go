package deponboardingsetting

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

type UnshareDepOnboardingSettingForSchoolDataSyncServiceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// UnshareDepOnboardingSettingForSchoolDataSyncService - Invoke action unshareForSchoolDataSyncService
func (c DepOnboardingSettingClient) UnshareDepOnboardingSettingForSchoolDataSyncService(ctx context.Context, id beta.DeviceManagementDepOnboardingSettingId) (result UnshareDepOnboardingSettingForSchoolDataSyncServiceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/unshareForSchoolDataSyncService", id.ID()),
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
