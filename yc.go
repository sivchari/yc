package yc

import (
	"bytes"
)

type RootBlock struct {
	Blocks []*Block
}

type Block struct {
	Name        string
	Blocks      []*Block
	Value       Elm
	ArrayValues []any
}

type Elm struct {
	Value any
}

func NewRootBlock() *RootBlock {
	return &RootBlock{}
}

func (r *RootBlock) AddBlock(block *Block) {
	r.Blocks = append(r.Blocks, block)
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

func (b *Block) AddArrayValues(values ...any) {
	b.ArrayValues = append(b.ArrayValues, values...)
}

func (b *Block) AddAnyValue(value any) {
	b.Value = Elm{Value: value}
}

func (b *RootBlock) YAML() string {
	buf := bytes.NewBuffer(nil)
	for _, block := range b.Blocks {
		block.YAML(buf, false)
	}

	return buf.String()
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

	var once bool
	length := len(b.ArrayValues)
	for i, val := range b.ArrayValues {
		block, ok := val.(*Block)
		if !ok {
			continue
		}
		if !once {
			buf.WriteString("  - ")
		}
		buf.WriteString(block.Name + ": " + block.Value.Value.(string) + "\n")
		if i+1 != length {
			buf.WriteString("    ")
		}
		once = true
	}

	return buf.String()
}
