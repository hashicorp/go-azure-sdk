package sourcecontrolsyncjobstreams

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SourceControlSyncJobStreamsClient struct {
	Client *resourcemanager.Client
}

func NewSourceControlSyncJobStreamsClientWithBaseURI(api environments.Api) (*SourceControlSyncJobStreamsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "sourcecontrolsyncjobstreams", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SourceControlSyncJobStreamsClient: %+v", err)
	}

	return &SourceControlSyncJobStreamsClient{
		Client: client,
	}, nil
}
