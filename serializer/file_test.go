package serializer_test

import (
	go_grpc "github.com/anujaneja/go_grpc/pb"
	"github.com/anujaneja/go_grpc/sample"
	"github.com/anujaneja/go_grpc/serializer"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()
	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"
	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &go_grpc.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = serializer.WriteProtobufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)
}
