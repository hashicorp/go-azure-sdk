package v2024_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/policyinsights/2024-10-01/attestations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/policyinsights/2024-10-01/checkpolicyrestrictions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/policyinsights/2024-10-01/componentpolicystates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/policyinsights/2024-10-01/policyevents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/policyinsights/2024-10-01/policymetadata"
	"github.com/hashicorp/go-azure-sdk/resource-manager/policyinsights/2024-10-01/policystates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/policyinsights/2024-10-01/remediations"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Attestations            *attestations.AttestationsClient
	CheckPolicyRestrictions *checkpolicyrestrictions.CheckPolicyRestrictionsClient
	ComponentPolicyStates   *componentpolicystates.ComponentPolicyStatesClient
	PolicyEvents            *policyevents.PolicyEventsClient
	PolicyMetadata          *policymetadata.PolicyMetadataClient
	PolicyStates            *policystates.PolicyStatesClient
	Remediations            *remediations.RemediationsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	attestationsClient, err := attestations.NewAttestationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Attestations client: %+v", err)
	}
	configureFunc(attestationsClient.Client)

	checkPolicyRestrictionsClient, err := checkpolicyrestrictions.NewCheckPolicyRestrictionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CheckPolicyRestrictions client: %+v", err)
	}
	configureFunc(checkPolicyRestrictionsClient.Client)

	componentPolicyStatesClient, err := componentpolicystates.NewComponentPolicyStatesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComponentPolicyStates client: %+v", err)
	}
	configureFunc(componentPolicyStatesClient.Client)

	policyEventsClient, err := policyevents.NewPolicyEventsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyEvents client: %+v", err)
	}
	configureFunc(policyEventsClient.Client)

	policyMetadataClient, err := policymetadata.NewPolicyMetadataClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyMetadata client: %+v", err)
	}
	configureFunc(policyMetadataClient.Client)

	policyStatesClient, err := policystates.NewPolicyStatesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyStates client: %+v", err)
	}
	configureFunc(policyStatesClient.Client)

	remediationsClient, err := remediations.NewRemediationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Remediations client: %+v", err)
	}
	configureFunc(remediationsClient.Client)

	return &Client{
		Attestations:            attestationsClient,
		CheckPolicyRestrictions: checkPolicyRestrictionsClient,
		ComponentPolicyStates:   componentPolicyStatesClient,
		PolicyEvents:            policyEventsClient,
		PolicyMetadata:          policyMetadataClient,
		PolicyStates:            policyStatesClient,
		Remediations:            remediationsClient,
	}, nil
}
