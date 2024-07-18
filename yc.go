package yc

import (
	"bytes"
)

type RootBlock struct {
	Blocks []*Block
}

func NewRootBlock() *RootBlock {
	return &RootBlock{}
}

func (r *RootBlock) AddBlock(block *Block) {
	r.Blocks = append(r.Blocks, block)
}

func (b *RootBlock) YAML() string {
	buf := bytes.NewBuffer(nil)
	for _, block := range b.Blocks {
		block.YAML(buf, false)
	}

	return buf.String()
}

type Block struct {
	Name        string
	Blocks      []*Block
	Value       Elm
	ArrayValues [][]*Block
}

type Elm struct {
	Value any
}

func NewBlock(name string) *Block {
	return &Block{Name: name}
}

func (b *Block) AddBlock(block *Block) {
	b.Blocks = append(b.Blocks, block)
}

func (b *Block) AddValue(elm Elm) {
	b.Value = elm
}

func (b *Block) AddArrayValues(values ...*Block) {
	b.ArrayValues = append(b.ArrayValues, values)
}

func (b *Block) YAML(buf *bytes.Buffer, child bool) string {
	if child {
		buf.WriteString("  ")
	}

	if b.Name != "" {
		buf.WriteString(b.Name + ":")
	}

	if b.Value.Value != nil {
		buf.WriteString(" " + b.Value.Value.(string) + "\n")
	} else {
		buf.WriteString("\n")
	}

	for _, block := range b.Blocks {
		block.YAML(buf, true)
	}

	for _, array := range b.ArrayValues {
		length := len(array)
		var once bool
		for i, block := range array {
			if !once {
				buf.WriteString("  - ")
			}
			buf.WriteString(block.Name + ": " + block.Value.Value.(string) + "\n")
			if i+1 != length {
				buf.WriteString("    ")
			}
			once = true
		}
	}

	return buf.String()
}
