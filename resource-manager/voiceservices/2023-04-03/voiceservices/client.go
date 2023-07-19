package voiceservices

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VoiceservicesClient struct {
	Client *resourcemanager.Client
}

func NewVoiceservicesClientWithBaseURI(api environments.Api) (*VoiceservicesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "voiceservices", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VoiceservicesClient: %+v", err)
	}

	return &VoiceservicesClient{
		Client: client,
	}, nil
}
