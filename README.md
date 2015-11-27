version: 0.1

    Usage of gofn:

    Gofn executes the passed -fn code. Therefore a go-binary will be
    created and stored. With a chance of 1 to 30 a cleaning-routine
    deletes the oldest gofn-binaries, if there are more than 200.
    Set the GOFN environment variable to change the gofn binary directory.
    Set GOFN_MAX to specify how many binaries should be kept.

      -debug
            Print generated code to stdout and exit
      -fn string
            Mandatory function-body without function declaration.
            Body is inserted into the function: fn(arr []string)
      -imports string
            Space-separated list of imports
      -list
            List all existing gofn binaries
      -name string
            Mandatory function name

Bash example, passing two arguments to `hello_bash`:

    #!/usr/bin/env bash

    cmd='
    s := strings.Join(arr, " ")
    fmt.Println(s)
    '

    gofn -name="hello_bash" -imports="fmt strings" -fn="$cmd" "Hello" "World"

fish shell example:

    #!/usr/bin/env fish

    set cmd '
    rand.Seed(time.Now().UnixNano())
    rollTheDice := func() int {
         return rand.Intn(6) +1
    }
    for i:=0; i<20; i++ {
        fmt.Println(rollTheDice())
    }
    '

    gofn -name="hello_fish" -imports="fmt math/rand time" -fn="$cmd"

### How it works:

Compiled gofn functions contain a hash and the compilation date. When gofn is executed, it looks for an existing gofn binary with the same name and hash. If there is one, the binary is executed, otherwise a new gofn binary is compiled.

### Tested on:

* Ubuntu 15.10

### TODO

* Write more tests
* Get gofn working on Windows & Mac
