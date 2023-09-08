package directoryadministrativeunit

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryAdministrativeUnitClient struct {
	Client *msgraph.Client
}

func NewDirectoryAdministrativeUnitClientWithBaseURI(api sdkEnv.Api) (*DirectoryAdministrativeUnitClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "directoryadministrativeunit", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryAdministrativeUnitClient: %+v", err)
	}

	return &DirectoryAdministrativeUnitClient{
		Client: client,
	}, nil
}
