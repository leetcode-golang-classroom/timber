package timber_test

import (
	"errors"
	"log"
	"testing"

	"github.com/leetcode-golang-classroom/timber/timber"
	timber_example "github.com/leetcode-golang-classroom/timber/timber/example"
	timber_interface "github.com/leetcode-golang-classroom/timber/timber/interface"
)

func init() {
	log.SetFlags(0)
}

func TestSuccess_complex(t *testing.T) {
	// given
	tree := timber_example.Branch{
		Label: "branch0",
		Items: []timber_interface.NodeTree{
			timber_example.Branch{
				Label: "branch2",
				Items: []timber_interface.NodeTree{
					timber_example.Leaf{Label: "leaf2"},
					timber_example.Leaf{Label: "leaf3"},
				},
			},
			timber_example.Leaf{Label: "leaf"},
		},
	}
	printer := timber.NewDefaultPrinter()
	// when
	result, err := printer.Print(tree)
	// then
	Equal(t, nil, err)
	castedResult := result.([]string)
	Equal(t, 5, len(castedResult))
	Equal(t, "└ branch0", castedResult[0])
	Equal(t, " ├ branch2", castedResult[1])
	Equal(t, " │ ├ leaf2", castedResult[2])
	Equal(t, " │ └ leaf3", castedResult[3])
	Equal(t, " └ leaf", castedResult[4])
}

// TODO: potential bug
func TestSuccess_emptyTree(t *testing.T) {
	// given
	tree := timber_example.Branch{}
	printer := timber.NewDefaultPrinter()
	// when
	result, err := printer.Print(tree)
	// then
	Equal(t, nil, err)
	castedResult := result.([]string)
	Equal(t, 1, len(castedResult))
}
func TestFailure_printFnFailure(t *testing.T) {
	// given
	tree := timber_example.Branch{}
	faultyPrintFn := func(display string) (interface{}, error) {
		log.Print(display)
		return display, errors.New("bam")
	}
	printer := timber.NewCustomPrinter(faultyPrintFn, make([]string, 0),
		timber.DefaultResultFn)

	// when
	result, err := printer.Print(tree)
	// then
	Equal(t, "nope", err.Error())
	Equal(t, nil, result)
}

func TestFailure_resultFnFailure(t *testing.T) {
	// given
	tree := timber_example.Branch{}
	faultyResultFn := func(previousResult, currentItem interface{}) (interface{}, error) {
		return nil, errors.New("boom")
	}
	printer := timber.NewCustomPrinter(timber.DefaultPrinterFn, make([]string, 0),
		faultyResultFn)

	// when
	result, err := printer.Print(tree)
	// then
	Equal(t, "boom", err.Error())
	Equal(t, nil, result)
}

func Equal(t *testing.T, wanted interface{},
	received interface{}) {
	if wanted != received {
		t.Errorf("\nwanted: \t%v\nreceived: \t%v", wanted, received)
	}
}
