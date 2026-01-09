package windowsinformationprotectionapplearningsummary

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionAppLearningSummaryClient struct {
	Client *msgraph.Client
}

func NewWindowsInformationProtectionAppLearningSummaryClientWithBaseURI(sdkApi sdkEnv.Api) (*WindowsInformationProtectionAppLearningSummaryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "windowsinformationprotectionapplearningsummary", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WindowsInformationProtectionAppLearningSummaryClient: %+v", err)
	}

	return &WindowsInformationProtectionAppLearningSummaryClient{
		Client: client,
	}, nil
}
