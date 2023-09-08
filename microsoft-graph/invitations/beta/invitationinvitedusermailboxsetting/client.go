package invitationinvitedusermailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvitationInvitedUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewInvitationInvitedUserMailboxSettingClientWithBaseURI(api sdkEnv.Api) (*InvitationInvitedUserMailboxSettingClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "invitationinvitedusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InvitationInvitedUserMailboxSettingClient: %+v", err)
	}

	return &InvitationInvitedUserMailboxSettingClient{
		Client: client,
	}, nil
}
