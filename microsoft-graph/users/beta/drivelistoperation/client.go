package drivelistoperation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListOperationClient struct {
	Client *msgraph.Client
}

func NewDriveListOperationClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListOperationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistoperation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListOperationClient: %+v", err)
	}

	return &DriveListOperationClient{
		Client: client,
	}, nil
}
