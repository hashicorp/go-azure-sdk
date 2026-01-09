package stable

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/apiconnector"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/authenticationeventlistener"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/authenticationeventsflow"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/authenticationeventsflowcondition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/authenticationeventsflowconditionapplicationincludeapplication"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/authenticationeventsflowexternalusersselfservicesignupeventsflowcondition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/authenticationeventsflowexternalusersselfservicesignupeventsflowconditionapplicationincludeapplication"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/authenticationeventsflowexternalusersselfservicesignupeventsflowonattributecollection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/authenticationeventsflowexternalusersselfservicesignupeventsflowonattributecollectiononattributecollectionexternalusersselfservicesignupattribute"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/authenticationeventsflowexternalusersselfservicesignupeventsflowonauthenticationmethodloadstart"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/authenticationeventsflowexternalusersselfservicesignupeventsflowonauthenticationmethodloadstartonauthenticationmethodloadstartexternalusersselfservicesignupidentityprovider"
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
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccessauthenticationcontextclassreference"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccessauthenticationstrength"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccessauthenticationstrengthauthenticationmethodmode"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccessauthenticationstrengthpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccessauthenticationstrengthpolicycombinationconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccessnamedlocation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccesspolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccesstemplate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/customauthenticationextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/identity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/identityprovider"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/userflowattribute"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ApiConnector                                                                                                                                                                 *apiconnector.ApiConnectorClient
	AuthenticationEventListener                                                                                                                                                  *authenticationeventlistener.AuthenticationEventListenerClient
	AuthenticationEventsFlow                                                                                                                                                     *authenticationeventsflow.AuthenticationEventsFlowClient
	AuthenticationEventsFlowCondition                                                                                                                                            *authenticationeventsflowcondition.AuthenticationEventsFlowConditionClient
	AuthenticationEventsFlowConditionApplicationIncludeApplication                                                                                                               *authenticationeventsflowconditionapplicationincludeapplication.AuthenticationEventsFlowConditionApplicationIncludeApplicationClient
	AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowCondition                                                                                                    *authenticationeventsflowexternalusersselfservicesignupeventsflowcondition.AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowConditionClient
	AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplication                                                                       *authenticationeventsflowexternalusersselfservicesignupeventsflowconditionapplicationincludeapplication.AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationClient
	AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollection                                                                                        *authenticationeventsflowexternalusersselfservicesignupeventsflowonattributecollection.AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionClient
	AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttribute                            *authenticationeventsflowexternalusersselfservicesignupeventsflowonattributecollectiononattributecollectionexternalusersselfservicesignupattribute.AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeClient
	AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStart                                                                              *authenticationeventsflowexternalusersselfservicesignupeventsflowonauthenticationmethodloadstart.AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartClient
	AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProvider *authenticationeventsflowexternalusersselfservicesignupeventsflowonauthenticationmethodloadstartonauthenticationmethodloadstartexternalusersselfservicesignupidentityprovider.AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderClient
	B2xUserFlow                                                                                                                                                                  *b2xuserflow.B2xUserFlowClient
	B2xUserFlowApiConnectorConfiguration                                                                                                                                         *b2xuserflowapiconnectorconfiguration.B2xUserFlowApiConnectorConfigurationClient
	B2xUserFlowApiConnectorConfigurationPostAttributeCollection                                                                                                                  *b2xuserflowapiconnectorconfigurationpostattributecollection.B2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient
	B2xUserFlowApiConnectorConfigurationPostFederationSignup                                                                                                                     *b2xuserflowapiconnectorconfigurationpostfederationsignup.B2xUserFlowApiConnectorConfigurationPostFederationSignupClient
	B2xUserFlowIdentityProvider                                                                                                                                                  *b2xuserflowidentityprovider.B2xUserFlowIdentityProviderClient
	B2xUserFlowLanguage                                                                                                                                                          *b2xuserflowlanguage.B2xUserFlowLanguageClient
	B2xUserFlowLanguageDefaultPage                                                                                                                                               *b2xuserflowlanguagedefaultpage.B2xUserFlowLanguageDefaultPageClient
	B2xUserFlowLanguageOverridesPage                                                                                                                                             *b2xuserflowlanguageoverridespage.B2xUserFlowLanguageOverridesPageClient
	B2xUserFlowUserAttributeAssignment                                                                                                                                           *b2xuserflowuserattributeassignment.B2xUserFlowUserAttributeAssignmentClient
	B2xUserFlowUserAttributeAssignmentUserAttribute                                                                                                                              *b2xuserflowuserattributeassignmentuserattribute.B2xUserFlowUserAttributeAssignmentUserAttributeClient
	B2xUserFlowUserFlowIdentityProvider                                                                                                                                          *b2xuserflowuserflowidentityprovider.B2xUserFlowUserFlowIdentityProviderClient
	ConditionalAccessAuthenticationContextClassReference                                                                                                                         *conditionalaccessauthenticationcontextclassreference.ConditionalAccessAuthenticationContextClassReferenceClient
	ConditionalAccessAuthenticationStrength                                                                                                                                      *conditionalaccessauthenticationstrength.ConditionalAccessAuthenticationStrengthClient
	ConditionalAccessAuthenticationStrengthAuthenticationMethodMode                                                                                                              *conditionalaccessauthenticationstrengthauthenticationmethodmode.ConditionalAccessAuthenticationStrengthAuthenticationMethodModeClient
	ConditionalAccessAuthenticationStrengthPolicy                                                                                                                                *conditionalaccessauthenticationstrengthpolicy.ConditionalAccessAuthenticationStrengthPolicyClient
	ConditionalAccessAuthenticationStrengthPolicyCombinationConfiguration                                                                                                        *conditionalaccessauthenticationstrengthpolicycombinationconfiguration.ConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationClient
	ConditionalAccessNamedLocation                                                                                                                                               *conditionalaccessnamedlocation.ConditionalAccessNamedLocationClient
	ConditionalAccessPolicy                                                                                                                                                      *conditionalaccesspolicy.ConditionalAccessPolicyClient
	ConditionalAccessTemplate                                                                                                                                                    *conditionalaccesstemplate.ConditionalAccessTemplateClient
	CustomAuthenticationExtension                                                                                                                                                *customauthenticationextension.CustomAuthenticationExtensionClient
	Identity                                                                                                                                                                     *identity.IdentityClient
	IdentityProvider                                                                                                                                                             *identityprovider.IdentityProviderClient
	UserFlowAttribute                                                                                                                                                            *userflowattribute.UserFlowAttributeClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	apiConnectorClient, err := apiconnector.NewApiConnectorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApiConnector client: %+v", err)
	}
	configureFunc(apiConnectorClient.Client)

	authenticationEventListenerClient, err := authenticationeventlistener.NewAuthenticationEventListenerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationEventListener client: %+v", err)
	}
	configureFunc(authenticationEventListenerClient.Client)

	authenticationEventsFlowClient, err := authenticationeventsflow.NewAuthenticationEventsFlowClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationEventsFlow client: %+v", err)
	}
	configureFunc(authenticationEventsFlowClient.Client)

	authenticationEventsFlowConditionApplicationIncludeApplicationClient, err := authenticationeventsflowconditionapplicationincludeapplication.NewAuthenticationEventsFlowConditionApplicationIncludeApplicationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationEventsFlowConditionApplicationIncludeApplication client: %+v", err)
	}
	configureFunc(authenticationEventsFlowConditionApplicationIncludeApplicationClient.Client)

	authenticationEventsFlowConditionClient, err := authenticationeventsflowcondition.NewAuthenticationEventsFlowConditionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationEventsFlowCondition client: %+v", err)
	}
	configureFunc(authenticationEventsFlowConditionClient.Client)

	authenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationClient, err := authenticationeventsflowexternalusersselfservicesignupeventsflowconditionapplicationincludeapplication.NewAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplication client: %+v", err)
	}
	configureFunc(authenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationClient.Client)

	authenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowConditionClient, err := authenticationeventsflowexternalusersselfservicesignupeventsflowcondition.NewAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowConditionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowCondition client: %+v", err)
	}
	configureFunc(authenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowConditionClient.Client)

	authenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionClient, err := authenticationeventsflowexternalusersselfservicesignupeventsflowonattributecollection.NewAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollection client: %+v", err)
	}
	configureFunc(authenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionClient.Client)

	authenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeClient, err := authenticationeventsflowexternalusersselfservicesignupeventsflowonattributecollectiononattributecollectionexternalusersselfservicesignupattribute.NewAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttribute client: %+v", err)
	}
	configureFunc(authenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeClient.Client)

	authenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartClient, err := authenticationeventsflowexternalusersselfservicesignupeventsflowonauthenticationmethodloadstart.NewAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStart client: %+v", err)
	}
	configureFunc(authenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartClient.Client)

	authenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderClient, err := authenticationeventsflowexternalusersselfservicesignupeventsflowonauthenticationmethodloadstartonauthenticationmethodloadstartexternalusersselfservicesignupidentityprovider.NewAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProvider client: %+v", err)
	}
	configureFunc(authenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderClient.Client)

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

	conditionalAccessAuthenticationContextClassReferenceClient, err := conditionalaccessauthenticationcontextclassreference.NewConditionalAccessAuthenticationContextClassReferenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConditionalAccessAuthenticationContextClassReference client: %+v", err)
	}
	configureFunc(conditionalAccessAuthenticationContextClassReferenceClient.Client)

	conditionalAccessAuthenticationStrengthAuthenticationMethodModeClient, err := conditionalaccessauthenticationstrengthauthenticationmethodmode.NewConditionalAccessAuthenticationStrengthAuthenticationMethodModeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConditionalAccessAuthenticationStrengthAuthenticationMethodMode client: %+v", err)
	}
	configureFunc(conditionalAccessAuthenticationStrengthAuthenticationMethodModeClient.Client)

	conditionalAccessAuthenticationStrengthClient, err := conditionalaccessauthenticationstrength.NewConditionalAccessAuthenticationStrengthClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConditionalAccessAuthenticationStrength client: %+v", err)
	}
	configureFunc(conditionalAccessAuthenticationStrengthClient.Client)

	conditionalAccessAuthenticationStrengthPolicyClient, err := conditionalaccessauthenticationstrengthpolicy.NewConditionalAccessAuthenticationStrengthPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConditionalAccessAuthenticationStrengthPolicy client: %+v", err)
	}
	configureFunc(conditionalAccessAuthenticationStrengthPolicyClient.Client)

	conditionalAccessAuthenticationStrengthPolicyCombinationConfigurationClient, err := conditionalaccessauthenticationstrengthpolicycombinationconfiguration.NewConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConditionalAccessAuthenticationStrengthPolicyCombinationConfiguration client: %+v", err)
	}
	configureFunc(conditionalAccessAuthenticationStrengthPolicyCombinationConfigurationClient.Client)

	conditionalAccessNamedLocationClient, err := conditionalaccessnamedlocation.NewConditionalAccessNamedLocationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConditionalAccessNamedLocation client: %+v", err)
	}
	configureFunc(conditionalAccessNamedLocationClient.Client)

	conditionalAccessPolicyClient, err := conditionalaccesspolicy.NewConditionalAccessPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConditionalAccessPolicy client: %+v", err)
	}
	configureFunc(conditionalAccessPolicyClient.Client)

	conditionalAccessTemplateClient, err := conditionalaccesstemplate.NewConditionalAccessTemplateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConditionalAccessTemplate client: %+v", err)
	}
	configureFunc(conditionalAccessTemplateClient.Client)

	customAuthenticationExtensionClient, err := customauthenticationextension.NewCustomAuthenticationExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CustomAuthenticationExtension client: %+v", err)
	}
	configureFunc(customAuthenticationExtensionClient.Client)

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
		ApiConnector:                      apiConnectorClient,
		AuthenticationEventListener:       authenticationEventListenerClient,
		AuthenticationEventsFlow:          authenticationEventsFlowClient,
		AuthenticationEventsFlowCondition: authenticationEventsFlowConditionClient,
		AuthenticationEventsFlowConditionApplicationIncludeApplication:                                                                                                               authenticationEventsFlowConditionApplicationIncludeApplicationClient,
		AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowCondition:                                                                                                    authenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowConditionClient,
		AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplication:                                                                       authenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowConditionApplicationIncludeApplicationClient,
		AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollection:                                                                                        authenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionClient,
		AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttribute:                            authenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeClient,
		AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStart:                                                                              authenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartClient,
		AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProvider: authenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderClient,
		B2xUserFlow:                          b2xUserFlowClient,
		B2xUserFlowApiConnectorConfiguration: b2xUserFlowApiConnectorConfigurationClient,
		B2xUserFlowApiConnectorConfigurationPostAttributeCollection:           b2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient,
		B2xUserFlowApiConnectorConfigurationPostFederationSignup:              b2xUserFlowApiConnectorConfigurationPostFederationSignupClient,
		B2xUserFlowIdentityProvider:                                           b2xUserFlowIdentityProviderClient,
		B2xUserFlowLanguage:                                                   b2xUserFlowLanguageClient,
		B2xUserFlowLanguageDefaultPage:                                        b2xUserFlowLanguageDefaultPageClient,
		B2xUserFlowLanguageOverridesPage:                                      b2xUserFlowLanguageOverridesPageClient,
		B2xUserFlowUserAttributeAssignment:                                    b2xUserFlowUserAttributeAssignmentClient,
		B2xUserFlowUserAttributeAssignmentUserAttribute:                       b2xUserFlowUserAttributeAssignmentUserAttributeClient,
		B2xUserFlowUserFlowIdentityProvider:                                   b2xUserFlowUserFlowIdentityProviderClient,
		ConditionalAccessAuthenticationContextClassReference:                  conditionalAccessAuthenticationContextClassReferenceClient,
		ConditionalAccessAuthenticationStrength:                               conditionalAccessAuthenticationStrengthClient,
		ConditionalAccessAuthenticationStrengthAuthenticationMethodMode:       conditionalAccessAuthenticationStrengthAuthenticationMethodModeClient,
		ConditionalAccessAuthenticationStrengthPolicy:                         conditionalAccessAuthenticationStrengthPolicyClient,
		ConditionalAccessAuthenticationStrengthPolicyCombinationConfiguration: conditionalAccessAuthenticationStrengthPolicyCombinationConfigurationClient,
		ConditionalAccessNamedLocation:                                        conditionalAccessNamedLocationClient,
		ConditionalAccessPolicy:                                               conditionalAccessPolicyClient,
		ConditionalAccessTemplate:                                             conditionalAccessTemplateClient,
		CustomAuthenticationExtension:                                         customAuthenticationExtensionClient,
		Identity:                                                              identityClient,
		IdentityProvider:                                                      identityProviderClient,
		UserFlowAttribute:                                                     userFlowAttributeClient,
	}, nil
}
