package windowsinformationprotectionnetworklearningsummary

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionNetworkLearningSummaryClient struct {
	Client *msgraph.Client
}

func NewWindowsInformationProtectionNetworkLearningSummaryClientWithBaseURI(sdkApi sdkEnv.Api) (*WindowsInformationProtectionNetworkLearningSummaryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "windowsinformationprotectionnetworklearningsummary", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WindowsInformationProtectionNetworkLearningSummaryClient: %+v", err)
	}

	return &WindowsInformationProtectionNetworkLearningSummaryClient{
		Client: client,
	}, nil
}
