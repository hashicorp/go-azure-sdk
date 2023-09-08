package v1_0

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/invitations/v1_0/invitation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/invitations/v1_0/invitationinviteduser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/invitations/v1_0/invitationinvitedusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
)

type Client struct {
	Invitation                          *invitation.InvitationClient
	InvitationInvitedUser               *invitationinviteduser.InvitationInvitedUserClient
	InvitationInvitedUserMailboxSetting *invitationinvitedusermailboxsetting.InvitationInvitedUserMailboxSettingClient
}

func NewClientWithBaseURI(api skdEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	invitationClient, err := invitation.NewInvitationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Invitation client: %+v", err)
	}
	configureFunc(invitationClient.Client)

	invitationInvitedUserClient, err := invitationinviteduser.NewInvitationInvitedUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InvitationInvitedUser client: %+v", err)
	}
	configureFunc(invitationInvitedUserClient.Client)

	invitationInvitedUserMailboxSettingClient, err := invitationinvitedusermailboxsetting.NewInvitationInvitedUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InvitationInvitedUserMailboxSetting client: %+v", err)
	}
	configureFunc(invitationInvitedUserMailboxSettingClient.Client)

	return &Client{
		Invitation:                          invitationClient,
		InvitationInvitedUser:               invitationInvitedUserClient,
		InvitationInvitedUserMailboxSetting: invitationInvitedUserMailboxSettingClient,
	}, nil
}
