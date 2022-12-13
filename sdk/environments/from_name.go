package environments

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
)

func FromName(name string) (*Environment, error) {
	switch strings.ToLower(name) {
	case "global", "public":
		return pointer.To(AzurePublic()), nil
	}

	return nil, fmt.Errorf("no environment was found with the name %q", name)
}

// TODO: From Name
//func FromNamed(env string) (*Environment, error) {
//	switch strings.ToLower(env) {
//	case "public", "global":
//		return &Global, nil
//	case "usgovernment", "usgovernmentl4":
//		return &USGovernmentL4, nil
//	case "dod", "usgovernmentl5":
//		return &USGovernmentL5, nil
//	case "canary":
//		return &Canary, nil
//	case "china":
//		return &China, nil
//	}
//
//	return nil, fmt.Errorf("invalid environment specified: %s", env)
//}
