package drivelistitemcreatedbyusermailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListItemCreatedByUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewDriveListItemCreatedByUserMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListItemCreatedByUserMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistitemcreatedbyusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListItemCreatedByUserMailboxSettingClient: %+v", err)
	}

	return &DriveListItemCreatedByUserMailboxSettingClient{
		Client: client,
	}, nil
}
