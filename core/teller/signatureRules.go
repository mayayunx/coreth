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
 	// calc_token_amount(address,uint256[4],bool)	861cdef0
 	// calc_token_amount(uint256[2],bool)	ed8e84f3
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
 	// t.WatchList = append(t.WatchList, WatchFunction{
 	// 	Signature: common.FromHex("0x861cdef0"),
 	// 	Address:   make(map[string]BreakPointType),
 	// })
 }

 func (t *tellerCore) loadKyberNetwork() {
 	// "0x809a9e55" getExpectedRate(address,address,uint256)
 	// "7cd44272": "getConversionRate(address,address,uint256,uint256)",
 	t.WatchList = append(t.WatchList, WatchFunction{
 		Signature: common.FromHex("0x809a9e55"),
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