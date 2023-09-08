package groupdrive

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupDriveClient struct {
	Client *msgraph.Client
}

func NewGroupDriveClientWithBaseURI(api sdkEnv.Api) (*GroupDriveClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupdrive", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupDriveClient: %+v", err)
	}

	return &GroupDriveClient{
		Client: client,
	}, nil
}
