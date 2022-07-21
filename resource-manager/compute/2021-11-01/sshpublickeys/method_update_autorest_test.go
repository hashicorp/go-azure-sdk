package sshpublickeys_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/go-azure-sdk/auth"
	"github.com/hashicorp/go-azure-sdk/environments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/sshpublickeys"
)

func TestUpdate(t *testing.T) {
	client := sshpublickeys.NewSshPublicKeysClientWithBaseURI(string(environments.Global.ResourceManager.Endpoint))
	client.Client2.Client.Authorizer = (&auth.ClientCredentialsConfig{
		Environment:  environments.Global,
		TenantID:     os.Getenv("ARM_TENANT_ID"),
		ClientID:     os.Getenv("ARM_CLIENT_ID"),
		ClientSecret: os.Getenv("ARM_CLIENT_SECRET"),
		Scopes:       []string{environments.Global.ResourceManager.DefaultScope()},
		TokenVersion: auth.TokenVersion2,
	}).TokenSource(context.TODO(), auth.ClientCredentialsSecretType)

	id := sshpublickeys.NewSshPublicKeyID(os.Getenv("ARM_SUBSCRIPTION_ID"), "aaacctest-tbamford", "aaa2")
	key := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC6IPaOit/uGk6NZ0bfLXi24+jxs3NfvPkajVqdlq9IYEGtqJSTstEaaxtOYyCDIZajv6K/C3b3sMxPL5adnR70bTu/PK7ge3jpHBjc0eVVSk8bP4/QYy2nDNszhwXuSP8xjJkBhVbq4gYSkQ/N1FPgoxOYIhL44RChXEx0A+GmWqzoqgh6stj0mMTV2L13pFDhLURLdGR4KRV+M+hI46CwpL5ZhmZ032Xy0ANXWYSl86AqGVbi8+JDGVCzIDXlmdw4CoGXbO4cqRfCqa/jbA4F8XiIpKtzG+5ZGQ69Z6dgiNcPiG5BNQOhFSs5BBrjKYb3k8RrmuJN62XFE7GBDQn0A1S98aE1umrI8+OpFeUE/Rl9D8yT7ItAyjYWYI0h0HM34HjCCLLgJMF6sSS/dB34sAFy9OWrl1Q6PJnqCgJwlFg+STQBjGT8grqqSJu+W2PHCnWMYeqqVuMB7DreY56sHfvhHZVq8idKYxisC+HMfu74fxUa00FbAgnGB9a4bVgNmrrkLhVYKa7xCFkACIBbo9uF72DwuwUBqbqQkytyK7FY93JwY3Fb9KDnMrk0+uNfS7UoEWs852Qkg1d2pumhJ/5/7ae7sM8PsSnArCDjvj9aBSlbVI8Ob82awu0p0MYa7TAC4uk2jLqi1/LIbLRqa/EjhHsXV9/ZyyH8k1abGQ== tom@bamford.io"
	input := sshpublickeys.SshPublicKeyUpdateResource{
		Properties: &sshpublickeys.SshPublicKeyResourceProperties{
			PublicKey: &key,
		},
	}

	res, err := client.Update(context.TODO(), id, input)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n", res)
}
