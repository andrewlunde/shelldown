# shelldown!

_Generate shell scripts from embedded comments and code blocks within markdown_

---


### Installation
```
go get github.com/rigelrozanski/shelldown
make all
```

### Usage

The premise of `shelldown` is to be able to generate a shell script
to test any bash scripts codeblocks you may have within your a markdown file. 
This is accomplished by loading a template shell script which is contained within
markdown comments at the top of you markdown file:

```
<!---
<All your regular shell commands go here>
-->
```

Within the body of your markdown file you may have codeblocks you wish to test.
To designate a codeblock for testing simply add `shelldown[index]` to the first line
designating a codeblock. Here `index` refers to any integer:

``` markdown 
    ```  shelldown[0]
    echo "my codeblock"
    echo "what a cool echo"
    ```
``` 

To reference from this codeblock within shell script template use
`#shelldown[index][line]` where the `index` is the same integer as above, and
`line` is the line number within the code block to be referenced. Here is a
very simple completed example:

```
<!---
#!/bin/bash

#shelldown[0][1]
-->

# My cool markdown file

here is some awesome code btw:

    ``` shelldown[0]
    echo "everything is okay"
    echo "everything is right"

    ```
```

To generate the shell script use the command:

```
shelldown path/to/markdown.md
```

if the bash script generated from the above example was run, it should output
`everythin is okay` as it was referencing the second line the first codeblock

### Full Examples

For full demonstration checkout the examples within `examples/`.  Here
`example_simple.md` is a simple demonstration of the principles of shelldown,
and `example_shunit2.md` is a more complex example which takes advantage of the
[shunit2](https://github.com/kward/shunit2) tool (highly recommended for bash
testing)

These examples can automatically be run by using calling `make test_example`. 
Check out the code within the Makefile for full detail. However the calls that 
are made effectively perform the following commands:

```
shelldown examples
bash example_simple.sh
bash example_shunit2.sh
```

Note that you must have [shunit2](https://github.com/kward/shunit2) installed.
Additionally, it's worth mentioning, that within the `.gitignore` file in this 
repo we prevent git from tracking the scripts which `shelldown` generates. 
This is important for ensuring that no antiquated data remains after shelldown 
markdown file has been updated  
 
### Contributing

1. Fork it
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create new Pull Request

### License

shelldown is released under the Apache 2.0 license.
