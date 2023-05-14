# how to run unit test in GO

## go unit testing

```bash
# run a unit test commonly
$ go test

# run a complete unit test (with verbose command flag) ==> by default
$ go test -v

# run a complete and specific unit test (with verbose command flag)
$ go test -v -run=TestYourFuncName

# run all complete unit test "from top module" folder with all their module or package
$ go test -v ./...

# run specified of Sub test on a module or package ==> better for Sub Test testing
$ go test -run=TestYourSubTestFuncName/SubTestName

# runs all the sub tests in the module or package with the named sub test `/SubTestName`
$ go test -run=/SubTestName

```

> **disclaimer**
> We must entire the directory if will be testing for specified package, expect run all testing in top of folder command sh

## go benchmark testing

```bash
# run all complete benchmark in specified module or package
$ go test -v -bench=.

# run all complete benchmark without unit test in specified module or package
$ go test -v -run=NotMatchUnitTest -bench=.

# run all complete Benchmark unit test "from top module" folder with all their module or package without unit test
$ go test -v -run=NotMatchUnitTest -bench=. ./...

# run all complete Benchmark unit test and unit test "from top module" folder with all their module or package
$ go test -v -bench=. ./...

# run a specified complete benchmark func without unit test in specified module or package
$ go test -v -run=NotMatchUnitTest -bench=BenchmarkUnitTestName

```
