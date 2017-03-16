package templates

import (
	"context"
	"github.com/go-qbit/template/filter"
	"github.com/go-qbit/template/utils"
	"io"
)

var (
	s96fe34f46084699530cb8cbd35895dda = []byte{0x22, 0x3c, 0x70, 0x3e, 0x22}
	s7b0f3af983749abb730d6180d0a43586 = []byte{0x22, 0x28, 0x22}
	sb653000c3bf94af208d7453ae8971fe5 = []byte{0x22, 0x4d, 0x61, 0x6e, 0x22}
	scd6c3f76caa0c7939d754b9b331e1c49 = []byte{0x22, 0x29, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x70, 0x3e, 0x22}
	sf3f1196886a8ba27a941f322228e4439 = []byte{0x22, 0x3c, 0x2f, 0x68, 0x31, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x68, 0x72, 0x3e, 0x22}
	s5d70ad090a95c7bb49b92212141809dc = []byte{0x22, 0x3a, 0x22}
	sd5b84c0a6a6a38e2acb82de2283f56b9 = []byte{0x22, 0x57, 0x6f, 0x6d, 0x61, 0x6e, 0x22}
	sfcc3d7489d15ef49dbbf735234234cf7 = []byte{0x22, 0x20, 0x22}
	sa5f6581eccd129593b2a978d9e4fc581 = []byte{0x22, 0x3c, 0x68, 0x31, 0x3e, 0x22}
)

func ProcessTest(ctx context.Context, w io.Writer, header string, users []User) {
	Wrapperpage(ctx, w, func() {
		w.Write(sa5f6581eccd129593b2a978d9e4fc581)
		io.WriteString(w, filter.HTML(header))
		w.Write(sf3f1196886a8ba27a941f322228e4439)
		for _, user := range users {
			w.Write(s96fe34f46084699530cb8cbd35895dda)
			ProcessUserName(ctx, w, user)
			w.Write(s5d70ad090a95c7bb49b92212141809dc)
			io.WriteString(w, utils.ToString(user.Age))
			w.Write(s7b0f3af983749abb730d6180d0a43586)
			if user.IsMan {
				w.Write(sb653000c3bf94af208d7453ae8971fe5)
			} else {
				w.Write(sd5b84c0a6a6a38e2acb82de2283f56b9)
			}
			w.Write(scd6c3f76caa0c7939d754b9b331e1c49)
		}
	}, header)
}

func ProcessUserName(ctx context.Context, w io.Writer, user User) {
	io.WriteString(w, filter.HTML(user.Name))
	w.Write(sfcc3d7489d15ef49dbbf735234234cf7)
	io.WriteString(w, filter.HTML(user.Lastname))
}
