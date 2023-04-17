# timber

## About

Printing an ASCII tree with Go(education purpose).

Thanks to https://github.com/Tufin/asciitree for the idea

Example of ASCII tree:
```shell
|- root1
| |- sibling1
| |- sibling2
|    |- sibling1
|    |- sibling2
|- root2
   |- sibling1
   |- sibling2
```
## Usage

```go
func main() {
	tree := timber_example.Branch{
		Label: "branch",
		Items: []timber_interface.NodeTree{
			timber_example.Leaf{
				Label: "leaf",
			},
			timber_example.Branch{
				Label: "branch",
				Items: []timber_interface.NodeTree{
					timber_example.Leaf{
						Label: "leaf",
					},
					timber_example.Leaf{
						Label: "leaf",
					},
					timber_example.Leaf{
						Label: "leaf",
					},
				},
			},
			timber_example.Leaf{
				Label: "leaf",
			},
		},
	}
	printer := timber.NewDefaultPrinter()
	printer.Print(tree)
}
```

will output

```shell
INFO[0000] └ branch                                     
INFO[0000]  ├ leaf                                      
INFO[0000]  ├ branch                                    
INFO[0000]  │ ├ leaf                                    
INFO[0000]  │ ├ leaf                                    
INFO[0000]  │ └ leaf                                    
INFO[0000]  └ leaf   
```

If a custom printer function or result aggergater function is needed, use:

```go
func NewCustomPrinter(printerFn PrinterFn, result interface{}, resultFn ResultFn) *customPrinter
```
## Installation

### Prequisite

```shell
$ go version
go version go1.20 linux/amd64
```
### Command

`go get github.com/leetcode-golang-classroom/timber`

## Contributing

Please free to fork, create a branch and open a pull request.

## License

This is under MIT License.
## Contact

Please contact:

[LINKEDIN](https://www.linkedin.com/in/json-liang-nctu/)

