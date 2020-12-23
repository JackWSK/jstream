package jstream

import (
    "reflect"
)

type Collector interface {
    sink
}

////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////// Array Collector ///////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////

// collect all element into array
type arrayCollector struct {
    output                interface{}
    assignableOutputValue reflect.Value
    outputValue           reflect.Value
}

func (th *arrayCollector) begin(size int) {
    sliceType, outputAddressValue := parseSlice(th.output)
    th.assignableOutputValue = outputAddressValue
    th.outputValue = reflect.MakeSlice(sliceType, 0, size)
}

func (th *arrayCollector) accept(element interface{}) {
    th.outputValue = reflect.Append(th.outputValue, reflect.ValueOf(element))
}

func (th *arrayCollector) end() {
    th.assignableOutputValue.Set(th.outputValue)
}

func parseSlice(output interface{}) (sliceType reflect.Type, assignableOutputValue reflect.Value) {
    value := reflect.ValueOf(output)
    value = extractCanSet(value)

    outputType := value.Type()
    mustBeArrayOrSlice(outputType)

    return outputType, value
}

////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////// Map Collector ///////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////

// collect all element into map
type MapFunction func(interface{}) interface{}

type mapCollector struct {
    output                interface{}
    assignableOutputValue reflect.Value
    outputValue           reflect.Value
    keySupplier           MapFunction
    // allow nil, use element itself if the value supplier is nil
    valueSupplier MapFunction
}

func (th *mapCollector) begin(size int) {
    mapType, assignableOutputValue := parseMap(th.output)
    th.assignableOutputValue = assignableOutputValue
    th.outputValue = reflect.MakeMap(mapType)
}

func (th *mapCollector) accept(element interface{}) {
    key := th.keySupplier(element)
    var value interface{}
    if th.valueSupplier == nil {
        value = element
    } else {
        value = th.valueSupplier(element)
    }

    th.outputValue.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(value))
}

func (th *mapCollector) end() {
    th.assignableOutputValue.Set(th.outputValue)
}

func parseMap(output interface{}) (mapType reflect.Type, assignableOutputValue reflect.Value) {
    value := reflect.ValueOf(output)
    value = extractCanSet(value)
    mustBeMap(value.Type())

    outputType := value.Type()
    return outputType, value
}

////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////// Group Collector ///////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////

// collect all element into map
type GroupFunction func(interface{}) interface{}

type groupCollector struct {
    output                interface{}
    assignableOutputValue reflect.Value
    outputValue           reflect.Value
    // type of map value
    sliceType             reflect.Type
    keySupplier           GroupFunction
    // allow nil, use element itself if the value supplier is nil
    valueSupplier GroupFunction
}

func (th *groupCollector) begin(size int) {
    mapType, assignableOutputValue := parseMap(th.output)
    th.assignableOutputValue = assignableOutputValue
    th.outputValue = reflect.MakeMap(mapType)
    th.sliceType = mapType.Elem()
}

func (th *groupCollector) accept(element interface{}) {
    // get key
    key := th.keySupplier(element)

    // get value
    var value interface{}
    if th.valueSupplier == nil {
        value = element
    } else {
        value = th.valueSupplier(element)
    }

    // get reflect key value and value value
    keyValue := reflect.ValueOf(key)
    valueValue := reflect.ValueOf(value)

    // reset list value in map
    sliceValue := th.outputValue.MapIndex(keyValue)
    if sliceValue.Kind() == reflect.Invalid {
       sliceValue = reflect.MakeSlice(th.sliceType, 0, 0)
    }
    sliceValue = reflect.Append(sliceValue, valueValue)

    th.outputValue.SetMapIndex(keyValue, sliceValue)
}

func (th *groupCollector) end() {
    th.assignableOutputValue.Set(th.outputValue)
}

func extractCanSet(value reflect.Value) reflect.Value {
    if value.Kind() != reflect.Ptr {
        panic("value must be ptr")
    }

    value = value.Elem()
    if !value.CanSet() {
        panic("outputValue must be can set")
    }

    return value
}

func mustBeArrayOrSlice(t reflect.Type) {
    if t.Kind() != reflect.Array &&
        t.Kind() != reflect.Slice {
        panic("only support array or slice")
    }
}

func mustBeMap(t reflect.Type) {
    if t.Kind() != reflect.Map {
        panic("only support map")
    }
}
