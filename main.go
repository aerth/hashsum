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
  "crypto/md5"
  "crypto/sha1"
  "crypto/sha256"
)
func readStdin() []byte {
  var rawBytes []byte
  return rawBytes
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

func Md5sum(filename string) string {
  var hashString string
  rawBytes := readFile(filename)
  md5sum := md5.Sum(rawBytes)
  hashString = fmt.Sprintf("%x", md5sum)
  return hashString
}

func Sha1sum(filename string) string {
  var hashString string
  rawBytes := readFile(filename)
  sha1sum := sha1.Sum(rawBytes)
  hashString = fmt.Sprintf("%x", sha1sum)
  return hashString
}

func Sha256sum(filename string) string {
  var hashString string
  rawBytes := readFile(filename)
  sha256sum := sha256.Sum256(rawBytes)
  hashString = fmt.Sprintf("%x", sha256sum)
  return hashString
}

func main() {
  if len(os.Args) > 1 {
    filename := os.Args[1]
    md5sum := Md5sum(filename)
    sha1sum := Sha1sum(filename)
    sha256sum := Sha256sum(filename)
    fmt.Printf("MD5:    %s  %s\n", md5sum, filename)
    fmt.Printf("SHA1:   %s  %s\n", sha1sum, filename)
    fmt.Printf("SHA256: %s  %s\n", sha256sum, filename)
    os.Exit(1)
}else

{
fmt.Printf("Usage: " + os.Args[0] + " [filename] \n")
os.Exit(1)

}
}
