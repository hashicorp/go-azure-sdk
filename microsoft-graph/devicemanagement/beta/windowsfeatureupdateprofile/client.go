package windowsfeatureupdateprofile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsFeatureUpdateProfileClient struct {
	Client *msgraph.Client
}

func NewWindowsFeatureUpdateProfileClientWithBaseURI(sdkApi sdkEnv.Api) (*WindowsFeatureUpdateProfileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "windowsfeatureupdateprofile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WindowsFeatureUpdateProfileClient: %+v", err)
	}

	return &WindowsFeatureUpdateProfileClient{
		Client: client,
	}, nil
}
