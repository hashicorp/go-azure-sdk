package test

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func MetadataStubServer(ctx context.Context, port int) chan bool {
	handler := http.NewServeMux()

	handler.HandleFunc("/metadata/endpoints", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		apiVersion := q.Get("api-version")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		switch apiVersion {
		case "1.0", "2015-01-01":
			fmt.Fprint(w, `{"galleryEndpoint":"https://gallery.azure.com/","graphEndpoint":"https://graph.windows.net/","portalEndpoint":"https://portal.azure.com/","authentication":{"loginEndpoint":"https://login.microsoftonline.com/","audiences":["https://management.core.windows.net/","https://management.azure.com/"]}}`)
		case "2018-01-01":
			fmt.Fprint(w, `{"galleryEndpoint":"https://gallery.azure.com/","graphEndpoint":"https://graph.windows.net/","portalEndpoint":"https://portal.azure.com/","graphAudience":"https://graph.windows.net/","authentication":{"loginEndpoint":"https://login.microsoftonline.com/","audiences":["https://management.core.windows.net/","https://management.azure.com/"],"tenant":"common","identityProvider":"AAD"},"cloudEndpoint":{"public":{"endpoint":"management.azure.com","locations":["australiacentral","australiacentral2","australiaeast","australiasoutheast","brazilsouth","brazilsoutheast","brazilus","canadacentral","canadaeast","centralindia","centralus","centraluseuap","eastasia","eastus","eastus2","eastus2euap","francecentral","francesouth","germanynorth","germanywestcentral","japaneast","japanwest","jioindiacentral","jioindiawest","koreacentral","koreasouth","northcentralus","northeurope","norwayeast","norwaywest","qatarcentral","southafricanorth","southafricawest","southcentralus","southeastasia","southindia","swedencentral","swedensouth","switzerlandnorth","switzerlandwest","uaecentral","uaenorth","uksouth","ukwest","westcentralus","westeurope","westindia","westus","westus2","westus3","israelcentral","italynorth","malaysiasouth","polandcentral","taiwannorth","taiwannorthwest"]},"chinaCloud":{"endpoint":"management.chinacloudapi.cn","locations":["chinaeast","chinanorth","chinanorth2","chinaeast2","chinanorth3","chinaeast3"]},"usGovCloud":{"endpoint":"management.usgovcloudapi.net","locations":["usgovvirginia","usgoviowa","usdodeast","usdodcentral","usgovtexas","usgovarizona"]},"germanCloud":{"endpoint":"management.microsoftazure.de","locations":["germanycentral","germanynortheast"]}}}`)
		case "2019-05-01":
			fmt.Fprint(w, `[{"portal":"https://portal.azure.com","authentication":{"loginEndpoint":"https://login.microsoftonline.com/","audiences":["https://management.core.windows.net/","https://management.azure.com/"],"tenant":"common","identityProvider":"AAD"},"media":"https://rest.media.azure.net","graphAudience":"https://graph.windows.net/","graph":"https://graph.windows.net/","name":"AzureCloud","suffixes":{"azureDataLakeStoreFileSystem":"azuredatalakestore.net","acrLoginServer":"azurecr.io","sqlServerHostname":"database.windows.net","azureDataLakeAnalyticsCatalogAndJob":"azuredatalakeanalytics.net","keyVaultDns":"vault.azure.net","storage":"core.windows.net","azureFrontDoorEndpointSuffix":"azurefd.net"},"batch":"https://batch.core.windows.net/","resourceManager":"https://management.azure.com/","vmImageAliasDoc":"https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/arm-compute/quickstart-templates/aliases.json","activeDirectoryDataLake":"https://datalake.azure.net/","sqlManagement":"https://management.core.windows.net:8443/","gallery":"https://gallery.azure.com/"},{"portal":"https://portal.azure.cn","authentication":{"loginEndpoint":"https://login.chinacloudapi.cn","audiences":["https://management.core.chinacloudapi.cn","https://management.chinacloudapi.cn"],"tenant":"common","identityProvider":"AAD"},"media":"https://rest.media.chinacloudapi.cn","graphAudience":"https://graph.chinacloudapi.cn","graph":"https://graph.chinacloudapi.cn","name":"AzureChinaCloud","suffixes":{"acrLoginServer":"azurecr.cn","sqlServerHostname":"database.chinacloudapi.cn","keyVaultDns":"vault.azure.cn","storage":"core.chinacloudapi.cn","azureFrontDoorEndpointSuffix":""},"batch":"https://batch.chinacloudapi.cn","resourceManager":"https://management.chinacloudapi.cn","vmImageAliasDoc":"https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/arm-compute/quickstart-templates/aliases.json","sqlManagement":"https://management.core.chinacloudapi.cn:8443","gallery":"https://gallery.chinacloudapi.cn"},{"portal":"https://portal.azure.us","authentication":{"loginEndpoint":"https://login.microsoftonline.us","audiences":["https://management.core.usgovcloudapi.net","https://management.usgovcloudapi.net"],"tenant":"common","identityProvider":"AAD"},"media":"https://rest.media.usgovcloudapi.net","graphAudience":"https://graph.windows.net","graph":"https://graph.windows.net","name":"AzureUSGovernment","suffixes":{"acrLoginServer":"azurecr.us","sqlServerHostname":"database.usgovcloudapi.net","keyVaultDns":"vault.usgovcloudapi.net","storage":"core.usgovcloudapi.net","azureFrontDoorEndpointSuffix":""},"batch":"https://batch.core.usgovcloudapi.net","resourceManager":"https://management.usgovcloudapi.net","vmImageAliasDoc":"https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/arm-compute/quickstart-templates/aliases.json","sqlManagement":"https://management.core.usgovcloudapi.net:8443","gallery":"https://gallery.usgovcloudapi.net"},{"portal":"https://portal.microsoftazure.de","authentication":{"loginEndpoint":"https://login.microsoftonline.de","audiences":["https://management.core.cloudapi.de","https://management.microsoftazure.de"],"tenant":"common","identityProvider":"AAD"},"media":"https://rest.media.cloudapi.de","graphAudience":"https://graph.cloudapi.de","graph":"https://graph.cloudapi.de","name":"AzureGermanCloud","suffixes":{"sqlServerHostname":"database.cloudapi.de","keyVaultDns":"vault.microsoftazure.de","storage":"core.cloudapi.de","azureFrontDoorEndpointSuffix":""},"batch":"https://batch.cloudapi.de","resourceManager":"https://management.microsoftazure.de","vmImageAliasDoc":"https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/arm-compute/quickstart-templates/aliases.json","sqlManagement":"https://management.core.cloudapi.de:8443","gallery":"https://gallery.cloudapi.de"}]`)
		default:
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, `{"error":"unrecognized api-version"}`)
			return
		}
	})

	done := make(chan bool, 1)
	server := &http.Server{
		Addr:    fmt.Sprintf("127.0.0.1:%d", port),
		Handler: handler,
	}

	go func() {
		<-done
		err := server.Shutdown(ctx)
		if err != nil {
			log.Fatalf("failed to gracefully shut down ARM Metadata stub server: %v", err)
		}
	}()

	go func() {
		log.Printf("ARM Metadata Stub Service listening on 127.0.0.1:%d\n", port)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("server.ListenAndServe: %v", err)
		}
	}()

	return done
}
