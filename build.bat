@echo off
IF NOT EXIST .\build mkdir build

pushd build
go build ..
popd
