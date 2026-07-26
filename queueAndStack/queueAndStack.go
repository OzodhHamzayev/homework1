package main	

import ("fmt")



type Stack struct {
	items []int
}


func (s *Stack) Push(num int) {
	s.items = append(s.items, num)
}

func (s *Stack) Pop() int {
	last := len(s.items)-1
	val := last
	s.items = s.items[:last]
	return val		
}



type Queue struct { 
	items []int
}

func (q *Queue) Enqueue(newElement int) { 
	q.items = append(q.items, newElement)
}


func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}





func main() {



	myQueue := &Queue{}
	myQueue.Enqueue(1)
	myQueue.Enqueue(2)
	myQueue.Enqueue(3)
	myQueue.Enqueue(4)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)

	myStack := &Stack{}
	myStack.Push(5)
	myStack.Push(6)
	myStack.Push(7)
	myStack.Push(8)
	myStack.Pop()
	fmt.Println(myStack)
	fmt.Println(myStack.Pop())


	   list := &LinkedList{}

    list.Append(10)
    list.Append(20)
    list.Append(30)
    list.Print() // 10 -> 20 -> 30 -> nil

    list.Prepend(5)
    list.Print() // 5 -> 10 -> 20 -> 30 -> nil

    list.Delete(20)
    list.Print() // 5 -> 10 -> 30 -> nil


}






// Node — ro'yxatning bitta elementi
type Node struct {
    Value int
    Next  *Node // keyingi node'ga pointer (yoki nil)
}

// LinkedList — butun ro'yxatni boshqaruvchi struktura
type LinkedList struct {
    Head *Node // birinchi node'ga pointer
}

func (l *LinkedList) Append(value int) {
    newNode := &Node{Value: value, Next: nil}

    // Agar ro'yxat bo'sh bo'lsa
    if l.Head == nil {
        l.Head = newNode
        return
    }

    // Oxirgi node'gacha yuramiz
    current := l.Head
    for current.Next != nil {
        current = current.Next
    }
    current.Next = newNode
}
func (l *LinkedList) Prepend(value int) {
    newNode := &Node{Value: value, Next: l.Head}
    l.Head = newNode
}

func (l *LinkedList) Print() {
    current := l.Head
    for current != nil {
        fmt.Printf("%d -> ", current.Value)
        current = current.Next
    }
    fmt.Println("nil")
}

func (l *LinkedList) Delete(value int) {
    if l.Head == nil {
        return
    }

    // Agar o'chiriladigan element Head bo'lsa
    if l.Head.Value == value {
        l.Head = l.Head.Next
        return
    }

    current := l.Head
    for current.Next != nil {
        if current.Next.Value == value {
            current.Next = current.Next.Next // "bog'lab o'tib ketamiz"
            return
        }
        current = current.Next
   }
}