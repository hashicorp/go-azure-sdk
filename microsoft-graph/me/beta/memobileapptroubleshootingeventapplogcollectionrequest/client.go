package memobileapptroubleshootingeventapplogcollectionrequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeMobileAppTroubleshootingEventAppLogCollectionRequestClient struct {
	Client *msgraph.Client
}

func NewMeMobileAppTroubleshootingEventAppLogCollectionRequestClientWithBaseURI(api sdkEnv.Api) (*MeMobileAppTroubleshootingEventAppLogCollectionRequestClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "memobileapptroubleshootingeventapplogcollectionrequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeMobileAppTroubleshootingEventAppLogCollectionRequestClient: %+v", err)
	}

	return &MeMobileAppTroubleshootingEventAppLogCollectionRequestClient{
		Client: client,
	}, nil
}
