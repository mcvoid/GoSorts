package sort

import (
    "math/rand"
    "testing"
    )

type intList []int

func (a intList) Less(i, j int) bool {
    return a[i] < a[j]
}

func (a intList) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}

func (a intList) Len() int {
    return len(a)
}

func (a intList) fillRandom(max int) {
    for i := 0; i < a.Len(); i++ {
        a[i] = rand.Intn(max)
    }
}

func (a intList) fillSequential() {
    for i := 0; i < a.Len(); i++ {
        a[i] = i
    }
}

func (a intList) fillReverseSequential() {
    for i := 0; i < a.Len(); i++ {
        a[i] = len(a) - i
    }
}

func (a intList) isSorted() bool {
    for i := 0; i < a.Len() - 1; i++ {
        if a.Less(i + 1, i) { return false }
    }
    return true
}

func testFunction(f func(Interface), t *testing.T) {
    var a intList
    // handle zero length without crashing
    a = intList(make([]int, 0))
    f(a)

    a = intList(make([]int, 2))
    a.fillSequential()
    f(a)
    if !a.isSorted() {
        t.Errorf("Failed on sequential input, size %d.", 2)
        t.Log(a)
    }
    a.fillReverseSequential()
    f(a)
    if !a.isSorted() {
        t.Errorf("Failed on reverse sequential input, size %d.", 2)
        t.Log(a)
    }
    a.fillRandom(5)
    f(a)
    if !a.isSorted() {
        t.Errorf("Failed on random input, size %d.", 2)
        t.Log(a)
    }

    a = intList(make([]int, 20))
    a.fillSequential()
    f(a)
    if !a.isSorted() {
        t.Errorf("Failed on sequential input, size %d.", 20)
        t.Log(a)
    }
    a.fillReverseSequential()
    f(a)
    if !a.isSorted() {
        t.Errorf("Failed on reverse sequential input, size %d.", 20)
        t.Log(a)
    }
    a.fillRandom(10)
    f(a)
    if !a.isSorted() {
        t.Errorf("Failed on random input, size %d.", 20)
        t.Log(a)
    }
    
    a = intList(make([]int, 2000))
    a.fillSequential()
    f(a)
    if !a.isSorted() {
        t.Errorf("Failed on sequential input, size %d.", 2000)
        t.Log("(contents too big to print)")
    }
    a.fillReverseSequential()
    f(a)
    if !a.isSorted() {
        t.Errorf("Failed on reverse sequential input, size %d.", 2000)
        t.Log("(contents too big to print)")
    }
    a.fillRandom(1000)
    f(a)
    if !a.isSorted() {
        t.Errorf("Failed on random input, size %d.", 2000)
        t.Log("(contents too big to print)")
    }
}

func benchmarkFunction(f func(Interface), b *testing.B) {
    var a intList
    a = intList(make([]int, 20000))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        b.StopTimer()
        a.fillRandom(1000)
        b.StartTimer()
        f(a)
    }
}

func TestBubble(t *testing.T) {
    testFunction(Bubble, t)
}

func TestGnome(t *testing.T) {
    testFunction(Gnome, t)
}

func TestQuicksort(t *testing.T) {
    testFunction(Quicksort, t)
}

func TestSelection(t *testing.T) {
    testFunction(Selection, t)
}

func TestInsertion(t *testing.T) {
    testFunction(Insertion, t)
}

func TestHeap(t *testing.T) {
    testFunction(Heap, t)
}

func TestParallelQuicksort(t *testing.T) {
    testFunction(ParallelQuicksort, t)
}

func BenchmarkBubble(b *testing.B) {
    benchmarkFunction(Bubble, b)
}

func BenchmarkGnome(b *testing.B) {
    benchmarkFunction(Gnome, b)
}

func BenchmarkSelection(b *testing.B) {
    benchmarkFunction(Selection, b)
}

func BenchmarkInsertion(b *testing.B) {
    benchmarkFunction(Insertion, b)
}

func BenchmarkHeap(b *testing.B) {
    benchmarkFunction(Heap, b)
}

func BenchmarkQuicksort(b *testing.B) {
    benchmarkFunction(Quicksort, b)
}

func BenchmarkParallelQuicksort(b *testing.B) {
    benchmarkFunction(ParallelQuicksort, b)
}
