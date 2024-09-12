package driverootlistitemcreatedbyusermailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootListItemCreatedByUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewDriveRootListItemCreatedByUserMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootListItemCreatedByUserMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootlistitemcreatedbyusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootListItemCreatedByUserMailboxSettingClient: %+v", err)
	}

	return &DriveRootListItemCreatedByUserMailboxSettingClient{
		Client: client,
	}, nil
}
