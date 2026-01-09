package manageddevicecleanuprule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceCleanupRuleClient struct {
	Client *msgraph.Client
}

func NewManagedDeviceCleanupRuleClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedDeviceCleanupRuleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "manageddevicecleanuprule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDeviceCleanupRuleClient: %+v", err)
	}

	return &ManagedDeviceCleanupRuleClient{
		Client: client,
	}, nil
}
