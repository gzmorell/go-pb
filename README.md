# go-pb

## Tutorial: <https://ednsquare.com/story/working-with-protocol-buffers-in-golang------8g4uuR>

Modified as an example of proto2

## Proto Code Generation

Run the next command to generate the proto code:

`> protoc -I . --go_out=.  .\addressbook.proto`

This will generate the tutorialpb/addressbookpb/addressbook.pb.go file.

Where tutorial is the package name, and addressbook the file name.

## Running the app

`> go run .\main.go`
