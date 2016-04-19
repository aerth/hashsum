// Package hashsum returns hashstrings for []rawBytes

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

package hashsum

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func Md5sum(rawBytes []byte) (hashString string) {
	md5sum := md5.Sum(rawBytes)
	hashString = fmt.Sprintf("%x", md5sum)
	return hashString
}

func Sha1sum(rawBytes []byte) (hashString string) {
	sha1sum := sha1.Sum(rawBytes)
	hashString = fmt.Sprintf("%x", sha1sum)
	return hashString
}

func Sha256sum(rawBytes []byte) (hashString string) {
	sha256sum := sha256.Sum256(rawBytes)
	hashString = fmt.Sprintf("%x", sha256sum)
	return hashString
}

// Return the SHA384 hash
func Sha384sum(rawBytes []byte) (hashString string) {
	sha384sum := sha512.Sum384(rawBytes)
	hashString = fmt.Sprintf("%x", sha384sum)
	return hashString
}

// Return the SHA512 hash
func Sha512sum(rawBytes []byte) (hashString string) {
	sha512sum := sha512.Sum512(rawBytes)
	hashString = fmt.Sprintf("%x", sha512sum)
	return hashString
}
