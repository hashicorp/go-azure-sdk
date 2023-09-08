package directoryfederationconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryFederationConfigurationClient struct {
	Client *msgraph.Client
}

func NewDirectoryFederationConfigurationClientWithBaseURI(api sdkEnv.Api) (*DirectoryFederationConfigurationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "directoryfederationconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryFederationConfigurationClient: %+v", err)
	}

	return &DirectoryFederationConfigurationClient{
		Client: client,
	}, nil
}
