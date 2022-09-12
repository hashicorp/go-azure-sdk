package environments

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/hashicorp/go-cleanhttp"
)

type metadataSchema struct {
	ActiveDirectoryDataLake string                 `json:"activeDirectoryDataLake"`
	Authentication          metadataAuthentication `json:"authentication"`
	Batch                   string                 `json:"batch"`
	Gallery                 string                 `json:"gallery"`
	Graph                   string                 `json:"graph"`
	GraphAudience           string                 `json:"graphAudience"`
	Media                   string                 `json:"media"`
	Name                    string                 `json:"name"`
	Portal                  string                 `json:"portal"`
	ResourceManager         string                 `json:"resourceManager"`
	SqlManagement           string                 `json:"sqlManagement"`
	Suffixes                metadataSuffixes       `json:"suffixes"`
	VmImageAliasDoc         string                 `json:"vmImageAliasDoc"`
}

type metadataAuthentication struct {
	Audiences        []string `json:"audiences"`
	IdentityProvider string   `json:"identityProvider"`
	LoginEndpoint    string   `json:"loginEndpoint"`
	Tenant           string   `json:"tenant"`
}

type metadataSuffixes struct {
	AcrLoginServer                      string `json:"acrLoginServer"`
	AzureDataLakeAnalyticsCatalogAndJob string `json:"azureDataLakeAnalyticsCatalogAndJob"`
	AzureDataLakeStoreFileSystem        string `json:"azureDataLakeStoreFileSystem"`
	AzureFrontDoorEndpointSuffix        string `json:"azureFrontDoorEndpointSuffix"`
	KeyVaultDns                         string `json:"keyVaultDns"`
	SqlServerHostname                   string `json:"sqlServerHostname"`
	Storage                             string `json:"storage"`
}

func FromMetadataEndpoint(ctx context.Context, metadataHost string, environmentName *string) (*Environment, error) {
	if metadataHost == "" {
		return nil, fmt.Errorf("custom metadata host was not specified")
	}

	proto := "https"
	for _, s := range []string{"localhost", "127."} {
		if strings.Index(metadataHost, s) == 0 {
			proto = "http"
			break
		}
	}
	endpoint := fmt.Sprintf("%s://%s/metadata/endpoints?api-version=2019-05-01", proto, metadataHost)
	client := cleanhttp.DefaultClient()

	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("retrieving environment metadata from custom metadata host %q: %+v", metadataHost, err)
	}

	if resp.StatusCode != http.StatusOK {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("unexpected status %d received from metadata service at %q, could not read response body", resp.StatusCode, metadataHost)
		}
		return nil, fmt.Errorf("unexpected status %d received from metadata service at %q with response: %s", resp.StatusCode, metadataHost, respBody)
	}

	var metadata []metadataSchema
	if err := json.NewDecoder(resp.Body).Decode(&metadata); err != nil {
		return nil, err
	}

	// while the array contains values
	for _, env := range metadata {
		if environmentName != nil && strings.EqualFold(env.Name, *environmentName) || (environmentName == nil && len(metadata) == 1) {
			// if resourceManager endpoint is empty, assume it's the same endpoint
			if env.ResourceManager == "" {
				env.ResourceManager = fmt.Sprintf("https://%s/", metadataHost)
			}

			return &Environment{
				AzureADEndpoint: azureAdEndpoint(env.Authentication.LoginEndpoint),

				AADGraph: Api{
					AppId:    PublishedApis["AzureActiveDirectoryGraph"],
					Endpoint: apiEndpoint(env.Graph),
				},
				BatchManagement: Api{
					AppId:    PublishedApis["AzureBatch"],
					Endpoint: apiEndpoint(env.Batch),
				},
				DataLake: Api{
					AppId:    PublishedApis["AzureDataLake"],
					Endpoint: apiEndpoint(env.ActiveDirectoryDataLake),
				},
				KeyVault: Api{
					AppId:    PublishedApis["AzureKeyVault"],
					Endpoint: apiEndpoint(fmt.Sprintf("https://%s", env.Suffixes.KeyVaultDns)),
				},
				ResourceManager: Api{
					AppId:    PublishedApis["AzureServiceManagement"],
					Endpoint: apiEndpoint(env.ResourceManager),
				},
				SQLDatabase: Api{
					AppId:    PublishedApis["AzureSqlDatabase"],
					Endpoint: apiEndpoint(fmt.Sprintf("https://%s", env.Suffixes.SqlServerHostname)),
				},

				MSGraph:             ApiUnavailable,
				OperationalInsights: ApiUnavailable,
				OSSRDBMS:            ApiUnavailable,
				ServiceBus:          ApiUnavailable,
				ServiceManagement:   ApiUnavailable,
				Storage:             StoragePublic,
				Synapse:             ApiUnavailable,
			}, nil

		}
	}

	e := ""
	if environmentName != nil {
		e = fmt.Sprintf(" for environment %q", *environmentName)
	}
	return nil, fmt.Errorf("unable to locate metadata%s from custom metadata host %q", e, metadataHost)
}
