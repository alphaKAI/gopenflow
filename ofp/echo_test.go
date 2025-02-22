package ofp

import (
	"testing"

	"github.com/alphakai/gopenflow/internal/encodingtest"
)

func TestEchoRequest(t *testing.T) {
	data := []byte{0x00, 0x01, 0x02, 0x03}
	tests := []encodingtest.MU{
		{ReadWriter: &EchoRequest{Data: data}, Bytes: data},
	}

	encodingtest.RunMU(t, tests)
}

func TestEchoReply(t *testing.T) {
	data := []byte{0x00, 0x01, 0x02, 0x03}
	tests := []encodingtest.MU{
		{ReadWriter: &EchoReply{Data: data}, Bytes: data},
	}

	encodingtest.RunMU(t, tests)
}
