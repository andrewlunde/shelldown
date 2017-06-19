<!--- shelldown script template, see github.com/rigelrozanski/shelldown
#!/bin/bash

echo "Testing out the first script:"
#shelldown[1][0]
echo "The above should have output: woohoo!"

echo "Testing out the second script, route stderr to stdout:"
$(#shelldown[2][0]) 2>&1 > /dev/null
echo "The above should have output: whoops!"
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
