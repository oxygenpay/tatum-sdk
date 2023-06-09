/*
 * Tatum API Reference
 *
 * # Welcome to the Tatum API Reference!  ## What is Tatum? ## What is Tatum?  Tatum offers a flexible framework to build, run, and scale blockchain apps fast. To learn more about the Tatum blockchain development framework, visit [our website](https://tatum.io/framework).  The Tatum API features powerful endpoints that simplify a complex blockchain into single API requests. Code for all supported blockchains using unified API calls.  ## Need help?  To chat with other developers, get help from the Support team, and engage with the thriving Tatum community, join  our [Discord server](https://discord.com/invite/tatum). For more information about how to work with Tatum, review the [online documentation](https://docs.tatum.io/).  ## About this API Reference  The Tatum API Reference is based on OpenAPI Specification v3.1.0 with a few [vendor extensions](https://github.com/Redocly/redoc/blob/master/docs/redoc-vendor-extensions.md) applied.  # Authentication  When using the Tatum API, you authenticate yourself with an **API key**. <SecurityDefinitions /> 
 *
 * API version: 3.16.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package tatum

// 1 point in the chart. This point represents the tick in the grapch based on the specified time frame.
type Chart struct {
	// Milliseconds in UTC of the time interval.
	Timestamp float64 `json:"timestamp"`
	// Highest trade value in the current interval.
	High string `json:"high"`
	// Lowest trade value in the current interval.
	Low string `json:"low"`
	// Open trade value in the current interval.
	Open string `json:"open"`
	// Close trade value in the current interval.
	Close string `json:"close"`
	// Total volume of assets traded in the current interval. Volume is in currency1 asset.
	Volume string `json:"volume"`
}
