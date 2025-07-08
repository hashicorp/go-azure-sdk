// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcemanager

import (
	"net/http"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
)

type expectedResponse struct {
	dropsConnection      bool
	httpStatusCode       *int
	status               *status
	useStatus            bool
	useProvisioningState bool
}

func responseThatDropsTheConnection() expectedResponse {
	return expectedResponse{
		dropsConnection: true,
	}
}

func responseWithStatusInProvisioningState(input status) expectedResponse {
	return expectedResponse{
		status:               pointer.To(input),
		useProvisioningState: true,
	}
}

func responseWithStatusInStatusField(input status) expectedResponse {
	return expectedResponse{
		status:    pointer.To(input),
		useStatus: true,
	}
}

func responseWithHttpStatusCode(input int) expectedResponse {
	return expectedResponse{
		httpStatusCode: pointer.To(input),
	}
}

func responseWithNoTerminalState() expectedResponse {
	return expectedResponse{
		status: pointer.ToEnum[status](""),
	}
}

func dropConnection(t *testing.T, w http.ResponseWriter) {
	httpHijacker, ok := w.(http.Hijacker)
	if !ok {
		t.Fatalf("`w` didn't implement `http.httpHijacker`")
	}
	conn, _, err := httpHijacker.Hijack()
	if err != nil {
		return
	}
	conn.Close()
}
