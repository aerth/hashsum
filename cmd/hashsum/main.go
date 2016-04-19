// Hashsum file integrity

/*

The MIT License (MIT)

Copyright (c) 2016 aerth

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/aerth/hashsum"
)

func main() {
	var logo = `

   __            __
  / /  ___ ____ / / _____ ____ _
 / _ / _  (_-</ _ \(_-/ // /  ' \
/_//_\_,_/___/_//_/___\_,_/_/_/_/

 [ https://github.com/aerth/hashsum ]

		`

	if len(os.Args) == 2 {

		if os.Args[1] == "-h" || os.Args[1] == "--help" || os.Args[1] == "help" || strings.HasPrefix(os.Args[1], "--") {

			fmt.Println(logo)
			fmt.Printf("Usage: " + os.Args[0] + " [filename(s)] \n")

			os.Exit(1)

		}

		filename := os.Args[1]
		md5sum := hashsum.Md5sum(readFile(filename))
		sha1sum := hashsum.Sha1sum(readFile(filename))
		sha256sum := hashsum.Sha256sum(readFile(filename))
		sha384sum := hashsum.Sha384sum(readFile(filename))
		sha512sum := hashsum.Sha512sum(readFile(filename))
		fmt.Printf("MD5:    %s  %s\n", md5sum, filename)
		fmt.Printf("SHA1:   %s  %s\n", sha1sum, filename)
		fmt.Printf("SHA256: %s  %s\n", sha256sum, filename)
		fmt.Printf("SHA384: %s  %s\n", sha384sum, filename)
		fmt.Printf("SHA512: %s  %s\n", sha512sum, filename)
		os.Exit(1)
	} else if len(os.Args) > 2 { // Multiple files
		fmt.Println()
		for file := range os.Args {
			if file == 0 {
				continue
			}
			filename := os.Args[file]
			md5sum := hashsum.Md5sum(readFile(filename))
			sha1sum := hashsum.Sha1sum(readFile(filename))
			sha256sum := hashsum.Sha256sum(readFile(filename))
			sha384sum := hashsum.Sha384sum(readFile(filename))
			sha512sum := hashsum.Sha512sum(readFile(filename))
			fmt.Println()
			fmt.Printf("MD5:    %s  %s\n", md5sum, filename)
			fmt.Printf("SHA1:   %s  %s\n", sha1sum, filename)
			fmt.Printf("SHA256: %s  %s\n", sha256sum, filename)
			fmt.Printf("SHA384: %s  %s\n", sha384sum, filename)
			fmt.Printf("SHA512: %s  %s\n", sha512sum, filename)

			fmt.Println()
		}

	} else {

		fmt.Println(logo)
		fmt.Printf("Usage: " + os.Args[0] + " [filename(s)] \n")

		os.Exit(1)

	}
}

func readFile(filename string) []byte {
	var rawBytes []byte
	var err error
	if filename != "" {
		rawBytes, err = ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	return rawBytes
}
