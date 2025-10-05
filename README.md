# powerbi-go

Go client for the Power BI REST API

# Example

```go
package main

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	
	"github.com/stpabhi/powerbi-go"
	"github.com/stpabhi/powerbi-go/types"
)

func main() {
	clientID := "<client_id>"
	clientSecret := "<client_secret>"
	tenantID := "<tenant_id>"

	ctx := context.Background()

	// Authenticate with Azure Go SDK
	cred, err := azidentity.NewClientSecretCredential(tenantID, clientID, clientSecret, nil)
	if err != nil {
		panic(err)
	}
	token, err := cred.GetToken(ctx, policy.TokenRequestOptions{
		Scopes: []string{"https://analysis.windows.net/powerbi/api/.default"},
	})
	if err != nil {
		panic(err)
	}

	// Create Power BI client
	pbi := powerbi.NewFromToken(token.Token)
	groups, err := pbi.Groups.List(ctx, types.ListGroupsOptions{})
	if err != nil {
		panic(err)
	}

	for _, g := range groups {
		println(g.String())
	}
}
```
