package timber

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	timber_interface "github.com/leetcode-golang-classroom/timber/timber/interface"
)

type PrinterFn func(display string) (interface{}, error)
type ResultFn func(previousResult, currentItem interface{}) (interface{}, error)

var DefaultPrinterFn = func(display string) (interface{}, error) {
	log.Print(display)
	return display, nil
}

var DefaultResultFn = func(previousResult, currentItem interface{}) (interface{}, error) {
	return append(previousResult.([]string), currentItem.(string)), nil
}

type customPrinter struct {
	printerFn PrinterFn
	result    interface{}
	resultFn  ResultFn
}

func NewCustomPrinter(printerFn PrinterFn, result interface{}, resultFn ResultFn) *customPrinter {
	return &customPrinter{printerFn: printerFn, result: result, resultFn: resultFn}
}

func NewDefaultPrinter() *customPrinter {
	return &customPrinter{
		printerFn: DefaultPrinterFn,
		result:    make([]string, 0),
		resultFn:  DefaultResultFn,
	}
}
func (printer *customPrinter) Print(tree timber_interface.NodeTree) (interface{}, error) {
	if err := printer.doPrint(tree, node{prefix: "", isLast: true}); err != nil {
		return nil, err
	}
	return printer.result, nil
}

func (printer *customPrinter) doPrint(tree timber_interface.NodeTree, node node) error {
	display := fmt.Sprint(node.display(), tree.Display())
	displayResult, errPrint := printer.printerFn(display)
	if errPrint != nil {
		return errPrint
	}
	updateResult, errResult := printer.resultFn(printer.result, displayResult)
	if errResult != nil {
		return errResult
	}
	printer.result = updateResult
	length := len(tree.Components())
	for idx, child := range tree.Components() {
		isLast := idx == length-1
		printer.doPrint(child, node.createChildNode(isLast))
	}
	return nil
}

// node
type node struct {
	prefix string
	isLast bool
}

func (n node) display() string {
	if n.isLast {
		return n.prefix + "\u2514 "
	}
	return n.prefix + "\u251c "
}

func (n node) createChildNode(isLast bool) node {
	return node{
		prefix: n.prefix + n.nextPrefix(),
		isLast: isLast,
	}
}

func (n node) nextPrefix() string {
	if n.isLast {
		return " "
	}
	return "\u2502 "
}
