package groupmemberswithlicenseerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupMembersWithLicenseErrorClient struct {
	Client *msgraph.Client
}

func NewGroupMembersWithLicenseErrorClientWithBaseURI(api sdkEnv.Api) (*GroupMembersWithLicenseErrorClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupmemberswithlicenseerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupMembersWithLicenseErrorClient: %+v", err)
	}

	return &GroupMembersWithLicenseErrorClient{
		Client: client,
	}, nil
}
