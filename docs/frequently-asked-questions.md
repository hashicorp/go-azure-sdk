# Frequently Asked Questions

## What's the source-data used for this SDK?

This SDK is not at all generated [from the Swagger/OpenAPI Definitions within the `Azure/azure-rest-api-specs` repository](https://github.com/Azure/azure-rest-api-specs).

## Which Services are Available?

At this time only support Resource Manager Services are supported - the services supported [can be found in the `./resource-manager` directory](../resource-manager).

## How can I request a new Service / API Version?

You can [request support for a new Service/API Version by opening a new GitHub Issue](https://github.com/hashicorp/go-azure-sdk/issues/new/choose) - please open an issue and we can look into adding support for this particular Service / API Version.

## I think I've found a bug in the SDK, where should I report it?

We'd really appreciate it if you'd [open a new GitHub Issue](https://github.com/hashicorp/go-azure-sdk/issues/new/choose) so that we can take a look.

Since we release new versions of the SDK fairly often, it'd be helpful to know:

1. Which version of the SDK you're using.
2. Which Service / API Version combination you're using.
3. An example snippet highlighting the issue.
4. Debug logs (if available).

## How is this SDK versioned?

New versions of this SDK are released as required - and are [based on the Swagger/OpenAPI Definitions within the `Azure/azure-rest-api-specs` repository](https://github.com/Azure/azure-rest-api-specs).

We follow the version strategy `v0.YYYYMMDD.1HHmmSS` (for example for an SDK released on `2022-06-30` at `09:30:00`, we'll use the version `v0.20220630.1093000`).
