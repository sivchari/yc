package yc

import (
	"embed"
	"testing"

	"github.com/google/go-cmp/cmp"
)

//go:embed testdata
var testdata embed.FS

func TestYC(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name     string
		process  func() *RootBlock
		wantFile string
	}{
		{
			name: "simple.yaml",
			process: func() *RootBlock {
				b := NewRootBlock()

				apiVersion := NewBlock("apiVersion")
				apiVersion.AddValue(Elm{Value: "v1"})

				kind := NewBlock("kind")
				kind.AddValue(Elm{Value: "Pod"})

				name := NewBlock("name")
				name.AddValue(Elm{Value: "hello"})

				metadata := NewBlock("metadata")
				metadata.AddBlock(name)

				image := NewBlock("image")
				image.AddValue(Elm{Value: "hello-world:latest"})

				containers := NewBlock("containers")
				containers.AddArrayValues(name, image)

				spec := NewBlock("spec")
				spec.AddBlock(containers)

				b.AddBlock(apiVersion)
				b.AddBlock(kind)
				b.AddBlock(metadata)
				b.AddBlock(spec)

				return b
			},
			wantFile: "testdata/simple.yaml",
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.process()
			want, err := testdata.ReadFile(tt.wantFile)
			if err != nil {
				t.Fatalf("testdata.ReadFile(%q) failed: %v", tt.wantFile, err)
			}
			if diff := cmp.Diff(string(want), got.YAML()); diff != "" {
				t.Errorf("(-want, +got)\n%s", diff)
			}
		})
	}
}
