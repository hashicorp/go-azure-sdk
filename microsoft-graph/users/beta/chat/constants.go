package chat

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemBodyContentType string

const (
	ItemBodyContentTypeHtml ItemBodyContentType = "html"
	ItemBodyContentTypeText ItemBodyContentType = "text"
)

func PossibleValuesForItemBodyContentType() []string {
	return []string{
		string(ItemBodyContentTypeHtml),
		string(ItemBodyContentTypeText),
	}
}

func (s *ItemBodyContentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseItemBodyContentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseItemBodyContentType(input string) (*ItemBodyContentType, error) {
	vals := map[string]ItemBodyContentType{
		"html": ItemBodyContentTypeHtml,
		"text": ItemBodyContentTypeText,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ItemBodyContentType(input)
	return &out, nil
}

type TeamworkActivityTopicSource string

const (
	TeamworkActivityTopicSourceEntityUrl TeamworkActivityTopicSource = "entityUrl"
	TeamworkActivityTopicSourceText      TeamworkActivityTopicSource = "text"
)

func PossibleValuesForTeamworkActivityTopicSource() []string {
	return []string{
		string(TeamworkActivityTopicSourceEntityUrl),
		string(TeamworkActivityTopicSourceText),
	}
}

func (s *TeamworkActivityTopicSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamworkActivityTopicSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamworkActivityTopicSource(input string) (*TeamworkActivityTopicSource, error) {
	vals := map[string]TeamworkActivityTopicSource{
		"entityurl": TeamworkActivityTopicSourceEntityUrl,
		"text":      TeamworkActivityTopicSourceText,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkActivityTopicSource(input)
	return &out, nil
}

type TeamworkUserIdentityUserIdentityType string

const (
	TeamworkUserIdentityUserIdentityTypeAadUser                        TeamworkUserIdentityUserIdentityType = "aadUser"
	TeamworkUserIdentityUserIdentityTypeAnonymousGuest                 TeamworkUserIdentityUserIdentityType = "anonymousGuest"
	TeamworkUserIdentityUserIdentityTypeAzureCommunicationServicesUser TeamworkUserIdentityUserIdentityType = "azureCommunicationServicesUser"
	TeamworkUserIdentityUserIdentityTypeEmailUser                      TeamworkUserIdentityUserIdentityType = "emailUser"
	TeamworkUserIdentityUserIdentityTypeFederatedUser                  TeamworkUserIdentityUserIdentityType = "federatedUser"
	TeamworkUserIdentityUserIdentityTypeOnPremiseAadUser               TeamworkUserIdentityUserIdentityType = "onPremiseAadUser"
	TeamworkUserIdentityUserIdentityTypePersonalMicrosoftAccountUser   TeamworkUserIdentityUserIdentityType = "personalMicrosoftAccountUser"
	TeamworkUserIdentityUserIdentityTypePhoneUser                      TeamworkUserIdentityUserIdentityType = "phoneUser"
	TeamworkUserIdentityUserIdentityTypeSkypeUser                      TeamworkUserIdentityUserIdentityType = "skypeUser"
)

func PossibleValuesForTeamworkUserIdentityUserIdentityType() []string {
	return []string{
		string(TeamworkUserIdentityUserIdentityTypeAadUser),
		string(TeamworkUserIdentityUserIdentityTypeAnonymousGuest),
		string(TeamworkUserIdentityUserIdentityTypeAzureCommunicationServicesUser),
		string(TeamworkUserIdentityUserIdentityTypeEmailUser),
		string(TeamworkUserIdentityUserIdentityTypeFederatedUser),
		string(TeamworkUserIdentityUserIdentityTypeOnPremiseAadUser),
		string(TeamworkUserIdentityUserIdentityTypePersonalMicrosoftAccountUser),
		string(TeamworkUserIdentityUserIdentityTypePhoneUser),
		string(TeamworkUserIdentityUserIdentityTypeSkypeUser),
	}
}

func (s *TeamworkUserIdentityUserIdentityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamworkUserIdentityUserIdentityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamworkUserIdentityUserIdentityType(input string) (*TeamworkUserIdentityUserIdentityType, error) {
	vals := map[string]TeamworkUserIdentityUserIdentityType{
		"aaduser":                        TeamworkUserIdentityUserIdentityTypeAadUser,
		"anonymousguest":                 TeamworkUserIdentityUserIdentityTypeAnonymousGuest,
		"azurecommunicationservicesuser": TeamworkUserIdentityUserIdentityTypeAzureCommunicationServicesUser,
		"emailuser":                      TeamworkUserIdentityUserIdentityTypeEmailUser,
		"federateduser":                  TeamworkUserIdentityUserIdentityTypeFederatedUser,
		"onpremiseaaduser":               TeamworkUserIdentityUserIdentityTypeOnPremiseAadUser,
		"personalmicrosoftaccountuser":   TeamworkUserIdentityUserIdentityTypePersonalMicrosoftAccountUser,
		"phoneuser":                      TeamworkUserIdentityUserIdentityTypePhoneUser,
		"skypeuser":                      TeamworkUserIdentityUserIdentityTypeSkypeUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkUserIdentityUserIdentityType(input)
	return &out, nil
}
