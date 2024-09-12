package remotedesktopsecurityconfigurationtargetdevicegroup

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateRemoteDesktopSecurityConfigurationTargetDeviceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// UpdateRemoteDesktopSecurityConfigurationTargetDeviceGroup - Update targetDeviceGroup. Update the properties of a
// targetDeviceGroup object for remoteDesktopSecurityConfiguration object on the servicePrincipal. You can configure a
// maximum of 10 target device groups for the remoteDesktopSecurityConfiguraiton object on the servicePrincipal.
func (c RemoteDesktopSecurityConfigurationTargetDeviceGroupClient) UpdateRemoteDesktopSecurityConfigurationTargetDeviceGroup(ctx context.Context, id stable.ServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupId, input stable.TargetDeviceGroup) (result UpdateRemoteDesktopSecurityConfigurationTargetDeviceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPatch,
		Path:       id.ID(),
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
