// Copyright (C) 2018  Qi Yin <qiyin@thinkeridea.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package helper

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/thinkeridea/go-extend/exbytes"
)

const k = 1 << 10

// PanicRecover 帮助把 panic 转为 error返回，并获取堆栈信息打印日志。
// 该方法会把错误信息打印到标准错误输出，并包含一段堆栈信息，帮助我们快速查找程序问题。
func PanicRecover(r interface{}) error {
	loggerStderr := log.New(os.Stderr, "", log.LstdFlags)

	if r != nil {
		buf := make([]byte, 4*k)
		n := runtime.Stack(buf, false)
		loggerStderr.Printf("[Recovery] panic recovered:\n%v\n%s\n", r, exbytes.ToString(buf[:n]))

		var err error
		switch x := r.(type) {
		case string:
			err = errors.New(x)
		case error:
			err = x
		default:
			err = fmt.Errorf("%v", r)
		}

		return err
	}

	return nil
}
