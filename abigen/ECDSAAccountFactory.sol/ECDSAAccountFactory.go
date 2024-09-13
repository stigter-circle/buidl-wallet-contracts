// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ECDSAAccountFactory

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ECDSAAccountFactoryMetaData contains all meta data concerning the ECDSAAccountFactory contract.
var ECDSAAccountFactoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_entryPoint\",\"type\":\"address\",\"internalType\":\"contractIEntryPoint\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"accountImplementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractECDSAAccount\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createAccount\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"contractECDSAAccount\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAccount\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"contractECDSAAccount\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAddress\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAddress\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AccountCreated\",\"inputs\":[{\"name\":\"proxy\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x60a0346100cc576001600160401b0390601f61225138819003918201601f1916830191848311848410176100b6578084926020946040528339810103126100cc57516001600160a01b038116908190036100cc5760405191611a2790818401908111848210176100b657602092849261082a843981520301906000f080156100aa5760805260405161075890816100d2823960805181818161017f015281816102ec01526103da0152f35b6040513d6000823e3d90fd5b634e487b7160e01b600052604160045260246000fd5b600080fdfe60806040818152600436101561001457600080fd5b600091823560e01c90816311464fbe1461016b575080637ac4ed64146101435780639859387b146100ec578063ae22c57d146100955763f14ddffc1461005957600080fd5b346100915780600319360112610091576020906100806100776101ae565b6024359061036a565b90516001600160a01b039091168152f35b5080fd5b5034610091576020366003190112610091576100806020926100b56101ae565b908351858101916bffffffffffffffffffffffff198460601b1683526034820152603481526100e3816101c9565b5190209061027b565b50346100915760203660031901126100915761008060209261010c6101ae565b908351858101916bffffffffffffffffffffffff198460601b16835260348201526034815261013a816101c9565b5190209061036a565b50346100915780600319360112610091576020906100806101626101ae565b6024359061027b565b8390346100915781600319360112610091577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b600435906001600160a01b03821682036101c457565b600080fd5b6060810190811067ffffffffffffffff8211176101e557604052565b634e487b7160e01b600052604160045260246000fd5b90601f8019910116810190811067ffffffffffffffff8211176101e557604052565b60005b8381106102305750506000910152565b8181015183820152602001610220565b909160609260018060a01b031682526040602083015261026f815180928160408601526020868601910161021d565b601f01601f1916010190565b600b906055926102d3602090610346610352836040968751906102a0838701836101fb565b858252828201956104508739885163189acdbd60e31b848201526001600160a01b03918216602480830191909152815261031290610320906102e1816101c9565b8b51928391878301957f00000000000000000000000000000000000000000000000000000000000000001686610240565b03601f1981018352826101fb565b8951958693610337868601998a925192839161021d565b8401915180938684019061021d565b010380845201826101fb565b5190208351938401528201523081520160ff81532090565b9190610376818461027b565b803b610441575060405163189acdbd60e31b60208201526001600160a01b03938416602480830182905282529391906103ae816101c9565b604051906102d38083019183831067ffffffffffffffff8411176101e5578392610400926104508539867f00000000000000000000000000000000000000000000000000000000000000001690610240565b03906000f58015610435571691827fac631f3001b55ea1509cf3d7e74898f85392a61a76e8149181ae1259622dabc8600080a3565b6040513d6000823e3d90fd5b6001600160a01b031692505056fe60806040526102d38038038061001481610194565b92833981019060408183031261018f5780516001600160a01b03811680820361018f5760208381015190936001600160401b03821161018f570184601f8201121561018f5780519061006d610068836101cf565b610194565b9582875285838301011161018f57849060005b83811061017b57505060009186010152813b15610163577f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b03191682179055604051907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b600080a28351156101455750600080848461012c96519101845af4903d1561013c573d61011c610068826101cf565b908152600081943d92013e6101ea565b505b6040516085908161024e8239f35b606092506101ea565b9250505034610154575061012e565b63b398979f60e01b8152600490fd5b60249060405190634c9c8ce360e01b82526004820152fd5b818101830151888201840152869201610080565b600080fd5b6040519190601f01601f191682016001600160401b038111838210176101b957604052565b634e487b7160e01b600052604160045260246000fd5b6001600160401b0381116101b957601f01601f191660200190565b9061021157508051156101ff57805190602001fd5b604051630a12f52160e11b8152600490fd5b81511580610244575b610222575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b1561021a56fe60806040527f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc54600090819081906001600160a01b0316368280378136915af43d82803e15604b573d90f35b3d90fdfea2646970667358221220d4df0841384884154f220c2629955c803d304e6525a53c5a1e6229a169328e9d64736f6c63430008180033a26469706673582212208542b3160b9ba032ed4f9f1f3bd33d4d6fc8f0aa7ecb2c26f1aa9cc90f4d63c464736f6c6343000818003360c034620000c257601f62001a2738819003918201601f19168301916001600160401b03831184841017620000c757808492602094604052833981010312620000c257516001600160a01b0381168103620000c25760805262000061620000dd565b3060a0526200006f620000dd565b6040516118aa90816200017d82396080518181816102f5015281816105a2015281816107b70152818161086101528181610d8801528181610f51015261139c015260a0518181816109070152610a400152f35b600080fd5b634e487b7160e01b600052604160045260246000fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805460ff8160401c166200016a576001600160401b036002600160401b0319828216016200012b57505050565b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1565b60405163f92ee8a960e01b8152600490fdfe60806040526004361015610023575b361561001957600080fd5b61002161124b565b005b60003560e01c806223de29146101bd57806301ffc9a7146101b857806306dc245c14610168578063150b7a02146101b35780631626ba7e146101ae57806318dfb3c7146101a957806319822f7c146101a45780633f4ba83a1461019f57806347e1da2a1461019a5780634a58db19146101955780634d44560d146101905780634f1ef2861461018b57806352d1902d146101865780635c975abb14610181578063715018a61461017c5780638456cb59146101775780638da5cb5b14610172578063ad3cb1cc1461016d578063b0d691fe14610168578063b61d27f614610163578063bc197c811461015e578063c399ec8814610159578063c4d66de814610154578063d087d2881461014f578063f23a6e611461014a5763f2fde38b0361000e57611016565b610fbc565b610f1e565b610de9565b610d5c565b610ccd565b610c70565b6102df565b610c18565b610ba2565b610b33565b610ac8565b610a98565b610a2d565b6108c5565b610828565b6107a0565b6106ca565b61065e565b61056c565b6104bd565b61043d565b610324565b610271565b610205565b6001600160a01b038116036101d357565b600080fd5b9181601f840112156101d3578235916001600160401b0383116101d357602083818601950101116101d357565b346101d35760c03660031901126101d3576102216004356101c2565b61022c6024356101c2565b6102376044356101c2565b6001600160401b036084358181116101d3576102579036906004016101d8565b505060a4359081116101d3576100219036906004016101d8565b346101d35760203660031901126101d35760043563ffffffff60e01b81168091036101d357602090630a85bd0160e11b81149081156102ce575b81156102bd575b506040519015158152f35b6301ffc9a760e01b149050386102b2565b630271189760e51b811491506102ab565b346101d35760003660031901126101d3576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346101d35760803660031901126101d3576103406004356101c2565b61034b6024356101c2565b6064356001600160401b0381116101d35761036a9036906004016101d8565b5050604051630a85bd0160e11b8152602090f35b634e487b7160e01b600052604160045260246000fd5b6001600160401b0381116103a757604052565b61037e565b90601f801991011681019081106001600160401b038211176103a757604052565b6001600160401b0381116103a757601f01601f191660200190565b9291926103f4826103cd565b9161040260405193846103ac565b8294818452818301116101d3578281602093846000960137010152565b9080601f830112156101d35781602061043a933591016103e8565b90565b346101d35760403660031901126101d3576024356001600160401b0381116101d35761047a610472602092369060040161041f565b600435611043565b6040516001600160e01b03199091168152f35b9181601f840112156101d3578235916001600160401b0383116101d3576020808501948460051b0101116101d357565b346101d35760403660031901126101d3576001600160401b036004358181116101d3576104ee90369060040161048d565b916024359081116101d35761050790369060040161048d565b9190610511611367565b610519611392565b6105248385146110b5565b60005b84811061053057005b80610566610541600193888761110d565b3561054b816101c2565b610560610559848988611154565b36916103e8565b9061142d565b01610527565b346101d3576003196060368201126101d357600435906001600160401b0382116101d3576101209082360301126101d3576044357f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163303610619576105e36105fb9260243590600401611497565b90806105ff575b506040519081529081906020820190565b0390f35b600080808093338219f15061061261150f565b50386105ea565b60405162461bcd60e51b815260206004820152601c60248201527f6163636f756e743a206e6f742066726f6d20456e747279506f696e74000000006044820152606490fd5b346101d35760003660031901126101d35761067761153f565b61067f6115b6565b6106876115b6565b60008051602061183583398151915260ff1981541690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b346101d35760603660031901126101d3576001600160401b036004358181116101d3576106fb90369060040161048d565b6024358381116101d35761071390369060040161048d565b936044359081116101d35761072c90369060040161048d565b92610735611367565b61073d611392565b6107488482146110b5565b6107538682146110b5565b60005b81811061075f57005b8061079a610770600193858a61110d565b3561077a816101c2565b610785838b8961110d565b35610794610559858b8a611154565b91611464565b01610756565b600080600319360112610825576107b5611367565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681813b156108255760405163b760faf960e01b8152306004820152918290602490829034905af1801561082057610814575080f35b61081d90610394565b80f35b61116f565b80fd5b346101d3576000604036600319011261082557600435610847816101c2565b61084f61153f565b610857611367565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811691839190833b156108c15760449083604051958694859363040b850f60e31b855216600484015260243560248401525af1801561082057610814575080f35b8280fd5b60403660031901126101d35760048035906108df826101c2565b6024356001600160401b0381116101d3576108fd903690830161041f565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116308114908115610a11575b50610a0057906020839261094561153f565b6040516352d1902d60e01b8152938491829088165afa600092816109cf575b50610992575050604051634c9c8ce360e01b81526001600160a01b0390921690820190815281906020010390fd5b838360008051602061181583398151915284036109b357610021838361161e565b604051632a87526960e21b815290810184815281906020010390fd5b6109f291935060203d6020116109f9575b6109ea81836103ac565b81019061117b565b9138610964565b503d6109e0565b60405163703e46dd60e11b81528390fd5b9050816000805160206118158339815191525416141538610933565b346101d35760003660031901126101d3577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163003610a865760206040516000805160206118158339815191528152f35b60405163703e46dd60e11b8152600490fd5b346101d35760003660031901126101d357602060ff60008051602061183583398151915254166040519015158152f35b346101d35760008060031936011261082557610ae261153f565b6000805160206117f583398151915280546001600160a01b0319811690915581906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b346101d35760003660031901126101d357610b4c61153f565b610b54611367565b610b5c611367565b600080516020611835833981519152600160ff198254161790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b346101d35760003660031901126101d3576000805160206117f5833981519152546040516001600160a01b039091168152602090f35b919082519283825260005b848110610c04575050826000602080949584010152601f8019910116010190565b602081830181015184830182015201610be3565b346101d35760003660031901126101d35760405160408101908082106001600160401b038311176103a7576105fb9160405260058152640352e302e360dc1b6020820152604051918291602083526020830190610bd8565b346101d35760603660031901126101d357600435610c8d816101c2565b604435906001600160401b0382116101d357610cc3610cb36100219336906004016101d8565b610cbb611367565b610559611392565b9060243590611464565b346101d35760a03660031901126101d357610ce96004356101c2565b610cf46024356101c2565b6001600160401b036044358181116101d357610d1490369060040161048d565b50506064358181116101d357610d2e90369060040161048d565b50506084359081116101d357610d489036906004016101d8565b505060405163bc197c8160e01b8152602090f35b346101d35760003660031901126101d3576040516370a0823160e01b81523060048201526020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa801561082057602091600091610dcc575b50604051908152f35b610de39150823d84116109f9576109ea81836103ac565b38610dc3565b346101d35760203660031901126101d357600435610e06816101c2565b60008051602061185583398151915254906001600160401b0360ff8360401c1615921680159081610f16575b6001149081610f0c575b159081610f03575b50610ef157600080516020611855833981519152805467ffffffffffffffff19166001179055610e789082610ec75761118a565b610e7e57005b600080516020611855833981519152805460ff60401b19169055604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b600080516020611855833981519152805460ff60401b19166801000000000000000017905561118a565b60405163f92ee8a960e01b8152600490fd5b90501538610e44565b303b159150610e3c565b839150610e32565b346101d35760003660031901126101d357604051631aab3f0d60e11b8152306004820152600060248201526020816044817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa8015610820576105fb91600091610f9d57506040519081529081906020820190565b610fb6915060203d6020116109f9576109ea81836103ac565b386105ea565b346101d35760a03660031901126101d357610fd86004356101c2565b610fe36024356101c2565b6084356001600160401b0381116101d3576110029036906004016101d8565b505060405163f23a6e6160e01b8152602090f35b346101d35760203660031901126101d357610021600435611036816101c2565b61103e61153f565b6111d7565b906110999161109360018060a01b036000805160206117f58339815191525416917f19457468657265756d205369676e6564204d6573736167653a0a333200000000600052601c52603c60002090565b90611279565b156110a957630b135d3f60e11b90565b6001600160e01b031990565b156110bc57565b60405162461bcd60e51b815260206004820152601360248201527277726f6e67206172726179206c656e6774687360681b6044820152606490fd5b634e487b7160e01b600052603260045260246000fd5b919081101561111d5760051b0190565b6110f7565b903590601e19813603018212156101d357018035906001600160401b0382116101d3576020019181360383136101d357565b9082101561111d5761116b9160051b810190611122565b9091565b6040513d6000823e3d90fd5b908160209103126101d3575190565b6111a6906111966116c5565b61119e6116c5565b61103e6116c5565b6111ae6116c5565b6111b66116c5565b600080516020611835833981519152805460ff191690556111d56116c5565b565b6001600160a01b03908116908115611232576000805160206117f583398151915280546001600160a01b031981168417909155167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3565b604051631e4fbdf760e01b815260006004820152602490fd5b6040513481527f6063d17f97b8837d6ec87876288b0dcba611bf6d68465a6e592be6f9d9ba4dc360203392a2565b61128383836115e2565b5060048195929510156113515715938461133b575b5083156112a6575b50505090565b60009293509082916040516112ed816112df6020820194630b135d3f60e11b998a87526024840152604060448401526064830190610bd8565b03601f1981018352826103ac565b51915afa906112fa61150f565b8261132d575b82611310575b50503880806112a0565b6113259192506020808251830101910161117b565b143880611306565b915060208251101591611300565b6001600160a01b03838116911614935038611298565b634e487b7160e01b600052602160045260246000fd5b60ff600080516020611835833981519152541661138057565b60405163d93c066560e01b8152600490fd5b60018060a01b03807f0000000000000000000000000000000000000000000000000000000000000000163314908115611412575b50156113ce57565b606460405162461bcd60e51b815260206004820152602060248201527f6163636f756e743a206e6f7420456e747279506f696e74206f72204f776e65726044820152fd5b90506000805160206117f583398151915254163314386113c6565b600091829182602083519301915af13d906040519060208383010160405282825260208201926000843e15611460575050565b5190fd5b916000928392602083519301915af13d906040519060208383010160405282825260208201926000843e15611460575050565b6114cd611500927f19457468657265756d205369676e6564204d6573736167653a0a333200000000600052601c52603c60002090565b6000805160206117f5833981519152546001600160a01b0316916114fa9061055990610100810190611122565b91611279565b1561150a57600090565b600190565b3d1561153a573d90611520826103cd565b9161152e60405193846103ac565b82523d6000602084013e565b606090565b6000805160206117f5833981519152546001600160a01b0316331480156115ad575b1561156857565b60405162461bcd60e51b815260206004820152601760248201527f43616c6c6572206973206e6f7420746865206f776e65720000000000000000006044820152606490fd5b50303314611561565b60ff6000805160206118358339815191525416156115d057565b604051638dfc202b60e01b8152600490fd5b81519190604183036116135761160c92506020820151906060604084015193015160001a906116f4565b9192909190565b505060009160029190565b90813b156116a45760008051602061181583398151915280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b600080a28051156116895761168691611778565b50565b50503461169257565b60405163b398979f60e01b8152600490fd5b604051634c9c8ce360e01b81526001600160a01b0383166004820152602490fd5b60ff6000805160206118558339815191525460401c16156116e257565b604051631afcd79f60e31b8152600490fd5b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841161176c57926020929160ff608095604051948552168484015260408301526060820152600092839182805260015afa156108205780516001600160a01b0381161561176357918190565b50809160019190565b50505060009160039190565b60008061043a93602081519101845af461179061150f565b91906117b857508051156117a657805190602001fd5b604051630a12f52160e11b8152600490fd5b815115806117eb575b6117c9575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b156117c156fe9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbccd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a2646970667358221220448dc9503740dd7944f3b2174822176887490cfad277ab777610394c94adc0ff64736f6c63430008180033",
}

// ECDSAAccountFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ECDSAAccountFactoryMetaData.ABI instead.
var ECDSAAccountFactoryABI = ECDSAAccountFactoryMetaData.ABI

// ECDSAAccountFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ECDSAAccountFactoryMetaData.Bin instead.
var ECDSAAccountFactoryBin = ECDSAAccountFactoryMetaData.Bin

// DeployECDSAAccountFactory deploys a new Ethereum contract, binding an instance of ECDSAAccountFactory to it.
func DeployECDSAAccountFactory(auth *bind.TransactOpts, backend bind.ContractBackend, _entryPoint common.Address) (common.Address, *types.Transaction, *ECDSAAccountFactory, error) {
	parsed, err := ECDSAAccountFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ECDSAAccountFactoryBin), backend, _entryPoint)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECDSAAccountFactory{ECDSAAccountFactoryCaller: ECDSAAccountFactoryCaller{contract: contract}, ECDSAAccountFactoryTransactor: ECDSAAccountFactoryTransactor{contract: contract}, ECDSAAccountFactoryFilterer: ECDSAAccountFactoryFilterer{contract: contract}}, nil
}

// ECDSAAccountFactory is an auto generated Go binding around an Ethereum contract.
type ECDSAAccountFactory struct {
	ECDSAAccountFactoryCaller     // Read-only binding to the contract
	ECDSAAccountFactoryTransactor // Write-only binding to the contract
	ECDSAAccountFactoryFilterer   // Log filterer for contract events
}

// ECDSAAccountFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ECDSAAccountFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAAccountFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECDSAAccountFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAAccountFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECDSAAccountFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAAccountFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECDSAAccountFactorySession struct {
	Contract     *ECDSAAccountFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ECDSAAccountFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECDSAAccountFactoryCallerSession struct {
	Contract *ECDSAAccountFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ECDSAAccountFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECDSAAccountFactoryTransactorSession struct {
	Contract     *ECDSAAccountFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ECDSAAccountFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ECDSAAccountFactoryRaw struct {
	Contract *ECDSAAccountFactory // Generic contract binding to access the raw methods on
}

// ECDSAAccountFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECDSAAccountFactoryCallerRaw struct {
	Contract *ECDSAAccountFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ECDSAAccountFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECDSAAccountFactoryTransactorRaw struct {
	Contract *ECDSAAccountFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewECDSAAccountFactory creates a new instance of ECDSAAccountFactory, bound to a specific deployed contract.
func NewECDSAAccountFactory(address common.Address, backend bind.ContractBackend) (*ECDSAAccountFactory, error) {
	contract, err := bindECDSAAccountFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECDSAAccountFactory{ECDSAAccountFactoryCaller: ECDSAAccountFactoryCaller{contract: contract}, ECDSAAccountFactoryTransactor: ECDSAAccountFactoryTransactor{contract: contract}, ECDSAAccountFactoryFilterer: ECDSAAccountFactoryFilterer{contract: contract}}, nil
}

// NewECDSAAccountFactoryCaller creates a new read-only instance of ECDSAAccountFactory, bound to a specific deployed contract.
func NewECDSAAccountFactoryCaller(address common.Address, caller bind.ContractCaller) (*ECDSAAccountFactoryCaller, error) {
	contract, err := bindECDSAAccountFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSAAccountFactoryCaller{contract: contract}, nil
}

// NewECDSAAccountFactoryTransactor creates a new write-only instance of ECDSAAccountFactory, bound to a specific deployed contract.
func NewECDSAAccountFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ECDSAAccountFactoryTransactor, error) {
	contract, err := bindECDSAAccountFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSAAccountFactoryTransactor{contract: contract}, nil
}

// NewECDSAAccountFactoryFilterer creates a new log filterer instance of ECDSAAccountFactory, bound to a specific deployed contract.
func NewECDSAAccountFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ECDSAAccountFactoryFilterer, error) {
	contract, err := bindECDSAAccountFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECDSAAccountFactoryFilterer{contract: contract}, nil
}

// bindECDSAAccountFactory binds a generic wrapper to an already deployed contract.
func bindECDSAAccountFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ECDSAAccountFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSAAccountFactory *ECDSAAccountFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSAAccountFactory.Contract.ECDSAAccountFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSAAccountFactory *ECDSAAccountFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSAAccountFactory.Contract.ECDSAAccountFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSAAccountFactory *ECDSAAccountFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSAAccountFactory.Contract.ECDSAAccountFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSAAccountFactory *ECDSAAccountFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSAAccountFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSAAccountFactory *ECDSAAccountFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSAAccountFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSAAccountFactory *ECDSAAccountFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSAAccountFactory.Contract.contract.Transact(opts, method, params...)
}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_ECDSAAccountFactory *ECDSAAccountFactoryCaller) AccountImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ECDSAAccountFactory.contract.Call(opts, &out, "accountImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_ECDSAAccountFactory *ECDSAAccountFactorySession) AccountImplementation() (common.Address, error) {
	return _ECDSAAccountFactory.Contract.AccountImplementation(&_ECDSAAccountFactory.CallOpts)
}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_ECDSAAccountFactory *ECDSAAccountFactoryCallerSession) AccountImplementation() (common.Address, error) {
	return _ECDSAAccountFactory.Contract.AccountImplementation(&_ECDSAAccountFactory.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0x7ac4ed64.
//
// Solidity: function getAddress(address _owner, bytes32 _salt) view returns(address)
func (_ECDSAAccountFactory *ECDSAAccountFactoryCaller) GetAddress(opts *bind.CallOpts, _owner common.Address, _salt [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ECDSAAccountFactory.contract.Call(opts, &out, "getAddress", _owner, _salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x7ac4ed64.
//
// Solidity: function getAddress(address _owner, bytes32 _salt) view returns(address)
func (_ECDSAAccountFactory *ECDSAAccountFactorySession) GetAddress(_owner common.Address, _salt [32]byte) (common.Address, error) {
	return _ECDSAAccountFactory.Contract.GetAddress(&_ECDSAAccountFactory.CallOpts, _owner, _salt)
}

// GetAddress is a free data retrieval call binding the contract method 0x7ac4ed64.
//
// Solidity: function getAddress(address _owner, bytes32 _salt) view returns(address)
func (_ECDSAAccountFactory *ECDSAAccountFactoryCallerSession) GetAddress(_owner common.Address, _salt [32]byte) (common.Address, error) {
	return _ECDSAAccountFactory.Contract.GetAddress(&_ECDSAAccountFactory.CallOpts, _owner, _salt)
}

// GetAddress0 is a free data retrieval call binding the contract method 0xae22c57d.
//
// Solidity: function getAddress(address _owner) view returns(address)
func (_ECDSAAccountFactory *ECDSAAccountFactoryCaller) GetAddress0(opts *bind.CallOpts, _owner common.Address) (common.Address, error) {
	var out []interface{}
	err := _ECDSAAccountFactory.contract.Call(opts, &out, "getAddress0", _owner)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress0 is a free data retrieval call binding the contract method 0xae22c57d.
//
// Solidity: function getAddress(address _owner) view returns(address)
func (_ECDSAAccountFactory *ECDSAAccountFactorySession) GetAddress0(_owner common.Address) (common.Address, error) {
	return _ECDSAAccountFactory.Contract.GetAddress0(&_ECDSAAccountFactory.CallOpts, _owner)
}

// GetAddress0 is a free data retrieval call binding the contract method 0xae22c57d.
//
// Solidity: function getAddress(address _owner) view returns(address)
func (_ECDSAAccountFactory *ECDSAAccountFactoryCallerSession) GetAddress0(_owner common.Address) (common.Address, error) {
	return _ECDSAAccountFactory.Contract.GetAddress0(&_ECDSAAccountFactory.CallOpts, _owner)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9859387b.
//
// Solidity: function createAccount(address _owner) returns(address account)
func (_ECDSAAccountFactory *ECDSAAccountFactoryTransactor) CreateAccount(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _ECDSAAccountFactory.contract.Transact(opts, "createAccount", _owner)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9859387b.
//
// Solidity: function createAccount(address _owner) returns(address account)
func (_ECDSAAccountFactory *ECDSAAccountFactorySession) CreateAccount(_owner common.Address) (*types.Transaction, error) {
	return _ECDSAAccountFactory.Contract.CreateAccount(&_ECDSAAccountFactory.TransactOpts, _owner)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9859387b.
//
// Solidity: function createAccount(address _owner) returns(address account)
func (_ECDSAAccountFactory *ECDSAAccountFactoryTransactorSession) CreateAccount(_owner common.Address) (*types.Transaction, error) {
	return _ECDSAAccountFactory.Contract.CreateAccount(&_ECDSAAccountFactory.TransactOpts, _owner)
}

// CreateAccount0 is a paid mutator transaction binding the contract method 0xf14ddffc.
//
// Solidity: function createAccount(address _owner, bytes32 _salt) returns(address account)
func (_ECDSAAccountFactory *ECDSAAccountFactoryTransactor) CreateAccount0(opts *bind.TransactOpts, _owner common.Address, _salt [32]byte) (*types.Transaction, error) {
	return _ECDSAAccountFactory.contract.Transact(opts, "createAccount0", _owner, _salt)
}

// CreateAccount0 is a paid mutator transaction binding the contract method 0xf14ddffc.
//
// Solidity: function createAccount(address _owner, bytes32 _salt) returns(address account)
func (_ECDSAAccountFactory *ECDSAAccountFactorySession) CreateAccount0(_owner common.Address, _salt [32]byte) (*types.Transaction, error) {
	return _ECDSAAccountFactory.Contract.CreateAccount0(&_ECDSAAccountFactory.TransactOpts, _owner, _salt)
}

// CreateAccount0 is a paid mutator transaction binding the contract method 0xf14ddffc.
//
// Solidity: function createAccount(address _owner, bytes32 _salt) returns(address account)
func (_ECDSAAccountFactory *ECDSAAccountFactoryTransactorSession) CreateAccount0(_owner common.Address, _salt [32]byte) (*types.Transaction, error) {
	return _ECDSAAccountFactory.Contract.CreateAccount0(&_ECDSAAccountFactory.TransactOpts, _owner, _salt)
}

// ECDSAAccountFactoryAccountCreatedIterator is returned from FilterAccountCreated and is used to iterate over the raw logs and unpacked data for AccountCreated events raised by the ECDSAAccountFactory contract.
type ECDSAAccountFactoryAccountCreatedIterator struct {
	Event *ECDSAAccountFactoryAccountCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ECDSAAccountFactoryAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ECDSAAccountFactoryAccountCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ECDSAAccountFactoryAccountCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ECDSAAccountFactoryAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ECDSAAccountFactoryAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ECDSAAccountFactoryAccountCreated represents a AccountCreated event raised by the ECDSAAccountFactory contract.
type ECDSAAccountFactoryAccountCreated struct {
	Proxy common.Address
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAccountCreated is a free log retrieval operation binding the contract event 0xac631f3001b55ea1509cf3d7e74898f85392a61a76e8149181ae1259622dabc8.
//
// Solidity: event AccountCreated(address indexed proxy, address indexed owner)
func (_ECDSAAccountFactory *ECDSAAccountFactoryFilterer) FilterAccountCreated(opts *bind.FilterOpts, proxy []common.Address, owner []common.Address) (*ECDSAAccountFactoryAccountCreatedIterator, error) {

	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ECDSAAccountFactory.contract.FilterLogs(opts, "AccountCreated", proxyRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ECDSAAccountFactoryAccountCreatedIterator{contract: _ECDSAAccountFactory.contract, event: "AccountCreated", logs: logs, sub: sub}, nil
}

// WatchAccountCreated is a free log subscription operation binding the contract event 0xac631f3001b55ea1509cf3d7e74898f85392a61a76e8149181ae1259622dabc8.
//
// Solidity: event AccountCreated(address indexed proxy, address indexed owner)
func (_ECDSAAccountFactory *ECDSAAccountFactoryFilterer) WatchAccountCreated(opts *bind.WatchOpts, sink chan<- *ECDSAAccountFactoryAccountCreated, proxy []common.Address, owner []common.Address) (event.Subscription, error) {

	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ECDSAAccountFactory.contract.WatchLogs(opts, "AccountCreated", proxyRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ECDSAAccountFactoryAccountCreated)
				if err := _ECDSAAccountFactory.contract.UnpackLog(event, "AccountCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccountCreated is a log parse operation binding the contract event 0xac631f3001b55ea1509cf3d7e74898f85392a61a76e8149181ae1259622dabc8.
//
// Solidity: event AccountCreated(address indexed proxy, address indexed owner)
func (_ECDSAAccountFactory *ECDSAAccountFactoryFilterer) ParseAccountCreated(log types.Log) (*ECDSAAccountFactoryAccountCreated, error) {
	event := new(ECDSAAccountFactoryAccountCreated)
	if err := _ECDSAAccountFactory.contract.UnpackLog(event, "AccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}