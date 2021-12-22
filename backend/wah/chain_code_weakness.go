package wah

var isDebug bool = false
var isFirstDetect = true
var hasRandomImport = false
var hasTimeImport = false
var inFuncDecl = false
var inBlockStmt = false
var inDeclStmt = false
var hasProbGotGlobal = false
var hasOsImport = false
var hasIoutilImport = false
//var osMethodMap = [...]string { "Chmod", "Chtimes", "CreateTemp", "Expand", "ExpandEnv", "FileMode", "Getenv", "LookupEnv", "MkdirTemp", "OpenFile", "ReadDir", "ReadFile", "Unsetenv", "WriteFile"}
var osMethodMap = map[string]bool { 
	"Chmod":true, 
	"Chtimes":true, 
	"CreateTemp":true, 
	"Expand":true, 
	"ExpandEnv":true, 
	"FileMode":true, 
	"Getenv":true, 
	"LookupEnv":true, 
	"MkdirTemp":true, 
	"OpenFile":true, 
	"ReadDir":true, 
	"ReadFile":true, 
	"Unsetenv":true, 
	"WriteFile":true,
	"Open":true,
	"Create":true,
}
type CCW int

const (
	MAP_STRUCTURE_ITER = iota + 1
	RANDOM_NUMBER_GENERATION
	GF_DECLARATION//
	UNCHECKED_INPUT_ARGUMENTS//
	UNHANDLED_ERROR
	USED_GOROUTINE
	PHANTOM_READS
	READ_YOUR_WRITE//
	RANGE_QUERY_RISK//
	USE_TIME_MODULE
	USED_GLOBAL_VARIABLE//1
	EXTERNAL_FILE_ACCESS//4
)

// RANGE_QUERY_RISK

func (c CCW) String() string {
	return [...]string{"MAP_STRUCTURE_ITER", "RANDOM_NUMBER_GENERATION", "GF_DECLARATION", 
					   "UNCHECKED_INPUT_ARGUMENTS", "UNHANDLED_ERROR", "USED_GOROUTINE",
					   "PHANTOM_READS", "READ_YOUR_WRITE", "RANGE_QUERY_RISK", "USE_TIME_MODULE", 
					   "USED_GLOBAL_VARIABLE", "EXTERNAL_FILE_ACCESS"}[c-1]
}
