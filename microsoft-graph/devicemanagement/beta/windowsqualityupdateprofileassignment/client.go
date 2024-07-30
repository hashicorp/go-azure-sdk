package windowsqualityupdateprofileassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsQualityUpdateProfileAssignmentClient struct {
	Client *msgraph.Client
}

func NewWindowsQualityUpdateProfileAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*WindowsQualityUpdateProfileAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "windowsqualityupdateprofileassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WindowsQualityUpdateProfileAssignmentClient: %+v", err)
	}

	return &WindowsQualityUpdateProfileAssignmentClient{
		Client: client,
	}, nil
}
