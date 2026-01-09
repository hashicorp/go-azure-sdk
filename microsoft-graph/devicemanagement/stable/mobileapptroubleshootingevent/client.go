package mobileapptroubleshootingevent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppTroubleshootingEventClient struct {
	Client *msgraph.Client
}

func NewMobileAppTroubleshootingEventClientWithBaseURI(sdkApi sdkEnv.Api) (*MobileAppTroubleshootingEventClient, error) {
	client, err := msgraph.NewClient(sdkApi, "mobileapptroubleshootingevent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MobileAppTroubleshootingEventClient: %+v", err)
	}

	return &MobileAppTroubleshootingEventClient{
		Client: client,
	}, nil
}
