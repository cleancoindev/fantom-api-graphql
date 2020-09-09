package repository

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// DefiToken loads details of a single DeFi token by it's address.
func (p *proxy) DefiToken(token *common.Address) (*types.DefiToken, error) {
	return p.rpc.DefiToken(token)
}

// DefiTokens resolves list of DeFi tokens available for the DeFi functions.
func (p *proxy) DefiTokens() ([]types.DefiToken, error) {
	return p.rpc.DefiTokens()
}

// DefiConfiguration resolves the current DeFi contract settings.
func (p *proxy) DefiConfiguration() (*types.DefiSettings, error) {
	return p.rpc.DefiConfiguration()
}

// DefiTokenPrice loads the current price of the given token
// from on-chain price oracle.
func (p *proxy) DefiTokenPrice(token *common.Address) (hexutil.Big, error) {
	return p.rpc.FMintTokenPrice(token)
}

// FMintAccount loads details of a DeFi/fMint account identified by the owner address.
func (p *proxy) FMintAccount(owner common.Address) (*types.FMintAccount, error) {
	return p.rpc.FMintAccount(&owner)
}

// FMintTokenBalance loads balance of a single DeFi token by it's address.
func (p *proxy) FMintTokenBalance(owner *common.Address, token *common.Address, tp types.DefiTokenType) (hexutil.Big, error) {
	return p.rpc.FMintTokenBalance(owner, token, tp)
}

// FMintTokenValue loads value of a single DeFi token by it's address in fUSD.
func (p *proxy) FMintTokenValue(owner *common.Address, token *common.Address, tp types.DefiTokenType) (hexutil.Big, error) {
	return p.rpc.FMintTokenValue(owner, token, tp)
}

// Erc20Balance load the current available balance of and ERC20 token identified by the token
// contract address for an identified owner address.
func (p *proxy) Erc20Balance(owner *common.Address, token *common.Address) (hexutil.Big, error) {
	return p.rpc.Erc20Balance(owner, token)
}

// Erc20Allowance loads the current amount of ERC20 tokens unlocked for DeFi
// contract by the token owner.
func (p *proxy) Erc20Allowance(owner *common.Address, token *common.Address) (hexutil.Big, error) {
	return p.rpc.Erc20Allowance(owner, token)
}
