package csmpublishingcredentialspoliciesentityftpallowedslot

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CsmPublishingCredentialsPoliciesEntityFtpAllowedSlotClient struct {
	Client *resourcemanager.Client
}

func NewCsmPublishingCredentialsPoliciesEntityFtpAllowedSlotClientWithBaseURI(sdkApi sdkEnv.Api) (*CsmPublishingCredentialsPoliciesEntityFtpAllowedSlotClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "csmpublishingcredentialspoliciesentityftpallowedslot", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CsmPublishingCredentialsPoliciesEntityFtpAllowedSlotClient: %+v", err)
	}

	return &CsmPublishingCredentialsPoliciesEntityFtpAllowedSlotClient{
		Client: client,
	}, nil
}
