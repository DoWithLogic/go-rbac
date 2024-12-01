package main

import "github.com/DoWithLogic/go-rbac/cmd"

// main is entrypoint of application
//
//	@title	Go Role and Scope Based Access Control (RBAC)
//	@description.markdown
//	@termsOfService				https://github.com/DoWithLogic/go-rbac
//	@BasePath					/api/v1/rbac
//
//	@securityDefinitions.apiKey	BearerToken
//	@in							header
//	@name						Authorization
func main() {
	cmd.Start()
}
