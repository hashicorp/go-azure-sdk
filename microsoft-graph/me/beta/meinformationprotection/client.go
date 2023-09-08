package meinformationprotection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeInformationProtectionClient struct {
	Client *msgraph.Client
}

func NewMeInformationProtectionClientWithBaseURI(api sdkEnv.Api) (*MeInformationProtectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meinformationprotection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeInformationProtectionClient: %+v", err)
	}

	return &MeInformationProtectionClient{
		Client: client,
	}, nil
}
