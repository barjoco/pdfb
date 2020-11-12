#!/bin/sh

#
#   Store base64 encoded font files in a Go string array
#   accessible from package main
#

gofile='package pdfb

var fontsBase64 = [][]string{\n'

for f in Inter/*; do
    b64=$(base64 "$f" | tr -d '\n')
    f=$(basename "$f")
    gofile+="\t[]string{\""${f%.*}"\", \"$b64\"},\n"
done

gofile+='}'

echo -e "$gofile" > base64_fonts.go
