package mapreduce

import(
	"os"
	"encoding/json"
	"log"
	//"fmt"
	"sort"
)
func doReduce(
	jobName string, // the name of the whole MapReduce job
	reduceTask int, // which reduce task this is
	outFile string, // write the output here
	nMap int, // the number of map tasks that were run ("M" in the paper)
	reduceF func(key string, values []string) string,
) {
	keyValues := make(map[string][]string, 0) //used for creating files

    for i := 0; i < nMap; i++ {
        fileName := reduceName(jobName, i, reduceTask)
        file, err := os.Open(fileName)
        if err != nil {
            log.Fatal("doReduce: open intermediate file ", fileName, " error: ", err)
        }
        defer file.Close()

        dec := json.NewDecoder(file)
        for {
            var kv KeyValue

            err := dec.Decode(&kv)
            if err != nil {
                break
            }

            _, ok := keyValues[kv.Key]
            if !ok {
                keyValues[kv.Key] = make([]string, 0)
            }
            keyValues[kv.Key] = append(keyValues[kv.Key], kv.Value)
        }
    }

    var keys []string

    for k, _ := range keyValues {
        keys = append(keys, k)
    }
    sort.Strings(keys)

    mergeFileName := mergeName(jobName, reduceTask)
    mergeFile, err := os.Create(mergeFileName)
    if err != nil {
        log.Fatal("doReduce: create merge file ", mergeFileName, " error: ", err)
    }
    defer mergeFile.Close()

    enc := json.NewEncoder(mergeFile)
    for _, k := range keys {
        res := reduceF(k, keyValues[k])
        /*fmt.Printf("k is")
        fmt.Println(k)
        fmt.Printf("keyValues is")
        fmt.Println(keyValues[k])
        fmt.Printf("res is")
        fmt.Println(res)*/
        err := enc.Encode(&KeyValue{k, res})
        if err != nil {
            log.Fatal("doReduce: encode error: ", err)
        }
    }
	//
	// doReduce manages one reduce task: it should read the intermediate
	// files for the task, sort the intermediate key/value pairs by key,
	// call the user-defined reduce function (reduceF) for each key, and
	// write reduceF's output to disk.
	//
	// You'll need to read one intermediate file from each map task;
	// reduceName(jobName, m, reduceTask) yields the file
	// name from map task m.
	//
	// Your doMap() encoded the key/value pairs in the intermediate
	// files, so you will need to decode them. If you used JSON, you can
	// read and decode by creating a decoder and repeatedly calling
	// .Decode(&kv) on it until it returns an error.
	//
	// You may find the first example in the golang sort package
	// documentation useful.
	//
	// reduceF() is the application's reduce function. You should
	// call it once per distinct key, with a slice of all the values
	// for that key. reduceF() returns the reduced value for that key.
	//
	// You should write the reduce output as JSON encoded KeyValue
	// objects to the file named outFile. We require you to use JSON
	// because that is what the merger than combines the output
	// from all the reduce tasks expects. There is nothing special about
	// JSON -- it is just the marshalling format we chose to use. Your
	// output code will look something like this:
	//
	// enc := json.NewEncoder(file)
	// for key := ... {
	// 	enc.Encode(KeyValue{key, reduceF(...)})
	// }
	// file.Close()
	//
	// Your code here (Part I).
	//
}
