package informationprotectionsensitivitylabelsublabel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationProtectionSensitivityLabelSublabelClient struct {
	Client *msgraph.Client
}

func NewInformationProtectionSensitivityLabelSublabelClientWithBaseURI(sdkApi sdkEnv.Api) (*InformationProtectionSensitivityLabelSublabelClient, error) {
	client, err := msgraph.NewClient(sdkApi, "informationprotectionsensitivitylabelsublabel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InformationProtectionSensitivityLabelSublabelClient: %+v", err)
	}

	return &InformationProtectionSensitivityLabelSublabelClient{
		Client: client,
	}, nil
}
