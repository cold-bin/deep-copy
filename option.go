// @author cold bin
// @date 2022/9/4

package deep_copy

// Duplicator 拷贝器
type Duplicator struct {
	// 当配置选项里 DuplicatorOption.IsOpenPtr 为true时，
	// 应当初始化这个配置选项为对应的函数
	OpenPtrFunc func() error
}

// DuplicatorOption 配置一些个性深度拷贝，指定哪些能拷贝
type DuplicatorOption struct {
	IsOpenPtr bool // 是否启用指针的深度拷贝
}
