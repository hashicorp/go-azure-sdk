package driveitemlistitemdocumentsetversion

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemListItemDocumentSetVersionClient struct {
	Client *msgraph.Client
}

func NewDriveItemListItemDocumentSetVersionClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemListItemDocumentSetVersionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemlistitemdocumentsetversion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemListItemDocumentSetVersionClient: %+v", err)
	}

	return &DriveItemListItemDocumentSetVersionClient{
		Client: client,
	}, nil
}
