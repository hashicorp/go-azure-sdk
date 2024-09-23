package driverootlastmodifiedbyusermailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootLastModifiedByUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewDriveRootLastModifiedByUserMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootLastModifiedByUserMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootlastmodifiedbyusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootLastModifiedByUserMailboxSettingClient: %+v", err)
	}

	return &DriveRootLastModifiedByUserMailboxSettingClient{
		Client: client,
	}, nil
}
