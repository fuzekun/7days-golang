package variable

import (
	"strings"
	"sync"
)

// StringBuffer 线程安全的字符串缓冲区
type StringBuffer struct {
	builder strings.Builder
	mu      sync.RWMutex
}

// New 创建一个新的 StringBuffer 实例
func New() *StringBuffer {
	return &StringBuffer{}
}

// WriteString 追加字符串
func (sb *StringBuffer) WriteString(s string) {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	sb.builder.WriteString(s)
}

// WriteByte 追加单个字节
func (sb *StringBuffer) WriteByte(b byte) {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	sb.builder.WriteByte(b)
}

// WriteRune 追加单个 Unicode 字符
func (sb *StringBuffer) WriteRune(r rune) {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	sb.builder.WriteRune(r)
}

// Write 追加字节切片
func (sb *StringBuffer) Write(p []byte) (n int, err error) {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	return sb.builder.Write(p)
}

// String 返回当前拼接的字符串
func (sb *StringBuffer) String() string {
	sb.mu.RLock()
	defer sb.mu.RUnlock()
	return sb.builder.String()
}

// Len 返回当前长度（字节数）
func (sb *StringBuffer) Len() int {
	sb.mu.RLock()
	defer sb.mu.RUnlock()
	return sb.builder.Len()
}

// Cap 返回当前容量
func (sb *StringBuffer) Cap() int {
	sb.mu.RLock()
	defer sb.mu.RUnlock()
	return sb.builder.Cap()
}

// Reset 清空缓冲区
func (sb *StringBuffer) Reset() {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	sb.builder.Reset()
}

// Grow 预分配更多容量
func (sb *StringBuffer) Grow(n int) {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	sb.builder.Grow(n)
}
