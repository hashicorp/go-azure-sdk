package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/administrativeunits/beta/administrativeunit"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/administrativeunits/beta/administrativeunitextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/administrativeunits/beta/administrativeunitmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/administrativeunits/beta/administrativeunitscopedrolemember"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
)

type Client struct {
	AdministrativeUnit                 *administrativeunit.AdministrativeUnitClient
	AdministrativeUnitExtension        *administrativeunitextension.AdministrativeUnitExtensionClient
	AdministrativeUnitMember           *administrativeunitmember.AdministrativeUnitMemberClient
	AdministrativeUnitScopedRoleMember *administrativeunitscopedrolemember.AdministrativeUnitScopedRoleMemberClient
}

func NewClientWithBaseURI(api skdEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	administrativeUnitClient, err := administrativeunit.NewAdministrativeUnitClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AdministrativeUnit client: %+v", err)
	}
	configureFunc(administrativeUnitClient.Client)

	administrativeUnitExtensionClient, err := administrativeunitextension.NewAdministrativeUnitExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AdministrativeUnitExtension client: %+v", err)
	}
	configureFunc(administrativeUnitExtensionClient.Client)

	administrativeUnitMemberClient, err := administrativeunitmember.NewAdministrativeUnitMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AdministrativeUnitMember client: %+v", err)
	}
	configureFunc(administrativeUnitMemberClient.Client)

	administrativeUnitScopedRoleMemberClient, err := administrativeunitscopedrolemember.NewAdministrativeUnitScopedRoleMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AdministrativeUnitScopedRoleMember client: %+v", err)
	}
	configureFunc(administrativeUnitScopedRoleMemberClient.Client)

	return &Client{
		AdministrativeUnit:                 administrativeUnitClient,
		AdministrativeUnitExtension:        administrativeUnitExtensionClient,
		AdministrativeUnitMember:           administrativeUnitMemberClient,
		AdministrativeUnitScopedRoleMember: administrativeUnitScopedRoleMemberClient,
	}, nil
}
