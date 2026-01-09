package mobileapptroubleshootingeventapplogcollectionrequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppTroubleshootingEventAppLogCollectionRequestClient struct {
	Client *msgraph.Client
}

func NewMobileAppTroubleshootingEventAppLogCollectionRequestClientWithBaseURI(sdkApi sdkEnv.Api) (*MobileAppTroubleshootingEventAppLogCollectionRequestClient, error) {
	client, err := msgraph.NewClient(sdkApi, "mobileapptroubleshootingeventapplogcollectionrequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MobileAppTroubleshootingEventAppLogCollectionRequestClient: %+v", err)
	}

	return &MobileAppTroubleshootingEventAppLogCollectionRequestClient{
		Client: client,
	}, nil
}
