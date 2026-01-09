package informationprotectionsensitivitylabelsublabelright

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationProtectionSensitivityLabelSublabelRightClient struct {
	Client *msgraph.Client
}

func NewInformationProtectionSensitivityLabelSublabelRightClientWithBaseURI(sdkApi sdkEnv.Api) (*InformationProtectionSensitivityLabelSublabelRightClient, error) {
	client, err := msgraph.NewClient(sdkApi, "informationprotectionsensitivitylabelsublabelright", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InformationProtectionSensitivityLabelSublabelRightClient: %+v", err)
	}

	return &InformationProtectionSensitivityLabelSublabelRightClient{
		Client: client,
	}, nil
}
