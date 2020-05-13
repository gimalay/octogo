cd "$GOPATH"/src || exit
gofiles=("aggregate/events" "aggregate/project" "aggregate/task" "command/command")

for f in "${gofiles[@]}"
do
protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf/protobuf --gofast_out=\
Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:. \
github.com/gimalay/octogo/todoapp/core/$f.proto
done

