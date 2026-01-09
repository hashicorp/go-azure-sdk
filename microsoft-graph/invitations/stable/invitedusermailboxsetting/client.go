package invitedusermailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvitedUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewInvitedUserMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*InvitedUserMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "invitedusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InvitedUserMailboxSettingClient: %+v", err)
	}

	return &InvitedUserMailboxSettingClient{
		Client: client,
	}, nil
}
