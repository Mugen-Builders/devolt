package integration

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/devolthq/devolt/pkg/rollups_contracts"
	"github.com/devolthq/devolt/test/integration/artifacts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rollmelette/rollmelette"
)

const applicationAddress = "0x70ac08179605AF2D9e75782b8DEcDD3c22aA4D0C"

var addressToPrivateKey = map[common.Address]string{
	common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"): "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
	common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8"): "0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d",
	common.HexToAddress("0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"): "0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a",
	common.HexToAddress("0x90F79bf6EB2c4f870365E785982E1f101E93b906"): "0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6",
	common.HexToAddress("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65"): "0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a",
	common.HexToAddress("0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc"): "0x8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba",
	common.HexToAddress("0x976EA74026E726554dB657fA54763abd0C3a0aa9"): "0x92db14e403b83dfe3df233f83dfa3a0d7096f21ca9b0d6d6b8d88b2b4ec1564e",
	common.HexToAddress("0x14dC79964da2C08b23698B3D3cc7Ca32193d9955"): "0x4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356",
	common.HexToAddress("0x23618e81E3f5cdF7f54C3d65f7FBc0aBf5B21E8f"): "0xdbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97",
	common.HexToAddress("0xa0Ee7A142d267C1f36714E4a8F75612F20a79720"): "0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6",
}

type TestToken struct {
	RPCUrl   string
	Instance *artifacts.Token
	Address  common.Address
	Book     rollmelette.AddressBook
}

func NewTestToken(rpcUrl string, initialOwner string) (*TestToken, error) {
	client, opts, err := setupClient(rpcUrl, initialOwner)
	if err != nil {
		return nil, err
	}
	address, tx, _, err := artifacts.DeployToken(opts, client)
	if err != nil {
		return nil, err
	}
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return nil, err
	}
	instance, err := artifacts.NewToken(address, client)
	if err != nil {
		return nil, err
	}
	return &TestToken{RPCUrl: rpcUrl, Instance: instance, Address: address, Book: rollmelette.NewAddressBook()}, nil
}

func (dc *TestToken) Mint(
	sender string,
	to string,
	amount *big.Int,
) error {
	client, opts, err := setupClient(dc.RPCUrl, sender)
	if err != nil {
		return err
	}
	tx, err := dc.Instance.Mint(opts, common.HexToAddress(to), amount)
	if err != nil {
		return err
	}
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return err
	}
	return nil
}

func (dc *TestToken) Deposit(
	sender string,
	amount *big.Int,
	_execLayerData []byte,
) error {
	client, opts, err := setupClient(dc.RPCUrl, sender)
	if err != nil {
		return err
	}

	// Approve tokens for the ERC20Portal
	approveTx, err := dc.Instance.Approve(opts, dc.Book.ERC20Portal, amount)
	if err != nil {
		return fmt.Errorf("failed to approve tokens: %w", err)
	}
	_, err = bind.WaitMined(context.Background(), client, approveTx)
	if err != nil {
		return fmt.Errorf("approval transaction failed: %w", err)
	}

	// Deposit tokens into the ERC20Portal
	erc20Portal, err := rollups_contracts.NewERC20Portal(dc.Book.ERC20Portal, client)
	if err != nil {
		return fmt.Errorf("failed to create ERC20Portal instance: %w", err)
	}

	depositTx, err := erc20Portal.DepositERC20Tokens(opts, dc.Address, common.HexToAddress(applicationAddress), amount, _execLayerData)
	if err != nil {
		return err
	}
	_, err = bind.WaitMined(context.Background(), client, depositTx)
	if err != nil {
		return fmt.Errorf("deposit transaction failed: %w", err)
	}
	return nil
}

func setupClient(
	rpcurl string,
	sender string,
) (*ethclient.Client, *bind.TransactOpts, error) {
	client, err := ethclient.Dial(rpcurl)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to blockchain: %v", err)
	}

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get chain id: %v", err)
	}
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(addressToPrivateKey[common.HexToAddress(sender)], "0x"))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load private key: %v", err)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create transactor: %v", err)
	}
	return client, opts, err
}
