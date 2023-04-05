package repositories

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RepoType string

const (
	RepoTypeDevOps RepoType = "DevOps"
	RepoTypeGithub RepoType = "Github"
)

func PossibleValuesForRepoType() []string {
	return []string{
		string(RepoTypeDevOps),
		string(RepoTypeGithub),
	}
}
