package memailfolderuserconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeMailFolderUserConfigurationClient struct {
	Client *msgraph.Client
}

func NewMeMailFolderUserConfigurationClientWithBaseURI(api sdkEnv.Api) (*MeMailFolderUserConfigurationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "memailfolderuserconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeMailFolderUserConfigurationClient: %+v", err)
	}

	return &MeMailFolderUserConfigurationClient{
		Client: client,
	}, nil
}
