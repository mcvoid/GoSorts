package sort

import "sync"

type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}

func Bubble(a Interface)  {
    for i := a.Len() - 1; i > 0; i-- {
        for j := 0; j < i; j++ {
            if a.Less(j + 1, j) {
                a.Swap(j, j + 1)
            }
        }
    }
}

func Gnome(a Interface) {
    i := 0
    for i < a.Len() {
        if i == 0 || !a.Less(i, i - 1) {
            i++
        } else {
            a.Swap(i - 1, i)
            i--
        }
    }
}

func Quicksort(a Interface)  {
    var qsort func(int, int)
    partition := func(left, right int) int {
        a.Swap(right, (left + right) / 2)
        pivot := right
        storeIndex := left
        for i := left; i < right; i++ {
            if a.Less(i, pivot) {
                a.Swap(i, storeIndex)
                storeIndex++
            }
        }
        a.Swap(storeIndex, pivot)
        return storeIndex
    }
    qsort = func(p, r int) {
        if p < r {
            q := partition(p, r)
            qsort(p, q - 1)
            qsort(q + 1, r)
        }
    }
    qsort(0, a.Len() - 1)
}

func Selection(a Interface) {
    least := func (start int) int {
        lowest := start
        for i := start + 1; i < a.Len(); i++ {
            if a.Less(i, lowest) {
                lowest = i
            }
        }
        return lowest
    }
    for i := 0; i < a.Len() - 1; i++ {
        a.Swap(i, least(i))
    }
}

func Insertion(a Interface) {
    insert := func(idx int) {
        for i := idx; i > 0 && a.Less(i, i - 1); i-- {
            a.Swap(i, i - 1)
        }
    }
    for i := 1; i < a.Len(); i++ {
        insert(i);
    }
}

func Heap(a Interface) {
    heapSize := a.Len()
    var maxHeapify func(int)
    maxHeapify = func(i int) {
        var largest int
        l, r := 2 * i + 1, 2 * i + 2
        if l <  heapSize && a.Less(i, l) {
            largest = l
        } else {
            largest = i
        }
        if r < heapSize && a.Less(largest, r) {
            largest = r
        }
        if largest != i {
            a.Swap(i, largest)
            maxHeapify(largest)
        }
    }
    for i := a.Len() / 2; i >= 0; i-- {
        maxHeapify(i)
    }
    for i := a.Len() - 1; i > 0; i-- {
        a.Swap(0, i)
        heapSize--
        maxHeapify(0)
    }
}

func ParallelQuicksort(a Interface)  {
    var wg sync.WaitGroup
    var qsort func(int, int)
    partition := func(left, right int) int {
        pivot := right
        storeIndex := left
        for i := left; i < right; i++ {
            if a.Less(i, pivot) {
                a.Swap(i, storeIndex)
                storeIndex++
            }
        }
        a.Swap(storeIndex, pivot)
        return storeIndex
    }
    qsort = func(p, r int) {
        defer wg.Done()
        if p < r {
            wg.Add(2)
            q := partition(p, r)
            qsort(p, q - 1)
            qsort(q + 1, r)
        }
    }
    wg.Add(1)
    qsort(0, a.Len() - 1)
    wg.Wait()
}
