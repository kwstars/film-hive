name: Go

on:
  push:
    branches:
      - main
  pull_request:
    branches:
      - main
  workflow_dispatch:

jobs:
  build:
    strategy:
      matrix:
        go: [ 1.19 ]
    name: build & test
    runs-on: ubuntu-22.04

    steps:
      - uses: actions/checkout@v3
      - name: Set up Go
        uses: actions/setup-go@v3.5.0
        with:
          go-version: ${{ matrix.go }}

      - name: Setup Environment
        run: |
          echo "GOPATH=$(go env GOPATH)" >> $GITHUB_ENV
          echo "$(go env GOPATH)/bin" >> $GITHUB_PATH

      - name: Module cache
        uses: actions/cache@v3
        with:
          path: |
            ~/.cache/go-build
            ~/go/pkg/mod
          key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}
          restore-keys: |
            ${{ runner.os }}-go

      - name: Test
        run: make test

