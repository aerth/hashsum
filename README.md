# Hashsum

## Usage:

```shell

$ hashsum ~/Downloads/pocorgtfo08.pdf

MD5:    257fc8f01fa20e21f8bd5577639ff596  /home/user/Downloads/pocorgtfo08.pdf
SHA1:   1c6e3200a1005f2acd5a8ebf3462be605c83f3a1  /home/user/Downloads/pocorgtfo08.pdf
SHA256: 7a942c425f471f99d8cba8da117cc4a53cddb3551e4b16c8b9feae31b5654a33  /home/user/Downloads/pocorgtfo08.pdf

```

## Built for use in a cross-compile release-style shell script like this one:

```shell
#!/bin/sh

for i in $(ls);do hashsum $i; echo ""; done &>>/tmp/SHSUM.txt

```


## Todo:

maybe: Add stdin functionality ( ``` cat cat.gif | hashsum ``` )
maybe: multiple files (``` hashsum file1 file2 file3 ```)
