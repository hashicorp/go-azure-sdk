package driveitemcreatedbyusermailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemCreatedByUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewDriveItemCreatedByUserMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemCreatedByUserMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemcreatedbyusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemCreatedByUserMailboxSettingClient: %+v", err)
	}

	return &DriveItemCreatedByUserMailboxSettingClient{
		Client: client,
	}, nil
}
