package templates

import (
	"context"
	"github.com/go-qbit/template/filter"
	"github.com/go-qbit/template/utils"
	"io"
)

var (
	s627661c621eab1b7b298abc47d1a250d = []byte{0x4d, 0x61, 0x6e}
	s921c704d576fc50eafc928f28d287019 = []byte{0x29, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x70, 0x3e}
	s7215ee9c7d9dc229d2921a40e899ec5f = []byte{0x20}
	sa5c0aa5a33c99f4179eb48443e8bb33c = []byte{0x3c, 0x68, 0x31, 0x3e}
	s4da1a46ec20cf93ee5c846a51e04f0ed = []byte{0x3c, 0x70, 0x3e}
	s853ae90f0351324bd73ea615e6487517 = []byte{0x3a}
	s84c40473414caf2ed4a7b1283e48bbf4 = []byte{0x28}
	s2b7c41aee73401e8cd970fba4c3bd4bd = []byte{0x3c, 0x2f, 0x68, 0x31, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x68, 0x72, 0x3e}
	s5276343758bbfb9e0a374c9dd2fa6ce7 = []byte{0x57, 0x6f, 0x6d, 0x61, 0x6e}
)

func ProcessTest(ctx context.Context, w io.Writer, header string, users []User) {
	Wrapperpage(ctx, w, func() {
		w.Write(sa5c0aa5a33c99f4179eb48443e8bb33c)
		io.WriteString(w, filter.HTML(header))
		w.Write(s2b7c41aee73401e8cd970fba4c3bd4bd)
		for _, user := range users {
			w.Write(s4da1a46ec20cf93ee5c846a51e04f0ed)
			ProcessUserName(ctx, w, user)
			w.Write(s853ae90f0351324bd73ea615e6487517)
			io.WriteString(w, utils.ToString(user.Age))
			w.Write(s84c40473414caf2ed4a7b1283e48bbf4)
			if user.IsMan {
				w.Write(s627661c621eab1b7b298abc47d1a250d)
			} else {
				w.Write(s5276343758bbfb9e0a374c9dd2fa6ce7)
			}
			w.Write(s921c704d576fc50eafc928f28d287019)
		}
	}, header)
}

func ProcessUserName(ctx context.Context, w io.Writer, user User) {
	io.WriteString(w, filter.HTML(user.Name))
	w.Write(s7215ee9c7d9dc229d2921a40e899ec5f)
	io.WriteString(w, filter.HTML(user.Lastname))
}
