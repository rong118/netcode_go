type Trie struct {
	children  [26]*Trie
	isWord bool
	root      *Trie
}

func Constructor() Trie {
	return Trie{
		children:  [26]*Trie{},
		isWord: false,
		root:  &Trie{},
  }
}

func (this *Trie) Insert(word string)  {
    cur := this.root
    for _, val := range(word){
      index := val - 'a'
			if cur.children[index] == nil {
				cur.children[index] = &Trie{}
			}
			cur = cur.children[index]
    }

    cur.isWord = true
}

func (this *Trie) Search(word string) bool {
		cur := this.root
    for _, val := range(word){
        index := val - 'a'
				if cur.children[index] == nil {
					return false
				}
				cur = cur.children[index]
		}

    return cur.isWord == true
}

func (this *Trie) StartsWith(prefix string) bool {
		cur := this.root
    for _, val := range(prefix){
        index := val - 'a'
				if cur.children[index] == nil {
					return false
				}
				cur = cur.children[index]
		}

    return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */