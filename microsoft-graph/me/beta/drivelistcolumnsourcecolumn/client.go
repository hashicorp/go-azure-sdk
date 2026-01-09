package drivelistcolumnsourcecolumn

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveListColumnSourceColumnClient struct {
	Client *msgraph.Client
}

func NewDriveListColumnSourceColumnClientWithBaseURI(sdkApi sdkEnv.Api) (*DriveListColumnSourceColumnClient, error) {
	client, err := msgraph.NewClient(sdkApi, "drivelistcolumnsourcecolumn", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DriveListColumnSourceColumnClient: %+v", err)
	}

	return &DriveListColumnSourceColumnClient{
		Client: client,
	}, nil
}
