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

// Cert must be at least 1024 bits long, and should be a base64 encoded PKCS12 certificate.
const dummyClientCertificate = "MIIHDwIBAzCCBsUGCSqGSIb3DQEHAaCCBrYEggayMIIGrjCCA2oGCSqGSIb3DQEHBqCCA1swggNXAgEAMIIDUAYJKoZIhvcNAQcBMF8GCSqGSIb3DQEFDTBSMDEGCSqGSIb3DQEFDDAkBBCiVw9+8AWoAiYYsSdT6tJOAgIIADAMBggqhkiG9w0CCQUAMB0GCWCGSAFlAwQBKgQQ2wK1uGYMi1PYXw48kRLyuoCCAuAZG8S/3Hl0FgHx5Sj+mjCGeTD4sppJLddw9dCAgbgIMy9lJxkzQron49rpjR4gtVCa2VmmrRtTN3I86yHOLOVV/dbvjHvKuzp4Ruak5qYJSj0MKXt7pJYxb3CoaSX4JTgjkQ2bUdwrmv+7UOeC95QFjQU6TmDDptfI4b/N3kaJiFl0mPpJc+zFN5Tw3rPWs/a7JHYjZYcgjKvefQm441ph09Pkp+v21iURAs6lgo4LLekiKP+kmhyiaOwwqHLFj7/sowL+v8oL01GddOPQkO6cz6UMVz2HVoEgxoSLWCl4IVmdqGps2CmCCgSGqRZOlbGG1CjXJnotjgEjCO7fuDswnYQVsJHj4/0XwdnvOdnylkDqt7KmvtZaPRX/Vx9B07SKGsf3wK7HlRGvXHq976byDFCq45jfXFIK69k4Wvr03q4meI94fPrkbqxTo29anbbTB46CsTUaz7xFg4UCDlgRXKizK2ZnXQ1xzvheg6vPlp3Rdstmo1cVpGZYvsT5X3nwMKMIwNbGxbM48y5Yiio2uDQWmMFlhfgYydozlyJ8LxGD32XmOEXMGaxB9qn2lROrUgCnDsY+jenWPP5LXmQtBW8PYjoJ/pIzzrR14Q80nd1n0au+PlSabPuKeMFmDZQDL6SIphB6sY8n+9Yq2JLpA27GonlBEBSen8mxa036oyXgvl5qtixACcyVCLzA5gNRhrEJu+GwfvCUZOss8xk6jUQ74K48P4/zselhPs9kscT4rulZdMAE9xzz60eElSOi2KKZZZClbR56KPk9RTLzbmkA/nZEPcZv484rnJCnxblV1vWJnXnmuHnBT4IQSg/HbURp+ST9EjBfSJ7v+B2jtorYpU3yjWMNgFF2rSbbdfdqeyFG3uXdLnIlUyPsNcvWM4fRRd2MJfO6nXA7RgmQq/hg0YeS2s45rUezz3DUcpjCijI3sGscLF/Qt6/oQIzhmLZtp4KTxZsorbwSHYY7MIIDPAYJKoZIhvcNAQcBoIIDLQSCAykwggMlMIIDIQYLKoZIhvcNAQwKAQKgggLpMIIC5TBfBgkqhkiG9w0BBQ0wUjAxBgkqhkiG9w0BBQwwJAQQoZN+nEXpMhnZi1UXw3je6AICCAAwDAYIKoZIhvcNAgkFADAdBglghkgBZQMEASoEEFBVIqGFWv0fw34+zJ+RJHwEggKAlQpKIr/XpjDMFc4FSVsCYm0Rfieek1xgLwlbnRtM40Gl6kteAhbaarTeqpoQXWhOEQgu4VZHBVojDK5M9yV5CiBN95DZDEhhyYFVVr3BME9YGwgIFinmJmFOaiGhG2W7PddYZEStXprTvEmq68FKKHsLELA/j1BgSGnTXZGKa3T9FnDH9PbWZqAWlvkay+i8ZHF0I9CXnm4O7cjiAlqqoWZaKeIY/q7iN4MOw95t2kU5h/++tBWnt/59ZehyqWf9h7zy64qeOiOx+j+x+JSBWKJez+n9cHfep4HWOt36y6anSdYcGUNyC5NkOHP8Rbea9Ov9FoHtDiGB6Mo2JSdF7HUbSV9CyHQUjvZ1b+AIBvhEpSyiVal260IIY3Mqx6Q5CUeU4vwCRXgGxjko4k6tMTnr0pqdzJeWTcR7Uo5e/Bph09k/+gsh1D482xuHNcZjOhE+PC/HP4Dd8LhORjsS80kL28odS9bYnVDQOgH4VGkhXoBbo8cvnyjy3WvwC5VceMd68Go/XlPJA3uXO/pGvGJDwuRL0p6UeX/1DwlZ905I1VbelvHkG9pKmXivk1t42NXdJqGpsXdiSUAXmqbKul2HH4HmFBdRjEM+d20ZvJN3+7/X0qKj7SaxPiJ++VlucrZzJJjYmDp3JbN+7Hxa7iC0AhHNU79J1WbLWiDffeq49onvFRUQwVlcZlU4QAk6OKm1aMfpLNYNuB3lv1Wz5Wg7UBFPj+YZTShKeCLVeqOIrFO68u8dHaWnRWWsnOku1EvnWDmWVwPGR3nQy7khOEfxHcjT0En6t8c2uiEDxzGZUXCbXPulYEoviy/WuOwMMTTkpGOflPLngOYnFxyseTElMCMGCSqGSIb3DQEJFTEWBBSDg33aQWb4YePRDLM+mxwY9j+b7TBBMDEwDQYJYIZIAWUDBAIBBQAEII28eubTebpTiOzuTDIuLXdwp0NQ3oqLQgzBYxZJRYC0BAhH3fm8hdm/1gICCAA="

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
