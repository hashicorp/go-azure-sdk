package meinformationprotectionbitlocker

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeInformationProtectionBitlockerClient struct {
	Client *msgraph.Client
}

func NewMeInformationProtectionBitlockerClientWithBaseURI(api sdkEnv.Api) (*MeInformationProtectionBitlockerClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meinformationprotectionbitlocker", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeInformationProtectionBitlockerClient: %+v", err)
	}

	return &MeInformationProtectionBitlockerClient{
		Client: client,
	}, nil
}
