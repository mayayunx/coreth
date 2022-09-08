package teller

import "github.com/ethereum/go-ethereum/common"



 func (t *tellerCore) loadPangolin() {
 	// 0902f1ac  =>  getReserves()
 	// 5909c0d5  =>  price0CumulativeLast()
 	// 5a3d5493  =>  price1CumulativeLast()

 	t.WatchList = append(t.WatchList, WatchFunction{
 		Signature: common.FromHex("0x0902f1ac"),
 		Address:   make(map[string]BreakPointType),
 	})
 	t.WatchList = append(t.WatchList, WatchFunction{
 		Signature: common.FromHex("0x5909c0d5"),
 		Address:   make(map[string]BreakPointType),
 	})
 	t.WatchList = append(t.WatchList, WatchFunction{
 		Signature: common.FromHex("0x5a3d5493"),
 		Address:   make(map[string]BreakPointType),
 	})
 	// t.WatchList = append(t.WatchList, WatchFunction{
 	// 	Signature: common.FromHex("0x7464fc3d"),
 	// 	Address:   make(map[string]BreakPointType),
 	// })
 }

 func (t *tellerCore) loadCrv() {
 	// calc_withdraw_one_coin(address,uint256,int128)	41b028f3
	// calc_withdraw_one_coin(uint256,int128)	cc2b27d7
 	// calc_token_amount(uint256[2],bool)	ed8e84f3
	// calc_token_amount(uint256[3],bool) 3883e119
	// get_virtual_price() bb7b8b80
	// get_dy(int128,int128,uint256)	5e0d443f
 	// get_dy_underlying(int128,int128,uint256)	07211ef7
 	t.WatchList = append(t.WatchList, WatchFunction{
		Signature: common.FromHex("0x3883e119"),
		Address:   make(map[string]BreakPointType),
	})
	t.WatchList = append(t.WatchList, WatchFunction{
		Signature: common.FromHex("0xed8e84f3"),
		Address:   make(map[string]BreakPointType),
	})
 	t.WatchList = append(t.WatchList, WatchFunction{
 		Signature: common.FromHex("0x41b028f3"),
 		Address:   make(map[string]BreakPointType),
 	})
	 t.WatchList = append(t.WatchList, WatchFunction{
		Signature: common.FromHex("0xcc2b27d7"),
		Address:   make(map[string]BreakPointType),
	})
	t.WatchList = append(t.WatchList, WatchFunction{
		Signature: common.FromHex("0xbb7b8b80"),
		Address:   make(map[string]BreakPointType),
	})
 	// t.WatchList = append(t.WatchList, WatchFunction{
 	// 	Signature: common.FromHex("0x861cdef0"),
 	// 	Address:   make(map[string]BreakPointType),
 	// })
 }

 func (t *tellerCore) loadKyberNetwork() {
 	// "0x809a9e55" getExpectedRate(address,address,uint256)
	// "0x217ac237" getPoolState()
 	// "7cd44272": "getConversionRate(address,address,uint256,uint256)",
 	t.WatchList = append(t.WatchList, WatchFunction{
 		Signature: common.FromHex("0x809a9e55"),
 		Address:   make(map[string]BreakPointType),
 	})
	 t.WatchList = append(t.WatchList, WatchFunction{
		Signature: common.FromHex("0x217ac237"),
		Address:   make(map[string]BreakPointType),
	})
 	// t.WatchList = append(t.WatchList, WatchFunction{
 	// 	Signature: common.FromHex("0x7cd44272"),
 	// 	Address:   make(map[string]BreakPointType),
 	// })
 }

 func (t *tellerCore) loadPlatypus() {
	// "43c2e2f5",  quotePotentialSwap(address,address,uint256)　
	// "907448ed", quotePotentialWithdraw(address,uint256)　　
	t.WatchList = append(t.WatchList, WatchFunction{
		Signature: common.FromHex("0x43c2e2f5"),
		Address:   make(map[string]BreakPointType),
	})
	t.WatchList = append(t.WatchList, WatchFunction{
		Signature: common.FromHex("0x907448ed"),
		Address:   make(map[string]BreakPointType),
	})
	// t.WatchList = append(t.WatchList, WatchFunction{
	// 	Signature: common.FromHex("0x7cd44272"),
	// 	Address:   make(map[string]BreakPointType),
	// })
}

func (t *tellerCore) loadTraderJoe() {
	// "ad615dec", quote(uint256,uint256,uint256)
	// "054d50d4", getAmountOut(uint256,uint256,uint256)
	// "85f8c259", getAmountIn(uint256,uint256,uint256)
	// "d06ca61f", getAmountsOut(uint256,address[])
	// "1f00ca74", getAmountsIn(uint256,address[])
	t.WatchList = append(t.WatchList, WatchFunction{
		Signature: common.FromHex("0xad615dec"),
		Address:   make(map[string]BreakPointType),
	})
	t.WatchList = append(t.WatchList, WatchFunction{
	   Signature: common.FromHex("0x054d50d4"),
	   Address:   make(map[string]BreakPointType),
	})
	t.WatchList = append(t.WatchList, WatchFunction{
		Signature: common.FromHex("0x85f8c259"),
		Address:   make(map[string]BreakPointType),
	})	
	t.WatchList = append(t.WatchList, WatchFunction{
		Signature: common.FromHex("0xd06ca61f"),
		Address:   make(map[string]BreakPointType),
	})	
	t.WatchList = append(t.WatchList, WatchFunction{
		Signature: common.FromHex("0x1f00ca74"),
		Address:   make(map[string]BreakPointType),
	})
	// t.WatchList = append(t.WatchList, WatchFunction{
	// 	Signature: common.FromHex("0x7cd44272"),
	// 	Address:   make(map[string]BreakPointType),
	// })
}

 
 func (t *tellerCore) loadConstantFunc() {
 	t.loadPangolin()
 	t.loadCrv()
	t.loadKyberNetwork()
 }