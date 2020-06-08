cd "$GOPATH"/src || exit
gofiles=("aggregate/events" "aggregate/project" "aggregate/task" "command/command" "viewModel/viewModel")

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

cd "$GOPATH"/src/github.com/gimalay/octogo/todoapp/core/command || exit
protoc -I=. -I=$GOPATH/src  --swift_out=../../ios/todoapp/src/Models command.proto
protoc --java_out=../../android/todoapp/app/src/main/java command.proto

cd "$GOPATH"/src/github.com/gimalay/octogo/todoapp/core/viewModel || exit
protoc -I=. -I=$GOPATH/src  --swift_out=../../ios/todoapp/src/Models viewModel.proto
protoc --java_out=../../android/todoapp/app/src/main/java viewModel.proto

