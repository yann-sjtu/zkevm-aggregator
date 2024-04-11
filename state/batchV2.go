package state

// BuildChangeL2Block returns a changeL2Block tx to use in the BatchL2Data
func (p *State) BuildChangeL2Block(deltaTimestamp uint32, l1InfoTreeIndex uint32) []byte {
	l2block := ChangeL2BlockHeader{
		DeltaTimestamp:  deltaTimestamp,
		IndexL1InfoTree: l1InfoTreeIndex,
	}
	var data []byte
	data = l2block.Encode(data)
	return data
}
