package driveitemlistitemdocumentsetversionfield

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemListItemDocumentSetVersionFieldClient struct {
	Client *msgraph.Client
}

func NewDriveItemListItemDocumentSetVersionFieldClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveItemListItemDocumentSetVersionFieldClient, error) {
	client, err := msgraph.NewClient(sdkApi, "driveitemlistitemdocumentsetversionfield", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveItemListItemDocumentSetVersionFieldClient: %+v", err)
	}

	return &DriveItemListItemDocumentSetVersionFieldClient{
		Client: client,
	}, nil
}
