/*
 * Tatum API Reference
 *
 * # Welcome to the Tatum API Reference!  ## What is Tatum? ## What is Tatum?  Tatum offers a flexible framework to build, run, and scale blockchain apps fast. To learn more about the Tatum blockchain development framework, visit [our website](https://tatum.io/framework).  The Tatum API features powerful endpoints that simplify a complex blockchain into single API requests. Code for all supported blockchains using unified API calls.  ## Need help?  To chat with other developers, get help from the Support team, and engage with the thriving Tatum community, join  our [Discord server](https://discord.com/invite/tatum). For more information about how to work with Tatum, review the [online documentation](https://docs.tatum.io/).  ## About this API Reference  The Tatum API Reference is based on OpenAPI Specification v3.1.0 with a few [vendor extensions](https://github.com/Redocly/redoc/blob/master/docs/redoc-vendor-extensions.md) applied.  # Authentication  When using the Tatum API, you authenticate yourself with an **API key**. <SecurityDefinitions /> 
 *
 * API version: 3.16.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package tatum

type DeployTrcOffchainMnemXpub struct {
	// Name of the TRC token - stored as a symbol on Blockchain
	Symbol string `json:"symbol"`
	// max supply of TRC token.
	Supply string `json:"supply"`
	// Number of decimal points of the token.
	Decimals float64 `json:"decimals"`
	// Type of TRC token to create.
	Type_ string `json:"type"`
	// Description of the TRC token
	Description string `json:"description"`
	// URL of the project. Applicable for TRC-10 only.
	Url string `json:"url,omitempty"`
	// Base pair for TRC token. 1 token will be equal to 1 unit of base pair. Transaction value will be calculated according to this base pair.
	BasePair string `json:"basePair"`
	// Exchange rate of the base pair. Each unit of the created curency will represent value of baseRate*1 basePair.
	BaseRate float64 `json:"baseRate,omitempty"`
	Customer *CustomerRegistration `json:"customer,omitempty"`
	// Extended public key (xpub), from which address, where all initial supply will be stored, will be generated. Either xpub and derivationIndex, or address must be present, not both.
	Xpub string `json:"xpub"`
	// Derivation index for xpub to generate specific deposit address.
	DerivationIndex int32 `json:"derivationIndex"`
	// Mnemonic to generate private key for the deploy account of TRC, from which the gas will be paid (index will be used). If address is not present, mnemonic is used to generate xpub and index is set to 1. Either mnemonic and index or privateKey and address must be present, not both.
	Mnemonic string `json:"mnemonic"`
	// derivation index of address to pay for deployment of TRC
	Index int32 `json:"index"`
}