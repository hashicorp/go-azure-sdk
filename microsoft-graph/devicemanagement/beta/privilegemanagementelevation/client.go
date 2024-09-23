package privilegemanagementelevation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegeManagementElevationClient struct {
	Client *msgraph.Client
}

func NewPrivilegeManagementElevationClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegeManagementElevationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegemanagementelevation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegeManagementElevationClient: %+v", err)
	}

	return &PrivilegeManagementElevationClient{
		Client: client,
	}, nil
}
