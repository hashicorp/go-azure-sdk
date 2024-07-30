package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/invitations/stable/invitation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/invitations/stable/inviteduser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/invitations/stable/invitedusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/invitations/stable/inviteduserserviceprovisioningerror"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	Invitation                          *invitation.InvitationClient
	InvitedUser                         *inviteduser.InvitedUserClient
	InvitedUserMailboxSetting           *invitedusermailboxsetting.InvitedUserMailboxSettingClient
	InvitedUserServiceProvisioningError *inviteduserserviceprovisioningerror.InvitedUserServiceProvisioningErrorClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	invitationClient, err := invitation.NewInvitationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Invitation client: %+v", err)
	}
	configureFunc(invitationClient.Client)

	invitedUserClient, err := inviteduser.NewInvitedUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InvitedUser client: %+v", err)
	}
	configureFunc(invitedUserClient.Client)

	invitedUserMailboxSettingClient, err := invitedusermailboxsetting.NewInvitedUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InvitedUserMailboxSetting client: %+v", err)
	}
	configureFunc(invitedUserMailboxSettingClient.Client)

	invitedUserServiceProvisioningErrorClient, err := inviteduserserviceprovisioningerror.NewInvitedUserServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InvitedUserServiceProvisioningError client: %+v", err)
	}
	configureFunc(invitedUserServiceProvisioningErrorClient.Client)

	return &Client{
		Invitation:                          invitationClient,
		InvitedUser:                         invitedUserClient,
		InvitedUserMailboxSetting:           invitedUserMailboxSettingClient,
		InvitedUserServiceProvisioningError: invitedUserServiceProvisioningErrorClient,
	}, nil
}
