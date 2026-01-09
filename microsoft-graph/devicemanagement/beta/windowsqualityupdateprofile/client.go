package windowsqualityupdateprofile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsQualityUpdateProfileClient struct {
	Client *msgraph.Client
}

func NewWindowsQualityUpdateProfileClientWithBaseURI(sdkApi sdkEnv.Api) (*WindowsQualityUpdateProfileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "windowsqualityupdateprofile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WindowsQualityUpdateProfileClient: %+v", err)
	}

	return &WindowsQualityUpdateProfileClient{
		Client: client,
	}, nil
}
