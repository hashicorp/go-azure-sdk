package directoryadministrativeunitextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryAdministrativeUnitExtensionClient struct {
	Client *msgraph.Client
}

func NewDirectoryAdministrativeUnitExtensionClientWithBaseURI(api sdkEnv.Api) (*DirectoryAdministrativeUnitExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "directoryadministrativeunitextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryAdministrativeUnitExtensionClient: %+v", err)
	}

	return &DirectoryAdministrativeUnitExtensionClient{
		Client: client,
	}, nil
}
