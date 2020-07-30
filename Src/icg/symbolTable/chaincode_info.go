package symbolTable

type ChaincodeType int

const (
	Chaincode = iota
	ChaincodeStubInterface
	CommonIteratorInterface
	StateQueryIteratorInterface
	HistoryQueryIteratorInterface
	MockQueryIteratorInterface
)

var TypeConvertMap = map[string]ChaincodeType{
	"Chaincode":                     Chaincode,
	"ChaincodeStubInterface":        ChaincodeStubInterface,
	"CommonIteratorInterface":       CommonIteratorInterface,
	"StateQueryIteratorInterface":   StateQueryIteratorInterface,
	"HistoryQueryIteratorInterface": HistoryQueryIteratorInterface,
	"MockQueryIteratorInterface":    MockQueryIteratorInterface,
}

func TypeToString(t ChaincodeType) string {
	names := [...]string {
		"Chaincode","ChaincodeStubInterface","CommonIteratorInterface","StateQueryIteratorInterface","HistoryQueryIteratorInterface",
		"MockQueryIteratorInterface",
	}

	return names[t]
}

type InterfaceFunction int

const (
	Init = iota + 345
	Invoke
	GetArgs
	GetStringArgs
	GetFunctionAndParameters
	GetArgsSlice
	GetTxID
	GetChannelID
	InvokeChaincode
	GetState
	PutState
	DelState
	SetStateValidationParameter
	GetStateValidationParameter
	GetStateByRange
	GetStateByRangeWithPagination
	GetStateByPartialCompositeKey
	GetStateByPartialCompositeKeyWithPagination
	CreateCompositeKey
	SplitCompositeKey
	GetQueryResult
	GetQueryResultWithPagination
	GetHistoryForKey
	GetPrivateData
	GetPrivateDataHash
	PutPrivateData
	DelPrivateData
	SetPrivateDataValidationParameter
	GetPrivateDataValidationParameter
	GetPrivateDataByRange
	GetPrivateDataByPartialCompositeKey
	GetPrivateDataQueryResult
	GetCreator
	GetTransient
	GetBinding
	GetDecorations
	GetSignedProposal
	GetTxTimestamp
	SetEvent
	HasNext
	Close
	NextOfStateQueryIteratorInterface
	NextOfHistoryQueryIteratorInterface
)

var FunctionConvertMap = map[string]InterfaceFunction{
	"Init":                          Init,
	"Invoke":                        Invoke,
	"GetArgs":                       GetArgs,
	"GetStringArgs":                 GetStringArgs,
	"GetFunctionAndParameters":      GetFunctionAndParameters,
	"GetArgsSlice":                  GetArgsSlice,
	"GetTxID":                       GetTxID,
	"GetChannelID":                  GetChannelID,
	"InvokeChaincode":               InvokeChaincode,
	"GetState":                      GetState,
	"PutState":                      PutState,
	"DelState":                      DelState,
	"SetStateValidationParameter":   SetStateValidationParameter,
	"GetStateValidationParameter":   GetStateValidationParameter,
	"GetStateByRange":               GetStateByRange,
	"GetStateByRangeWithPagination": GetStateByRangeWithPagination,
	"GetStateByPartialCompositeKey": GetStateByPartialCompositeKey,
	"GetStateByPartialCompositeKeyWithPagination": GetStateByPartialCompositeKeyWithPagination,
	"CreateCompositeKey":                          CreateCompositeKey,
	"SplitCompositeKey":                           SplitCompositeKey,
	"GetQueryResult":                              GetQueryResult,
	"GetQueryResultWithPagination":                GetQueryResultWithPagination,
	"GetHistoryForKey":                            GetHistoryForKey,
	"GetPrivateData":                              GetPrivateData,
	"GetPrivateDataHash":                          GetPrivateDataHash,
	"PutPrivateData":                              PutPrivateData,
	"DelPrivateData":                              DelPrivateData,
	"SetPrivateDataValidationParameter":           SetPrivateDataValidationParameter,
	"GetPrivateDataValidationParameter":           GetPrivateDataValidationParameter,
	"GetPrivateDataByRange":                       GetPrivateDataByRange,
	"GetPrivateDataByPartialCompositeKey":         GetPrivateDataByPartialCompositeKey,
	"GetPrivateDataQueryResult":                   GetPrivateDataQueryResult,
	"GetCreator":                                  GetCreator,
	"GetTransient":                                GetTransient,
	"GetBinding":                                  GetBinding,
	"GetDecorations":                              GetDecorations,
	"GetSignedProposal":                           GetSignedProposal,
	"GetTxTimestamp":                              GetTxTimestamp,
	"SetEvent":                                    SetEvent,
	"HasNext":                                     HasNext,
	"Close":                                       Close,
	"NextOfStateQueryIteratorInterface":           NextOfStateQueryIteratorInterface,
	"NextOfHistoryQueryIteratorInterface":         NextOfHistoryQueryIteratorInterface,
}

type InterfaceFunctionList []InterfaceFunction

var ChaincodeMethodTable = map[ChaincodeType]InterfaceFunctionList{
	Chaincode: {Init, Invoke},
	ChaincodeStubInterface: {
		GetArgs, GetStringArgs, GetFunctionAndParameters, GetArgsSlice,
		GetTxID, GetChannelID, InvokeChaincode, GetState,
		PutState, DelState, SetStateValidationParameter, GetStateValidationParameter,
		GetStateByRange, GetStateByRangeWithPagination, GetStateByPartialCompositeKey, GetStateByPartialCompositeKeyWithPagination,
		CreateCompositeKey, SplitCompositeKey, GetQueryResult, GetQueryResultWithPagination,
		GetHistoryForKey, GetPrivateData, GetPrivateDataHash, PutPrivateData,
		DelPrivateData, SetPrivateDataValidationParameter, GetPrivateDataValidationParameter, GetPrivateDataByRange,
		GetPrivateDataByPartialCompositeKey, GetPrivateDataQueryResult, GetCreator,
		GetTransient, GetBinding, GetDecorations, GetSignedProposal,
		GetTxTimestamp, SetEvent},
	CommonIteratorInterface:       {HasNext, Close},
	StateQueryIteratorInterface:   {HasNext, Close, NextOfStateQueryIteratorInterface},
	HistoryQueryIteratorInterface: {HasNext, Close, NextOfHistoryQueryIteratorInterface},
	MockQueryIteratorInterface:    {HasNext, Close, NextOfStateQueryIteratorInterface},
}

type BlockChaincodeSymbolTable struct {
	table map[int]*ChaincodeSymbolTable
}

func (table *BlockChaincodeSymbolTable) Init() {
	table.table = make(map[int]*ChaincodeSymbolTable)
}
func (o *BlockChaincodeSymbolTable) Insert(k int, v *ChaincodeSymbolTable) {
	if _, ok := o.table[k]; !ok {
		o.table[k] = v
	}
}
func (o *BlockChaincodeSymbolTable) Table() map[int]*ChaincodeSymbolTable {
	return o.table
}
func (o *BlockChaincodeSymbolTable) GetSymTable(k int) *ChaincodeSymbolTable {
	if val, ok := o.table[k]; ok {
		return val
	}
	return nil
}

type ChaincodeSymbolTable struct {
	table       map[string]ChaincodeType
	ParentBlock int
	BlockNum    int
}

func (table *ChaincodeSymbolTable) Init() {
	table.table = make(map[string]ChaincodeType)
	table.ParentBlock = -1
	table.BlockNum = 0
}
func (table *ChaincodeSymbolTable) Insert(k string, v ChaincodeType) {
	if _, ok := table.table[k]; !ok {
		table.table[k] = v
	}
}
func (table *ChaincodeSymbolTable) GetType(sym string) (ChaincodeType, bool) {
	var res ChaincodeType
	resbool := false
	if v, ok := table.table[sym]; ok {
		res = v
		resbool = ok
	}

	return res, resbool
}
func (table *ChaincodeSymbolTable) GetTable() map[string]ChaincodeType {
	return table.table
}
