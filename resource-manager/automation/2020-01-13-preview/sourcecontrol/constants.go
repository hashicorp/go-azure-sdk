package sourcecontrol

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
