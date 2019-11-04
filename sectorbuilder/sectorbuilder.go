package sectorbuilder

/*
This module was ported from https://github.com/filecoin-project/lotus/tree/master/lib/sectorbuilder
*/
import (
	"unsafe"

	sectorbuilder "github.com/filecoin-project/go-sectorbuilder"
)

type SectorSealingStatus = sectorbuilder.SectorSealingStatus

type StagedSectorMetadata = sectorbuilder.StagedSectorMetadata

type SortedSectorInfo = sectorbuilder.SortedSectorInfo

type SectorInfo = sectorbuilder.SectorInfo

type SealTicket = sectorbuilder.SealTicket

type SealedSectorMetadata = sectorbuilder.SealedSectorMetadata

const CommLen = sectorbuilder.CommitmentBytesLen

type Sectorbuilder struct {
	handle unsafe.Pointer
}
