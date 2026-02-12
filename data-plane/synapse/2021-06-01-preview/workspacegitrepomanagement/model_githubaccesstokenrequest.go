package workspacegitrepomanagement

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GitHubAccessTokenRequest struct {
	GitHubAccessCode         string `json:"gitHubAccessCode"`
	GitHubAccessTokenBaseURL string `json:"gitHubAccessTokenBaseUrl"`
	GitHubClientId           string `json:"gitHubClientId"`
}
