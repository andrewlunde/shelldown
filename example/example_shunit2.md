<!--- shelldown script template, see github.com/andrewlunde/shelldown
#!/bin/bash

test00EchoStd() {
  echo "Running standard out test"
  RES=$(#shelldown[1][0])
  assertEquals "stdout not woohoo!" "woohoo!" "$RES" 
}

test01EchoErr() {
  echo "Running standard err test"
  RES=$((#shelldown[2][0]) 2>&1 > /dev/null)
  assertEquals "stderr not whoops!" "whoops!" "$RES"  
}

# load and run these tests with shunit2!
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )" #get this files directory
. $DIR/shunit2
-->

# Example Markdown File!

Hey welcome to the example markdown file for shelldown!
These following codeblocks will be executed and tested
but can also serve as tutorial information for anyone viewing 
this markdown file.

### Some Simple Scripts

The following terminal command should output to `stdout`:

``` shelldown[1]
echo "woohoo!"
```

This next line of code should reroute the output to `stderr`:

``` shelldown[2]
echo "whoops!" 1>&2
```

Enjoy!
