package utils

import (
	"bytes"
	"encoding/gob"
	"github.com/wegoteam/wepkg/copy"
)

/*
*
bean属性src拷贝给dst
*/
func BeanCopy(dst, src interface{}) (err error) {
	return copy.New(src).To(dst)
}

/*
*
小写变量，函数拷贝成功
*/
func DeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}
