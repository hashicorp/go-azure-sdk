package environments

import (
	"fmt"
	"strings"
)

func FromName(name string) (*Environment, error) {
	switch strings.ToLower(name) {
	case "china":
		return AzureChina(), nil

	case "canary":
		return AzurePublicCanary(), nil

	case "global", "public":
		return AzurePublic(), nil

	case "usgovernment", "usgovernmentl4":
		return AzureUSGovernment(), nil

	case "dod", "usgovernmentl5":
		return AzureUSGovernmentL5(), nil
	}

	return nil, fmt.Errorf("no environment was found with the name %q", name)
}

// TODO: From Name
//func FromNamed(env string) (*Environment, error) {
//	switch strings.ToLower(env) {
//	case "canary":
//		return &Canary, nil
//	}
//
//	return nil, fmt.Errorf("invalid environment specified: %s", env)
//}
