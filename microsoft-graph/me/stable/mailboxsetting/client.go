package mailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailboxSettingClient struct {
	Client *msgraph.Client
}

func NewMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*MailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "mailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MailboxSettingClient: %+v", err)
	}

	return &MailboxSettingClient{
		Client: client,
	}, nil
}
