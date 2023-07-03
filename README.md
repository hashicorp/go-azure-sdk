# `hashicorp/go-azure-sdk`

This SDK is an opinionated Go SDK for Azure Resource Manager, primarily intended to be used in the [Terraform AzureRM Provider](https://github.com/hashicorp/terraform-provider-azurerm) and the [Packer Azure Plugin](https://github.com/hashicorp/packer-plugin-azure).

The SDKs within this repository are generated using the data in [the Azure Rest API Specifications repository](https://github.com/Azure/azure-rest-api-specs).

At this time this SDK makes use of [the `Azure/go-autorest` base layer](https://github.com/Azure/go-autorest), however this will change in a future release.

## Repository Structure

This repository contains:

* `./docs` - usage documentation.
* `./resource-manager` - the Resource Manager Services supported by this SDK.

In addition, a `README.md` file can be found within each Service / API Version in the `./resource-manager` directory highlighting how that SDK can be used ([example](https://github.com/hashicorp/go-azure-sdk/tree/main/resource-manager/aadb2c/2021-04-01-preview/tenants)).

## Further Documentation

Further information [can be found in the `./docs` directory](./docs), and also the `README.md` file within each Service / API Version in the `./resource-manager` directory ([example](https://github.com/hashicorp/go-azure-sdk/tree/main/resource-manager/aadb2c/2021-04-01-preview/tenants)).

