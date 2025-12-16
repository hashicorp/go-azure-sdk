package diagnosticremotesupportsettingsoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiagnosticRemoteSupportSettingsOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewDiagnosticRemoteSupportSettingsOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*DiagnosticRemoteSupportSettingsOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "diagnosticremotesupportsettingsoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DiagnosticRemoteSupportSettingsOperationGroupClient: %+v", err)
	}

	return &DiagnosticRemoteSupportSettingsOperationGroupClient{
		Client: client,
	}, nil
}
