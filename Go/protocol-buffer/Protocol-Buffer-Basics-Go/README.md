protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto

$SRC_DIR - folder where addressbook.proto is located
$DST_DIR - go path + src - / Users/dpranjic/go/src

https://developers.google.com/protocol-buffers/docs/gotutorial