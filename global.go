package jstream

import (
    "errors"
    "reflect"
)

// create stream from array
func FromArray(array interface{}) JStream {

    t := reflect.TypeOf(array)
    mustBeArrayOrSlice(t)

    p := &arrayPipeline{
        pipeline: pipeline{
            source: nil,
            prev:   nil,
            wrapSink: func(next sink) sink {
                panic(errors.New("head pipeline can not call wrap sink method"))
            },
        },
        value: reflect.ValueOf(array),
    }
    p.source = p
    return p
}

// Create array collector
// Collect all element into array
func ToArray(output interface{}) Collector {
    return &arrayCollector{
        output: output,
    }
}

// Create array collector
// Collect all element into map
func ToMap(output interface{}, keySupplier MapFunction) Collector {
    return ToMapAndChangeValue(output, keySupplier, nil)
}

func ToMapAndChangeValue(output interface{}, keySupplier MapFunction, valueSupplier MapFunction) Collector {
    return &mapCollector{
        output:        output,
        keySupplier:   keySupplier,
        valueSupplier: valueSupplier,
    }
}

// Create array collector
// Collect all element and group into map
func Group(output interface{}, keySupplier GroupFunction) Collector {
    return GroupAndChangeValue(output, keySupplier, nil)
}

func GroupAndChangeValue(output interface{}, keySupplier GroupFunction, valueSupplier GroupFunction) Collector {
    return &groupCollector{
        output:        output,
        keySupplier:   keySupplier,
        valueSupplier: valueSupplier,
    }
}
