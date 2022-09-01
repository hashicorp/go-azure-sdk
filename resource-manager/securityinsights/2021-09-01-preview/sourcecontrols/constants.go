package sourcecontrols

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentType string

const (
	ContentTypeAnalyticRule ContentType = "AnalyticRule"
	ContentTypeWorkbook     ContentType = "Workbook"
)

func PossibleValuesForContentType() []string {
	return []string{
		string(ContentTypeAnalyticRule),
		string(ContentTypeWorkbook),
	}
}

func parseContentType(input string) (*ContentType, error) {
	vals := map[string]ContentType{
		"analyticrule": ContentTypeAnalyticRule,
		"workbook":     ContentTypeWorkbook,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContentType(input)
	return &out, nil
}

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

func parseRepoType(input string) (*RepoType, error) {
	vals := map[string]RepoType{
		"devops": RepoTypeDevOps,
		"github": RepoTypeGithub,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RepoType(input)
	return &out, nil
}
