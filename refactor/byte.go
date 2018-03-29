package refactor

import "bytes"

//多个[]byte合并成一个[]byte
func BytesMerge(baseBytes []byte, bytesSlice ...[]byte) (result []byte, err error) {
	var b []byte
	buf := bytes.NewBuffer(b)
	_, err = buf.Write(baseBytes)
	if err != nil {
		return
	}
	for k := range bytesSlice {
		_, err = buf.Write(bytesSlice[k])
		if err != nil {
			return
		}
	}
	result = buf.Bytes()
	return
}
