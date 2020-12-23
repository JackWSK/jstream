package jstream

type sink interface {
    //  start iterate
    // - size: size of elements for initial collection
    begin(size int)

    // accept element from previous pipeline
    accept(element interface{})

    // iterate end
    end()
}

type baseSink struct {
    next sink
}

func (th *baseSink) send(element interface{}) {
    if th.next != nil {
        th.next.accept(element)
    }
}

func (th *baseSink) begin(size int) {
    if th.next != nil {
        th.next.begin(size)
    }
}

func (th *baseSink) end() {
    if th.next != nil {
        th.next.end()
    }
}

/// filter element
type FilterHandler func(interface{}) bool

type filterSink struct {
    baseSink
    f FilterHandler
}

func (th *filterSink) accept(element interface{}) {
    if th.f != nil && th.f(element) {
        th.send(element)
    }
}

/// map one object to another object
type MapHandler func(interface{}) interface{}

type mapSink struct {
    baseSink
    f MapHandler
}

func (th *mapSink) accept(element interface{}) {
    if th.f != nil {
        mapped := th.f(element)
        th.send(mapped)
    }
}
