package driverootlistitemlastmodifiedbyusermailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootListItemLastModifiedByUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewDriveRootListItemLastModifiedByUserMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootListItemLastModifiedByUserMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootlistitemlastmodifiedbyusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootListItemLastModifiedByUserMailboxSettingClient: %+v", err)
	}

	return &DriveRootListItemLastModifiedByUserMailboxSettingClient{
		Client: client,
	}, nil
}
