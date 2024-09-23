package drivelistcreatedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListCreatedByUserClient struct {
	Client *msgraph.Client
}

func NewDriveListCreatedByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListCreatedByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistcreatedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListCreatedByUserClient: %+v", err)
	}

	return &DriveListCreatedByUserClient{
		Client: client,
	}, nil
}
