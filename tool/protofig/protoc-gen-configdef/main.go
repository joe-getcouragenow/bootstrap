package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"google.golang.org/protobuf/proto"
	plugin "google.golang.org/protobuf/types/pluginpb"
)

func init() {
	log.SetOutput(os.Stderr)
}

func main() {
	req := readGenRequest(os.Stdin)

	params, err := parseCommandLineParams(req.GetParameter())
	if err != nil {
		Error(err, "could not parse parameters")
	}

	g := newGenerator(params)

	resp := g.Generate(req)

	writeResponse(os.Stdout, resp)
}

func readGenRequest(r io.Reader) *plugin.CodeGeneratorRequest {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		Error(err, "reading input")
	}

	req := new(plugin.CodeGeneratorRequest)
	if err = proto.Unmarshal(data, req); err != nil {
		Error(err, "parsing input proto")
	}

	if len(req.FileToGenerate) == 0 {
		Fail("no files to generate")
	}

	return req
}

func writeResponse(w io.Writer, resp *plugin.CodeGeneratorResponse) {
	data, err := proto.Marshal(resp)
	if err != nil {
		Error(err, "marshaling response")
	}
	_, err = w.Write(data)
	if err != nil {
		Error(err, "writing response")
	}
}

// Error print error and exit
func Error(err error, msgs ...string) {
	s := strings.Join(msgs, " ") + ":" + err.Error()
	log.Print("error:", s)
	os.Exit(1)
}

// Fail print error and exit
func Fail(msgs ...string) {
	s := strings.Join(msgs, " ")
	log.Print("error:", s)
	os.Exit(1)
}
