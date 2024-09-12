package exportdevicemanagementvirtualendpointreportsjobs

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExportDeviceManagementVirtualEndpointReportsJobsClient struct {
	Client *msgraph.Client
}

func NewExportDeviceManagementVirtualEndpointReportsJobsClientWithBaseURI(sdkApi sdkEnv.Api) (*ExportDeviceManagementVirtualEndpointReportsJobsClient, error) {
	client, err := msgraph.NewClient(sdkApi, "exportdevicemanagementvirtualendpointreportsjobs", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExportDeviceManagementVirtualEndpointReportsJobsClient: %+v", err)
	}

	return &ExportDeviceManagementVirtualEndpointReportsJobsClient{
		Client: client,
	}, nil
}
