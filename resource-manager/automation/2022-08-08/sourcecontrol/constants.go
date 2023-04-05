package sourcecontrol

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SourceType string

const (
	SourceTypeGitHub  SourceType = "GitHub"
	SourceTypeVsoGit  SourceType = "VsoGit"
	SourceTypeVsoTfvc SourceType = "VsoTfvc"
)

func PossibleValuesForSourceType() []string {
	return []string{
		string(SourceTypeGitHub),
		string(SourceTypeVsoGit),
		string(SourceTypeVsoTfvc),
	}
}

func (s *SourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSourceType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SourceType(decoded)
	return nil
}

type TokenType string

const (
	TokenTypeOauth               TokenType = "Oauth"
	TokenTypePersonalAccessToken TokenType = "PersonalAccessToken"
)

func PossibleValuesForTokenType() []string {
	return []string{
		string(TokenTypeOauth),
		string(TokenTypePersonalAccessToken),
	}
}

func (s *TokenType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForTokenType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = TokenType(decoded)
	return nil
}
