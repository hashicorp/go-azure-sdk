package favoritesapis

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FavoriteSourceType string

const (
	FavoriteSourceTypeEvents       FavoriteSourceType = "events"
	FavoriteSourceTypeFunnel       FavoriteSourceType = "funnel"
	FavoriteSourceTypeImpact       FavoriteSourceType = "impact"
	FavoriteSourceTypeNotebook     FavoriteSourceType = "notebook"
	FavoriteSourceTypeRetention    FavoriteSourceType = "retention"
	FavoriteSourceTypeSegmentation FavoriteSourceType = "segmentation"
	FavoriteSourceTypeSessions     FavoriteSourceType = "sessions"
	FavoriteSourceTypeUserflows    FavoriteSourceType = "userflows"
)

func PossibleValuesForFavoriteSourceType() []string {
	return []string{
		string(FavoriteSourceTypeEvents),
		string(FavoriteSourceTypeFunnel),
		string(FavoriteSourceTypeImpact),
		string(FavoriteSourceTypeNotebook),
		string(FavoriteSourceTypeRetention),
		string(FavoriteSourceTypeSegmentation),
		string(FavoriteSourceTypeSessions),
		string(FavoriteSourceTypeUserflows),
	}
}

func (s *FavoriteSourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFavoriteSourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFavoriteSourceType(input string) (*FavoriteSourceType, error) {
	vals := map[string]FavoriteSourceType{
		"events":       FavoriteSourceTypeEvents,
		"funnel":       FavoriteSourceTypeFunnel,
		"impact":       FavoriteSourceTypeImpact,
		"notebook":     FavoriteSourceTypeNotebook,
		"retention":    FavoriteSourceTypeRetention,
		"segmentation": FavoriteSourceTypeSegmentation,
		"sessions":     FavoriteSourceTypeSessions,
		"userflows":    FavoriteSourceTypeUserflows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FavoriteSourceType(input)
	return &out, nil
}

type FavoriteType string

const (
	FavoriteTypeShared FavoriteType = "shared"
	FavoriteTypeUser   FavoriteType = "user"
)

func PossibleValuesForFavoriteType() []string {
	return []string{
		string(FavoriteTypeShared),
		string(FavoriteTypeUser),
	}
}

func (s *FavoriteType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFavoriteType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFavoriteType(input string) (*FavoriteType, error) {
	vals := map[string]FavoriteType{
		"shared": FavoriteTypeShared,
		"user":   FavoriteTypeUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FavoriteType(input)
	return &out, nil
}
