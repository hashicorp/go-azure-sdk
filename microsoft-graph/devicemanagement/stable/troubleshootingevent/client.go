package troubleshootingevent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TroubleshootingEventClient struct {
	Client *msgraph.Client
}

func NewTroubleshootingEventClientWithBaseURI(sdkApi sdkEnv.Api) (*TroubleshootingEventClient, error) {
	client, err := msgraph.NewClient(sdkApi, "troubleshootingevent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TroubleshootingEventClient: %+v", err)
	}

	return &TroubleshootingEventClient{
		Client: client,
	}, nil
}
