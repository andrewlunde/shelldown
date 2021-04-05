<!--- shelldown script template, see github.com/andrewlunde/shelldown
#!/bin/bash

echo "Testing out the first script:"
echo "The above should have output: woohoo!"

echo "Testing out the second script, route stderr to stdout:"
$(#shelldown[2][0]) 2>&1 > /dev/null
echo "The above should have output: whoops!"

echo "Printing an entire codeblock"
echo #shelldown[1][-1]
#shelldown[1][-1]
echo "The above should have output: woohoo! then hoohoo!"

-->

# Example Markdown File!

Hey welcome to the example markdown file for shelldown!
These following codeblocks will be executed and tested
but can also serve as tutorial information for anyone viewing 
this markdown file.

### Some Simple Scripts

The following terminal commands should output to `stdout`:

``` shelldown[1]
echo "woohoo!"

echo "hoohoo!"
```

This next line of code should reroute the output to `stderr`:

``` shelldown[2]
echo "whoops!" 1>&2
```

Enjoy!
