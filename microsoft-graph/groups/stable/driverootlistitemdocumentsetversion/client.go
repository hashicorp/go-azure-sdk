package driverootlistitemdocumentsetversion

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootListItemDocumentSetVersionClient struct {
	Client *msgraph.Client
}

func NewDriveRootListItemDocumentSetVersionClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootListItemDocumentSetVersionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootlistitemdocumentsetversion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootListItemDocumentSetVersionClient: %+v", err)
	}

	return &DriveRootListItemDocumentSetVersionClient{
		Client: client,
	}, nil
}
