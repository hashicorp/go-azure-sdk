package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/administrativeunits/beta/administrativeunit"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/administrativeunits/beta/deletedmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/administrativeunits/beta/extension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/administrativeunits/beta/member"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/administrativeunits/beta/scopedrolemember"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AdministrativeUnit *administrativeunit.AdministrativeUnitClient
	DeletedMember      *deletedmember.DeletedMemberClient
	Extension          *extension.ExtensionClient
	Member             *member.MemberClient
	ScopedRoleMember   *scopedrolemember.ScopedRoleMemberClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	administrativeUnitClient, err := administrativeunit.NewAdministrativeUnitClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AdministrativeUnit client: %+v", err)
	}
	configureFunc(administrativeUnitClient.Client)

	deletedMemberClient, err := deletedmember.NewDeletedMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeletedMember client: %+v", err)
	}
	configureFunc(deletedMemberClient.Client)

	extensionClient, err := extension.NewExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Extension client: %+v", err)
	}
	configureFunc(extensionClient.Client)

	memberClient, err := member.NewMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Member client: %+v", err)
	}
	configureFunc(memberClient.Client)

	scopedRoleMemberClient, err := scopedrolemember.NewScopedRoleMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ScopedRoleMember client: %+v", err)
	}
	configureFunc(scopedRoleMemberClient.Client)

	return &Client{
		AdministrativeUnit: administrativeUnitClient,
		DeletedMember:      deletedMemberClient,
		Extension:          extensionClient,
		Member:             memberClient,
		ScopedRoleMember:   scopedRoleMemberClient,
	}, nil
}
