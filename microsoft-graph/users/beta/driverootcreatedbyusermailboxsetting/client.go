package driverootcreatedbyusermailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootCreatedByUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewDriveRootCreatedByUserMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootCreatedByUserMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootcreatedbyusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootCreatedByUserMailboxSettingClient: %+v", err)
	}

	return &DriveRootCreatedByUserMailboxSettingClient{
		Client: client,
	}, nil
}
