package importedwindowsautopilotdeviceidentity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedWindowsAutopilotDeviceIdentityClient struct {
	Client *msgraph.Client
}

func NewImportedWindowsAutopilotDeviceIdentityClientWithBaseURI(sdkApi sdkEnv.Api) (*ImportedWindowsAutopilotDeviceIdentityClient, error) {
	client, err := msgraph.NewClient(sdkApi, "importedwindowsautopilotdeviceidentity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ImportedWindowsAutopilotDeviceIdentityClient: %+v", err)
	}

	return &ImportedWindowsAutopilotDeviceIdentityClient{
		Client: client,
	}, nil
}
