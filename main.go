/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

//https://towardsdatascience.com/how-to-create-a-cli-in-golang-with-cobra-d729641c7177

//go install github.com/spf13/cobra-cli@latest
//cd netflix
//Users/borakasmer/go/bin/cobra-cli init
//Users/borakasmer/go/bin/cobra-cli add add => We added "add" function
// Create "addInt()" in add.go
//*go install exchange-cli
//go env -w GO111MODULE=on NOT NEEEDD

//**** If there is no Go => Install Go https://go.dev/dl/
//go version (Check Version => 1.18 :) )
//go env (Check GOPATH)

//*** change go.mod "module netfix" =>"module github.com/borakasmer/netflix"
//*** change main.go "import netflix/cmd" => "import "github.com/borakasmer/netflix/cmd"
//*** change get.go "import netflix/parser" => "import github.com/borakasmer/netflix/parser"
//*** change root.go "import netflix/core" => "import github.com/borakasmer/netflix/core
////*** change root.go "import netflix/parser" => "import github.com/borakasmer/netflix/parser"

//**** => go install github.com/borakasmer/netflix@latest
//**** FOR GO.1.18 => MOVE "exchange-cli" and "cobra-cli" from "/Users/borakasmer/go/bin" to => "/usr/local/go/bin"
package main

import "github.com/borakasmer/netflix/cmd"

func main() {
	cmd.Execute()
}
