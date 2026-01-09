package drivelistitemdocumentsetversion

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListItemDocumentSetVersionClient struct {
	Client *msgraph.Client
}

func NewDriveListItemDocumentSetVersionClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListItemDocumentSetVersionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistitemdocumentsetversion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListItemDocumentSetVersionClient: %+v", err)
	}

	return &DriveListItemDocumentSetVersionClient{
		Client: client,
	}, nil
}
