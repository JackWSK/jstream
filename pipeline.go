package jstream

import "reflect"

type headPipeline interface {

    exec(next sink)

    size() int

}

type wrapSinkFunc func(next sink) sink

type pipeline struct {

    // source pipeline
    source headPipeline

    // previous stage
    prev *pipeline

    // wrap the sink, decorator design pattern
    wrapSink wrapSinkFunc

    // depth 深度
    depth int
}

func (th *pipeline) Filter(f FilterHandler) JStream {
    return th.newPipeline(func(next sink) sink {
        return &filterSink{
            baseSink: baseSink{
                next: next,
            },
            f: f,
        }
    })
}

func (th *pipeline) Map(f MapHandler) JStream {
    return th.newPipeline(func(next sink) sink {
        return &mapSink{
            baseSink: baseSink{
                next: next,
            },
            f: f,
        }
    })
}

func (th *pipeline) Collect(collector Collector) {

    if th.source == nil {
        panic("source is nil")
    }

    // iterate all pipeline from last to second
    // use wrap sink method make the call chain invert
    var s sink = collector
    for p := th; p.depth > 0; p = p.prev {
        s = p.wrapSink(s)
    }

    if s != nil {
        s.begin(th.source.size())
        th.source.exec(s)
        s.end()
    }
}

func (th *pipeline) newPipeline(wrapSinkFunc wrapSinkFunc) JStream {
    return &pipeline{
        source:   th.source,
        prev:     th,
        depth:    th.depth + 1,
        wrapSink: wrapSinkFunc,
    }
}

/// array pipeline
type arrayPipeline struct {
    pipeline
    value       reflect.Value
}

func (th *arrayPipeline) exec(next sink) {
    for i := 0; i < th.value.Len(); i++ {
        if next != nil {
            next.accept(th.value.Index(i).Interface())
        }
    }
}

func (th *arrayPipeline) size() int {
    return th.value.Len()
}
