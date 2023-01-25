package environments

import (
	"context"
	"testing"
)

func TestFromEndpoint_Public(t *testing.T) {
	actual, err := FromEndpoint(context.Background(), "https://management.azure.com")
	if err != nil {
		t.Fatalf("loading from endpoint: %+v", err)
	}

	if actual.Name != "AzureCloud" {
		t.Fatalf("expected the Environment name to be `AzureCloud` but got %q", actual.Name)
	}
}
