package usermobileapptroubleshootingevent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserMobileAppTroubleshootingEventClient struct {
	Client *msgraph.Client
}

func NewUserMobileAppTroubleshootingEventClientWithBaseURI(api sdkEnv.Api) (*UserMobileAppTroubleshootingEventClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermobileapptroubleshootingevent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserMobileAppTroubleshootingEventClient: %+v", err)
	}

	return &UserMobileAppTroubleshootingEventClient{
		Client: client,
	}, nil
}
