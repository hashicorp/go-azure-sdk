package driverootlistitemdocumentsetversionfield

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRootListItemDocumentSetVersionFieldClient struct {
	Client *msgraph.Client
}

func NewDriveRootListItemDocumentSetVersionFieldClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveRootListItemDocumentSetVersionFieldClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driverootlistitemdocumentsetversionfield", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveRootListItemDocumentSetVersionFieldClient: %+v", err)
	}

	return &DriveRootListItemDocumentSetVersionFieldClient{
		Client: client,
	}, nil
}
