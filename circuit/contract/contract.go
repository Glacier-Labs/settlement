// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// GlacierVerifierMetaData contains all meta data concerning the GlacierVerifier contract.
var GlacierVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[1]\",\"name\":\"input\",\"type\":\"uint256[1]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611b8c806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806343753b4d14610030575b600080fd5b61004a600480360381019061004591906113e6565b610060565b604051610057919061146a565b60405180910390f35b600061006a61104c565b60405180604001604052808760006002811061008957610088611485565b5b60200201518152602001876001600281106100a7576100a6611485565b5b6020020151815250816000018190525060405180604001604052806040518060400160405280886000600281106100e1576100e0611485565b5b60200201516000600281106100f9576100f8611485565b5b602002015181526020018860006002811061011757610116611485565b5b602002015160016002811061012f5761012e611485565b5b6020020151815250815260200160405180604001604052808860016002811061015b5761015a611485565b5b602002015160006002811061017357610172611485565b5b602002015181526020018860016002811061019157610190611485565b5b60200201516001600281106101a9576101a8611485565b5b602002015181525081525081602001819052506040518060400160405280856000600281106101db576101da611485565b5b60200201518152602001856001600281106101f9576101f8611485565b5b602002015181525081604001819052507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4781600001516000015110610273576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026a90611511565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47816000015160200151106102dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102d49061157d565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4781602001516000015160006002811061031a57610319611485565b5b60200201511061035f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610356906115e9565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4781602001516020015160006002811061039c5761039b611485565b5b6020020151106103e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103d890611655565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4781602001516000015160016002811061041e5761041d611485565b5b602002015110610463576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161045a906116c1565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478160200151602001516001600281106104a05761049f611485565b5b6020020151106104e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104dc9061172d565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478160400151600001511061054f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054690611799565b60405180910390fd5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47816040015160200151106105b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105b090611805565b60405180910390fd5b60005b6001811015610651577f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018482600181106105f9576105f8611485565b5b60200201351061063e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161063590611871565b60405180910390fd5b8080610649906118c0565b9150506105bc565b50600061065c6107fa565b905060006040518060400160405280600081526020016000815250905061068161107f565b6106896110a1565b6000604051806040016040528060008152602001600081525090507f15442a2ac58b1fbf9053991305a5cbf0e22767b430638d3de24076ed5dc7ab928460000181815250507f01f9882633f0960584deaa18f403dfb81d953b9f62533d61e9ed5e65baf1b28e8460200181815250507f1a711dc85d2f511c7bb95242baad83a0d0e78f8078434b6636ade1d9ceea95d98260006003811061072d5761072c611485565b5b6020020181815250507f042ec05c7989f6f226eeb999dd4b2f0f08c55765296f998a69d313dff03f53448260016003811061076b5761076a611485565b5b6020020181815250508760006001811061078857610787611485565b5b6020020135826002600381106107a1576107a0611485565b5b6020020181815250506107b682828587610ab0565b6107ea6107c68760000151610b52565b876020015187600001518860200151888a604001518c604001518c60600151610c10565b9650505050505050949350505050565b6108026110c3565b60405180604001604052807e3920c2cd4a8f0a4548afb9b83dea1abef965a443104bf3463e5e887959c08381526020017f192ea1aa688ea207c98738985c5a3ffa7b92906fc57ae57af997b5f9791e49438152508160000181905250604051806040016040528060405180604001604052807f0b278a0b8ce9cbdfde2f4b6dbc86a85249df1166cf0a46472e0a3c3dc30da9b481526020017f2b31d4adf57d2e7810cf92c3326b07046f6a8e576f81dc8da7f177545f768f4a815250815260200160405180604001604052807f116402f87a5976c6babc1ee43a0c6cc17f4de5222846c7bb6d6ce70d2341e5ae81526020017f113d99d3963ba1844328309f8eabe5fa300f5b2d135cb862af2944baed099d8c8152508152508160200181905250604051806040016040528060405180604001604052807f04730be7aa1c8e25da079667f0442a5018d4c3f035090261f6854d1774d2418081526020017f1b59f096f71438e35465a66b4d3aacf4b05a24efade63aad8fe9a80c089b9458815250815260200160405180604001604052807f02e0b5c692f3cccc21c49e84981c9298956e5759e8d18e3b3ae2932167fac64c81526020017f169034f497f2e04b228ab4a9f45a01e1f77d8361644c44e01ee8979f8b7f45538152508152508160400181905250604051806040016040528060405180604001604052807f2a686464c19686d4f160c5c28892c6c4305d62dde97db8731bcf63c9da2b212a81526020017f0d42c3a41d6058f73a75dc0d29708293f1be974a2a0f93d829cce8acb0e73379815250815260200160405180604001604052807f17144e385de4b80dea0b636a3213bff70b31139df026a80c0fa5c551ea4aa2d081526020017f04bce3a0c3e4b592adde274dc885bf60ecb5a857ce723a388eb515938f519038815250815250816060018190525090565b610aba8484610f86565b806000015182600060048110610ad357610ad2611485565b5b602002018181525050806020015182600160048110610af557610af4611485565b5b602002018181525050826000015182600260048110610b1757610b16611485565b5b602002018181525050826020015182600360048110610b3957610b38611485565b5b602002018181525050610b4c8282610fe9565b50505050565b610b5a611103565b60008260000151148015610b72575060008260200151145b15610b955760405180604001604052806000815260200160008152509050610c0b565b6040518060400160405280836000015181526020017f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478460200151610bda9190611937565b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47610c059190611968565b81525090505b919050565b60008060405180608001604052808b8152602001898152602001878152602001858152509050600060405180608001604052808b815260200189815260200187815260200185815250905060006018905060008167ffffffffffffffff811115610c7d57610c7c6111ac565b5b604051908082528060200260200182016040528015610cab5781602001602082028036833780820191505090505b50905060005b6004811015610ee9576000600682610cc9919061199c565b9050858260048110610cde57610cdd611485565b5b60200201516000015183600083610cf591906119de565b81518110610d0657610d05611485565b5b602002602001018181525050858260048110610d2557610d24611485565b5b60200201516020015183600183610d3c91906119de565b81518110610d4d57610d4c611485565b5b602002602001018181525050848260048110610d6c57610d6b611485565b5b602002015160000151600060028110610d8857610d87611485565b5b602002015183600283610d9b91906119de565b81518110610dac57610dab611485565b5b602002602001018181525050848260048110610dcb57610dca611485565b5b602002015160000151600160028110610de757610de6611485565b5b602002015183600383610dfa91906119de565b81518110610e0b57610e0a611485565b5b602002602001018181525050848260048110610e2a57610e29611485565b5b602002015160200151600060028110610e4657610e45611485565b5b602002015183600483610e5991906119de565b81518110610e6a57610e69611485565b5b602002602001018181525050848260048110610e8957610e88611485565b5b602002015160200151600160028110610ea557610ea4611485565b5b602002015183600583610eb891906119de565b81518110610ec957610ec8611485565b5b602002602001018181525050508080610ee1906118c0565b915050610cb1565b50610ef261111d565b6000602082602086026020860160086107d05a03fa90508060008103610f1457fe5b5080610f55576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f4c90611a5e565b60405180910390fd5b600082600060018110610f6b57610f6a611485565b5b60200201511415965050505050505098975050505050505050565b600060608260808560076107d05a03fa90508060008103610fa357fe5b5080610fe4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fdb90611aca565b60405180910390fd5b505050565b600060608260c08560066107d05a03fa9050806000810361100657fe5b5080611047576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161103e90611b36565b60405180910390fd5b505050565b604051806060016040528061105f611103565b815260200161106c61113f565b8152602001611079611103565b81525090565b6040518060800160405280600490602082028036833780820191505090505090565b6040518060600160405280600390602082028036833780820191505090505090565b60405180608001604052806110d6611103565b81526020016110e361113f565b81526020016110f061113f565b81526020016110fd61113f565b81525090565b604051806040016040528060008152602001600081525090565b6040518060200160405280600190602082028036833780820191505090505090565b6040518060400160405280611152611165565b815260200161115f611165565b81525090565b6040518060400160405280600290602082028036833780820191505090505090565b6000604051905090565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6111e48261119b565b810181811067ffffffffffffffff82111715611203576112026111ac565b5b80604052505050565b6000611216611187565b905061122282826111db565b919050565b600067ffffffffffffffff821115611242576112416111ac565b5b602082029050919050565b600080fd5b6000819050919050565b61126581611252565b811461127057600080fd5b50565b6000813590506112828161125c565b92915050565b600061129b61129684611227565b61120c565b905080602084028301858111156112b5576112b461124d565b5b835b818110156112de57806112ca8882611273565b8452602084019350506020810190506112b7565b5050509392505050565b600082601f8301126112fd576112fc611196565b5b600261130a848285611288565b91505092915050565b600067ffffffffffffffff82111561132e5761132d6111ac565b5b602082029050919050565b600061134c61134784611313565b61120c565b905080604084028301858111156113665761136561124d565b5b835b8181101561138f578061137b88826112e8565b845260208401935050604081019050611368565b5050509392505050565b600082601f8301126113ae576113ad611196565b5b60026113bb848285611339565b91505092915050565b6000819050826020600102820111156113e0576113df61124d565b5b92915050565b600080600080610120858703121561140157611400611191565b5b600061140f878288016112e8565b945050604061142087828801611399565b93505060c0611431878288016112e8565b925050610100611443878288016113c4565b91505092959194509250565b60008115159050919050565b6114648161144f565b82525050565b600060208201905061147f600083018461145b565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082825260208201905092915050565b7f76657269666965722d61582d6774652d7072696d652d71000000000000000000600082015250565b60006114fb6017836114b4565b9150611506826114c5565b602082019050919050565b6000602082019050818103600083015261152a816114ee565b9050919050565b7f76657269666965722d61592d6774652d7072696d652d71000000000000000000600082015250565b60006115676017836114b4565b915061157282611531565b602082019050919050565b600060208201905081810360008301526115968161155a565b9050919050565b7f76657269666965722d6258302d6774652d7072696d652d710000000000000000600082015250565b60006115d36018836114b4565b91506115de8261159d565b602082019050919050565b60006020820190508181036000830152611602816115c6565b9050919050565b7f76657269666965722d6259302d6774652d7072696d652d710000000000000000600082015250565b600061163f6018836114b4565b915061164a82611609565b602082019050919050565b6000602082019050818103600083015261166e81611632565b9050919050565b7f76657269666965722d6258312d6774652d7072696d652d710000000000000000600082015250565b60006116ab6018836114b4565b91506116b682611675565b602082019050919050565b600060208201905081810360008301526116da8161169e565b9050919050565b7f76657269666965722d6259312d6774652d7072696d652d710000000000000000600082015250565b60006117176018836114b4565b9150611722826116e1565b602082019050919050565b600060208201905081810360008301526117468161170a565b9050919050565b7f76657269666965722d63582d6774652d7072696d652d71000000000000000000600082015250565b60006117836017836114b4565b915061178e8261174d565b602082019050919050565b600060208201905081810360008301526117b281611776565b9050919050565b7f76657269666965722d63592d6774652d7072696d652d71000000000000000000600082015250565b60006117ef6017836114b4565b91506117fa826117b9565b602082019050919050565b6000602082019050818103600083015261181e816117e2565b9050919050565b7f76657269666965722d6774652d736e61726b2d7363616c61722d6669656c6400600082015250565b600061185b601f836114b4565b915061186682611825565b602082019050919050565b6000602082019050818103600083015261188a8161184e565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006118cb82611252565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036118fd576118fc611891565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061194282611252565b915061194d83611252565b92508261195d5761195c611908565b5b828206905092915050565b600061197382611252565b915061197e83611252565b925082820390508181111561199657611995611891565b5b92915050565b60006119a782611252565b91506119b283611252565b92508282026119c081611252565b915082820484148315176119d7576119d6611891565b5b5092915050565b60006119e982611252565b91506119f483611252565b9250828201905080821115611a0c57611a0b611891565b5b92915050565b7f70616972696e672d6f70636f64652d6661696c65640000000000000000000000600082015250565b6000611a486015836114b4565b9150611a5382611a12565b602082019050919050565b60006020820190508181036000830152611a7781611a3b565b9050919050565b7f70616972696e672d6d756c2d6661696c65640000000000000000000000000000600082015250565b6000611ab46012836114b4565b9150611abf82611a7e565b602082019050919050565b60006020820190508181036000830152611ae381611aa7565b9050919050565b7f70616972696e672d6164642d6661696c65640000000000000000000000000000600082015250565b6000611b206012836114b4565b9150611b2b82611aea565b602082019050919050565b60006020820190508181036000830152611b4f81611b13565b905091905056fea26469706673582212201bf7c14c96e97c201197393c32e19aff3b8cd3f3a52ce3acf5403d9693f0b98764736f6c63430008120033",
}

// GlacierVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use GlacierVerifierMetaData.ABI instead.
var GlacierVerifierABI = GlacierVerifierMetaData.ABI

// GlacierVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GlacierVerifierMetaData.Bin instead.
var GlacierVerifierBin = GlacierVerifierMetaData.Bin

// DeployGlacierVerifier deploys a new Ethereum contract, binding an instance of GlacierVerifier to it.
func DeployGlacierVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GlacierVerifier, error) {
	parsed, err := GlacierVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GlacierVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GlacierVerifier{GlacierVerifierCaller: GlacierVerifierCaller{contract: contract}, GlacierVerifierTransactor: GlacierVerifierTransactor{contract: contract}, GlacierVerifierFilterer: GlacierVerifierFilterer{contract: contract}}, nil
}

// GlacierVerifier is an auto generated Go binding around an Ethereum contract.
type GlacierVerifier struct {
	GlacierVerifierCaller     // Read-only binding to the contract
	GlacierVerifierTransactor // Write-only binding to the contract
	GlacierVerifierFilterer   // Log filterer for contract events
}

// GlacierVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type GlacierVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlacierVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GlacierVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlacierVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GlacierVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlacierVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GlacierVerifierSession struct {
	Contract     *GlacierVerifier  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GlacierVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GlacierVerifierCallerSession struct {
	Contract *GlacierVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// GlacierVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GlacierVerifierTransactorSession struct {
	Contract     *GlacierVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// GlacierVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type GlacierVerifierRaw struct {
	Contract *GlacierVerifier // Generic contract binding to access the raw methods on
}

// GlacierVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GlacierVerifierCallerRaw struct {
	Contract *GlacierVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// GlacierVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GlacierVerifierTransactorRaw struct {
	Contract *GlacierVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGlacierVerifier creates a new instance of GlacierVerifier, bound to a specific deployed contract.
func NewGlacierVerifier(address common.Address, backend bind.ContractBackend) (*GlacierVerifier, error) {
	contract, err := bindGlacierVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GlacierVerifier{GlacierVerifierCaller: GlacierVerifierCaller{contract: contract}, GlacierVerifierTransactor: GlacierVerifierTransactor{contract: contract}, GlacierVerifierFilterer: GlacierVerifierFilterer{contract: contract}}, nil
}

// NewGlacierVerifierCaller creates a new read-only instance of GlacierVerifier, bound to a specific deployed contract.
func NewGlacierVerifierCaller(address common.Address, caller bind.ContractCaller) (*GlacierVerifierCaller, error) {
	contract, err := bindGlacierVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GlacierVerifierCaller{contract: contract}, nil
}

// NewGlacierVerifierTransactor creates a new write-only instance of GlacierVerifier, bound to a specific deployed contract.
func NewGlacierVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*GlacierVerifierTransactor, error) {
	contract, err := bindGlacierVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GlacierVerifierTransactor{contract: contract}, nil
}

// NewGlacierVerifierFilterer creates a new log filterer instance of GlacierVerifier, bound to a specific deployed contract.
func NewGlacierVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*GlacierVerifierFilterer, error) {
	contract, err := bindGlacierVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GlacierVerifierFilterer{contract: contract}, nil
}

// bindGlacierVerifier binds a generic wrapper to an already deployed contract.
func bindGlacierVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GlacierVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlacierVerifier *GlacierVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GlacierVerifier.Contract.GlacierVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlacierVerifier *GlacierVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlacierVerifier.Contract.GlacierVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlacierVerifier *GlacierVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlacierVerifier.Contract.GlacierVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlacierVerifier *GlacierVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GlacierVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlacierVerifier *GlacierVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlacierVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlacierVerifier *GlacierVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlacierVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyProof is a free data retrieval call binding the contract method 0x43753b4d.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[1] input) view returns(bool r)
func (_GlacierVerifier *GlacierVerifierCaller) VerifyProof(opts *bind.CallOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [1]*big.Int) (bool, error) {
	var out []interface{}
	err := _GlacierVerifier.contract.Call(opts, &out, "verifyProof", a, b, c, input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x43753b4d.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[1] input) view returns(bool r)
func (_GlacierVerifier *GlacierVerifierSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [1]*big.Int) (bool, error) {
	return _GlacierVerifier.Contract.VerifyProof(&_GlacierVerifier.CallOpts, a, b, c, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x43753b4d.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[1] input) view returns(bool r)
func (_GlacierVerifier *GlacierVerifierCallerSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [1]*big.Int) (bool, error) {
	return _GlacierVerifier.Contract.VerifyProof(&_GlacierVerifier.CallOpts, a, b, c, input)
}
