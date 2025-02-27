// Copyright 2019-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package store

import (
	"encoding/json"
	"fmt"
	"github.com/onosproject/onos-config/pkg/store/change"
	"gotest.tools/assert"
	log "k8s.io/klog"
	"os"
	"strconv"
	"testing"
)

const (
	Test1Cont1A                  = "/test1:cont1a"
	Test1Cont1ACont2A            = "/test1:cont1a/cont2a"
	Test1Cont1ACont2ALeaf2A      = "/test1:cont1a/cont2a/leaf2a"
	Test1Cont1ACont2ALeaf2B      = "/test1:cont1a/cont2a/leaf2b"
	Test1Cont1ACont2ALeaf2C      = "/test1:cont1a/cont2a/leaf2c"
	Test1Cont1ALeaf1A            = "/test1:cont1a/leaf1a"
	Test1Cont1AList2ATxout1      = "/test1:cont1a/list2a[name=txout1]"
	Test1Cont1AList2ATxout1Txpwr = "/test1:cont1a/list2a[name=txout1]/tx-power"
	Test1Cont1AList2ATxout2      = "/test1:cont1a/list2a[name=txout2]"
	Test1Cont1AList2ATxout2Txpwr = "/test1:cont1a/list2a[name=txout2]/tx-power"
	Test1Cont1AList2ATxout3      = "/test1:cont1a/list2a[name=txout3]"
	Test1Cont1AList2ATxout3Txpwr = "/test1:cont1a/list2a[name=txout3]/tx-power"
	Test1Leaftoplevel            = "/test1:leafAtTopLevel"
)

const (
	ValueEmpty          = ""
	ValueLeaf2A13       = "13"
	ValueLeaf2B159      = "1.579"
	ValueLeaf2B314      = "3.14159"
	ValueLeaf2CAbc      = "abc"
	ValueLeaf2CDef      = "def"
	ValueLeaf2CGhi      = "ghi"
	ValueLeaf1AAbcdef   = "abcdef"
	ValueTxout1Txpwr8   = "8"
	ValueTxout2Txpwr10  = "10"
	ValueTxout3Txpwr16  = "16"
	ValueLeaftopWxy1234 = "WXY-1234"
)

var Config1Paths = [11]string{
	Test1Cont1A, // 0
	Test1Cont1ACont2A,
	Test1Cont1ACont2ALeaf2A,
	Test1Cont1ACont2ALeaf2B, // 3
	Test1Cont1ACont2ALeaf2C,
	Test1Cont1ALeaf1A, // 5
	Test1Cont1AList2ATxout1,
	Test1Cont1AList2ATxout1Txpwr,
	Test1Cont1AList2ATxout3, //10
	Test1Cont1AList2ATxout3Txpwr,
	Test1Leaftoplevel,
}

var Config1Values = [11]string{
	ValueEmpty, // 0
	ValueEmpty,
	ValueLeaf2A13,
	ValueLeaf2B314, // 3
	ValueLeaf2CDef,
	ValueLeaf1AAbcdef, // 5
	ValueEmpty,
	ValueTxout1Txpwr8,
	ValueEmpty, // 10
	ValueTxout3Txpwr16,
	ValueLeaftopWxy1234,
}

var Config1PreviousPaths = [13]string{
	Test1Cont1A, // 0
	Test1Cont1ACont2A,
	Test1Cont1ACont2ALeaf2A,
	Test1Cont1ACont2ALeaf2B, // 3
	Test1Cont1ACont2ALeaf2C,
	Test1Cont1ALeaf1A, // 5
	Test1Cont1AList2ATxout1,
	Test1Cont1AList2ATxout1Txpwr,
	Test1Cont1AList2ATxout2,
	Test1Cont1AList2ATxout2Txpwr,
	Test1Cont1AList2ATxout3, //10
	Test1Cont1AList2ATxout3Txpwr,
	Test1Leaftoplevel,
}

var Config1PreviousValues = [13]string{
	ValueEmpty, // 0
	ValueEmpty,
	ValueLeaf2A13,
	ValueLeaf2B314, // 3
	ValueLeaf2CAbc,
	ValueLeaf1AAbcdef, // 5
	ValueEmpty,
	ValueTxout1Txpwr8,
	ValueEmpty,
	ValueTxout2Txpwr10,
	ValueEmpty, // 10
	ValueTxout3Txpwr16,
	ValueLeaftopWxy1234,
}

var Config1FirstPaths = [11]string{
	Test1Cont1A, // 0
	Test1Cont1ACont2A,
	Test1Cont1ACont2ALeaf2A,
	Test1Cont1ACont2ALeaf2B, // 3
	Test1Cont1ACont2ALeaf2C,
	Test1Cont1ALeaf1A, // 5
	Test1Cont1AList2ATxout1,
	Test1Cont1AList2ATxout1Txpwr,
	Test1Cont1AList2ATxout2,
	Test1Cont1AList2ATxout2Txpwr,
	Test1Leaftoplevel, //10
}

var Config1FirstValues = [11]string{
	ValueEmpty, // 0
	ValueEmpty,
	ValueLeaf2A13,
	ValueLeaf2B159, // 3
	ValueLeaf2CAbc,
	ValueLeaf1AAbcdef, // 5
	ValueEmpty,
	ValueTxout1Txpwr8,
	ValueEmpty,
	ValueTxout2Txpwr10,
	ValueLeaftopWxy1234, //10
}

var Config2Paths = [11]string{
	Test1Cont1A, // 0
	Test1Cont1ACont2A,
	Test1Cont1ACont2ALeaf2A,
	Test1Cont1ACont2ALeaf2B, // 3
	Test1Cont1ACont2ALeaf2C,
	Test1Cont1ALeaf1A, // 5
	Test1Cont1AList2ATxout2,
	Test1Cont1AList2ATxout2Txpwr,
	Test1Cont1AList2ATxout3, //10
	Test1Cont1AList2ATxout3Txpwr,
	Test1Leaftoplevel,
}

var Config2Values = [11]string{
	ValueEmpty, // 0
	ValueEmpty,
	ValueLeaf2A13,
	ValueLeaf2B314, // 3
	ValueLeaf2CGhi,
	ValueLeaf1AAbcdef, // 5
	ValueEmpty,
	ValueTxout2Txpwr10,
	ValueEmpty, // 10
	ValueTxout3Txpwr16,
	ValueLeaftopWxy1234,
}

func TestMain(m *testing.M) {
	log.SetOutput(os.Stdout)
	os.Exit(m.Run())
}

func setUp() (device1V, device2V *Configuration, changeStore map[string]*change.Change) {
	var err error

	var (
		change1, change2, change3, change4 *change.Change
	)

	config1Value01, _ := change.CreateChangeValue(Test1Cont1A, ValueEmpty, false)
	config1Value02, _ := change.CreateChangeValue(Test1Cont1ACont2A, ValueEmpty, false)
	config1Value03, _ := change.CreateChangeValue(Test1Cont1ACont2ALeaf2A, ValueLeaf2A13, false)
	config1Value04, _ := change.CreateChangeValue(Test1Cont1ACont2ALeaf2B, ValueLeaf2B159, false)
	config1Value05, _ := change.CreateChangeValue(Test1Cont1ACont2ALeaf2C, ValueLeaf2CAbc, false)
	config1Value06, _ := change.CreateChangeValue(Test1Cont1ALeaf1A, ValueLeaf1AAbcdef, false)
	config1Value07, _ := change.CreateChangeValue(Test1Cont1AList2ATxout1, ValueEmpty, false)
	config1Value08, _ := change.CreateChangeValue(Test1Cont1AList2ATxout1Txpwr, ValueTxout1Txpwr8, false)
	config1Value09, _ := change.CreateChangeValue(Test1Cont1AList2ATxout2, ValueEmpty, false)
	config1Value10, _ := change.CreateChangeValue(Test1Cont1AList2ATxout2Txpwr, ValueTxout2Txpwr10, false)
	config1Value11, _ := change.CreateChangeValue(Test1Leaftoplevel, ValueLeaftopWxy1234, false)
	change1, err = change.CreateChange(change.ValueCollections{
		config1Value01, config1Value02, config1Value03, config1Value04, config1Value05,
		config1Value06, config1Value07, config1Value08, config1Value09, config1Value10,
		config1Value11,
	}, "Original Config for test switch")
	if err != nil {
		log.Error(err)
		os.Exit(-1)
	}

	config2Value01, _ := change.CreateChangeValue(Test1Cont1ACont2ALeaf2B, ValueLeaf2B314, false)
	config2Value02, _ := change.CreateChangeValue(Test1Cont1AList2ATxout3, ValueEmpty, false)
	config2Value03, _ := change.CreateChangeValue(Test1Cont1AList2ATxout3Txpwr, ValueTxout3Txpwr16, false)
	change2, err = change.CreateChange(change.ValueCollections{
		config2Value01, config2Value02, config2Value03,
	}, "Trim power level")
	if err != nil {
		log.Error(err)
		os.Exit(-1)
	}

	config3Value01, _ := change.CreateChangeValue(Test1Cont1ACont2ALeaf2C, ValueLeaf2CDef, false)
	config3Value02, _ := change.CreateChangeValue(Test1Cont1AList2ATxout2, ValueEmpty, true)
	change3, err = change.CreateChange(change.ValueCollections{
		config3Value01, config3Value02,
	}, "Remove txout 2")
	if err != nil {
		log.Error(err)
		os.Exit(-1)
	}

	changeStore = make(map[string]*change.Change)
	changeStore[B64(change1.ID)] = change1
	changeStore[B64(change2.ID)] = change2
	changeStore[B64(change3.ID)] = change3

	device1V, err = CreateConfiguration("Device1", "1.0.0", "TestDevice",
		[]change.ID{change1.ID, change2.ID, change3.ID})
	if err != nil {
		log.Error(err)
		os.Exit(-1)
	}

	config4Value01, _ := change.CreateChangeValue(Test1Cont1ACont2ALeaf2C, ValueLeaf2CGhi, false)
	config4Value02, _ := change.CreateChangeValue(Test1Cont1AList2ATxout1, ValueEmpty, true)
	change4, err = change.CreateChange(change.ValueCollections{
		config4Value01, config4Value02,
	}, "Remove txout 1")
	if err != nil {
		log.Error(err)
		os.Exit(-1)
	}
	changeStore[B64(change4.ID)] = change4

	device2V, err = CreateConfiguration("Device2", "1.0.0", "TestDevice",
		[]change.ID{change1.ID, change2.ID, change4.ID})
	if err != nil {
		log.Error(err)
		os.Exit(-1)
	}

	return device1V, device2V, changeStore
}

func Test_device1_version(t *testing.T) {
	device1V, _, changeStore := setUp()

	log.Info("Configuration ", device1V.Name, " (latest) Changes:")
	for idx, cid := range device1V.Changes {
		log.Infof("%d: %s\n", idx, B64([]byte(cid)))
	}

	assert.Equal(t, device1V.Name, ConfigName("Device1-1.0.0"))

	config := device1V.ExtractFullConfig(changeStore, 0)
	for _, c := range config {
		log.Infof("Path %s = %s\n", c.Path, c.Value)
	}

	for i := 0; i < len(Config1Paths); i++ {
		checkPathvalue(t, config, i,
			Config1Paths[0:11], Config1Values[0:11])
	}
}

func Test_device1_prev_version(t *testing.T) {
	device1V, _, changeStore := setUp()

	const changePrevious = 1
	log.Info("Configuration ", device1V.Name, " (n-1) Changes:")
	for idx, cid := range device1V.Changes[0 : len(device1V.Changes)-changePrevious] {
		log.Infof("%d: %s\n", idx, B64([]byte(cid)))
	}

	assert.Equal(t, device1V.Name, ConfigName("Device1-1.0.0"))

	config := device1V.ExtractFullConfig(changeStore, changePrevious)
	for _, c := range config {
		log.Infof("Path %s = %s\n", c.Path, c.Value)
	}

	for i := 0; i < len(Config1PreviousPaths); i++ {
		checkPathvalue(t, config, i,
			Config1PreviousPaths[0:13], Config1PreviousValues[0:13])
	}
}

func Test_device1_first_version(t *testing.T) {
	device1V, _, changeStore := setUp()
	const changePrevious = 2
	log.Info("Configuration ", device1V.Name, " (n-2) Changes:")
	for idx, cid := range device1V.Changes[0 : len(device1V.Changes)-changePrevious] {
		log.Infof("%d: %s\n", idx, B64([]byte(cid)))
	}

	assert.Equal(t, device1V.Name, ConfigName("Device1-1.0.0"))

	config := device1V.ExtractFullConfig(changeStore, changePrevious)
	for _, c := range config {
		log.Infof("Path %s = %s\n", c.Path, c.Value)
	}

	for i := 0; i < len(Config1FirstPaths); i++ {
		checkPathvalue(t, config, i,
			Config1FirstPaths[0:11], Config1FirstValues[0:11])
	}
}

func Test_device1_invalid_version(t *testing.T) {
	device1V, _, changeStore := setUp()
	const changePrevious = 3
	log.Info("Configuration ", device1V.Name, " (n-3) Changes:")
	for idx, cid := range device1V.Changes[0 : len(device1V.Changes)-changePrevious] {
		log.Infof("%d: %s\n", idx, B64([]byte(cid)))
	}

	assert.Equal(t, device1V.Name, ConfigName("Device1-1.0.0"))

	config := device1V.ExtractFullConfig(changeStore, changePrevious)
	if len(config) > 0 {
		t.Errorf("Not expecting any values for change (n-3). Got %d", len(config))
	}

}

func Test_device2_version(t *testing.T) {
	_, device2V, changeStore := setUp()
	log.Info("Configuration ", device2V.Name, " (latest) Changes:")
	for idx, cid := range device2V.Changes {
		log.Infof("%d: %s\n", idx, B64([]byte(cid)))
	}

	assert.Equal(t, device2V.Name, ConfigName("Device2-1.0.0"))

	config := device2V.ExtractFullConfig(changeStore, 0)
	for _, c := range config {
		log.Infof("Path %s = %s\n", c.Path, c.Value)
	}

	for i := 0; i < len(Config2Paths); i++ {
		checkPathvalue(t, config, i,
			Config2Paths[0:11], Config2Values[0:11])
	}
}

func checkPathvalue(t *testing.T, config []*change.ConfigValue, index int,
	expPaths []string, expValues []string) {

	// Check that they are kept in a consistent order
	if config[index].Path != expPaths[index] {
		t.Errorf("Unexpected change %d Exp: %s, Got %s", index,
			expPaths[index], config[index].Path)
	}

	if config[index].Value != expValues[index] {
		t.Errorf("Unexpected change %d Exp: %s, Got %s", index,
			expValues[index], config[index].Value)
	}
}

func Test_convertChangeToGnmi(t *testing.T) {

	config3Value01, _ := change.CreateChangeValue(Test1Cont1ACont2ALeaf2C, ValueLeaf2CDef, false)
	config3Value02, _ := change.CreateChangeValue(Test1Cont1AList2ATxout2, ValueEmpty, true)
	change3, err := change.CreateChange(change.ValueCollections{
		config3Value01, config3Value02,
	}, "Remove txout 2")
	if err != nil {
		log.Error(err)
		os.Exit(-1)
	}

	setRequest, parseError := change3.GnmiChange()

	assert.NilError(t, parseError, "Parsing error for Gnmi change request")
	assert.Equal(t, len(setRequest.Update), 1)

	jsonstr, _ := json.Marshal(setRequest.Update[0])

	expectedStr := "{\"path\":{\"elem\":[{\"name\":\"cont1a\"},{\"name\":\"cont2a\"}," +
		"{\"name\":\"leaf2c\"}]},\"val\":{\"Value\":{\"StringVal\":\"def\"}}}"

	assert.Equal(t, string(jsonstr), expectedStr)

	assert.Equal(t, len(setRequest.Delete), 1)

	jsonstr2, _ := json.Marshal(setRequest.Delete[0])

	expectedStr2 := "{\"elem\":[{\"name\":\"cont1a\"},{\"name\":\"list2a\",\"key\":{\"name\":\"txout2\"}}]}"
	assert.Equal(t, string(jsonstr2), expectedStr2)
}

func Test_writeOutChangeFile(t *testing.T) {
	_, _, changeStore := setUp()
	if _, err := os.Stat("testout"); os.IsNotExist(err) {
		os.Mkdir("testout", os.ModePerm)
	}
	changeStoreFile, err := os.Create("testout/changeStore-sample.json")
	if err != nil {
		t.Errorf("%s", err)
	}

	jsonEncoder := json.NewEncoder(changeStoreFile)
	var csf = ChangeStore{Version: StoreVersion,
		Storetype: StoreTypeChange, Store: changeStore}
	err = jsonEncoder.Encode(csf)
	if err != nil {
		log.Error(err)
		os.Exit(-1)
	}
	defer changeStoreFile.Close()
}

func Test_loadChangeStoreFile(t *testing.T) {
	changeStore, err := LoadChangeStore("testout/changeStore-sample.json")
	assert.NilError(t, err, "Unexpected error when loading Change Store from file %s", err)
	assert.Equal(t, changeStore.Version, StoreVersion)
}

func Test_loadChangeStoreFileError(t *testing.T) {
	changeStore, err := LoadChangeStore("nonexistent.json")
	assert.Assert(t, err != nil, "Expected an error when loading Change Store from invalid file")
	assert.Equal(t, changeStore.Version, "")
}

func Test_loadConfigStoreFileBadVersion(t *testing.T) {
	_, err := LoadConfigStore("testdata/configStore-badVersion.json")
	assert.ErrorContains(t, err, "Store version invalid")
}

func Test_loadConfigStoreFileBadType(t *testing.T) {
	_, err := LoadConfigStore("testdata/configStore-badType.json")
	assert.ErrorContains(t, err, "Store type invalid")
}

func Test_loadChangeStoreFileBadVersion(t *testing.T) {
	_, err := LoadChangeStore("testdata/changeStore-badVersion.json")
	assert.ErrorContains(t, err, "Store version invalid")
}

func Test_loadChangeStoreFileBadType(t *testing.T) {
	_, err := LoadChangeStore("testdata/changeStore-badType.json")
	assert.ErrorContains(t, err, "Store type invalid")
}

func Test_loadNetworkStoreFileError(t *testing.T) {
	networkStore, err := LoadNetworkStore("nonexistent.json")
	assert.Assert(t, err != nil, "Expected an error when loading Change Store from invalid file")
	assert.Assert(t, networkStore == nil, "")
}

func Test_loadNetworkStoreFileBadVersion(t *testing.T) {
	_, err := LoadNetworkStore("testdata/networkStore-badVersion.json")
	assert.ErrorContains(t, err, "Store version invalid")
}

func Test_loadNetworkStoreFileBadType(t *testing.T) {
	_, err := LoadNetworkStore("testdata/networkStore-badType.json")
	assert.ErrorContains(t, err, "Store type invalid")
}

func TestCreateConfiguration_badname(t *testing.T) {
	_, err :=
		CreateConfiguration("", "1.0.0", "TestDevice",
			[]change.ID{})
	assert.ErrorContains(t, err, "deviceName, version and deviceType must have values", "Empty")

	_, err =
		CreateConfiguration("abc", "1.0.0", "TestDevice",
			[]change.ID{})
	assert.ErrorContains(t, err, "name abc does not match pattern", "Too short")

	_, err =
		CreateConfiguration("abc???", "1.0.0", "TestDevice",
			[]change.ID{})
	assert.ErrorContains(t, err, "name abc??? does not match pattern", "Illegal char")

	_, err =
		CreateConfiguration("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNO",
			"1.0.0", "TestDevice",
			[]change.ID{})
	assert.ErrorContains(t, err, "name abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNO does not match pattern", "Too long")
}

func TestCreateConfiguration_badversion(t *testing.T) {
	_, err :=
		CreateConfiguration("localhost-1", "1.234567890", "TestDevice",
			[]change.ID{})
	assert.ErrorContains(t, err, "version 1.234567890 does not match pattern", "Too long")

	_, err =
		CreateConfiguration("localhost-1", "a", "TestDevice",
			[]change.ID{})
	assert.ErrorContains(t, err, "version a does not match pattern", "Too short")

	_, err =
		CreateConfiguration("localhost-1", "1:0:0", "TestDevice",
			[]change.ID{})
	assert.ErrorContains(t, err, "version 1:0:0 does not match pattern", "Illegal char")
}

func TestCreateConfiguration_badtype(t *testing.T) {
	_, err :=
		CreateConfiguration("localhost-1", "1.0.0", "TestDeviceType",
			[]change.ID{})
	assert.ErrorContains(t, err, "deviceType TestDeviceType does not match pattern", "Too long")
}

func Test_writeOutConfigFile(t *testing.T) {
	device1V, device2V, _ := setUp()
	configurationStore := make(map[ConfigName]Configuration)
	configurationStore[device1V.Name] = *device1V
	configurationStore[device2V.Name] = *device2V

	configStoreFile, _ := os.Create("testout/configStore-sample.json")
	jsonEncoder := json.NewEncoder(configStoreFile)
	err := jsonEncoder.Encode(ConfigurationStore{Version: StoreVersion,
		Storetype: StoreTypeConfig, Store: configurationStore})
	if err != nil {
		log.Error(err)
		os.Exit(-1)
	}
	defer configStoreFile.Close()
}

func Test_loadConfigStoreFile(t *testing.T) {
	configStore, err := LoadConfigStore("testout/configStore-sample.json")

	assert.NilError(t, err, "Unexpected error when loading Config Store from file %s", err)
	assert.Equal(t, configStore.Version, StoreVersion)
}

func Test_loadConfigStoreFileError(t *testing.T) {
	configStore, err := LoadConfigStore("nonexistent.json")
	assert.Assert(t, err != nil, "Expected an error when loading Config Store from invalid file")
	assert.Equal(t, configStore.Version, "")
}

func Test_writeOutNetworkFile(t *testing.T) {
	networkStore := make([]NetworkConfiguration, 0)
	ccs := make(map[ConfigName]change.ID)
	ccs["Device2VersionMain"] = []byte("DCuMG07l01g2BvMdEta+7DyxMxk=")
	ccs["Device2VersionMain"] = []byte("LsDuwm2XJjdOq+u9QEcUJo/HxaM=")
	nw1, err := NewNetworkConfiguration("testChange", "nw1", ccs)
	if err != nil {
		log.Error(err)
		os.Exit(-1)
	}
	networkStore = append(networkStore, *nw1)

	networkStoreFile, _ := os.Create("testout/networkStore-sample.json")
	jsonEncoder := json.NewEncoder(networkStoreFile)
	err = jsonEncoder.Encode(NetworkStore{Version: StoreVersion,
		Storetype: StoreTypeNetwork, Store: networkStore})
	if err != nil {
		log.Error(err)
		os.Exit(-1)
	}
	defer networkStoreFile.Close()
}

// Test_createnetStore tests that a valid network config name is accepted
// Note: the testing against duplicate names is done in northbound/set_test.go
func Test_createnetStore(t *testing.T) {
	nwStore, err := NewNetworkConfiguration("testnwstore", "onos", make(map[ConfigName]change.ID))
	assert.NilError(t, err, "Unexpected error")

	assert.Equal(t, nwStore.User, "onos")
	assert.Equal(t, nwStore.Name, "testnwstore")
	assert.Equal(t, len(nwStore.ConfigurationChanges), 0)
}

func Test_createnetStore_badname(t *testing.T) {
	nwStore, err := NewNetworkConfiguration("???????", "onos", make(map[ConfigName]change.ID))
	assert.ErrorContains(t, err, "Error in name ???????")
	assert.Check(t, nwStore == nil, "unexpected result", nwStore)
}

// Test_createnetStore tests that a valid network config name is accepted
// Note: the testing of name with no extension 100 is done in northbound/set_test.go
func Test_createnetStore_noname(t *testing.T) {
	var noname string
	_, err := NewNetworkConfiguration(noname, "onos", make(map[ConfigName]change.ID))
	assert.ErrorContains(t, err, "Empty name not allowed")
}

func Test_loadNetworkStoreFile(t *testing.T) {
	networkStore, err := LoadNetworkStore("testout/networkStore-sample.json")
	assert.NilError(t, err, "Unexpected error when loading Network Store from file %s", err)
	assert.Equal(t, networkStore.Version, StoreVersion)
}

func BenchmarkCreateChangeValue(b *testing.B) {

	for i := 0; i < b.N; i++ {
		path := fmt.Sprintf("/test-%s", strconv.Itoa(b.N))
		cv, _ := change.CreateChangeValue(path, strconv.Itoa(i), false)
		err := cv.IsPathValid()
		assert.NilError(b, err, "path not valid %s", err)

	}
}

func BenchmarkCreateChange(b *testing.B) {

	changeValues := change.ValueCollections{}
	for i := 0; i < b.N; i++ {
		path := fmt.Sprintf("/test%d", i)
		cv, _ := change.CreateChangeValue(path, strconv.Itoa(i), false)
		changeValues = append(changeValues, cv)
	}

	change, _ := change.CreateChange(changeValues, "Benchmarked Change")

	err := change.IsValid()
	assert.NilError(b, err, "Invalid change %s", err)
}
