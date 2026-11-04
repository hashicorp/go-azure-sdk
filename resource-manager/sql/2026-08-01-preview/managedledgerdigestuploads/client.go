package managedledgerdigestuploads

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedLedgerDigestUploadsClient struct {
	Client *resourcemanager.Client
}

func NewManagedLedgerDigestUploadsClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedLedgerDigestUploadsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "managedledgerdigestuploads", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedLedgerDigestUploadsClient: %+v", err)
	}

	return &ManagedLedgerDigestUploadsClient{
		Client: client,
	}, nil
}
