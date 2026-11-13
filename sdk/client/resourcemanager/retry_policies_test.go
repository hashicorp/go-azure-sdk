// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcemanager

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
)

func TestRequest_ResourceProviderNotRegistered(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()

	message := "The subscription is not registered to use namespace 'Microsoft.Foo'. See https://aka.ms/rps-not-found for how to register subscriptions."
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		_, _ = w.Write([]byte(`{"error":{"code":"MissingSubscriptionRegistration","message":"` + message + `"}}`))
	}))
	defer server.Close()

	resourceManagerClient := &Client{
		Client:     client.NewClient(server.URL, "Example", "2020-01-01"),
		apiVersion: "2020-01-01",
	}
	req, err := resourceManagerClient.NewRequest(ctx, client.RequestOptions{
		ContentType:         "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{http.StatusOK},
		HttpMethod:          http.MethodPut,
		Path:                "/subscriptions/000/resourceGroups/rg/providers/Microsoft.Foo/things/t1",
	})
	if err != nil {
		t.Fatalf("building request: %v", err)
	}

	// the advice on registering the Resource Provider should be returned, rather than only the status and the API's message
	_, err = req.Execute(ctx)
	if err == nil {
		t.Fatal("expected an error but didn't get one")
	}

	expected := resourceProviderNotRegisteredError(message).Error()
	if err.Error() != expected {
		t.Fatalf("expected error:\n%s\n\ngot:\n%s", expected, err.Error())
	}
}
