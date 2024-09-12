package windowsautopilotsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotSettingClient struct {
	Client *msgraph.Client
}

func NewWindowsAutopilotSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*WindowsAutopilotSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "windowsautopilotsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WindowsAutopilotSettingClient: %+v", err)
	}

	return &WindowsAutopilotSettingClient{
		Client: client,
	}, nil
}
