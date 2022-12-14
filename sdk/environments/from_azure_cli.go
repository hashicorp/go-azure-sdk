package environments

import "fmt"

/*
func FromCliName(environmentName string) (*Environment, error) {
	// check az-cli minimum version
	if err := azurecli.CheckAzVersion(azureCliMinimumVersion, nil); err != nil {
		return nil, err
	}

	var metadata azureCliSchema
	if err := azurecli.JSONUnmarshalAzCmd(&metadata, "cloud", "show", "--name", environmentName); err != nil {
		return nil, err
	}
*/

func FromAzureCLI(name string) (*Environment, error) {
	return nil, fmt.Errorf("TODO: not implemented")
}
