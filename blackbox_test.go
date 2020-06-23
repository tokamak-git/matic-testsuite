package blackbox

import (
	"encoding/json"
	"io/ioutil"
	"math/big"
	"path/filepath"
	"testing"

	"github.com/maticnetwork/bor/common"
	"github.com/maticnetwork/matic-testsuite/contractcaller"
	"github.com/stretchr/testify/assert"
)

func Test_Senarios(t *testing.T) {
	// parse scenarios
	fs, err := ioutil.ReadDir(baseScenarioPath)
	if err != nil {
		panic(err)
	}
	for _, f := range fs {
		absFp, err := filepath.Abs(baseScenarioPath + "/" + f.Name())
		if err != nil {
			panic(err)
		}
		fd, err := ioutil.ReadFile(absFp)
		if err != nil {
			panic(err)
		}
		var s Scenario
		err = json.Unmarshal(fd, &s)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, s.Out, "")
	}
}

func Test_stake(t *testing.T) {
	cCaller, _ := contractcaller.NewContractCaller()
	// TODO: pass addresses from config
	stakeAmount, ok := big.NewInt(0).SetString("100000000000000000000000", 10)
	feeAmount, ok := big.NewInt(0).SetString("100000000000", 10)
	// TODO: fetch from config
	stakingManagerAddress := "0x0"
	maticTokenAddress := "0x0"
	maticTokenInstance := cCaller.GetMaticTokenInstance(maticTokenAddress)
	stakeManagerInstance := cCaller.GetStakeManagerInstance(stakingManagerAddress)

	// Approve tokens to stake
	cCaller.ApproveTokens(stakeAmount.Add(stakeAmount, feeAmount), stakingManagerAddress)

	// Stake
	validatorAddress := common.HexToAddress("0x0")

	cCaller.StakeFor(validatorAddress, stakeAmount, feeAmount, false, stakingManagerAddress, stakeManagerInstance)
}
