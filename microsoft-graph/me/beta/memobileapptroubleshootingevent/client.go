package memobileapptroubleshootingevent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeMobileAppTroubleshootingEventClient struct {
	Client *msgraph.Client
}

func NewMeMobileAppTroubleshootingEventClientWithBaseURI(api sdkEnv.Api) (*MeMobileAppTroubleshootingEventClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "memobileapptroubleshootingevent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeMobileAppTroubleshootingEventClient: %+v", err)
	}

	return &MeMobileAppTroubleshootingEventClient{
		Client: client,
	}, nil
}
