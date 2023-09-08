package usermobileapptroubleshootingeventapplogcollectionrequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserMobileAppTroubleshootingEventAppLogCollectionRequestClient struct {
	Client *msgraph.Client
}

func NewUserMobileAppTroubleshootingEventAppLogCollectionRequestClientWithBaseURI(api sdkEnv.Api) (*UserMobileAppTroubleshootingEventAppLogCollectionRequestClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermobileapptroubleshootingeventapplogcollectionrequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserMobileAppTroubleshootingEventAppLogCollectionRequestClient: %+v", err)
	}

	return &UserMobileAppTroubleshootingEventAppLogCollectionRequestClient{
		Client: client,
	}, nil
}
