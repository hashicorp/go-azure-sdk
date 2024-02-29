package managementgroupdiagnosticsettings

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagementGroupDiagnosticSettingsClient struct {
	Client *resourcemanager.Client
}

func NewManagementGroupDiagnosticSettingsClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagementGroupDiagnosticSettingsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "managementgroupdiagnosticsettings", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagementGroupDiagnosticSettingsClient: %+v", err)
	}

	return &ManagementGroupDiagnosticSettingsClient{
		Client: client,
	}, nil
}
