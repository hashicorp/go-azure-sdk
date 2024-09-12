package exportdevicemanagementreportsjobs

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExportDeviceManagementReportsJobsClient struct {
	Client *msgraph.Client
}

func NewExportDeviceManagementReportsJobsClientWithBaseURI(sdkApi sdkEnv.Api) (*ExportDeviceManagementReportsJobsClient, error) {
	client, err := msgraph.NewClient(sdkApi, "exportdevicemanagementreportsjobs", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExportDeviceManagementReportsJobsClient: %+v", err)
	}

	return &ExportDeviceManagementReportsJobsClient{
		Client: client,
	}, nil
}
