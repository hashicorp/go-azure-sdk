// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcemanager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/internal/stringfmt"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// TODO: return a typed error here so that we could potentially change this error/expose this in the Provider
// TODO: the error should return the default error message shown below

var defaultRetryFunctions = []client.RequestRetryFunc{
	// NOTE: 429 is handled by the base library
	handleResourceProviderNotRegistered,
	RetryOn409ConflictFunc,
}

func handleResourceProviderNotRegistered(r *http.Response, o *odata.OData) (bool, error) {
	if o != nil && o.Error != nil && o.Error.Code != nil && strings.EqualFold(*o.Error.Code, "MissingSubscriptionRegistration") {
		return false, resourceProviderNotRegisteredError(*o.Error.Message)
	}

	return false, nil
}

func resourceProviderNotRegisteredError(message string) error {
	messageSplit := stringfmt.QuoteAndSplitString(">", message, 100)
	return fmt.Errorf(`The Resource Provider was not registered

Resource Providers (APIs) in Azure need to be registered before they can be used - however the Resource
Provider was not registered, and calling the API returned the following error:

%[1]s

The Azure Provider by default will automatically register certain Resource Providers at launch-time,
whilst it's possible to opt-out of this (which you may have done) 

Please ensure that this Resource Provider is properly registered, you can do this using the Azure CLI
for example to register the Resource Provider "Some.ResourceProvider" is registered run:

> az provider register --namespace "Some.ResourceProvider"

Resource Providers can take a while to register, you can check the status by running:

> az provider show --namespace "Some.ResourceProvider" --query "registrationState"

Once this outputs "Registered" the Resource Provider is available for use and you can re-run Terraform.
`, strings.Join(messageSplit, "\n"))
}

// RetryOn409ConflictFunc is a RequestRetryFunc that inspects HTTP 409 Conflict responses.
// It will always retry if a Retry-After header is provided by the server.
// If the Retry-After header is omitted, the response payload is evaluated to determine if the resource is in a
// non-terminal provisioning state (such as Updating, Creating, Deleting, or InProgress).
// If it is in a non-terminal state, it returns true, indicating a retry should be performed using the default backoff.
// Known terminal states (Failed, Canceled, Cancelled, Succeeded) will not be retried unless a Retry-After header is present.
func RetryOn409ConflictFunc(resp *http.Response, o *odata.OData) (bool, error) {
	if resp == nil || resp.StatusCode != http.StatusConflict {
		return false, nil
	}

	if _, ok := resp.Header["Retry-After"]; ok {
		return true, nil
	}

	if resp.Body != nil {
		respBody, err := io.ReadAll(resp.Body)
		if err == nil {
			// Reassign the body so it can be consumed later
			resp.Body = io.NopCloser(bytes.NewBuffer(respBody))

			var state struct {
				Status     string `json:"status"`
				Properties struct {
					ProvisioningState string `json:"provisioningState"`
				} `json:"properties"`
			}

			if err := json.Unmarshal(respBody, &state); err == nil {
				s := state.Status
				if state.Properties.ProvisioningState != "" {
					s = state.Properties.ProvisioningState
				}

				if s != "" {
					// Check for known terminal states
					if strings.EqualFold(s, "Failed") || strings.EqualFold(s, "Canceled") || strings.EqualFold(s, "Cancelled") || strings.EqualFold(s, "Succeeded") {
						return false, nil
					}
					// If it's a non-terminal state, retry
					return true, nil
				}
			}
		}
	}

	return false, nil
}
