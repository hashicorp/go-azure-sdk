package environments_test

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-azure-sdk/internal/test"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"math/rand"
	"testing"
)

func TestFromMetadataEndpoint(t *testing.T) {
	ctx := context.Background()
	port := 8000 + rand.Intn(999)
	metadataHost := fmt.Sprintf("localhost:%d", port)

	done := test.MetadataStubServer(ctx, port)
	defer func() {
		done <- true
	}()

	envName := "AzureCloud"
	_, err := environments.FromMetadataEndpoint(ctx, metadataHost, &envName)
	if err != nil {
		t.Fatal(err)
	}
}
