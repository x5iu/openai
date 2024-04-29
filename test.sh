#!/bin/bash

go test -cover ./
go test -tags=legacy -cover ./