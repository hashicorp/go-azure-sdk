package v1_0

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/v1_0/identity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/v1_0/identityapiconnector"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/v1_0/identityb2xuserflow"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/v1_0/identityb2xuserflowapiconnectorconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/v1_0/identityb2xuserflowapiconnectorconfigurationpostattributecollection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/v1_0/identityb2xuserflowapiconnectorconfigurationpostfederationsignup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/v1_0/identityb2xuserflowidentityprovider"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/v1_0/identityb2xuserflowlanguage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/v1_0/identityb2xuserflowlanguagedefaultpage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/v1_0/identityb2xuserflowlanguageoverridespage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/v1_0/identityb2xuserflowuserattributeassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/v1_0/identityb2xuserflowuserattributeassignmentuserattribute"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/v1_0/identityb2xuserflowuserflowidentityprovider"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/v1_0/identityconditionalacces"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/v1_0/identityconditionalaccesauthenticationcontextclassreference"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/v1_0/identityconditionalaccesauthenticationstrength"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/v1_0/identityconditionalaccesauthenticationstrengthauthenticationmethodmode"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/v1_0/identityconditionalaccesauthenticationstrengthpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/v1_0/identityconditionalaccesauthenticationstrengthpolicycombinationconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/v1_0/identityconditionalaccesnamedlocation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/v1_0/identityconditionalaccespolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/v1_0/identityconditionalaccestemplate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/v1_0/identityidentityprovider"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/v1_0/identityuserflowattribute"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
)

type Client struct {
	Identity                                                                     *identity.IdentityClient
	IdentityApiConnector                                                         *identityapiconnector.IdentityApiConnectorClient
	IdentityB2xUserFlow                                                          *identityb2xuserflow.IdentityB2xUserFlowClient
	IdentityB2xUserFlowApiConnectorConfiguration                                 *identityb2xuserflowapiconnectorconfiguration.IdentityB2xUserFlowApiConnectorConfigurationClient
	IdentityB2xUserFlowApiConnectorConfigurationPostAttributeCollection          *identityb2xuserflowapiconnectorconfigurationpostattributecollection.IdentityB2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient
	IdentityB2xUserFlowApiConnectorConfigurationPostFederationSignup             *identityb2xuserflowapiconnectorconfigurationpostfederationsignup.IdentityB2xUserFlowApiConnectorConfigurationPostFederationSignupClient
	IdentityB2xUserFlowIdentityProvider                                          *identityb2xuserflowidentityprovider.IdentityB2xUserFlowIdentityProviderClient
	IdentityB2xUserFlowLanguage                                                  *identityb2xuserflowlanguage.IdentityB2xUserFlowLanguageClient
	IdentityB2xUserFlowLanguageDefaultPage                                       *identityb2xuserflowlanguagedefaultpage.IdentityB2xUserFlowLanguageDefaultPageClient
	IdentityB2xUserFlowLanguageOverridesPage                                     *identityb2xuserflowlanguageoverridespage.IdentityB2xUserFlowLanguageOverridesPageClient
	IdentityB2xUserFlowUserAttributeAssignment                                   *identityb2xuserflowuserattributeassignment.IdentityB2xUserFlowUserAttributeAssignmentClient
	IdentityB2xUserFlowUserAttributeAssignmentUserAttribute                      *identityb2xuserflowuserattributeassignmentuserattribute.IdentityB2xUserFlowUserAttributeAssignmentUserAttributeClient
	IdentityB2xUserFlowUserFlowIdentityProvider                                  *identityb2xuserflowuserflowidentityprovider.IdentityB2xUserFlowUserFlowIdentityProviderClient
	IdentityConditionalAcces                                                     *identityconditionalacces.IdentityConditionalAccesClient
	IdentityConditionalAccesAuthenticationContextClassReference                  *identityconditionalaccesauthenticationcontextclassreference.IdentityConditionalAccesAuthenticationContextClassReferenceClient
	IdentityConditionalAccesAuthenticationStrength                               *identityconditionalaccesauthenticationstrength.IdentityConditionalAccesAuthenticationStrengthClient
	IdentityConditionalAccesAuthenticationStrengthAuthenticationMethodMode       *identityconditionalaccesauthenticationstrengthauthenticationmethodmode.IdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeClient
	IdentityConditionalAccesAuthenticationStrengthPolicy                         *identityconditionalaccesauthenticationstrengthpolicy.IdentityConditionalAccesAuthenticationStrengthPolicyClient
	IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfiguration *identityconditionalaccesauthenticationstrengthpolicycombinationconfiguration.IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationClient
	IdentityConditionalAccesNamedLocation                                        *identityconditionalaccesnamedlocation.IdentityConditionalAccesNamedLocationClient
	IdentityConditionalAccesPolicy                                               *identityconditionalaccespolicy.IdentityConditionalAccesPolicyClient
	IdentityConditionalAccesTemplate                                             *identityconditionalaccestemplate.IdentityConditionalAccesTemplateClient
	IdentityIdentityProvider                                                     *identityidentityprovider.IdentityIdentityProviderClient
	IdentityUserFlowAttribute                                                    *identityuserflowattribute.IdentityUserFlowAttributeClient
}

func NewClientWithBaseURI(api skdEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	identityApiConnectorClient, err := identityapiconnector.NewIdentityApiConnectorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IdentityApiConnector client: %+v", err)
	}
	configureFunc(identityApiConnectorClient.Client)

	identityB2xUserFlowApiConnectorConfigurationClient, err := identityb2xuserflowapiconnectorconfiguration.NewIdentityB2xUserFlowApiConnectorConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IdentityB2xUserFlowApiConnectorConfiguration client: %+v", err)
	}
	configureFunc(identityB2xUserFlowApiConnectorConfigurationClient.Client)

	identityB2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient, err := identityb2xuserflowapiconnectorconfigurationpostattributecollection.NewIdentityB2xUserFlowApiConnectorConfigurationPostAttributeCollectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IdentityB2xUserFlowApiConnectorConfigurationPostAttributeCollection client: %+v", err)
	}
	configureFunc(identityB2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient.Client)

	identityB2xUserFlowApiConnectorConfigurationPostFederationSignupClient, err := identityb2xuserflowapiconnectorconfigurationpostfederationsignup.NewIdentityB2xUserFlowApiConnectorConfigurationPostFederationSignupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IdentityB2xUserFlowApiConnectorConfigurationPostFederationSignup client: %+v", err)
	}
	configureFunc(identityB2xUserFlowApiConnectorConfigurationPostFederationSignupClient.Client)

	identityB2xUserFlowClient, err := identityb2xuserflow.NewIdentityB2xUserFlowClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IdentityB2xUserFlow client: %+v", err)
	}
	configureFunc(identityB2xUserFlowClient.Client)

	identityB2xUserFlowIdentityProviderClient, err := identityb2xuserflowidentityprovider.NewIdentityB2xUserFlowIdentityProviderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IdentityB2xUserFlowIdentityProvider client: %+v", err)
	}
	configureFunc(identityB2xUserFlowIdentityProviderClient.Client)

	identityB2xUserFlowLanguageClient, err := identityb2xuserflowlanguage.NewIdentityB2xUserFlowLanguageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IdentityB2xUserFlowLanguage client: %+v", err)
	}
	configureFunc(identityB2xUserFlowLanguageClient.Client)

	identityB2xUserFlowLanguageDefaultPageClient, err := identityb2xuserflowlanguagedefaultpage.NewIdentityB2xUserFlowLanguageDefaultPageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IdentityB2xUserFlowLanguageDefaultPage client: %+v", err)
	}
	configureFunc(identityB2xUserFlowLanguageDefaultPageClient.Client)

	identityB2xUserFlowLanguageOverridesPageClient, err := identityb2xuserflowlanguageoverridespage.NewIdentityB2xUserFlowLanguageOverridesPageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IdentityB2xUserFlowLanguageOverridesPage client: %+v", err)
	}
	configureFunc(identityB2xUserFlowLanguageOverridesPageClient.Client)

	identityB2xUserFlowUserAttributeAssignmentClient, err := identityb2xuserflowuserattributeassignment.NewIdentityB2xUserFlowUserAttributeAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IdentityB2xUserFlowUserAttributeAssignment client: %+v", err)
	}
	configureFunc(identityB2xUserFlowUserAttributeAssignmentClient.Client)

	identityB2xUserFlowUserAttributeAssignmentUserAttributeClient, err := identityb2xuserflowuserattributeassignmentuserattribute.NewIdentityB2xUserFlowUserAttributeAssignmentUserAttributeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IdentityB2xUserFlowUserAttributeAssignmentUserAttribute client: %+v", err)
	}
	configureFunc(identityB2xUserFlowUserAttributeAssignmentUserAttributeClient.Client)

	identityB2xUserFlowUserFlowIdentityProviderClient, err := identityb2xuserflowuserflowidentityprovider.NewIdentityB2xUserFlowUserFlowIdentityProviderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IdentityB2xUserFlowUserFlowIdentityProvider client: %+v", err)
	}
	configureFunc(identityB2xUserFlowUserFlowIdentityProviderClient.Client)

	identityClient, err := identity.NewIdentityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Identity client: %+v", err)
	}
	configureFunc(identityClient.Client)

	identityConditionalAccesAuthenticationContextClassReferenceClient, err := identityconditionalaccesauthenticationcontextclassreference.NewIdentityConditionalAccesAuthenticationContextClassReferenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IdentityConditionalAccesAuthenticationContextClassReference client: %+v", err)
	}
	configureFunc(identityConditionalAccesAuthenticationContextClassReferenceClient.Client)

	identityConditionalAccesAuthenticationStrengthAuthenticationMethodModeClient, err := identityconditionalaccesauthenticationstrengthauthenticationmethodmode.NewIdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IdentityConditionalAccesAuthenticationStrengthAuthenticationMethodMode client: %+v", err)
	}
	configureFunc(identityConditionalAccesAuthenticationStrengthAuthenticationMethodModeClient.Client)

	identityConditionalAccesAuthenticationStrengthClient, err := identityconditionalaccesauthenticationstrength.NewIdentityConditionalAccesAuthenticationStrengthClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IdentityConditionalAccesAuthenticationStrength client: %+v", err)
	}
	configureFunc(identityConditionalAccesAuthenticationStrengthClient.Client)

	identityConditionalAccesAuthenticationStrengthPolicyClient, err := identityconditionalaccesauthenticationstrengthpolicy.NewIdentityConditionalAccesAuthenticationStrengthPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IdentityConditionalAccesAuthenticationStrengthPolicy client: %+v", err)
	}
	configureFunc(identityConditionalAccesAuthenticationStrengthPolicyClient.Client)

	identityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationClient, err := identityconditionalaccesauthenticationstrengthpolicycombinationconfiguration.NewIdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfiguration client: %+v", err)
	}
	configureFunc(identityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationClient.Client)

	identityConditionalAccesClient, err := identityconditionalacces.NewIdentityConditionalAccesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IdentityConditionalAcces client: %+v", err)
	}
	configureFunc(identityConditionalAccesClient.Client)

	identityConditionalAccesNamedLocationClient, err := identityconditionalaccesnamedlocation.NewIdentityConditionalAccesNamedLocationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IdentityConditionalAccesNamedLocation client: %+v", err)
	}
	configureFunc(identityConditionalAccesNamedLocationClient.Client)

	identityConditionalAccesPolicyClient, err := identityconditionalaccespolicy.NewIdentityConditionalAccesPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IdentityConditionalAccesPolicy client: %+v", err)
	}
	configureFunc(identityConditionalAccesPolicyClient.Client)

	identityConditionalAccesTemplateClient, err := identityconditionalaccestemplate.NewIdentityConditionalAccesTemplateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IdentityConditionalAccesTemplate client: %+v", err)
	}
	configureFunc(identityConditionalAccesTemplateClient.Client)

	identityIdentityProviderClient, err := identityidentityprovider.NewIdentityIdentityProviderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IdentityIdentityProvider client: %+v", err)
	}
	configureFunc(identityIdentityProviderClient.Client)

	identityUserFlowAttributeClient, err := identityuserflowattribute.NewIdentityUserFlowAttributeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IdentityUserFlowAttribute client: %+v", err)
	}
	configureFunc(identityUserFlowAttributeClient.Client)

	return &Client{
		Identity:             identityClient,
		IdentityApiConnector: identityApiConnectorClient,
		IdentityB2xUserFlow:  identityB2xUserFlowClient,
		IdentityB2xUserFlowApiConnectorConfiguration:                                 identityB2xUserFlowApiConnectorConfigurationClient,
		IdentityB2xUserFlowApiConnectorConfigurationPostAttributeCollection:          identityB2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient,
		IdentityB2xUserFlowApiConnectorConfigurationPostFederationSignup:             identityB2xUserFlowApiConnectorConfigurationPostFederationSignupClient,
		IdentityB2xUserFlowIdentityProvider:                                          identityB2xUserFlowIdentityProviderClient,
		IdentityB2xUserFlowLanguage:                                                  identityB2xUserFlowLanguageClient,
		IdentityB2xUserFlowLanguageDefaultPage:                                       identityB2xUserFlowLanguageDefaultPageClient,
		IdentityB2xUserFlowLanguageOverridesPage:                                     identityB2xUserFlowLanguageOverridesPageClient,
		IdentityB2xUserFlowUserAttributeAssignment:                                   identityB2xUserFlowUserAttributeAssignmentClient,
		IdentityB2xUserFlowUserAttributeAssignmentUserAttribute:                      identityB2xUserFlowUserAttributeAssignmentUserAttributeClient,
		IdentityB2xUserFlowUserFlowIdentityProvider:                                  identityB2xUserFlowUserFlowIdentityProviderClient,
		IdentityConditionalAcces:                                                     identityConditionalAccesClient,
		IdentityConditionalAccesAuthenticationContextClassReference:                  identityConditionalAccesAuthenticationContextClassReferenceClient,
		IdentityConditionalAccesAuthenticationStrength:                               identityConditionalAccesAuthenticationStrengthClient,
		IdentityConditionalAccesAuthenticationStrengthAuthenticationMethodMode:       identityConditionalAccesAuthenticationStrengthAuthenticationMethodModeClient,
		IdentityConditionalAccesAuthenticationStrengthPolicy:                         identityConditionalAccesAuthenticationStrengthPolicyClient,
		IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfiguration: identityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationClient,
		IdentityConditionalAccesNamedLocation:                                        identityConditionalAccesNamedLocationClient,
		IdentityConditionalAccesPolicy:                                               identityConditionalAccesPolicyClient,
		IdentityConditionalAccesTemplate:                                             identityConditionalAccesTemplateClient,
		IdentityIdentityProvider:                                                     identityIdentityProviderClient,
		IdentityUserFlowAttribute:                                                    identityUserFlowAttributeClient,
	}, nil
}
