package drivelistitemdocumentsetversionfield

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListItemDocumentSetVersionFieldClient struct {
	Client *msgraph.Client
}

func NewDriveListItemDocumentSetVersionFieldClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListItemDocumentSetVersionFieldClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistitemdocumentsetversionfield", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListItemDocumentSetVersionFieldClient: %+v", err)
	}

	return &DriveListItemDocumentSetVersionFieldClient{
		Client: client,
	}, nil
}
