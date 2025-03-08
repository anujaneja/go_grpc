package serializer

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"os"
)

// WriteProtobufToBinaryFile write the file with proto message in binary format.
func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)

	if err != nil {
		return fmt.Errorf("unable to marshal proto message to binary: %w", err)
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("unable to write protobuf to binary file: %w", err)
	}
	return nil
}

// ReadProtobufFromBinaryFile read binary file and create protobuf message
func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("unable to read protobuf from binary file: %w", err)
	}
	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("unable to unmarshal protobuf from binary file: %w", err)
	}

	return nil
}

func WriteProtobufToJSONFile(message proto.Message, filename string) error {
	data, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf("unable to write protobuf to json file: %w", err)
	}
	err = os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("unable to write protobuf to json file: %w", err)
	}

	return nil
}
