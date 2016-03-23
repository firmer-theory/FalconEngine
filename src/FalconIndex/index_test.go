package FalconIndex

import (
	//"fmt"
	"fmt"
	//"math/rand"
	//"math/rand"
	"testing"
	"time"
	"utils"
	// "math/rand"
	//"tree"
	//fi "FalconIndex"
)
/*
func Test_NewIndex(t *testing.T) {
	logger, _ := utils.New("test_log")

	var err error
	idx := NewEmptyIndex("test_index", "./", logger)

	idx.AddField(utils.SimpleFieldInfo{FieldName: "AAA", FieldType: utils.IDX_TYPE_PK})
	idx.AddField(utils.SimpleFieldInfo{FieldName: "BBB", FieldType: utils.IDX_TYPE_STRING})
	idx.AddField(utils.SimpleFieldInfo{FieldName: "AAA", FieldType: utils.IDX_TYPE_STRING})
	idx.AddField(utils.SimpleFieldInfo{FieldName: "CCC", FieldType: utils.IDX_TYPE_NUMBER})

	if err != nil {
		t.Error("Fail...", err)
	} else {
		t.Log("UnSubscribeEmail OK")
	}

}

func Test_MergeIndex(t *testing.T) {
	logger, _ := utils.New("test_log")

	var err error
	idx := NewIndexWithLocalFile("test_index", "./", logger)

	idx.AddField(utils.SimpleFieldInfo{FieldName: "AAA", FieldType: utils.IDX_TYPE_STRING})
	idx.AddField(utils.SimpleFieldInfo{FieldName: "BBB", FieldType: utils.IDX_TYPE_STRING})
	idx.AddField(utils.SimpleFieldInfo{FieldName: "AAA", FieldType: utils.IDX_TYPE_STRING})
	idx.AddField(utils.SimpleFieldInfo{FieldName: "CCC", FieldType: utils.IDX_TYPE_NUMBER})

	for docid := uint32(1); docid < 1000000; docid++ {
		content := make(map[string]string)
		content["AAA"] = fmt.Sprintf("%v", docid)
		content["BBB"] = fmt.Sprintf("%v", rand.Intn(200))
		content["YYY"] = fmt.Sprintf("%v", rand.Intn(10))
		content["CCC"] = fmt.Sprintf("%v", rand.Intn(2000))
		err := idx.UpdateDocument(content)
		if err != nil {
			fmt.Errorf("%v", err)
		}

		if docid%100000 == 0 {
			idx.SyncMemorySegment()
		}
	}
	idx.SyncMemorySegment()
	start := time.Now()
	idx.MergeSegments()
	fmt.Printf(">>>>>>>>>>  MERGE COST TIME:%v <<<<<<<<<\n", time.Now().Sub(start))

	if err != nil {
		t.Error("Fail...", err)
	} else {
		t.Log("UnSubscribeEmail OK")
	}

}
*/
func Test_LoadIndex(t *testing.T) {

	logger, _ := utils.New("test_log")

	var err error
	idx := NewIndexWithLocalFile("test_index", "./", logger)
	start := time.Now()
	res, _ := idx.SimpleSearch([]utils.FSSearchQuery{utils.FSSearchQuery{FieldName: "BBB", Value: "50"}},
		[]utils.FSSearchFilted{utils.FSSearchFilted{FieldName: "CCC", Start: 1970, Type: utils.FILT_OVER}},
		10, 1)
	fmt.Printf(">>>>>>>>>>  SimpleSearch COST TIME:%v <<<<<<<<<\n", time.Now().Sub(start))
	fmt.Printf("[TEST] >>>>>  SimpleSearch PG[1] : %v\n", res)

	res, _ = idx.SimpleSearch([]utils.FSSearchQuery{utils.FSSearchQuery{FieldName: "BBB", Value: "50"}},
		[]utils.FSSearchFilted{utils.FSSearchFilted{FieldName: "CCC", Start: 1970, Type: utils.FILT_OVER}},
		10, 2)
	fmt.Printf(">>>>>>>>>>  SimpleSearch COST TIME:%v <<<<<<<<<\n", time.Now().Sub(start))
	fmt.Printf("[TEST] >>>>>  SimpleSearch PG[2] : %v\n", res)

	res, _ = idx.SimpleSearch([]utils.FSSearchQuery{utils.FSSearchQuery{FieldName: "BBB", Value: "50"}},
		[]utils.FSSearchFilted{utils.FSSearchFilted{FieldName: "CCC", Start: 1970, Type: utils.FILT_OVER}},
		10, 3)
	fmt.Printf(">>>>>>>>>>  SimpleSearch COST TIME:%v <<<<<<<<<\n", time.Now().Sub(start))
	fmt.Printf("[TEST] >>>>>  SimpleSearch PG[3] : %v\n", res)

	if err != nil {
		t.Error("Fail...", err)
	} else {
		t.Log("UnSubscribeEmail OK")
	}
}
