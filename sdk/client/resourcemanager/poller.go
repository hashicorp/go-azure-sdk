package resourcemanager

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
)

func PollerFromResponse(response *client.Response, client *Client) (poller pollers.Poller, err error) {
	// TODO: support for checking both the LRO and then the ProvisioningState when this is an ARM Resource
	// presumably that ^ would be a `unified` poller or something?
	lroPoller, lroErr := longRunningOperationPollerFromResponse(response, client.Client)
	if lroErr != nil {
		return pollers.Poller{}, fmt.Errorf("building long-running-operation poller: %+v", err)
	}

	return pollers.NewPoller(lroPoller, lroPoller.initialRetryDuration), nil
}
