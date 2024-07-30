package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/apiconnector"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/b2xuserflow"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/b2xuserflowapiconnectorconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/b2xuserflowapiconnectorconfigurationpostattributecollection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/b2xuserflowapiconnectorconfigurationpostfederationsignup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/b2xuserflowidentityprovider"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/b2xuserflowlanguage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/b2xuserflowlanguagedefaultpage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/b2xuserflowlanguageoverridespage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/b2xuserflowuserattributeassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/b2xuserflowuserattributeassignmentuserattribute"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/b2xuserflowuserflowidentityprovider"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalacces"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccesauthenticationcontextclassreference"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccesauthenticationstrength"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccesauthenticationstrengthauthenticationmethodmode"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccesauthenticationstrengthpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccesauthenticationstrengthpolicycombinationconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccesnamedlocation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccespolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccestemplate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/identity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/identityprovider"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/userflowattribute"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	ApiConnector                                                         *apiconnector.ApiConnectorClient
	B2xUserFlow                                                          *b2xuserflow.B2xUserFlowClient
	B2xUserFlowApiConnectorConfiguration                                 *b2xuserflowapiconnectorconfiguration.B2xUserFlowApiConnectorConfigurationClient
	B2xUserFlowApiConnectorConfigurationPostAttributeCollection          *b2xuserflowapiconnectorconfigurationpostattributecollection.B2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient
	B2xUserFlowApiConnectorConfigurationPostFederationSignup             *b2xuserflowapiconnectorconfigurationpostfederationsignup.B2xUserFlowApiConnectorConfigurationPostFederationSignupClient
	B2xUserFlowIdentityProvider                                          *b2xuserflowidentityprovider.B2xUserFlowIdentityProviderClient
	B2xUserFlowLanguage                                                  *b2xuserflowlanguage.B2xUserFlowLanguageClient
	B2xUserFlowLanguageDefaultPage                                       *b2xuserflowlanguagedefaultpage.B2xUserFlowLanguageDefaultPageClient
	B2xUserFlowLanguageOverridesPage                                     *b2xuserflowlanguageoverridespage.B2xUserFlowLanguageOverridesPageClient
	B2xUserFlowUserAttributeAssignment                                   *b2xuserflowuserattributeassignment.B2xUserFlowUserAttributeAssignmentClient
	B2xUserFlowUserAttributeAssignmentUserAttribute                      *b2xuserflowuserattributeassignmentuserattribute.B2xUserFlowUserAttributeAssignmentUserAttributeClient
	B2xUserFlowUserFlowIdentityProvider                                  *b2xuserflowuserflowidentityprovider.B2xUserFlowUserFlowIdentityProviderClient
	ConditionalAcces                                                     *conditionalacces.ConditionalAccesClient
	ConditionalAccesAuthenticationContextClassReference                  *conditionalaccesauthenticationcontextclassreference.ConditionalAccesAuthenticationContextClassReferenceClient
	ConditionalAccesAuthenticationStrength                               *conditionalaccesauthenticationstrength.ConditionalAccesAuthenticationStrengthClient
	ConditionalAccesAuthenticationStrengthAuthenticationMethodMode       *conditionalaccesauthenticationstrengthauthenticationmethodmode.ConditionalAccesAuthenticationStrengthAuthenticationMethodModeClient
	ConditionalAccesAuthenticationStrengthPolicy                         *conditionalaccesauthenticationstrengthpolicy.ConditionalAccesAuthenticationStrengthPolicyClient
	ConditionalAccesAuthenticationStrengthPolicyCombinationConfiguration *conditionalaccesauthenticationstrengthpolicycombinationconfiguration.ConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationClient
	ConditionalAccesNamedLocation                                        *conditionalaccesnamedlocation.ConditionalAccesNamedLocationClient
	ConditionalAccesPolicy                                               *conditionalaccespolicy.ConditionalAccesPolicyClient
	ConditionalAccesTemplate                                             *conditionalaccestemplate.ConditionalAccesTemplateClient
	Identity                                                             *identity.IdentityClient
	IdentityProvider                                                     *identityprovider.IdentityProviderClient
	UserFlowAttribute                                                    *userflowattribute.UserFlowAttributeClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	apiConnectorClient, err := apiconnector.NewApiConnectorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApiConnector client: %+v", err)
	}
	configureFunc(apiConnectorClient.Client)

	b2xUserFlowApiConnectorConfigurationClient, err := b2xuserflowapiconnectorconfiguration.NewB2xUserFlowApiConnectorConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building B2xUserFlowApiConnectorConfiguration client: %+v", err)
	}
	configureFunc(b2xUserFlowApiConnectorConfigurationClient.Client)

	b2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient, err := b2xuserflowapiconnectorconfigurationpostattributecollection.NewB2xUserFlowApiConnectorConfigurationPostAttributeCollectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building B2xUserFlowApiConnectorConfigurationPostAttributeCollection client: %+v", err)
	}
	configureFunc(b2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient.Client)

	b2xUserFlowApiConnectorConfigurationPostFederationSignupClient, err := b2xuserflowapiconnectorconfigurationpostfederationsignup.NewB2xUserFlowApiConnectorConfigurationPostFederationSignupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building B2xUserFlowApiConnectorConfigurationPostFederationSignup client: %+v", err)
	}
	configureFunc(b2xUserFlowApiConnectorConfigurationPostFederationSignupClient.Client)

	b2xUserFlowClient, err := b2xuserflow.NewB2xUserFlowClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building B2xUserFlow client: %+v", err)
	}
	configureFunc(b2xUserFlowClient.Client)

	b2xUserFlowIdentityProviderClient, err := b2xuserflowidentityprovider.NewB2xUserFlowIdentityProviderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building B2xUserFlowIdentityProvider client: %+v", err)
	}
	configureFunc(b2xUserFlowIdentityProviderClient.Client)

	b2xUserFlowLanguageClient, err := b2xuserflowlanguage.NewB2xUserFlowLanguageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building B2xUserFlowLanguage client: %+v", err)
	}
	configureFunc(b2xUserFlowLanguageClient.Client)

	b2xUserFlowLanguageDefaultPageClient, err := b2xuserflowlanguagedefaultpage.NewB2xUserFlowLanguageDefaultPageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building B2xUserFlowLanguageDefaultPage client: %+v", err)
	}
	configureFunc(b2xUserFlowLanguageDefaultPageClient.Client)

	b2xUserFlowLanguageOverridesPageClient, err := b2xuserflowlanguageoverridespage.NewB2xUserFlowLanguageOverridesPageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building B2xUserFlowLanguageOverridesPage client: %+v", err)
	}
	configureFunc(b2xUserFlowLanguageOverridesPageClient.Client)

	b2xUserFlowUserAttributeAssignmentClient, err := b2xuserflowuserattributeassignment.NewB2xUserFlowUserAttributeAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building B2xUserFlowUserAttributeAssignment client: %+v", err)
	}
	configureFunc(b2xUserFlowUserAttributeAssignmentClient.Client)

	b2xUserFlowUserAttributeAssignmentUserAttributeClient, err := b2xuserflowuserattributeassignmentuserattribute.NewB2xUserFlowUserAttributeAssignmentUserAttributeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building B2xUserFlowUserAttributeAssignmentUserAttribute client: %+v", err)
	}
	configureFunc(b2xUserFlowUserAttributeAssignmentUserAttributeClient.Client)

	b2xUserFlowUserFlowIdentityProviderClient, err := b2xuserflowuserflowidentityprovider.NewB2xUserFlowUserFlowIdentityProviderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building B2xUserFlowUserFlowIdentityProvider client: %+v", err)
	}
	configureFunc(b2xUserFlowUserFlowIdentityProviderClient.Client)

	conditionalAccesAuthenticationContextClassReferenceClient, err := conditionalaccesauthenticationcontextclassreference.NewConditionalAccesAuthenticationContextClassReferenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConditionalAccesAuthenticationContextClassReference client: %+v", err)
	}
	configureFunc(conditionalAccesAuthenticationContextClassReferenceClient.Client)

	conditionalAccesAuthenticationStrengthAuthenticationMethodModeClient, err := conditionalaccesauthenticationstrengthauthenticationmethodmode.NewConditionalAccesAuthenticationStrengthAuthenticationMethodModeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConditionalAccesAuthenticationStrengthAuthenticationMethodMode client: %+v", err)
	}
	configureFunc(conditionalAccesAuthenticationStrengthAuthenticationMethodModeClient.Client)

	conditionalAccesAuthenticationStrengthClient, err := conditionalaccesauthenticationstrength.NewConditionalAccesAuthenticationStrengthClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConditionalAccesAuthenticationStrength client: %+v", err)
	}
	configureFunc(conditionalAccesAuthenticationStrengthClient.Client)

	conditionalAccesAuthenticationStrengthPolicyClient, err := conditionalaccesauthenticationstrengthpolicy.NewConditionalAccesAuthenticationStrengthPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConditionalAccesAuthenticationStrengthPolicy client: %+v", err)
	}
	configureFunc(conditionalAccesAuthenticationStrengthPolicyClient.Client)

	conditionalAccesAuthenticationStrengthPolicyCombinationConfigurationClient, err := conditionalaccesauthenticationstrengthpolicycombinationconfiguration.NewConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConditionalAccesAuthenticationStrengthPolicyCombinationConfiguration client: %+v", err)
	}
	configureFunc(conditionalAccesAuthenticationStrengthPolicyCombinationConfigurationClient.Client)

	conditionalAccesClient, err := conditionalacces.NewConditionalAccesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConditionalAcces client: %+v", err)
	}
	configureFunc(conditionalAccesClient.Client)

	conditionalAccesNamedLocationClient, err := conditionalaccesnamedlocation.NewConditionalAccesNamedLocationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConditionalAccesNamedLocation client: %+v", err)
	}
	configureFunc(conditionalAccesNamedLocationClient.Client)

	conditionalAccesPolicyClient, err := conditionalaccespolicy.NewConditionalAccesPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConditionalAccesPolicy client: %+v", err)
	}
	configureFunc(conditionalAccesPolicyClient.Client)

	conditionalAccesTemplateClient, err := conditionalaccestemplate.NewConditionalAccesTemplateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConditionalAccesTemplate client: %+v", err)
	}
	configureFunc(conditionalAccesTemplateClient.Client)

	identityClient, err := identity.NewIdentityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Identity client: %+v", err)
	}
	configureFunc(identityClient.Client)

	identityProviderClient, err := identityprovider.NewIdentityProviderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IdentityProvider client: %+v", err)
	}
	configureFunc(identityProviderClient.Client)

	userFlowAttributeClient, err := userflowattribute.NewUserFlowAttributeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserFlowAttribute client: %+v", err)
	}
	configureFunc(userFlowAttributeClient.Client)

	return &Client{
		ApiConnector:                         apiConnectorClient,
		B2xUserFlow:                          b2xUserFlowClient,
		B2xUserFlowApiConnectorConfiguration: b2xUserFlowApiConnectorConfigurationClient,
		B2xUserFlowApiConnectorConfigurationPostAttributeCollection:          b2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient,
		B2xUserFlowApiConnectorConfigurationPostFederationSignup:             b2xUserFlowApiConnectorConfigurationPostFederationSignupClient,
		B2xUserFlowIdentityProvider:                                          b2xUserFlowIdentityProviderClient,
		B2xUserFlowLanguage:                                                  b2xUserFlowLanguageClient,
		B2xUserFlowLanguageDefaultPage:                                       b2xUserFlowLanguageDefaultPageClient,
		B2xUserFlowLanguageOverridesPage:                                     b2xUserFlowLanguageOverridesPageClient,
		B2xUserFlowUserAttributeAssignment:                                   b2xUserFlowUserAttributeAssignmentClient,
		B2xUserFlowUserAttributeAssignmentUserAttribute:                      b2xUserFlowUserAttributeAssignmentUserAttributeClient,
		B2xUserFlowUserFlowIdentityProvider:                                  b2xUserFlowUserFlowIdentityProviderClient,
		ConditionalAcces:                                                     conditionalAccesClient,
		ConditionalAccesAuthenticationContextClassReference:                  conditionalAccesAuthenticationContextClassReferenceClient,
		ConditionalAccesAuthenticationStrength:                               conditionalAccesAuthenticationStrengthClient,
		ConditionalAccesAuthenticationStrengthAuthenticationMethodMode:       conditionalAccesAuthenticationStrengthAuthenticationMethodModeClient,
		ConditionalAccesAuthenticationStrengthPolicy:                         conditionalAccesAuthenticationStrengthPolicyClient,
		ConditionalAccesAuthenticationStrengthPolicyCombinationConfiguration: conditionalAccesAuthenticationStrengthPolicyCombinationConfigurationClient,
		ConditionalAccesNamedLocation:                                        conditionalAccesNamedLocationClient,
		ConditionalAccesPolicy:                                               conditionalAccesPolicyClient,
		ConditionalAccesTemplate:                                             conditionalAccesTemplateClient,
		Identity:                                                             identityClient,
		IdentityProvider:                                                     identityProviderClient,
		UserFlowAttribute:                                                    userFlowAttributeClient,
	}, nil
}
