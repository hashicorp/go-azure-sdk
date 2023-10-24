// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package auth_test

import (
	"context"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/go-azure-sdk/sdk/internal/test"
)

const dummyClientCertificate = "MIIEWQIBAzCCBB8GCSqGSIb3DQEHAaCCBBAEggQMMIIECDCCAicGCSqGSIb3DQEHBqCCAhgwggIUAgEAMIICDQYJKoZIhvcNAQcBMBwGCiqGSIb3DQEMAQYwDgQIVeyTunvg5bwCAggAgIIB4GHlNvzJm/847/2WxzcBnQentkW7NbxOs9JazFVSiP1X+uQs2Ef2OzOw5ZdqPctEwhyeJKwmxNv1MHzwhxurYLpEJ7l+01eCLt395abdcyyopzzkbUx28xkexe5Or5OrJUJju0YkT29jqcOuDBD46DKdZGc2/imlV+9/EzIv8a+P+dT+LLWl3ZJBuVSbzWCuiGZemzYfw8tiOYhIwCno5GWC0P14/nD1vmgn8/rhL9YPPu7HqCQqf9HJb36MVvfC/gUsNaW/QuIrGh7CxBfTUdj+esFRsP/A5HsURuronq7wS76xvx4uKAMgGfqxcvwvH/3jV5ROEp1iPRcrIKLH8aQQ8orgmvzodHZr7X4p41RtnJcmRF4894R1e5t61LdFjqAo/F8WcKo1/cPdTpyY2v9mKJQCqGAY47CxsCgo7XxKOA0tmhutyBxHlHupz0G1AZ4dOoJihuj71Rui59MyHGI38BpUHPxxRP7O8Hnx12QkfqxHNucqV9lsVWhOIe8G6llTD5eMMazfip4jpzS2uSlba/7UPstWuYtXSUqQlPTu+tXyVuEUokZbg2dzI3e4RCG8YuLYfjYIpvW0Pi3aioXmQkrJs4ViAwwYDCp4vmbY2IuFVwQGug+cXrpSAD5MKjCCAdkGCSqGSIb3DQEHAaCCAcoEggHGMIIBwjCCAb4GCyqGSIb3DQEMCgECoIIBhjCCAYIwHAYKKoZIhvcNAQwBAzAOBAjkIUiRa+TRrAICCAAEggFgHYhMxkmd5ZNTEmyLB8MRwlZmG0/shVPKNVTLxX7WLPxSAIW7PRQKOF2NiIsKDoZaznj4ie1qU0b14MAL9XN0aQaB1uN6QMt9H818qAF3rRdTj2RchIhFgLuxeYFKrMZcKwl5//IqZ0Rm/4fjVBP2HK7VRkgmtqjPf00wpoMQKMd+8UXqljhl0ydDb5Fdk7lGhCV9SV+jqY1qHZeMTGn23+ScJKCVsaW7tY1wjJWp9+tJhpHnXNKTI5hhgyHpx/Wy9x+W4deCb1aVFEM5DSKgXN5jiCywcJ8fnbbU3pb1QVYyWsFwFHXEWAUHut0E9b4uzdVlGoHE2JOA0Mx7l4yfbRN0nYF2olHg/Ppz7G6eTceM0KAvg2kcjAcrAXmm/0Z/KFLyBJZFY6p+zv3UDQ+UmyVHI+/QPrtyetFdHkEuZBm7OB5c1BFegQ9rzjdPauk85imDwOCPrb/87tfLDopCMjElMCMGCSqGSIb3DQEJFTEWBBS1BAW9xO/B/015SD/UniWesRmJtDAxMCEwCQYFKw4DAhoFAAQUzDlelwCxOEwh25GbS+rFBIPMzKAECLPBXdnOWTFsAgIIAA=="

func TestClientCertificateAuthorizer(t *testing.T) {
	ctx := context.Background()
	env := environments.AzurePublic()

	auth.Client = &test.AzureADAccessTokenMockClient{
		Authorization: *env.Authorization,
	}

	opts := auth.ClientCertificateAuthorizerOptions{
		Environment:  *env,
		Api:          env.MicrosoftGraph,
		TenantId:     "00000000-1111-0000-0000-000000000000",
		AuxTenantIds: test.AuxiliaryTenantIds,
		ClientId:     "11111111-0000-0000-0000-000000000000",
		Pkcs12Data:   test.Base64DecodeCertificate(t, dummyClientCertificate),
		Pkcs12Pass:   "certpassword",
	}

	authorizer, err := auth.NewClientCertificateAuthorizer(ctx, opts)
	if err != nil {
		t.Fatalf("NewClientCertificateAuthorizer(): %v", err)
	}

	if authorizer == nil {
		t.Fatal("authorizer is nil, expected Authorizer")
	}

	if _, err = testObtainAccessToken(ctx, authorizer); err != nil {
		t.Fatal(err)
	}
}

func TestAccClientCertificateAuthorizer(t *testing.T) {
	test.AccTest(t)

	ctx := context.Background()

	env, err := environments.FromName(test.Environment)
	if err != nil {
		t.Fatal(err)
	}

	opts := auth.ClientCertificateAuthorizerOptions{
		Environment:  *env,
		Api:          env.MicrosoftGraph,
		TenantId:     test.TenantId,
		AuxTenantIds: test.AuxiliaryTenantIds,
		ClientId:     test.ClientId,
		Pkcs12Data:   test.Base64DecodeCertificate(t, test.ClientCertificate),
		Pkcs12Path:   test.ClientCertificatePath,
		Pkcs12Pass:   test.ClientCertPassword,
	}

	authorizer, err := auth.NewClientCertificateAuthorizer(ctx, opts)
	if err != nil {
		t.Fatalf("NewClientCertificateAuthorizer(): %v", err)
	}

	if authorizer == nil {
		t.Fatal("authorizer is nil, expected Authorizer")
	}

	if _, err = testObtainAccessToken(ctx, authorizer); err != nil {
		t.Fatal(err)
	}
}

func TestAccClientCertificateAuthorizerKeychain(t *testing.T) {
	test.AccTest(t)

	ctx := context.Background()

	env, err := environments.FromName(test.Environment)
	if err != nil {
		t.Fatal(err)
	}

	opts := auth.ClientCertificateAuthorizerOptions{
		Environment:      *env,
		Api:              env.MicrosoftGraph,
		TenantId:         test.TenantId,
		AuxTenantIds:     test.AuxiliaryTenantIds,
		ClientId:         test.ClientId,
		SignerCommonName: test.SignerCommonName,
	}

	authorizer, err := auth.NewClientCertificateAuthorizer(ctx, opts)
	if err != nil {
		t.Fatalf("NewClientCertificateAuthorizer(): %v", err)
	}

	if authorizer == nil {
		t.Fatal("authorizer is nil, expected Authorizer")
	}

	if _, err = testObtainAccessToken(ctx, authorizer); err != nil {
		t.Fatal(err)
	}
}
