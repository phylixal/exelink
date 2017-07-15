#! /bin/sh

go build -ldflags "-s -w" 
upx --best -o git.exe exelink.exe
rm exelink.exe
