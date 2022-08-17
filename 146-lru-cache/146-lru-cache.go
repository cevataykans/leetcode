
type LRUCache struct {
    cur int
    max int
    lookup map[int]*list.Element
    keyMap map[*list.Element]int
    items *list.List
}


func Constructor(capacity int) LRUCache {
    
    return LRUCache{
        cur: 0,
        max: capacity,
        lookup: make(map[int]*list.Element),
        keyMap: make(map[*list.Element]int),
        items: list.New(),
    }
}


func (this *LRUCache) Get(key int) int {
    
    if e, ok := this.lookup[ key]; ok {
        this.items.MoveToBack( e)
        num, _ := e.Value.(int)
        return num
    }
    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    
    if _, ok := this.lookup[ key]; ok {
        this.lookup[ key].Value = value
        this.items.MoveToBack( this.lookup[ key])
        return
    }
    
    if this.cur == this.max {
        temp := this.items.Front()
        this.items.Remove( temp)
        
        keyToDelete := this.keyMap[ temp]
        delete( this.keyMap, temp)
        delete( this.lookup, keyToDelete)
    } else {
        this.cur++
    }
    e := this.items.PushBack( value)
    this.lookup[key] = e
    this.keyMap[e] = key
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */