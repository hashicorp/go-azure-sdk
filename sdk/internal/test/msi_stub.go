// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package test

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func ManagedIdentityStubServer(ctx context.Context, port int, token string) chan bool {
	handler := http.NewServeMux()

	handler.HandleFunc("/metadata/identity/oauth2/token", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		clientId := q.Get("client_id")
		resource := q.Get("resource")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		if r.Header.Get("Metadata") != "true" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, `{}`)
			return
		}

		fmt.Fprintf(w, `{"access_token":"%s","client_id":"%s","expires_in":"86391","expires_on":"1611701390","ext_expires_in":"86399","not_before":"1611614690","resource":"%s","token_type":"Bearer"}`,
			token, clientId, resource)
	})

	handler.HandleFunc("/metadata", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, `
`)
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
			log.Fatalf("failed to gracefully shut down MSI stub server: %v", err)
		}
	}()

	go func() {
		log.Printf("MSI Stub Service listening on 127.0.0.1:%d\n", port)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("server.ListenAndServe: %v", err)
		}
	}()

	return done
}
