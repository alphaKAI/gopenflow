package ofp

import (
	"encoding/gob"
	"testing"

	"github.com/alphakai/gopenflow/internal/encodingtest"
)

func TestTableMod(t *testing.T) {
	tests := []encodingtest.MU{
		{ReadWriter: &TableMod{
			Table:  TableMax,
			Config: TableConfigDeprecatedMask,
		}, Bytes: []byte{
			0xfe,             // Table identifier.
			0x00, 0x00, 0x00, // 3-byte padding.
			0x00, 0x00, 0x00, 0x03, // Configuration.
		}},
	}

	encodingtest.RunMU(t, tests)
}

func TestTableStats(t *testing.T) {
	tests := []encodingtest.MU{
		{ReadWriter: &TableStats{
			Table:        TableMax,
			ActiveCount:  267,
			LookupCount:  132,
			MatchedCount: 54,
		}, Bytes: []byte{
			0xfe,             // Table identifier.
			0x00, 0x00, 0x00, // 3-byte padding.
			0x00, 0x00, 0x01, 0x0b, // Active count.
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x84, // Lookup count.
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x36, // Matched count.
		}},
	}

	encodingtest.RunMU(t, tests)
}

func TestTablePropInstructions(t *testing.T) {
	tests := []encodingtest.MU{
		{ReadWriter: &TablePropInstructions{Instructions: []InstructionType{
			InstructionTypeMeter,
			InstructionTypeGotoTable,
		}}, Bytes: []byte{
			0x00, 0x00, // Property type.
			0x00, 0x0c, // Property length.

			// Instructions.
			0x00, 0x06, // Instruction type.
			0x00, 0x04, // Instruction length.

			0x00, 0x01, // Instruction type.
			0x00, 0x04, // Instruction length.

			0x00, 0x00, 0x00, 0x00, // 4-byte padding.
		}},
	}

	gob.Register(InstructionMeter{})
	gob.Register(InstructionGotoTable{})

	encodingtest.RunMU(t, tests)
}

func TestTablePropNextTables(t *testing.T) {
	tests := []encodingtest.MU{
		{ReadWriter: &TablePropNextTables{
			NextTables: []Table{1, 2, 3},
		}, Bytes: []byte{
			0x00, 0x02, // Property type.
			0x00, 0x07, // Property length.

			// Next tables.
			0x01, 0x02, 0x03,

			// Alignment.
			0x00,
		}},
	}

	encodingtest.RunMU(t, tests)
}

func TestTablePropWriteActions(t *testing.T) {
	tests := []encodingtest.MU{
		{ReadWriter: &TablePropWriteActions{Actions: []ActionType{
			ActionTypeCopyTTLOut,
			ActionTypeCopyTTLIn,
		}}, Bytes: []byte{
			0x00, 0x04, // Property type.
			0x00, 0x0c, // Property length.

			// Actions.
			0x00, 0xb, // Action type.
			0x00, 0x04, // Action length.

			0x00, 0xc, // Action type.
			0x00, 0x04, // Action length.

			0x00, 0x00, 0x00, 0x00, // 4-byte padding.
		}},
	}

	encodingtest.RunMU(t, tests)
}

func TestTablePropApplyActions(t *testing.T) {
	tests := []encodingtest.MU{
		{ReadWriter: &TablePropApplyActions{
			Miss: true, Actions: []ActionType{ActionTypeGroup},
		}, Bytes: []byte{
			0x00, 0x07, // Property type.
			0x00, 0x08, // Property length.

			// Actions.
			0x00, 0x16, // Action type.
			0x00, 0x04, // Action length.
		}},
	}

	encodingtest.RunMU(t, tests)
}

var fields = []XM{
	{
		Class: XMClassOpenflowBasic,
		Type:  XMTypeUDPSrc,
	},
	{
		Class: XMClassOpenflowBasic,
		Type:  XMTypeInPort,
	},
}

var fieldsBytes = []byte{
	0x80, 0x00, // OpenFlow basic.
	0x1e, // Match field + Mask flag.
	0x00, // Match length.

	0x80, 0x00, // OpenFlow basic.
	0x00, // Match field + Mask flag.
	0x00, // Match length.

	0x00, 0x00, 0x00, 0x00, // 4-byte padding.
}

func TestTablePropMatch(t *testing.T) {
	tests := []encodingtest.MU{
		{ReadWriter: &TablePropMatch{Fields: fields},
			Bytes: append([]byte{
				0x00, 0x08, // Property type.
				0x00, 0x0c, // Property length.
			}, fieldsBytes...)},
	}

	encodingtest.RunMU(t, tests)
}

func TestTablePropWildcards(t *testing.T) {
	tests := []encodingtest.MU{
		{ReadWriter: &TablePropWildcards{Fields: fields}, Bytes: append([]byte{
			0x00, 0x0a, // Property type.
			0x00, 0x0c, // Property length.
		}, fieldsBytes...)},
	}

	encodingtest.RunMU(t, tests)
}

func TestTablePropWriteSetField(t *testing.T) {
	tests := []encodingtest.MU{
		{
			ReadWriter: &TablePropWriteSetField{Fields: fields},
			Bytes: append([]byte{
				0x00, 0x0c, // Property type.
				0x00, 0x0c, // Property length.
			}, fieldsBytes...)},
	}

	encodingtest.RunMU(t, tests)
}

func TestTablePropApplySetField(t *testing.T) {
	tests := []encodingtest.MU{
		{
			ReadWriter: &TablePropApplySetField{
				Miss:   true,
				Fields: fields,
			},
			Bytes: append([]byte{
				0x00, 0x0f, // Property type.
				0x00, 0x0c, // Property length.
			}, fieldsBytes...)},
	}

	encodingtest.RunMU(t, tests)
}

func TestTablePropExperimenter(t *testing.T) {
	tests := []encodingtest.MU{
		{ReadWriter: &TablePropExperimenter{
			Experimenter: 42,
			ExpType:      43,
			Data:         []byte{0x11, 0x22},
		}, Bytes: []byte{
			0xff, 0xfe, // Property type.
			0x00, 0x0e, // Property length.
			0x00, 0x00, 0x00, 0x2a, // Experimenter.
			0x00, 0x00, 0x00, 0x2b, // Experimenter type.
			0x11, 0x22, // Experimenter data.

			// Alignment.
			0x00, 0x00,
		}},
	}

	encodingtest.RunMU(t, tests)
}

func TestTableFeatures(t *testing.T) {
	name := make([]byte, maxTableNameLen)
	copy(name, []byte("table-1"))

	properties := []TableProp{
		&TablePropApplyActions{Actions: []ActionType{
			ActionTypeGroup,
		}},
		&TablePropInstructions{Instructions: []InstructionType{
			InstructionTypeMeter,
		}},
	}

	bytes := []byte{
		0x00, 0x50, // Length.
		0x03,                         // Table identifier.
		0x00, 0x00, 0x00, 0x00, 0x00, // 5-byte padding.
	}

	bytes = append(bytes, name...)
	bytes = append(bytes,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, // Metadata match.
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xfe, // Metadata write.
		0x00, 0x00, 0x00, 0x00, // Table configuration.
		0x00, 0x00, 0x02, 0x00, // Maximum entries.

		// Properties.
		0x00, 0x06, // Property type.
		0x00, 0x08, // Property length.
		// Actions.
		0x00, 0x16, // Action type.
		0x00, 0x04, // Action length.

		0x00, 0x00, // Property type.
		0x00, 0x08, // Property length.
		// Instructions.
		0x00, 0x06, // Instruction type.
		0x00, 0x04, // Instruction length.
	)

	tests := []encodingtest.MU{
		{ReadWriter: &TableFeatures{
			Table:         Table(3),
			Name:          string(name),
			MetadataMatch: 0xff,
			MetadataWrite: 0xfe,
			MaxEntries:    512,
			Properties:    properties,
		}, Bytes: bytes},
	}

	gob.Register(TablePropApplyActions{})
	gob.Register(TablePropInstructions{})

	encodingtest.RunMU(t, tests)
}
