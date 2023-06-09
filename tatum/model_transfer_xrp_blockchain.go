/*
 * Tatum API Reference
 *
 * # Welcome to the Tatum API Reference!  ## What is Tatum? ## What is Tatum?  Tatum offers a flexible framework to build, run, and scale blockchain apps fast. To learn more about the Tatum blockchain development framework, visit [our website](https://tatum.io/framework).  The Tatum API features powerful endpoints that simplify a complex blockchain into single API requests. Code for all supported blockchains using unified API calls.  ## Need help?  To chat with other developers, get help from the Support team, and engage with the thriving Tatum community, join  our [Discord server](https://discord.com/invite/tatum). For more information about how to work with Tatum, review the [online documentation](https://docs.tatum.io/).  ## About this API Reference  The Tatum API Reference is based on OpenAPI Specification v3.1.0 with a few [vendor extensions](https://github.com/Redocly/redoc/blob/master/docs/redoc-vendor-extensions.md) applied.  # Authentication  When using the Tatum API, you authenticate yourself with an **API key**. <SecurityDefinitions /> 
 *
 * API version: 3.16.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package tatum

type TransferXrpBlockchain struct {
	// XRP account address. Must be the one used for generating deposit tags.
	FromAccount string `json:"fromAccount"`
	// Blockchain address to send assets
	To string `json:"to"`
	// Amount to be sent, in XRP.
	Amount string `json:"amount"`
	// Secret for account. Secret, or signature Id must be present.
	FromSecret string `json:"fromSecret"`
	// Fee to be paid, in XRP. If omitted, current fee will be calculated.
	Fee string `json:"fee,omitempty"`
	// Source tag of sender account, if any.
	SourceTag int32 `json:"sourceTag,omitempty"`
	// Destination tag of recipient account, if any.
	DestinationTag int32 `json:"destinationTag,omitempty"`
}
