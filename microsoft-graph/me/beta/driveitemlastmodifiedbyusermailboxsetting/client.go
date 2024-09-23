package driveitemlastmodifiedbyusermailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemLastModifiedByUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewDriveItemLastModifiedByUserMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemLastModifiedByUserMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemlastmodifiedbyusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemLastModifiedByUserMailboxSettingClient: %+v", err)
	}

	return &DriveItemLastModifiedByUserMailboxSettingClient{
		Client: client,
	}, nil
}
