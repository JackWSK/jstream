package jstream

type JStream interface {
    // return true if the element you want to keep
    Filter(f FilterHandler) JStream

    // map element to anther element
    Map(f MapHandler) JStream

    // collect all element
    Collect(c Collector)
}
