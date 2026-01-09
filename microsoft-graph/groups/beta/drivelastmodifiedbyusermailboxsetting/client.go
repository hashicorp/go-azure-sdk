package drivelastmodifiedbyusermailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveLastModifiedByUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewDriveLastModifiedByUserMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveLastModifiedByUserMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelastmodifiedbyusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveLastModifiedByUserMailboxSettingClient: %+v", err)
	}

	return &DriveLastModifiedByUserMailboxSettingClient{
		Client: client,
	}, nil
}
