package driveitemchildcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemChildContentClient struct {
	Client *msgraph.Client
}

func NewDriveItemChildContentClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemChildContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemchildcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemChildContentClient: %+v", err)
	}

	return &DriveItemChildContentClient{
		Client: client,
	}, nil
}
