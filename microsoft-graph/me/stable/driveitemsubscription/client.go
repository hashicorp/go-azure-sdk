package driveitemsubscription

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemSubscriptionClient struct {
	Client *msgraph.Client
}

func NewDriveItemSubscriptionClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemSubscriptionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemsubscription", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemSubscriptionClient: %+v", err)
	}

	return &DriveItemSubscriptionClient{
		Client: client,
	}, nil
}
