package environments

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
)

func FromName(name string) (*Environment, error) {
	switch strings.ToLower(name) {
	case "china":
		return pointer.To(AzureChina()), nil

	case "canary":
		return pointer.To(AzurePublicCanary()), nil

	case "global", "public":
		return pointer.To(AzurePublic()), nil

	case "usgovernment", "usgovernmentl4":
		return pointer.To(AzureGovernment()), nil

	case "dod", "usgovernmentl5":
		return pointer.To(AzureGovernmentL5()), nil
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
