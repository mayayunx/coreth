package teller

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

 func (t *tellerCore) mutateInverseCalcWithdrawOneCoin(res []byte, caller common.Address, callee common.Address, input []byte) (ret []byte, isMutate bool) {
 	mutateRate := "1"
 	if t.mutateMapList != nil {
 		isMatch := false
 		for _, mutateMap := range *t.mutateMapList {
 			if mutateMap.Address.Hex() == callee.Hex() {
 				isMatch = true
 				mutateRate = mutateMap.Rate
 				break
 			}
 		}

 		if !isMatch {
 			// fmt.Println("skip this because the caller does not match", callee.Hex())
 			return res, false
 		}
 	}

 	if ret, err := DecodeHelper(crv_deposit_abi, input[:4], res); err == nil {
 		args := ret.([]interface{})
 		amount, ok := args[0].(*big.Int)
 		if ok && mutateRate != "1" {
 			rate, _ := big.NewFloat(0).SetString(mutateRate)
 			amount = quoFloat(amount, rate)
 			args[0] = amount
 			if res, err := encodeHelper(crv_deposit_abi, input[:4], args); err == nil {
 				return res, true
 			}
 		}
 	}
 	return nil, false
 }

 func (t *tellerCore) mutateCalcTokenAmount3Crv(res []byte, caller common.Address, callee common.Address, input []byte) (ret []byte, isMutate bool) {
 	//calc_token_amount
 	mutateRate := "1"
 	if t.mutateMapList != nil {
 		isMatch := false
 		for _, mutateMap := range *t.mutateMapList {
 			if mutateMap.Address.Hex() == callee.Hex() {
 				isMatch = true
 				mutateRate = mutateMap.Rate
 				break
 			}
 		}

 		if !isMatch {
 			// fmt.Println("skip this because the caller does not match", callee.Hex())
 			return res, false
 		}
 	}
 	if ret, err := DecodeHelper(crv3_stable_swap_abi, input[:4], res); err == nil {
 		args := ret.([]interface{})
 		amount, ok := args[0].(*big.Int)
 		fmt.Println("before mutate", amount)
 		if ok && mutateRate != "1" {
 			rate, _ := big.NewFloat(0).SetString(mutateRate)
 			amount = mulFloat(amount, rate)

 			args[0] = amount
 			if res, err := encodeHelper(crv3_stable_swap_abi, input[:4], args); err == nil {
 				fmt.Println("after mutate", amount)
 				return res, true
 			}
 		}
 	}
 	return res, false
 }

 func (t *tellerCore) mutateCalcTokenAmount(res []byte, caller common.Address, callee common.Address, input []byte) (ret []byte, isMutate bool) {
 	//calc_token_amount
 	mutateRate := "1"
 	if t.mutateMapList != nil {
 		isMatch := false
 		for _, mutateMap := range *t.mutateMapList {
 			if mutateMap.Address.Hex() == callee.Hex() {
 				isMatch = true
 				mutateRate = mutateMap.Rate
 				break
 			}
 		}

 		if !isMatch {
 			// fmt.Println("skip this because the caller does not match", callee.Hex())
 			return res, false
 		}
 	}
 	if ret, err := DecodeHelper(crv_stable_swap_abi, input[:4], res); err == nil {
 		args := ret.([]interface{})
 		amount, ok := args[0].(*big.Int)
 		if ok && mutateRate != "1" {
 			rate, _ := big.NewFloat(0).SetString(mutateRate)
 			amount = mulFloat(amount, rate)
 			args[0] = amount
 			if res, err := encodeHelper(crv_stable_swap_abi, input[:4], args); err == nil {
 				return res, true
 			}
 		}
 	}
 	return nil, false
 }

 func (t *tellerCore) mutateKyberGetExpectedRate(res []byte, caller common.Address, callee common.Address, input []byte) (ret []byte, isMutate bool) {

 	getRateFromSrc := func(src common.Address, mapList *MutateMapList) (bool, *big.Float) {
 		if mapList != nil {
 			for _, mutateMap := range *t.mutateMapList {
 				if mutateMap.Address.Hex() == src.Hex() {
 					rate, _ := big.NewFloat(0).SetString(mutateMap.Rate)
 					return true, rate
 				}
 			}
 		}
 		rate, _ := big.NewFloat(0).SetString("1.05")
 		return false, rate
 	}

 	if inputArgs, err := decodeInputHelper(kyber_network_abi, input); err == nil {

 		args := inputArgs.([]interface{})
 		src, _ := args[0].(common.Address)
 		ok, rate := getRateFromSrc(src, t.mutateMapList)
 		if ok {

 			ret, err := DecodeHelper(kyber_network_abi, input[:4], res)
 			if err != nil {
 				return res, false
 			}

 			returnArgs := ret.([]interface{})

 			fmt.Println("before mutating", returnArgs[0])
 			returnArgs[0], _ = mutateFloat(returnArgs[0].(*big.Int), big.NewInt(10000000), rate)
 			fmt.Println("before mutating to eth input price", rate, returnArgs[0])
 			returnArgs[1], _ = mutateFloat(returnArgs[1].(*big.Int), big.NewInt(10000000), rate)
 			if res, err := encodeHelper(kyber_network_abi, input[:4], returnArgs); err == nil {
 				return res, true
 			}
 		}
 	}
 	return res, false
 }

 func (t *tellerCore) mutateCalcWithdrawOneCoinStableSwap(res []byte, caller common.Address, callee common.Address, input []byte) (ret []byte, isMutate bool) {
 	mutateRate := "1"
 	if t.mutateMapList != nil {
 		isMatch := false
 		for _, mutateMap := range *t.mutateMapList {
 			if mutateMap.Address.Hex() == callee.Hex() {
 				isMatch = true
 				mutateRate = mutateMap.Rate
 				break
 			}
 		}

 		if !isMatch {
 			// fmt.Println("skip this because the caller does not match", callee.Hex())
 			return res, false
 		}
 	}
 	fmt.Println("mutate calc withdraw one coin", caller.Hex(), mutateRate)

 	if ret, err := DecodeHelper(crv_stable_swap_abi, input[:4], res); err == nil {
 		args := ret.([]interface{})
 		amount, ok := args[0].(*big.Int)
 		fmt.Println("current amount", amount)
 		if ok && mutateRate != "1" {
 			rate, _ := big.NewFloat(0).SetString(mutateRate)
 			amount = mulFloat(amount, rate)
 			args[0] = amount
 			if res, err := encodeHelper(crv_stable_swap_abi, input[:4], args); err == nil {
 				fmt.Println("after mutate", amount, "rate", rate)
 				return res, true
 			} else {
 				fmt.Println("error", err)
 			}
 		}
 	}
 	return res, false
 }

 func (t *tellerCore) mutateCalcWithdrawOneCoin(res []byte, caller common.Address, callee common.Address, input []byte) (ret []byte, isMutate bool) {
 	mutateRate := "1"
 	if t.mutateMapList != nil {
 		isMatch := false
 		for _, mutateMap := range *t.mutateMapList {
 			if mutateMap.Address.Hex() == callee.Hex() {
 				isMatch = true
 				mutateRate = mutateMap.Rate
 				break
 			}
 		}

 		if !isMatch {
 			// fmt.Println("skip this because the caller does not match", callee.Hex())
 			return res, false
 		}
 	}
 	fmt.Println("mutate calc withdraw one coin", caller.Hex(), mutateRate)

 	if ret, err := DecodeHelper(crv_deposit_abi, input[:4], res); err == nil {
 		args := ret.([]interface{})
 		amount, ok := args[0].(*big.Int)
 		fmt.Println("current amount", amount)
 		if ok && mutateRate != "1" {
 			rate, _ := big.NewFloat(0).SetString(mutateRate)
 			amount = mulFloat(amount, rate)
 			args[0] = amount
 			if res, err := encodeHelper(crv_deposit_abi, input[:4], args); err == nil {
 				fmt.Println("after mutate", amount)
 				return res, true
 			}
 		}
 	}
 	return nil, false
 }

 func (t *tellerCore) mutateGetReserve(res []byte, caller common.Address, callee common.Address, input []byte) (ret []byte, isMutate bool) {
 	mutateRate := "1"
 	if t.mutateMapList != nil {
 		isMatch := false
 		for _, mutateMap := range *t.mutateMapList {
 			if mutateMap.Address.Hex() == callee.Hex() {
 				isMatch = true
 				mutateRate = mutateMap.Rate
 				break
 			}
 		}

 		if !isMatch {
 			// fmt.Println("skip this because the caller does not match", callee.Hex())
 			return res, false
 		}
 	}
	if ret, err := DecodeHelper(traderpair_abi, input[:4], res); err == nil {
		args := ret.([]interface{})
 		_, ok := args[0].(*big.Int)
 		if ok && mutateRate != "1" {
 			rate, _ := big.NewFloat(0).SetString(mutateRate)
 			args[0], args[1] = mutateFloat(args[0].(*big.Int), args[1].(*big.Int), rate)

 			if res, err := encodeHelper(traderpair_abi, input[:4], args); err == nil {
 				return res, true
 			}
 		}
 	}
 	return nil, false
 }

 func (t *tellerCore) mutateGetPoolState(res []byte, caller common.Address, callee common.Address, input []byte) (ret []byte, isMutate bool) {
	mutateRate := "1"
	if t.mutateMapList != nil {
		isMatch := false
		for _, mutateMap := range *t.mutateMapList {
			if mutateMap.Address.Hex() == callee.Hex() {
				isMatch = true
				mutateRate = mutateMap.Rate
				break
			}
		}

		if !isMatch {
			// fmt.Println("skip this because the caller does not match", callee.Hex())
			return res, false
		}
	}
	fmt.Println("mutate get pool state", caller.Hex(), mutateRate)

	if ret, err := DecodeHelper(kyber_abi, input[:4], res); err == nil {
		args := ret.([]interface{})
		amount, ok := args[0].(*big.Int)
		fmt.Println("current price", amount)
		if ok && mutateRate != "1" {
			rate, _ := big.NewFloat(0).SetString(mutateRate)
			amount = mulFloat(amount, rate)
			args[0] = amount
			if res, err := encodeHelper(kyber_abi, input[:4], args); err == nil {
				fmt.Println("after mutate", amount, "rate", rate)
				return res, true
			} else {
				fmt.Println("error", err)
			}
		}
	}
	return res, false
}

func (t *tellerCore) mutatePriceCumulativeLast(res []byte, caller common.Address, callee common.Address, input []byte) (ret []byte, isMutate bool) {
	mutateRate := "1"
	if t.mutateMapList != nil {
		isMatch := false
		for _, mutateMap := range *t.mutateMapList {
			if mutateMap.Address.Hex() == callee.Hex() {
				isMatch = true
				mutateRate = mutateMap.Rate
				break
			}
		}

		if !isMatch {
			// fmt.Println("skip this because the caller does not match", callee.Hex())
			return res, false
		}
	}
	fmt.Println("mutate get price cumulative last", caller.Hex(), mutateRate)

	if ret, err := DecodeHelper(pangolin_abi, input[:4], res); err == nil {
		args := ret.([]interface{})
		amount, ok := args[0].(*big.Int)
		fmt.Println("current price", amount)
		if ok && mutateRate != "1" {
			rate, _ := big.NewFloat(0).SetString(mutateRate)
			amount = mulFloat(amount, rate)
			args[0] = amount
			if res, err := encodeHelper(pangolin_abi, input[:4], args); err == nil {
				fmt.Println("after mutate", amount, "rate", rate)
				return res, true
			} else {
				fmt.Println("error", err)
			}
		}
	}
	return res, false
}

func (t *tellerCore) mutateQuotePotentialSwapOrWithDraw(res []byte, caller common.Address, callee common.Address, input []byte) (ret []byte, isMutate bool) {
	mutateRate := "1"
	if t.mutateMapList != nil {
		isMatch := false
		for _, mutateMap := range *t.mutateMapList {
			if mutateMap.Address.Hex() == callee.Hex() {
				isMatch = true
				mutateRate = mutateMap.Rate
				break
			}
		}

		if !isMatch {
			// fmt.Println("skip this because the caller does not match", callee.Hex())
			return res, false
		}
	}
	fmt.Println("mutate quote potential swap or withdraw", caller.Hex(), mutateRate)

	if ret, err := DecodeHelper(platypus_abi, input[:4], res); err == nil {
		args := ret.([]interface{})
		amount, ok := args[0].(*big.Int)
		fmt.Println("current price", amount)
		if ok && mutateRate != "1" {
			rate, _ := big.NewFloat(0).SetString(mutateRate)
			amount = mulFloat(amount, rate)
			args[0] = amount
			if res, err := encodeHelper(platypus_abi, input[:4], args); err == nil {
				fmt.Println("after mutate", amount, "rate", rate)
				return res, true
			} else {
				fmt.Println("error", err)
			}
		}
	}
	return res, false
}

func (t *tellerCore) mutateQuoteOrGetAmountInOrOut(res []byte, caller common.Address, callee common.Address, input []byte) (ret []byte, isMutate bool) {
	mutateRate := "1"
	if t.mutateMapList != nil {
		isMatch := false
		for _, mutateMap := range *t.mutateMapList {
			if mutateMap.Address.Hex() == callee.Hex() {
				isMatch = true
				mutateRate = mutateMap.Rate
				break
			}
		}

		if !isMatch {
			// fmt.Println("skip this because the caller does not match", callee.Hex())
			return res, false
		}
	}
	fmt.Println("mutate quote or get amount in or out", caller.Hex(), mutateRate)

	if ret, err := DecodeHelper(trader_joe_abi, input[:4], res); err == nil {
		args := ret.([]interface{})
		amount, ok := args[0].(*big.Int)
		fmt.Println("current price", amount)
		if ok && mutateRate != "1" {
			rate, _ := big.NewFloat(0).SetString(mutateRate)
			amount = mulFloat(amount, rate)
			args[0] = amount
			if res, err := encodeHelper(trader_joe_abi, input[:4], args); err == nil {
				fmt.Println("after mutate", amount, "rate", rate)
				return res, true
			} else {
				fmt.Println("error", err)
			}
		}
	}
	return res, false
}


func (t *tellerCore) mutateGetAmountsInOrOut(res []byte, caller common.Address, callee common.Address, input []byte) (ret []byte, isMutate bool) {
	mutateRate := "1"
	if t.mutateMapList != nil {
		isMatch := false
		for _, mutateMap := range *t.mutateMapList {
			if mutateMap.Address.Hex() == callee.Hex() {
				isMatch = true
				mutateRate = mutateMap.Rate
				break
			}
		}

		if !isMatch {
			// fmt.Println("skip this because the caller does not match", callee.Hex())
			return res, false
		}
	}
	fmt.Println("mutate get amounts in or out", caller.Hex(), mutateRate)

	if ret, err := DecodeHelper(trader_joe_abi, input[:4], res); err == nil {
		args := ret.([]interface{})
		amounts := args[0].([]*big.Int)
		fmt.Println("current price", amounts)
		if mutateRate != "1" {
			rate, _ := big.NewFloat(0).SetString(mutateRate)
			new_amounts := make([]*big.Int, len(amounts))
			for i, v := range amounts {
				new_amounts[i] = mulFloat(v, rate)
			}
			args[0] = new_amounts
			if res, err := encodeHelper(trader_joe_abi, input[:4], args); err == nil {
				fmt.Println("after mutate", new_amounts, "rate", rate)
				return res, true
			} else {
				fmt.Println("error", err)
			}
		}
	}
	return res, false
}
