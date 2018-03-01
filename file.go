package ecoff

import (
	"encoding/binary"
	"fmt"
	"io"
)

type FileHeader struct {
	Magic          uint16
	NumSections    uint16
	TimeDate       int32
	SymbolsPointer uint32
	NumSymbols     int32
	OptionalHeader uint16
	Flags          uint16
}

func (f FileHeader) String() string {
	return fmt.Sprintf("{Magic: 0x%X, NumSections: %d, TimeDate: 0x%X, SymbolsPointer: 0x%X, NumSymbols: %d, OptionalHeader: 0x%X, Flags: 0x%X}", f.Magic, f.NumSections, f.TimeDate, f.SymbolsPointer, f.NumSymbols, f.OptionalHeader, f.Flags)
}

type ObjectHeader struct {
	Magic     int16
	Vstamp    int16
	TextSize  int32
	DataSize  int32
	BssSize   int32
	Entry     uint32
	TextStart uint32
	DataStart uint32
	BssStart  uint32
	GprMask   uint32
	CprMask   [4]uint32
	GpValue   uint32
}

func (h ObjectHeader) String() string {
	return fmt.Sprintf("{Magic: 0x%X, Vstamp: 0x%X, TextSize: %d, DataSize: %d, BssSize: %d, Entry: 0x%X, TextStart: 0x%X, DataStart: 0x%X, BssStart: 0x%X, GprMask: 0x%X, CprMask: 0x%X, GpValue: 0x%X}", h.Magic, h.Vstamp, h.TextSize, h.DataSize, h.BssSize, h.Entry, h.TextStart, h.DataStart, h.BssStart, h.GprMask, h.CprMask, h.GpValue)
}

type SectionHeader struct {
	Name               [8]uint8
	PhysicalAddress    uint32
	VirtualAddress     uint32
	Size               int32
	SectionPointer     uint32
	RelocationsPointer uint32
	LnnoPtr            int32
	NumRelocations     uint16
	NumLnno            uint16
	Flags              int32
}

func (h SectionHeader) String() string {
	return fmt.Sprintf("{Name: %s, PhysicalAddress: 0x%X, VirtualAddress: 0x%X, Size: %d, SectionPointer: 0x%X, RelocationsPointer: 0x%X, LnnoPtr: 0x%X, NumRelocations: %d, NumLnno: %d, Flags: 0x%X}", string(h.Name[:]), h.PhysicalAddress, h.VirtualAddress, h.Size, h.SectionPointer, h.RelocationsPointer, h.LnnoPtr, h.NumRelocations, h.NumLnno, h.Flags)
}

type Header struct {
	FileHeader     FileHeader
	ObjectHeader   ObjectHeader
	SectionHeaders []SectionHeader
}

// Parse all headers out from a ECOFF file.
func ParseHeader(r io.Reader, bo binary.ByteOrder) (Header, error) {
	out := Header{}
	err := binary.Read(r, bo, &out.FileHeader)
	if err != nil {
		return out, err
	}

	err = binary.Read(r, bo, &out.ObjectHeader)
	if err != nil {
		return out, err
	}

	var i uint16
	for i = 0; i < out.FileHeader.NumSections; i++ {
		out.SectionHeaders = append(out.SectionHeaders, SectionHeader{})
		err = binary.Read(r, bo, &out.SectionHeaders[i])
		if err != nil {
			return out, err
		}
	}
	return out, nil
}
