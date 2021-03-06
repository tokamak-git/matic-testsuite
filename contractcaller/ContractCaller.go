package contractcaller

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/maticnetwork/bor"
	"github.com/maticnetwork/bor/accounts/abi"
	"github.com/maticnetwork/bor/accounts/abi/bind"
	"github.com/maticnetwork/bor/common"
	"github.com/maticnetwork/bor/crypto"
	"github.com/maticnetwork/bor/ethclient"
	"github.com/maticnetwork/heimdall/contracts/erc20"
	"github.com/maticnetwork/heimdall/contracts/rootchain"
	"github.com/maticnetwork/heimdall/contracts/slashmanager"
	"github.com/maticnetwork/heimdall/contracts/stakemanager"
	"github.com/maticnetwork/heimdall/contracts/stakinginfo"
	"github.com/maticnetwork/heimdall/contracts/statereceiver"
	"github.com/maticnetwork/heimdall/contracts/statesender"
	"github.com/maticnetwork/heimdall/contracts/validatorset"
	helper "github.com/maticnetwork/matic-testsuite/helper"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

// ContractCaller contract caller
type ContractCaller struct {
	MainChainClient  *ethclient.Client
	MaticChainClient *ethclient.Client

	RootChainABI     abi.ABI
	StakingInfoABI   abi.ABI
	ValidatorSetABI  abi.ABI
	StateReceiverABI abi.ABI
	StateSenderABI   abi.ABI
	StakeManagerABI  abi.ABI
	SlashManagerABI  abi.ABI
	MaticTokenABI    abi.ABI

	ContractInstanceCache map[common.Address]interface{}
}

// NewContractCaller contract caller
func NewContractCaller() (contractCallerObj ContractCaller, err error) {
	contractCallerObj.MainChainClient = helper.GetMainClient()
	contractCallerObj.MaticChainClient = helper.GetMaticClient()

	//
	// ABIs
	//

	if contractCallerObj.RootChainABI, err = getABI(string(rootchain.RootchainABI)); err != nil {
		return
	}

	if contractCallerObj.StakingInfoABI, err = getABI(string(stakinginfo.StakinginfoABI)); err != nil {
		return
	}

	if contractCallerObj.ValidatorSetABI, err = getABI(string(validatorset.ValidatorsetABI)); err != nil {
		return
	}

	if contractCallerObj.StateReceiverABI, err = getABI(string(statereceiver.StatereceiverABI)); err != nil {
		return
	}

	if contractCallerObj.StateSenderABI, err = getABI(string(statesender.StatesenderABI)); err != nil {
		return
	}

	if contractCallerObj.StakeManagerABI, err = getABI(string(stakemanager.StakemanagerABI)); err != nil {
		return
	}

	if contractCallerObj.SlashManagerABI, err = getABI(string(slashmanager.SlashmanagerABI)); err != nil {
		return
	}

	if contractCallerObj.MaticTokenABI, err = getABI(string(erc20.Erc20ABI)); err != nil {
		return
	}

	contractCallerObj.ContractInstanceCache = make(map[common.Address]interface{})

	return
}

// GetRootChainInstance returns RootChain contract instance for selected base chain
func (c *ContractCaller) GetRootChainInstance(rootchainAddress common.Address) (*rootchain.Rootchain, error) {
	contractInstance, ok := c.ContractInstanceCache[rootchainAddress]
	if !ok {
		ci, err := rootchain.NewRootchain(rootchainAddress, c.MainChainClient)
		c.ContractInstanceCache[rootchainAddress] = ci
		return ci, err
	}
	return contractInstance.(*rootchain.Rootchain), nil
}

// GetStakeManagerInstance returns stakinginfo contract instance for selected base chain
func (c *ContractCaller) GetStakeManagerInstance(stakingManagerAddress common.Address) (*stakemanager.Stakemanager, error) {
	return stakemanager.NewStakemanager(stakingManagerAddress, c.MainChainClient)
}

// GetMaticTokenInstance returns stakinginfo contract instance for selected base chain
func (c *ContractCaller) GetMaticTokenInstance(maticTokenAddress common.Address) (*erc20.Erc20, error) {
	return erc20.NewErc20(maticTokenAddress, c.MainChainClient)
}

// StakeFor stakes for a validator
func (c *ContractCaller) StakeFor(val common.Address, stakeAmount *big.Int, feeAmount *big.Int, acceptDelegation bool, stakeManagerAddress common.Address, stakeManagerInstance *stakemanager.Stakemanager) error {
	// TODO: pass pubkey
	// signerPubkey := ""
	signerPubkeyBytes := ""

	// pack data based on method definition
	data, err := c.StakeManagerABI.Pack("stakeFor", val, stakeAmount, feeAmount, acceptDelegation, signerPubkeyBytes)
	if err != nil {
		return err
	}

	auth, err := GenerateAuthObj(c.MainChainClient, stakeManagerAddress, data)
	if err != nil {
		return err
	}

	// stake for stake manager
	tx, err := stakeManagerInstance.StakeFor(
		auth,
		val,
		stakeAmount,
		feeAmount,
		acceptDelegation,
		[]byte(signerPubkeyBytes),
	)

	if err != nil {
		return err
	}

	fmt.Println("Submitted stake sucessfully", "txHash", tx.Hash().String())
	return nil
}

// ApproveTokens approves matic token for stake
func (c *ContractCaller) ApproveTokens(amount *big.Int, stakeManager common.Address, tokenAddress common.Address, maticTokenInstance *erc20.Erc20) error {
	data, err := c.MaticTokenABI.Pack("approve", stakeManager, amount)
	if err != nil {
		return err
	}

	auth, err := GenerateAuthObj(c.MainChainClient, tokenAddress, data)
	if err != nil {
		return err
	}

	tx, err := maticTokenInstance.Approve(auth, stakeManager, amount)
	if err != nil {
		return err
	}

	fmt.Println("Sent approve tx sucessfully", "txHash", tx.Hash().String())
	return nil
}

func GenerateAuthObj(client *ethclient.Client, address common.Address, data []byte) (auth *bind.TransactOpts, err error) {
	// generate call msg
	callMsg := ethereum.CallMsg{
		To:   &address,
		Data: data,
	}

	// get priv key
	pkObject := secp256k1.GenPrivKey()

	// create ecdsa private key
	ecdsaPrivateKey, err := crypto.ToECDSA(pkObject[:])
	if err != nil {
		return
	}

	// from address
	fromAddress := common.BytesToAddress(pkObject.PubKey().Address().Bytes())
	// fetch gas price
	gasprice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}
	// fetch nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return
	}

	// fetch gas limit
	callMsg.From = fromAddress
	gasLimit, err := client.EstimateGas(context.Background(), callMsg)

	// create auth
	auth = bind.NewKeyedTransactor(ecdsaPrivateKey)
	auth.GasPrice = gasprice
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(gasLimit) // uint64(gasLimit)

	return
}

//
// private abi methods
//

func getABI(data string) (abi.ABI, error) {
	return abi.JSON(strings.NewReader(data))
}
