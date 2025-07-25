package data_structure

// ...类型的可以传递给数组，但是数组不能传递给...类型的数据
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	// 核心，使用切片[m:n]，将数据复制到其中
	copy(slice[m:n], data)
	return slice
}
