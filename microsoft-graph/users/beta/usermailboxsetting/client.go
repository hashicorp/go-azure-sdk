package usermailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewUserMailboxSettingClientWithBaseURI(api sdkEnv.Api) (*UserMailboxSettingClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserMailboxSettingClient: %+v", err)
	}

	return &UserMailboxSettingClient{
		Client: client,
	}, nil
}
