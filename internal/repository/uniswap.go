package repository

import (
	"fantom-api-graphql/internal/repository/rpc/contracts"
	"fantom-api-graphql/internal/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// NativeTokenAddress returns address of the native token wrapper, if available.
func (p *proxy) NativeTokenAddress() (*common.Address, error) {
	return p.rpc.NativeTokenAddress()
}

// UniswapPairs returns list of all token pairs managed by Uniswap core.
func (p *proxy) UniswapPairs() ([]common.Address, error) {
	return p.rpc.UniswapPairs()
}

// UniswapPair returns an address of an Uniswap pair for the given tokens.
func (p *proxy) UniswapPair(tokenA *common.Address, tokenB *common.Address) (*common.Address, error) {
	return p.rpc.UniswapPair(tokenA, tokenB)
}

// UniswapAmountsOut resolves a list of output amounts for the given
// input amount and a list of tokens to be used to make the swap operation.
func (p *proxy) UniswapAmountsOut(amountIn hexutil.Big, tokens []common.Address) ([]hexutil.Big, error) {
	return p.rpc.UniswapAmountsOut(amountIn, tokens)
}

// UniswapAmountsIn resolves a list of input amounts for the given
// output amount and a list of tokens to be used to make the swap operation.
func (p *proxy) UniswapAmountsIn(amountOut hexutil.Big, tokens []common.Address) ([]hexutil.Big, error) {
	return p.rpc.UniswapAmountsIn(amountOut, tokens)
}

// UniswapQuoteInput calculates optimal input on sibling token based on input amount and
// self reserves of the analyzed token.
func (p *proxy) UniswapQuoteInput(
	amountIn hexutil.Big,
	reserveMy hexutil.Big,
	reserveSibling hexutil.Big,
) (hexutil.Big, error) {
	return p.rpc.UniswapQuoteInput(amountIn, reserveMy, reserveSibling)
}

// UniswapTokens returns list of addresses of tokens involved in a Uniswap pair.
func (p *proxy) UniswapTokens(pair *common.Address) ([]common.Address, error) {
	return p.rpc.UniswapTokens(pair)
}

// UniswapReserves returns list of token reserve amounts in a Uniswap pair.
func (p *proxy) UniswapReserves(pair *common.Address) ([]hexutil.Big, error) {
	return p.rpc.UniswapReserves(pair)
}

// UniswapReservesTimeStamp returns the timestamp of the reserves of a Uniswap pair.
func (p *proxy) UniswapReservesTimeStamp(pair *common.Address) (hexutil.Uint64, error) {
	return p.rpc.UniswapReservesTimeStamp(pair)
}

// UniswapCumulativePrices returns list of token cumulative prices of a Uniswap pair.
func (p *proxy) UniswapCumulativePrices(pair *common.Address) ([]hexutil.Big, error) {
	return p.rpc.UniswapCumulativePrices(pair)
}

// UniswapLastKValue returns the last value of the pool control coefficient.
func (p *proxy) UniswapLastKValue(pair *common.Address) (hexutil.Big, error) {
	return p.rpc.UniswapLastKValue(pair)
}

// UniswapPairContract returns instance of this contract according to given pair address
func (p *proxy) UniswapPairContract(pairAddres *common.Address) (*contracts.UniswapPair, error) {
	return p.rpc.UniswapPairContract(pairAddres)
}

// UniswapAdd notifies a new incoming swap from blockchain to the repository.
func (p *proxy) UniswapAdd(swap *types.Swap) error {
	return p.db.UniswapAdd(swap)
}

// LastKnownSwapBlock returns number of the last block known to the repository with the swap event.
func (p *proxy) LastKnownSwapBlock() (uint64, error) {
	return p.db.LastKnownSwapBlock()
}

// UniswapFactoryContract returns an instance of an Uniswap factory
func (p *proxy) UniswapFactoryContract() (*contracts.UniswapFactory, error) {
	return p.rpc.UniswapFactoryContract()
}

// UniswapVolume returns swap volume for specified uniswap pair
// If toTime = 0, then it resolves volumes till now
func (p *proxy) UniswapVolume(pairAddress *common.Address, fromTime int64, toTime int64) (types.DefiSwapVolume, error) {
	return p.db.UniswapVolume(pairAddress, fromTime, toTime)
}

// UniswapTimeVolumes returns daily swap volume for specified uniswap pair and period of time
// If toTime = 0, then it resolves volumes till now
func (p *proxy) UniswapTimeVolumes(pairAddress *common.Address, resolution string, fromTime int64, toTime int64) ([]types.DefiSwapVolume, error) {
	return p.db.UniswapTimeVolumes(pairAddress, resolution, fromTime, toTime)
}
