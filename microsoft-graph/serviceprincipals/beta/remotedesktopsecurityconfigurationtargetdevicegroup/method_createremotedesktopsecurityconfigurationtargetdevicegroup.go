package remotedesktopsecurityconfigurationtargetdevicegroup

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

type CreateRemoteDesktopSecurityConfigurationTargetDeviceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.TargetDeviceGroup
}

// CreateRemoteDesktopSecurityConfigurationTargetDeviceGroup - Create targetDeviceGroup. Create a new targetDeviceGroup
// object for the remoteDesktopSecurityConfiguration object on the servicePrincipal. You can configure a maximum of 10
// target device groups for the remoteDesktopSecurityConfiguraiton object on the servicePrincipal.
func (c RemoteDesktopSecurityConfigurationTargetDeviceGroupClient) CreateRemoteDesktopSecurityConfigurationTargetDeviceGroup(ctx context.Context, id beta.ServicePrincipalId, input beta.TargetDeviceGroup) (result CreateRemoteDesktopSecurityConfigurationTargetDeviceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/remoteDesktopSecurityConfiguration/targetDeviceGroups", id.ID()),
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

	var model beta.TargetDeviceGroup
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
