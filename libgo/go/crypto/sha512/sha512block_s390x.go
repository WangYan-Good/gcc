// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package sha512

import "internal/cpu"

var useAsm = cpu.S390X.HasSHA512
