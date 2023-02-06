package fuzzImagebuilder

import "strconv"
import "strings"
import "github.com/openshift/imagebuilder/dockerfile/parser"
import "github.com/openshift/imagebuilder/strslice"
import "github.com/openshift/imagebuilder/signal"
import fuzz "github.com/AdaLogics/go-fuzz-headers"


func mayhemit(bytes []byte) int {

    var num int
    if len(bytes) > 2 {
        num, _ = strconv.Atoi(string(bytes[0]))
        bytes = bytes[1:]

        switch num {
    
        case 0:
            var test strslice.StrSlice
            test.UnmarshalJSON(bytes)
            return 0

        case 1:
            fuzzConsumer := fuzz.NewConsumer(bytes)
            var data string
            err := fuzzConsumer.CreateSlice(&data)
            if err != nil {
                return 0
            }

            content := strings.NewReader(data)
            parser.Parse(content)
            return 0

        default:
            fuzzConsumer := fuzz.NewConsumer(bytes)
            var content string
            err := fuzzConsumer.CreateSlice(&content)
            if err != nil {
                return 0
            }
            
            signal.CheckSignal(content)
            return 0

        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}