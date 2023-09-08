package directoryadministrativeunitscopedrolemember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryAdministrativeUnitScopedRoleMemberClient struct {
	Client *msgraph.Client
}

func NewDirectoryAdministrativeUnitScopedRoleMemberClientWithBaseURI(api sdkEnv.Api) (*DirectoryAdministrativeUnitScopedRoleMemberClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "directoryadministrativeunitscopedrolemember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryAdministrativeUnitScopedRoleMemberClient: %+v", err)
	}

	return &DirectoryAdministrativeUnitScopedRoleMemberClient{
		Client: client,
	}, nil
}
