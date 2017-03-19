package templates

import (
	"context"
	"github.com/go-qbit/template/filter"
	"github.com/go-qbit/template/utils"
	"io"
)

var (
	s9444e86b936ba53fdc56632949cb3038 = []byte{0x4d, 0x61, 0x6e}
	s9ed6dfeeb854d5f5e292f977637a729b = []byte{0x57, 0x6f, 0x6d, 0x61, 0x6e}
	sf715636fb0391a8fdd43473a43b1ea72 = []byte{0x3c, 0x70, 0x3e}
	sd2c38510206af487d1a4984a8af4ca49 = []byte{0x28}
	sf8de4bf55e29f51642893aea2c376d30 = []byte{0x3a}
	s9e23173c90723882045e02a5a6c79a9b = []byte{0x29, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x70, 0x3e}
	s3c00a1c100edee938dfcd27113fd1f93 = []byte{0x20}
	sedbe56ac7404d66602a627bd637ef2c1 = []byte{0x3c, 0x68, 0x31, 0x3e}
	s0f9baed348e2ded156c71ed2555fb1da = []byte{0x3c, 0x2f, 0x68, 0x31, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x68, 0x72, 0x3e}
)

func ProcessTest(ctx context.Context, w io.Writer, header string, users []User) {
	Wrapperpage(ctx, w, func() {
		w.Write(sedbe56ac7404d66602a627bd637ef2c1)
		io.WriteString(w, filter.HTML(header))
		w.Write(s0f9baed348e2ded156c71ed2555fb1da)
		for _, user := range users {
			w.Write(sf715636fb0391a8fdd43473a43b1ea72)
			ProcessUserName(ctx, w, user)
			w.Write(sf8de4bf55e29f51642893aea2c376d30)
			io.WriteString(w, utils.ToString(user.Age))
			w.Write(sd2c38510206af487d1a4984a8af4ca49)
			if user.IsMan {
				w.Write(s9444e86b936ba53fdc56632949cb3038)
			} else {
				w.Write(s9ed6dfeeb854d5f5e292f977637a729b)
			}
			w.Write(s9e23173c90723882045e02a5a6c79a9b)
		}
	}, header)
}

func ProcessUserName(ctx context.Context, w io.Writer, user User) {
	io.WriteString(w, filter.HTML(user.Name))
	w.Write(s3c00a1c100edee938dfcd27113fd1f93)
	io.WriteString(w, filter.HTML(user.Lastname))
}
