package main

import (
	"./icg"
	"./icg/symbolTable"
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"os"
)

func SAFHeaderGen(litTable *symbolTable.LiteralTable, source string) string {
	table := litTable.GetTable()
	litCount := len(table)
	sources := strings.Split(source, "/")
	fileName := sources[len(sources)-1]

	buffer := strings.Builder{}
	buffer.WriteString("%%HeaderSectionStart\n")
	buffer.WriteString("\t%" + fmt.Sprintf("DefinedLiteralCount\t%d", litCount) + "\n")
	buffer.WriteString("\t%" + fmt.Sprintf("InitializedVariableCount\t%d", 0) + "\n")
	buffer.WriteString("\t%" + fmt.Sprintf("ExternalVariableCount\t%d", 0) + "\n")
	buffer.WriteString("\t%" + fmt.Sprintf("ExternalFunctionCount\t%d", 0) + "\n")
	buffer.WriteString("\t%" + fmt.Sprintf("InitFunctionName\t%d", 0) + "\n")
	buffer.WriteString("\t%" + fmt.Sprintf("EntryFunctionName\t%s", "&main") + "\n")
	buffer.WriteString("\t%" + fmt.Sprintf("SourceFileName\t%s", fileName) + "\n")
	buffer.WriteString("\t%" + fmt.Sprintf("DebugMode\t%s", "Release") + "\n")
	buffer.WriteString("%%HeaderSectionEnd\n")

	return buffer.String()

}
func SAFFileGen(silTable *icg.SILTable, litTable *symbolTable.LiteralTable, sourceFile string) {
	//Header gen
	buffer := strings.Builder{}
	buffer.WriteString(SAFHeaderGen(litTable, sourceFile))
	buffer.WriteString(CodeSectionGen(silTable, "main"))
	buffer.WriteString(DataSectionGen(litTable))

	data := []byte(buffer.String())
	sources := strings.Split(sourceFile, "/")
	fileName := sources[len(sources)-1]
	fileName = "./SILResult/" + fileName + ".saf"
	ioutil.WriteFile(fileName, data, 0644)
}
func DataSectionGen(litTable *symbolTable.LiteralTable) string {
	buffer := strings.Builder{}
	buffer.WriteString("%%DataSectionStart\n")

	buffer.WriteString("\t%LiteralTableStart\n")
	for k, v := range litTable.GetTable() {
		byteStr := []byte(k)
		buffer.WriteString(fmt.Sprintf("\t\t.literal_start\t@%d\t%d\t%d\n", v, 0, len(k)-1))
		buffer.WriteString("\t\t\t")
		for i, b := range byteStr {
			if i > 0 && i < len(byteStr)-1 {
				buffer.WriteString(fmt.Sprintf("0x%x", b))
				buffer.WriteString(",")
			}
			if i == len(byteStr)-1 {
				buffer.WriteString("0x00")
			}
		}
		buffer.WriteString("\n")
		buffer.WriteString("\t\t.literal_end\n")
	}
	buffer.WriteString("\t%LiteralTableEnd\n")

	//Internal Symbol table
	buffer.WriteString("\t%InternalSymbolTableStart\n")
	buffer.WriteString("\t%InternalSymbolTableEnd\n")
	//ExternalSymbolTable
	buffer.WriteString("\t%ExternalSymbolTableStart\n")
	buffer.WriteString("\t%ExternalSymbolTableEnd\n")
	buffer.WriteString("%%DataSectionEnd\n")
	return buffer.String()
}
func SAFFuncGen(paramCount int, funcName string, funcCode []icg.CodeInfo) string {
	buffer := strings.Builder{}
	//function info
	buffer.WriteString("\t%FunctionStart\n")
	buffer.WriteString("\t\t" + fmt.Sprintf(".func_name\t&%s\n", funcName))
	buffer.WriteString("\t\t" + fmt.Sprintf(".func_type\t%d\n", 2))
	buffer.WriteString("\t\t" + fmt.Sprintf(".param_count\t%d\n", paramCount))

	//opcode gen
	buffer.WriteString("\t\t.opcode_start\n")

	for _, info := range funcCode {

		if stackInfo, ok := info.(*icg.StackOpcode); ok {
			buffer.WriteString("\t\t\t" + stackInfo.String() + "\n")
		}
		if arithinfo, ok := info.(*icg.ArithmeticOpcode); ok {
			buffer.WriteString("\t\t\t" + arithinfo.String() + "\n")
		}
		if cinfo, ok := info.(*icg.ControlOpcode); ok {
			if cinfo.Opcode() == icg.Label {
				buffer.WriteString("\t\t%" + cinfo.String() + "\n")
			} else {
				buffer.WriteString("\t\t\t" + cinfo.String() + "\n")
			}
		}

	}
	buffer.WriteString("\t\t.opcode_end\n")
	buffer.WriteString("\t%FunctionEnd\n")

	return buffer.String()
}
func CodeSectionGen(table *icg.SILTable, entryFunction string) string {
	buffer := strings.Builder{}

	for k, v := range table.FunctionCodeTable() {
		funcName := table.StringPool().LookupSymbolName(k)
		paramCount := table.FunctionParamCountTable()[k]
		funcCode := SAFFuncGen(paramCount, funcName, v)

		if funcName == entryFunction {
			currentFuncCode := buffer.String()
			buffer.Reset()

			currentFuncCode = "%%CodeSectionStart\n" +funcCode + currentFuncCode
			buffer.WriteString(currentFuncCode)
		} else {
			buffer.WriteString(funcCode)
		}
	}
	buffer.WriteString("%%CodeSectionEnd\n")

	return buffer.String()
}

func PrintSIL(table *icg.SILTable) {
	buffer := strings.Builder{}
	for k, v := range table.FunctionCodeTable() {
		paramTable := table.FunctionParamCountTable()
		buffer.WriteString(fmt.Sprintf("paramCount : %d\n", paramTable[k]))
		buffer.WriteString("------------------------------\n")
		buffer.WriteString(fmt.Sprintf("\tFunction : %s \n", table.StringPool().LookupSymbolName(k)))
		buffer.WriteString("------------------------------\n")

		for _, info := range v {

			if stackInfo, ok := info.(*icg.StackOpcode); ok {
				buffer.WriteString("\t" + strconv.Itoa(stackInfo.GetLine()) + ": " + stackInfo.String() + "\n")
			}
			if arithinfo, ok := info.(*icg.ArithmeticOpcode); ok {
				buffer.WriteString("\t" + strconv.Itoa(arithinfo.GetLine()) + ": " + arithinfo.String() + "\n")
			}
			if cinfo, ok := info.(*icg.ControlOpcode); ok {
				if cinfo.Opcode() == icg.Label {
					buffer.WriteString(strconv.Itoa(cinfo.GetLine()) + ": " + cinfo.String() + "\n")
				} else {
					buffer.WriteString("\t" + strconv.Itoa(cinfo.GetLine()) + ": " + cinfo.String() + "\n")
				}
			}

		}
		buffer.WriteString("\n")
	}

	fmt.Println(buffer.String())
}
func main() {
	var goFile string
	var fs *token.FileSet
	var f *ast.File
	var err error
	var conf types.Config
	var info *types.Info
	var strPoolGenerator *symbolTable.StringPoolGenerator
	var strPool *symbolTable.StringPool
	var symTble *symbolTable.BlockSymbolTable
	var litTable *symbolTable.LiteralTable
	var silTable *icg.SILTable
	var importLibList []string
	var chaincodeTable *symbolTable.BlockChaincodeSymbolTable

	if len(os.Args) >= 2 {
		for i, args := range os.Args {
			if i == 1 && args != "-h" {
				goFile = args
				fs = token.NewFileSet()

				f, err = parser.ParseFile(fs, goFile, nil, parser.AllErrors)
				if err != nil {
					log.Printf("could not parse %s: %v", goFile, err)
				}
				conf = types.Config{Importer: importer.ForCompiler(fs, "source", nil)}
				info = &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
				if _, err := conf.Check("", fs, []*ast.File{f}, info); err != nil {
					log.Fatal(err) // type error
				}

				strPoolGenerator = &symbolTable.StringPoolGenerator{}
				strPoolGenerator.Init(info)

				strPool, symTble, litTable, importLibList,chaincodeTable = strPoolGenerator.Gen(f)
				silTable = icg.CodeGen(f, fs, info, strPool, symTble, litTable, importLibList,chaincodeTable)

				continue
			}

			switch args {
			// help option
			case "-h":
				fmt.Println("Usage : ICG AnalysisTarget.go")
				os.Exit(0)
				// print option
			}

		}
	} else {
		log.Fatal(fmt.Errorf("Please run it with reference to the -h option "))
	}

	SAFFileGen(silTable, litTable, goFile)
	//PrintSIL(silTable)

}
