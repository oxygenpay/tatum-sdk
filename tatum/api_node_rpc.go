
/*
 * Tatum API Reference
 *
 * # Welcome to the Tatum API Reference!  ## What is Tatum? ## What is Tatum?  Tatum offers a flexible framework to build, run, and scale blockchain apps fast. To learn more about the Tatum blockchain development framework, visit [our website](https://tatum.io/framework).  The Tatum API features powerful endpoints that simplify a complex blockchain into single API requests. Code for all supported blockchains using unified API calls.  ## Need help?  To chat with other developers, get help from the Support team, and engage with the thriving Tatum community, join  our [Discord server](https://discord.com/invite/tatum). For more information about how to work with Tatum, review the [online documentation](https://docs.tatum.io/).  ## About this API Reference  The Tatum API Reference is based on OpenAPI Specification v3.1.0 with a few [vendor extensions](https://github.com/Redocly/redoc/blob/master/docs/redoc-vendor-extensions.md) applied.  # Authentication  When using the Tatum API, you authenticate yourself with an **API key**. <SecurityDefinitions /> 
 *
 * API version: 3.16.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package tatum

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type NodeRPCApiService service
/*
NodeRPCApiService Connect to the blockchain node through an RPC driver
&lt;p&gt;&lt;b&gt;The number of credits consumed depends on the number of methods submitted in an API call:&lt;br/&gt; * 50 credits per debug*_/trace* method (for EVM-based blockchains)&lt;br/&gt; * 5 credits per eth_call method (for EVM-based blockchains)&lt;br/&gt; * 2 credits per any other RPC method&lt;/b&gt;&lt;/p&gt; &lt;p&gt;Connect directly to the blockchain node provided by Tatum.&lt;/p&gt; &lt;p&gt;The &lt;code&gt;POST&lt;/code&gt; method is used. The API endpoint URL acts as an HTTP-based RPC driver.&lt;/p&gt; &lt;p&gt;In the request body, provide valid Web3 RPC method content, for example:&lt;/p&gt; &lt;pre&gt; {   \&quot;jsonrpc\&quot;: \&quot;2.0\&quot;,   \&quot;id\&quot;: 1,   \&quot;method\&quot;: \&quot;method_name\&quot;,   \&quot;params\&quot;: [] } &lt;/pre&gt; &lt;p&gt;For the blockchains using the JSON-RPC 2.0 specification, you can submit multiple RPC methods in one API call:&lt;/p&gt; &lt;pre&gt; [   {     \&quot;jsonrpc\&quot;: \&quot;2.0\&quot;,     \&quot;id\&quot;: 1,     \&quot;method\&quot;: \&quot;method_1_name\&quot;,     \&quot;params\&quot;: []   },   {     \&quot;jsonrpc\&quot;: \&quot;2.0\&quot;,     \&quot;id\&quot;: 2,     \&quot;method\&quot;: \&quot;method_2_name\&quot;,     \&quot;params\&quot;: []   },   ... ] &lt;/pre&gt; &lt;p&gt;This API is supported for the following blockchains:&lt;/p&gt; &lt;ul&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://developer.algorand.org/docs/rest-apis/restendpoints/\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Algorand&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://ethereum.org/en/developers/docs/apis/json-rpc/\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Arbitrum&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://doc.aurora.dev/compat/rpc/\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Aurora&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://developer.bitcoin.org/reference/rpc/index.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Bitcoin&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://docs.bitcoincashnode.org/doc/json-rpc/\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Bitcoin Cash&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://docs.bnbchain.org/docs/rpc\&quot; target&#x3D;\&quot;_blank\&quot;&gt;BNB Smart Chain&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://docs.cardano.org/cardano-components/cardano-graphql\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Cardano&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://explorer.celo.org/api-docs\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Celo&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://cronos.org/docs/resources/chain-integration.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Cronos&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://dogecoin.com/\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Dogecoin&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://docs.elrond.com/sdk-and-tools/rest-api/nodes/\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Elrond&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://developers.eos.io/welcome/latest/reference/nodeos-rpc-api-reference\&quot; target&#x3D;\&quot;_blank\&quot;&gt;EOSIO&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://ethereum.org/en/developers/docs/apis/json-rpc/\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Ethereum&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://docs.fantom.foundation/api/public-api-endpoints\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Fantom&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://developers.gnosischain.com/for-developers/developer-resources\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Gnosis&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://docs.harmony.one/home/developers/api\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Harmony&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://docs.klaytn.foundation/dapp/json-rpc\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Klaytn&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://docs.kcc.io/\&quot; target&#x3D;\&quot;_blank\&quot;&gt;KuCoin Community Chain&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://guide.kusama.network/docs/build-node-interaction/\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Kusama&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://lisk.com/documentation/lisk-service/references/api.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Lisk&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://litecoin.org/\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Litecoin&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://docs.near.org/api/rpc/introduction\&quot; target&#x3D;\&quot;_blank\&quot;&gt;NEAR&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://docs.neo.org/docs/en-us/reference/rpc/latest-version/api.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Neo&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://docs.oasis.dev/oasis-core/oasis-node/rpc/\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Oasis Network&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://community.optimism.io/docs/developers/build/json-rpc/\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Optimism&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://docs.palm.io/Get-Started/Connect/Overview/\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Palm&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://wiki.polkadot.network/docs/build-node-interaction\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Polkadot&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://polygon.technology/developers\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Polygon&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://developers.rsk.co/rsk/node/architecture/json-rpc/\&quot; target&#x3D;\&quot;_blank\&quot;&gt;RSK&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://docs.solana.com/developing/clients/jsonrpc-api\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Solana&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://developers.stellar.org/api\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Stellar&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://opentezos.com/tezos-basics/cli-and-rpc/\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Tezos&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://zcash-rpc.github.io/\&quot; target&#x3D;\&quot;_blank\&quot;&gt;ZCash&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://developers.tron.network/reference/full-node-api-overview\&quot; target&#x3D;\&quot;_blank\&quot;&gt;TRON&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://www.vechain.org/\&quot; target&#x3D;\&quot;_blank\&quot;&gt;VeChain&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://apidocs.xinfin.network/docs/\&quot; target&#x3D;\&quot;_blank\&quot;&gt;XinFin&lt;/a&gt;&lt;/li&gt; &lt;/ul&gt; 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body
 * @param chain Blockchain to communicate with.
 * @param rpcPath Optional path of rpc call for non EVM nodes, e.g. Algorand or Stellar.
 * @param optional nil or *NodeRPCApiNodeJsonPostRpcDriverOpts - Optional Parameters:
     * @param "NodeType" (optional.String) -  Type of the node to access for Algorand.
     * @param "TestnetType" (optional.String) -  Type of Ethereum testnet. Defaults to ethereum-sepolia.
@return interface{}
*/

type NodeRPCApiNodeJsonPostRpcDriverOpts struct {
    NodeType optional.String
    TestnetType optional.String
}

func (a *NodeRPCApiService) NodeJsonPostRpcDriver(ctx context.Context, body interface{}, chain string, rpcPath string, localVarOptionals *NodeRPCApiNodeJsonPostRpcDriverOpts) (interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v3/blockchain/node/{chain}/{rpcPath}"
	localVarPath = strings.Replace(localVarPath, "{"+"chain"+"}", fmt.Sprintf("%v", chain), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rpcPath"+"}", fmt.Sprintf("%v", rpcPath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.NodeType.IsSet() {
		localVarQueryParams.Add("nodeType", parameterToString(localVarOptionals.NodeType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TestnetType.IsSet() {
		localVarQueryParams.Add("testnetType", parameterToString(localVarOptionals.TestnetType.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v Error400
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v InlineResponse401
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
NodeRPCApiService Connect to the blockchain node through an RPC driver
&lt;p&gt;&lt;b&gt;2 credits per RPC method in an API call&lt;/b&gt;&lt;/p&gt; &lt;p&gt;Connect directly to the blockchain node provided by Tatum.&lt;/p&gt; &lt;p&gt;The &lt;code&gt;GET&lt;/code&gt; method is used. The API endpoint URL acts as an HTTP-based RPC driver.&lt;/p&gt; &lt;p&gt;This API is supported for the following blockchains:&lt;/p&gt; &lt;ul&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://developer.algorand.org/docs/rest-apis/restendpoints/\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Algorand&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://docs.elrond.com/sdk-and-tools/rest-api/nodes/\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Elrond&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://lisk.com/documentation/lisk-service/references/api.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Lisk&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://developers.stellar.org/api\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Stellar&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://opentezos.com/tezos-basics/cli-and-rpc/\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Tezos&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://developers.tron.network/reference/full-node-api-overview\&quot; target&#x3D;\&quot;_blank\&quot;&gt;TRON&lt;/a&gt;&lt;/li&gt; &lt;/ul&gt; 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param chain Blockchain to communicate with.
 * @param rpcPath Optional path of rpc call for non EVM nodes, e.g. Algorand or Stellar.
 * @param optional nil or *NodeRPCApiNodeJsonRpcGetDriverOpts - Optional Parameters:
     * @param "NodeType" (optional.String) -  Type of the node to access for Algorand.
@return interface{}
*/

type NodeRPCApiNodeJsonRpcGetDriverOpts struct {
    NodeType optional.String
}

func (a *NodeRPCApiService) NodeJsonRpcGetDriver(ctx context.Context, chain string, rpcPath string, localVarOptionals *NodeRPCApiNodeJsonRpcGetDriverOpts) (interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v3/blockchain/node/{chain}/{rpcPath}"
	localVarPath = strings.Replace(localVarPath, "{"+"chain"+"}", fmt.Sprintf("%v", chain), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rpcPath"+"}", fmt.Sprintf("%v", rpcPath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.NodeType.IsSet() {
		localVarQueryParams.Add("nodeType", parameterToString(localVarOptionals.NodeType.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v Error400
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v InlineResponse401
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
NodeRPCApiService Connect to the blockchain node through an RPC driver
&lt;p&gt;&lt;b&gt;2 credits per RPC method in an API call&lt;/b&gt;&lt;/p&gt; &lt;p&gt;Connect directly to the blockchain node provided by Tatum.&lt;/p&gt; &lt;p&gt;The &lt;code&gt;PUT&lt;/code&gt; method is used. The API endpoint URL acts as an HTTP-based RPC driver.&lt;/p&gt; &lt;p&gt;In the request body, provide valid Web3 RPC method content, for example:&lt;/p&gt; &lt;pre&gt; {   \&quot;jsonrpc\&quot;: \&quot;2.0\&quot;,   \&quot;id\&quot;: 1,   \&quot;method\&quot;: \&quot;method_name\&quot;,   \&quot;params\&quot;: [] } &lt;/pre&gt; &lt;p&gt;For the blockchains using the JSON-RPC 2.0 specification, you can submit multiple RPC methods in one API call:&lt;/p&gt; &lt;pre&gt; [   {     \&quot;jsonrpc\&quot;: \&quot;2.0\&quot;,     \&quot;id\&quot;: 1,     \&quot;method\&quot;: \&quot;method_1_name\&quot;,     \&quot;params\&quot;: []   },   {     \&quot;jsonrpc\&quot;: \&quot;2.0\&quot;,     \&quot;id\&quot;: 2,     \&quot;method\&quot;: \&quot;method_2_name\&quot;,     \&quot;params\&quot;: []   },   ... ] &lt;/pre&gt; &lt;p&gt;This API is supported for the following blockchains:&lt;/p&gt; &lt;ul&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://developer.algorand.org/docs/rest-apis/restendpoints/\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Algorand&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://docs.elrond.com/sdk-and-tools/rest-api/nodes/\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Elrond&lt;/a&gt;&lt;/li&gt;   &lt;li&gt;&lt;a href&#x3D;\&quot;https://developers.stellar.org/api\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Stellar&lt;/a&gt;&lt;/li&gt; &lt;/ul&gt; 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body
 * @param chain Blockchain to communicate with.
 * @param rpcPath Optional path of rpc call for non EVM nodes, e.g. Algorand or Stellar.
 * @param optional nil or *NodeRPCApiNodeJsonRpcPutDriverOpts - Optional Parameters:
     * @param "NodeType" (optional.String) -  Type of the node to access for Algorand.
@return interface{}
*/

type NodeRPCApiNodeJsonRpcPutDriverOpts struct {
    NodeType optional.String
}

func (a *NodeRPCApiService) NodeJsonRpcPutDriver(ctx context.Context, body interface{}, chain string, rpcPath string, localVarOptionals *NodeRPCApiNodeJsonRpcPutDriverOpts) (interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v3/blockchain/node/{chain}/{rpcPath}"
	localVarPath = strings.Replace(localVarPath, "{"+"chain"+"}", fmt.Sprintf("%v", chain), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rpcPath"+"}", fmt.Sprintf("%v", rpcPath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.NodeType.IsSet() {
		localVarQueryParams.Add("nodeType", parameterToString(localVarOptionals.NodeType.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v Error400
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v InlineResponse401
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
