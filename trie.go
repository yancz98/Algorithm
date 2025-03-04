package Algorithm

import (
	"errors"
)

/**
 * 剑指 Offer II 062. 实现前缀树
 *
 * 前缀树：是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。
 * 这一数据结构有相当多的应用情景，例如自动补完和拼写检查。
 */

type Trie0 struct {
	Pass int        // 经过该字符的数量
	End  int        // 以该字符结尾的数量
	Next [26]*Trie0 // 局限：只能承载 a~z 字母的字符串
}

func InitTrie0() Trie0 {
	return Trie0{}
}

// 1.将字符串中的单个字符从前到后加入到一颗多叉树中；
// 2.字符放在路上，节点上有专属的数据项（pass、end）；
// 3.添加过程中，如果没有路就新建，有路就复用；
// 4.沿途节点的 pass 值加 1，每个字符串结束时的节点 end 加 1。
func (t *Trie0) Insert(word string) error {
	for _, ch := range word {
		// 计算 a~z 的下标，对应 0~25
		if ch < 97 || ch > 122 {
			return errors.New("字符串只能由小写英文字母组成")
		}
		idx := ch - 97

		// 该路不存在，则新建
		if t.Next[idx] == nil {
			t.Next[idx] = &Trie0{}
		}

		// 有路则复用
		t = t.Next[idx]

		// 沿途节点的 pass 加 1
		t.Pass++
	}

	// 遍历结束，末尾节点的 end 加 1
	t.End++

	return nil
}

// 查询前缀树中是否存在字符串
func (t *Trie0) Search(word string) bool {
	for _, ch := range word {
		// 计算 a~z 的下标，对应 0~25
		if ch < 97 || ch > 122 {
			// panic("字符串只能由小写英文字母组成")
			return false
		}
		idx := ch - 97

		// 任何一个字符不存在
		if t.Next[idx] == nil {
			return false
		}

		t = t.Next[idx]
	}

	// 遍历结束，end 值即为树中 word 的数量
	if t.End == 0 {
		return false
	}

	return true
}

// 查询前缀树中是否存在指定前缀
func (t *Trie0) StartWith(prefix string) bool {
	for _, ch := range prefix {
		// 计算 a~z 的下标，对应 0~25
		if ch < 97 || ch > 122 {
			// panic("字符串只能由小写英文字母组成")
			return false
		}
		idx := ch - 97

		// 任何一个字符不存在
		if t.Next[idx] == nil {
			return false
		}

		t = t.Next[idx]
	}

	// 遍历结束，t.Pass 即为以 prefix 为前缀的字符串数量

	return true
}

/**
 * 升级前缀树：能添加任意字符串
 */
type Trie struct {
	Pass int
	End  int
	Next map[rune]*Trie // rune <=> int32
}

func InitTrie() Trie {
	return Trie{
		Next: make(map[rune]*Trie, 0),
	}
}

// 往前缀树中插入字符串
func (t *Trie) Insert(word string) {
	for _, ch := range word {
		// 该路不存在，则新建
		if t.Next[ch] == nil {
			t.Next[ch] = &Trie{
				Next: make(map[rune]*Trie, 0),
			}
		}

		// 有路则复用
		t = t.Next[ch]

		// 沿途节点的 pass 加 1
		t.Pass++
	}

	// 遍历结束，末尾节点的 end 加 1
	t.End++
}

// 查询前缀树中是否存在字符串
func (t *Trie) Search(word string) bool {
	for _, ch := range word {
		// 任何一个字符不存在
		if t.Next[ch] == nil {
			return false
		}

		t = t.Next[ch]
	}

	// 遍历结束，end 值即为树中 word 的数量
	if t.End == 0 {
		return false
	}

	return true
}

// 查询前缀树中是否存在指定前缀
func (t *Trie) StartWith(prefix string) bool {
	for _, ch := range prefix {
		// 任何一个字符不存在
		if t.Next[ch] == nil {
			return false
		}

		t = t.Next[ch]
	}

	// 遍历结束，t.Pass 即为以 prefix 为前缀的字符串数量

	return true
}

// 删除
// 1.先判断 word 是否存在
// 2.若某个节点的 pass 值为 1，则后续节点无需遍历，直接将父节点置空即可
func (t *Trie) Delete(word string) {
	if t.Search(word) {
		// 存在才删除
		for _, ch := range word {
			// 若下一条路仅有这个 word 经过，则后续路也只有该 word 经过
			// 直接将父节点与与后续断开即可
			if t.Next[ch].Pass == 1 {
				t.Next[ch] = nil
				// 后续逻辑不用处理，直接结束
				return
			}

			t = t.Next[ch]
			t.Pass--
		}

		// 结尾时
		t.End--
	}
}
