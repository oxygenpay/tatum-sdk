/*
 * Tatum API Reference
 *
 * # Welcome to the Tatum API Reference!  ## What is Tatum? ## What is Tatum?  Tatum offers a flexible framework to build, run, and scale blockchain apps fast. To learn more about the Tatum blockchain development framework, visit [our website](https://tatum.io/framework).  The Tatum API features powerful endpoints that simplify a complex blockchain into single API requests. Code for all supported blockchains using unified API calls.  ## Need help?  To chat with other developers, get help from the Support team, and engage with the thriving Tatum community, join  our [Discord server](https://discord.com/invite/tatum). For more information about how to work with Tatum, review the [online documentation](https://docs.tatum.io/).  ## About this API Reference  The Tatum API Reference is based on OpenAPI Specification v3.1.0 with a few [vendor extensions](https://github.com/Redocly/redoc/blob/master/docs/redoc-vendor-extensions.md) applied.  # Authentication  When using the Tatum API, you authenticate yourself with an **API key**. <SecurityDefinitions /> 
 *
 * API version: 3.16.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package tatum

type ApproveNftSpendingKms struct {
	// Blockchain to work with.
	Chain string `json:"chain"`
	// Address of the ERC20 token, which is used for buying NFT asset from the marketplace.
	ContractAddress string `json:"contractAddress"`
	// Address of the auction smart contract - new spender.
	Spender string `json:"spender"`
	// True if asset is NFT of type ERC721, false if ERC1155.
	IsErc721 bool `json:"isErc721"`
	// ID of token, if transaction is for ERC-721 or ERC-1155.
	TokenId string `json:"tokenId"`
	// Identifier of the private key associated in signing application. Private key, or signature Id must be present.
	SignatureId string `json:"signatureId"`
	// If signatureId is mnemonic-based, this is the index to the specific address from that mnemonic.
	Index float64 `json:"index,omitempty"`
	// Nonce to be set to Ethereum transaction. If not present, last known nonce will be used.
	Nonce float64 `json:"nonce,omitempty"`
	Fee *DeployErc20Fee `json:"fee,omitempty"`
}
