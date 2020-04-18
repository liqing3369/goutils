package data

import mapset "github.com/deckarep/golang-set"

/**
直接使用golang-set即可
由于golang内置的数据结构中并未包含Set，因此我们可以使用github上提供的Set，
Set也分为线程安全的、非线程安全的。
线程安全的Set，使用读写锁对读写操作进行加锁，从而保证线程安全。
*/
type Set struct {
	mapset.Set
}
