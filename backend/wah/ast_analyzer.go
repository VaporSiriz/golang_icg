package wah

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	icg "go_icg/icg"
)

type ASTAnalyzer struct {
	analysisFile       string
	FuncRetTable       map[string]int // functionName , error 위치  error리턴이 없는경우 -1
	fs                 *token.FileSet
	ueCompleteFuncList []string
	analysisCount      int
	error_str          string
}

func (analyzer ASTAnalyzer) GetErrorString() string {
	return analyzer.error_str
}

func (analyzer *ASTAnalyzer) Init(analysisFile string, fset *token.FileSet) {
	analyzer.analysisFile = analysisFile
	analyzer.FuncRetTable = make(map[string]int)
	analyzer.fs = fset
	analyzer.analysisCount = 0
	analyzer.error_str = ""
}

func (analyzer *ASTAnalyzer) addErrorStr(ccw CCW, linenum int, add_str ...string) {
	analyzer.error_str += fmt.Sprintf("\t CCW-%03d : %s\n", ccw, ccw.String())
	for _, str := range add_str {
		analyzer.error_str += str
	}
	analyzer.error_str += fmt.Sprintf("\t %s : %d\n\n", analyzer.analysisFile, linenum)
}

func dealFirstDetect(analyzer *ASTAnalyzer) {
	if isFirstDetect {
		analyzer.error_str += fmt.Sprintf("chaincode weakness detected:\n")
		fmt.Printf("chaincode weakness detected:\n")
		isFirstDetect = false
	}
}

//MapStructureIteration...
func (analyzer *ASTAnalyzer) MSIAnalysis(node ast.Node, info *types.Info) {
	var ccw CCW = MAP_STRUCTURE_ITER
	var position token.Position

	if node != nil {
		position = analyzer.fs.Position(node.Pos())
	}

	if rangeFor, ok := node.(*ast.RangeStmt); ok {
		if tv, ok := info.Types[rangeFor.X]; ok {
			_, isMap := tv.Type.(*types.Map)
			if isMap {
				linenum := position.Line
				dealFirstDetect(analyzer)
				analyzer.error_str += fmt.Sprintf("\t CCW-%03d : %s\n", ccw, ccw.String())
				analyzer.error_str += fmt.Sprintf("\t not use a map type \"%s\" in loop range\n", rangeFor.X)
				analyzer.error_str += fmt.Sprintf("\t %s : %d\n\n", analyzer.analysisFile, linenum)
				fmt.Printf("\t CCW-%03d : %s\n", ccw, ccw.String())
				fmt.Printf("\t not use a map type \"%s\" in loop range\n", rangeFor.X)
				fmt.Printf("\t %s : %d\n\n", analyzer.analysisFile, linenum)
				analyzer.analysisCount++
			}
		}
	}
}

func (analyzer *ASTAnalyzer) UsedGoroutineAnalysis(node ast.Node, info *types.Info) {
	var ccw CCW = USED_GOROUTINE
	var position token.Position
	if node != nil {
		position = analyzer.fs.Position(node.Pos())
	}
	if goStmt, ok := node.(*ast.GoStmt); ok {
		if tv, ok := info.Types[goStmt.Call.Fun]; ok {
			linenum := position.Line
			dealFirstDetect(analyzer)
			analyzer.addErrorStr(ccw, linenum, fmt.Sprintf("\t not use go routine \"go %v\"\n", tv.Type.Underlying()))
			fmt.Printf("\t CCW-%03d : %s\n", ccw, ccw.String())
			fmt.Printf("\t %s : %d\n\n", analyzer.analysisFile, linenum)
			analyzer.analysisCount++
		}
	}
}

func (analyzer *ASTAnalyzer) UnhandledErrorsAnalysis(node ast.Node, info *types.Info) int {
	errLocation := -1
	var ccw CCW = UNHANDLED_ERROR
	var position token.Position

	if node != nil {
		position = analyzer.fs.Position(node.Pos())
	}
	if assign, ok := node.(*ast.AssignStmt); ok {
		for i, rhs := range assign.Rhs {
			funcName := ""
			if call, ok := rhs.(*ast.CallExpr); ok {
				funcName = icg.NodeString(analyzer.fs, call.Fun)
				if results, ok := info.Types[rhs].Type.(*types.Tuple); ok {
					for j := 0; j < results.Len(); j++ {
						res := results.At(j)
						resType := res.Type()
						if resType.String() == "error" {
							errLocation = i + j
						}
					}
				} else if res, ok := info.Types[rhs].Type.(*types.Named); ok {
					if res.String() == "error" {
						errLocation = i
					}
				}

				if errLocation != -1 {
					if ident, ok := assign.Lhs[errLocation].(*ast.Ident); ok {
						if ident.Name == "_" {
							linenum := position.Line
							dealFirstDetect(analyzer)
							analyzer.addErrorStr(ccw, linenum)
							fmt.Printf("\t CCW-%03d : %s\n", ccw, ccw.String())
							fmt.Printf("\t %s : %d\n\n", analyzer.analysisFile, linenum)
							analyzer.analysisCount++
							//fmt.Printf("\t   The %d return type of rhs( %s ) is error, but is not assigned to the %d lhs ( _ ).\n\n", errLocation, funcName, errLocation)
							errLocation = -1
						}
					}
				}
				analyzer.FuncRetTable[funcName] = errLocation
			}
		}

	}

	return errLocation
}
func (analyzer *ASTAnalyzer) PhantomReadAnalysis(node ast.Node, info *types.Info) {
	var ccw CCW = PHANTOM_READS
	var position token.Position

	if node != nil {
		position = analyzer.fs.Position(node.Pos())
	}
	if assign, ok := node.(*ast.AssignStmt); ok {
		for _, rhs := range assign.Rhs {
			funcName := ""
			if call, ok := rhs.(*ast.CallExpr); ok {
				funcName = icg.NodeString(analyzer.fs, call.Fun)

				if strings.Contains(funcName, "GetHistoryForKey") || strings.Contains(funcName, "GetQueryResult") {
					linenum := position.Line
					dealFirstDetect(analyzer)
					fmt.Printf("\t CCW-%03d : %s\n", ccw, ccw.String())

					fmt.Printf("\t %s : %d\n\n", analyzer.analysisFile, linenum)
					analyzer.analysisCount++
				}
			}
		}
	}
}

func (analyzer *ASTAnalyzer) RandomNumberGenerationAnalysis(node ast.Node, info *types.Info) {
	var ccw CCW = RANDOM_NUMBER_GENERATION
	var position token.Position

	if node != nil {
		position = analyzer.fs.Position(node.Pos())
	}

	if imprt, ok := node.(*ast.BasicLit); ok {
		if imprt.Value == "\"math/rand\"" {
			linenum := position.Line
			dealFirstDetect(analyzer)
			fmt.Printf("\t CCW-%03d : %s\n", ccw, ccw.String())
			fmt.Printf("\t %s : %d\n\n", analyzer.analysisFile, linenum)
			analyzer.addErrorStr(ccw, linenum)
			analyzer.analysisCount++
			hasRandomImport = true
		}
	} else if selexp, ok := node.(*ast.SelectorExpr); ok && hasRandomImport {
		is_rand := (selexp.X.(*ast.Ident).Name == "rand")
		if is_rand {
			linenum := position.Line
			dealFirstDetect(analyzer)
			fmt.Printf("\t CCW-%03d : %s\n", ccw, ccw.String())
			fmt.Printf("\t %s : %d\n\n", analyzer.analysisFile, linenum)
			analyzer.addErrorStr(ccw, linenum)
			analyzer.analysisCount++
		}
	}
}

func (analyzer *ASTAnalyzer) UseTimeMouleAnalysis(node ast.Node, info *types.Info) {
	var ccw CCW = USE_TIME_MODULE
	var position token.Position

	if node != nil {
		position = analyzer.fs.Position(node.Pos())
	}
	if imprt, ok := node.(*ast.BasicLit); ok {
		if imprt.Value == "\"time\"" {
			linenum := position.Line
			dealFirstDetect(analyzer)
			fmt.Printf("\t CCW-%03d : %s\n", ccw, ccw.String())

			fmt.Printf("\t %s : %d\n\n", analyzer.analysisFile, linenum)
			analyzer.analysisCount++
			hasRandomImport = true
		}
	} else if selexp, ok := node.(*ast.SelectorExpr); ok && hasRandomImport {
		is_rand := (selexp.X.(*ast.Ident).Name == "time")
		if is_rand {
			linenum := position.Line
			dealFirstDetect(analyzer)
			fmt.Printf("\t CCW-%03d : %s\n", ccw, ccw.String())

			fmt.Printf("\t %s : %d\n\n", analyzer.analysisFile, linenum)
			analyzer.analysisCount++
		}
	}
}

//Analyze ...
func (analyzer *ASTAnalyzer) Analysis(f *ast.File, info *types.Info) string {

	ast.Inspect(f, func(node ast.Node) bool {
		analyzer.MSIAnalysis(node, info)
		analyzer.UsedGoroutineAnalysis(node, info)
		analyzer.UnhandledErrorsAnalysis(node, info)
		analyzer.PhantomReadAnalysis(node, info)
		analyzer.RandomNumberGenerationAnalysis(node, info)
		analyzer.UseTimeMouleAnalysis(node, info)
		return true
	})

	analyzer.error_str += fmt.Sprintf("\t Total error count : %d\n", analyzer.analysisCount)
	return analyzer.error_str
}
