/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"bytes"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"math"
)

type dm_build_867 struct{}

var Dm_build_868 = &dm_build_867{}

func (Dm_build_870 *dm_build_867) Dm_build_869(dm_build_871 []byte, dm_build_872 int, dm_build_873 byte) int {
	dm_build_871[dm_build_872] = dm_build_873
	return 1
}

func (Dm_build_875 *dm_build_867) Dm_build_874(dm_build_876 []byte, dm_build_877 int, dm_build_878 int8) int {
	dm_build_876[dm_build_877] = byte(dm_build_878)
	return 1
}

func (Dm_build_880 *dm_build_867) Dm_build_879(dm_build_881 []byte, dm_build_882 int, dm_build_883 int16) int {
	dm_build_881[dm_build_882] = byte(dm_build_883)
	dm_build_882++
	dm_build_881[dm_build_882] = byte(dm_build_883 >> 8)
	return 2
}

func (Dm_build_885 *dm_build_867) Dm_build_884(dm_build_886 []byte, dm_build_887 int, dm_build_888 int32) int {
	dm_build_886[dm_build_887] = byte(dm_build_888)
	dm_build_887++
	dm_build_886[dm_build_887] = byte(dm_build_888 >> 8)
	dm_build_887++
	dm_build_886[dm_build_887] = byte(dm_build_888 >> 16)
	dm_build_887++
	dm_build_886[dm_build_887] = byte(dm_build_888 >> 24)
	dm_build_887++
	return 4
}

func (Dm_build_890 *dm_build_867) Dm_build_889(dm_build_891 []byte, dm_build_892 int, dm_build_893 int64) int {
	dm_build_891[dm_build_892] = byte(dm_build_893)
	dm_build_892++
	dm_build_891[dm_build_892] = byte(dm_build_893 >> 8)
	dm_build_892++
	dm_build_891[dm_build_892] = byte(dm_build_893 >> 16)
	dm_build_892++
	dm_build_891[dm_build_892] = byte(dm_build_893 >> 24)
	dm_build_892++
	dm_build_891[dm_build_892] = byte(dm_build_893 >> 32)
	dm_build_892++
	dm_build_891[dm_build_892] = byte(dm_build_893 >> 40)
	dm_build_892++
	dm_build_891[dm_build_892] = byte(dm_build_893 >> 48)
	dm_build_892++
	dm_build_891[dm_build_892] = byte(dm_build_893 >> 56)
	return 8
}

func (Dm_build_895 *dm_build_867) Dm_build_894(dm_build_896 []byte, dm_build_897 int, dm_build_898 float32) int {
	return Dm_build_895.Dm_build_914(dm_build_896, dm_build_897, math.Float32bits(dm_build_898))
}

func (Dm_build_900 *dm_build_867) Dm_build_899(dm_build_901 []byte, dm_build_902 int, dm_build_903 float64) int {
	return Dm_build_900.Dm_build_919(dm_build_901, dm_build_902, math.Float64bits(dm_build_903))
}

func (Dm_build_905 *dm_build_867) Dm_build_904(dm_build_906 []byte, dm_build_907 int, dm_build_908 uint8) int {
	dm_build_906[dm_build_907] = byte(dm_build_908)
	return 1
}

func (Dm_build_910 *dm_build_867) Dm_build_909(dm_build_911 []byte, dm_build_912 int, dm_build_913 uint16) int {
	dm_build_911[dm_build_912] = byte(dm_build_913)
	dm_build_912++
	dm_build_911[dm_build_912] = byte(dm_build_913 >> 8)
	return 2
}

func (Dm_build_915 *dm_build_867) Dm_build_914(dm_build_916 []byte, dm_build_917 int, dm_build_918 uint32) int {
	dm_build_916[dm_build_917] = byte(dm_build_918)
	dm_build_917++
	dm_build_916[dm_build_917] = byte(dm_build_918 >> 8)
	dm_build_917++
	dm_build_916[dm_build_917] = byte(dm_build_918 >> 16)
	dm_build_917++
	dm_build_916[dm_build_917] = byte(dm_build_918 >> 24)
	return 3
}

func (Dm_build_920 *dm_build_867) Dm_build_919(dm_build_921 []byte, dm_build_922 int, dm_build_923 uint64) int {
	dm_build_921[dm_build_922] = byte(dm_build_923)
	dm_build_922++
	dm_build_921[dm_build_922] = byte(dm_build_923 >> 8)
	dm_build_922++
	dm_build_921[dm_build_922] = byte(dm_build_923 >> 16)
	dm_build_922++
	dm_build_921[dm_build_922] = byte(dm_build_923 >> 24)
	dm_build_922++
	dm_build_921[dm_build_922] = byte(dm_build_923 >> 32)
	dm_build_922++
	dm_build_921[dm_build_922] = byte(dm_build_923 >> 40)
	dm_build_922++
	dm_build_921[dm_build_922] = byte(dm_build_923 >> 48)
	dm_build_922++
	dm_build_921[dm_build_922] = byte(dm_build_923 >> 56)
	return 3
}

func (Dm_build_925 *dm_build_867) Dm_build_924(dm_build_926 []byte, dm_build_927 int, dm_build_928 []byte, dm_build_929 int, dm_build_930 int) int {
	copy(dm_build_926[dm_build_927:dm_build_927+dm_build_930], dm_build_928[dm_build_929:dm_build_929+dm_build_930])
	return dm_build_930
}

func (Dm_build_932 *dm_build_867) Dm_build_931(dm_build_933 []byte, dm_build_934 int, dm_build_935 []byte, dm_build_936 int, dm_build_937 int) int {
	dm_build_934 += Dm_build_932.Dm_build_914(dm_build_933, dm_build_934, uint32(dm_build_937))
	return 4 + Dm_build_932.Dm_build_924(dm_build_933, dm_build_934, dm_build_935, dm_build_936, dm_build_937)
}

func (Dm_build_939 *dm_build_867) Dm_build_938(dm_build_940 []byte, dm_build_941 int, dm_build_942 []byte, dm_build_943 int, dm_build_944 int) int {
	dm_build_941 += Dm_build_939.Dm_build_909(dm_build_940, dm_build_941, uint16(dm_build_944))
	return 2 + Dm_build_939.Dm_build_924(dm_build_940, dm_build_941, dm_build_942, dm_build_943, dm_build_944)
}

func (Dm_build_946 *dm_build_867) Dm_build_945(dm_build_947 []byte, dm_build_948 int, dm_build_949 string, dm_build_950 string, dm_build_951 *DmConnection) int {
	dm_build_952 := Dm_build_946.Dm_build_1078(dm_build_949, dm_build_950, dm_build_951)
	dm_build_948 += Dm_build_946.Dm_build_914(dm_build_947, dm_build_948, uint32(len(dm_build_952)))
	return 4 + Dm_build_946.Dm_build_924(dm_build_947, dm_build_948, dm_build_952, 0, len(dm_build_952))
}

func (Dm_build_954 *dm_build_867) Dm_build_953(dm_build_955 []byte, dm_build_956 int, dm_build_957 string, dm_build_958 string, dm_build_959 *DmConnection) int {
	dm_build_960 := Dm_build_954.Dm_build_1078(dm_build_957, dm_build_958, dm_build_959)

	dm_build_956 += Dm_build_954.Dm_build_909(dm_build_955, dm_build_956, uint16(len(dm_build_960)))
	return 2 + Dm_build_954.Dm_build_924(dm_build_955, dm_build_956, dm_build_960, 0, len(dm_build_960))
}

func (Dm_build_962 *dm_build_867) Dm_build_961(dm_build_963 []byte, dm_build_964 int) byte {
	return dm_build_963[dm_build_964]
}

func (Dm_build_966 *dm_build_867) Dm_build_965(dm_build_967 []byte, dm_build_968 int) int16 {
	var dm_build_969 int16
	dm_build_969 = int16(dm_build_967[dm_build_968] & 0xff)
	dm_build_968++
	dm_build_969 |= int16(dm_build_967[dm_build_968]&0xff) << 8
	return dm_build_969
}

func (Dm_build_971 *dm_build_867) Dm_build_970(dm_build_972 []byte, dm_build_973 int) int32 {
	var dm_build_974 int32
	dm_build_974 = int32(dm_build_972[dm_build_973] & 0xff)
	dm_build_973++
	dm_build_974 |= int32(dm_build_972[dm_build_973]&0xff) << 8
	dm_build_973++
	dm_build_974 |= int32(dm_build_972[dm_build_973]&0xff) << 16
	dm_build_973++
	dm_build_974 |= int32(dm_build_972[dm_build_973]&0xff) << 24
	return dm_build_974
}

func (Dm_build_976 *dm_build_867) Dm_build_975(dm_build_977 []byte, dm_build_978 int) int64 {
	var dm_build_979 int64
	dm_build_979 = int64(dm_build_977[dm_build_978] & 0xff)
	dm_build_978++
	dm_build_979 |= int64(dm_build_977[dm_build_978]&0xff) << 8
	dm_build_978++
	dm_build_979 |= int64(dm_build_977[dm_build_978]&0xff) << 16
	dm_build_978++
	dm_build_979 |= int64(dm_build_977[dm_build_978]&0xff) << 24
	dm_build_978++
	dm_build_979 |= int64(dm_build_977[dm_build_978]&0xff) << 32
	dm_build_978++
	dm_build_979 |= int64(dm_build_977[dm_build_978]&0xff) << 40
	dm_build_978++
	dm_build_979 |= int64(dm_build_977[dm_build_978]&0xff) << 48
	dm_build_978++
	dm_build_979 |= int64(dm_build_977[dm_build_978]&0xff) << 56
	return dm_build_979
}

func (Dm_build_981 *dm_build_867) Dm_build_980(dm_build_982 []byte, dm_build_983 int) float32 {
	return math.Float32frombits(Dm_build_981.Dm_build_997(dm_build_982, dm_build_983))
}

func (Dm_build_985 *dm_build_867) Dm_build_984(dm_build_986 []byte, dm_build_987 int) float64 {
	return math.Float64frombits(Dm_build_985.Dm_build_1002(dm_build_986, dm_build_987))
}

func (Dm_build_989 *dm_build_867) Dm_build_988(dm_build_990 []byte, dm_build_991 int) uint8 {
	return uint8(dm_build_990[dm_build_991] & 0xff)
}

func (Dm_build_993 *dm_build_867) Dm_build_992(dm_build_994 []byte, dm_build_995 int) uint16 {
	var dm_build_996 uint16
	dm_build_996 = uint16(dm_build_994[dm_build_995] & 0xff)
	dm_build_995++
	dm_build_996 |= uint16(dm_build_994[dm_build_995]&0xff) << 8
	return dm_build_996
}

func (Dm_build_998 *dm_build_867) Dm_build_997(dm_build_999 []byte, dm_build_1000 int) uint32 {
	var dm_build_1001 uint32
	dm_build_1001 = uint32(dm_build_999[dm_build_1000] & 0xff)
	dm_build_1000++
	dm_build_1001 |= uint32(dm_build_999[dm_build_1000]&0xff) << 8
	dm_build_1000++
	dm_build_1001 |= uint32(dm_build_999[dm_build_1000]&0xff) << 16
	dm_build_1000++
	dm_build_1001 |= uint32(dm_build_999[dm_build_1000]&0xff) << 24
	return dm_build_1001
}

func (Dm_build_1003 *dm_build_867) Dm_build_1002(dm_build_1004 []byte, dm_build_1005 int) uint64 {
	var dm_build_1006 uint64
	dm_build_1006 = uint64(dm_build_1004[dm_build_1005] & 0xff)
	dm_build_1005++
	dm_build_1006 |= uint64(dm_build_1004[dm_build_1005]&0xff) << 8
	dm_build_1005++
	dm_build_1006 |= uint64(dm_build_1004[dm_build_1005]&0xff) << 16
	dm_build_1005++
	dm_build_1006 |= uint64(dm_build_1004[dm_build_1005]&0xff) << 24
	dm_build_1005++
	dm_build_1006 |= uint64(dm_build_1004[dm_build_1005]&0xff) << 32
	dm_build_1005++
	dm_build_1006 |= uint64(dm_build_1004[dm_build_1005]&0xff) << 40
	dm_build_1005++
	dm_build_1006 |= uint64(dm_build_1004[dm_build_1005]&0xff) << 48
	dm_build_1005++
	dm_build_1006 |= uint64(dm_build_1004[dm_build_1005]&0xff) << 56
	return dm_build_1006
}

func (Dm_build_1008 *dm_build_867) Dm_build_1007(dm_build_1009 []byte, dm_build_1010 int) []byte {
	dm_build_1011 := Dm_build_1008.Dm_build_997(dm_build_1009, dm_build_1010)
	return dm_build_1009[dm_build_1010+4 : dm_build_1010+4+int(dm_build_1011)]

}

func (Dm_build_1013 *dm_build_867) Dm_build_1012(dm_build_1014 []byte, dm_build_1015 int) []byte {
	dm_build_1016 := Dm_build_1013.Dm_build_992(dm_build_1014, dm_build_1015)
	return dm_build_1014[dm_build_1015+2 : dm_build_1015+2+int(dm_build_1016)]

}

func (Dm_build_1018 *dm_build_867) Dm_build_1017(dm_build_1019 []byte, dm_build_1020 int, dm_build_1021 int) []byte {
	return dm_build_1019[dm_build_1020 : dm_build_1020+dm_build_1021]

}

func (Dm_build_1023 *dm_build_867) Dm_build_1022(dm_build_1024 []byte, dm_build_1025 int, dm_build_1026 int, dm_build_1027 string, dm_build_1028 *DmConnection) string {
	return Dm_build_1023.Dm_build_1115(dm_build_1024[dm_build_1025:dm_build_1025+dm_build_1026], dm_build_1027, dm_build_1028)
}

func (Dm_build_1030 *dm_build_867) Dm_build_1029(dm_build_1031 []byte, dm_build_1032 int, dm_build_1033 string, dm_build_1034 *DmConnection) string {
	dm_build_1035 := Dm_build_1030.Dm_build_997(dm_build_1031, dm_build_1032)
	dm_build_1032 += 4
	return Dm_build_1030.Dm_build_1022(dm_build_1031, dm_build_1032, int(dm_build_1035), dm_build_1033, dm_build_1034)
}

func (Dm_build_1037 *dm_build_867) Dm_build_1036(dm_build_1038 []byte, dm_build_1039 int, dm_build_1040 string, dm_build_1041 *DmConnection) string {
	dm_build_1042 := Dm_build_1037.Dm_build_992(dm_build_1038, dm_build_1039)
	dm_build_1039 += 2
	return Dm_build_1037.Dm_build_1022(dm_build_1038, dm_build_1039, int(dm_build_1042), dm_build_1040, dm_build_1041)
}

func (Dm_build_1044 *dm_build_867) Dm_build_1043(dm_build_1045 byte) []byte {
	return []byte{dm_build_1045}
}

func (Dm_build_1047 *dm_build_867) Dm_build_1046(dm_build_1048 int16) []byte {
	return []byte{byte(dm_build_1048), byte(dm_build_1048 >> 8)}
}

func (Dm_build_1050 *dm_build_867) Dm_build_1049(dm_build_1051 int32) []byte {
	return []byte{byte(dm_build_1051), byte(dm_build_1051 >> 8), byte(dm_build_1051 >> 16), byte(dm_build_1051 >> 24)}
}

func (Dm_build_1053 *dm_build_867) Dm_build_1052(dm_build_1054 int64) []byte {
	return []byte{byte(dm_build_1054), byte(dm_build_1054 >> 8), byte(dm_build_1054 >> 16), byte(dm_build_1054 >> 24), byte(dm_build_1054 >> 32),
		byte(dm_build_1054 >> 40), byte(dm_build_1054 >> 48), byte(dm_build_1054 >> 56)}
}

func (Dm_build_1056 *dm_build_867) Dm_build_1055(dm_build_1057 float32) []byte {
	return Dm_build_1056.Dm_build_1067(math.Float32bits(dm_build_1057))
}

func (Dm_build_1059 *dm_build_867) Dm_build_1058(dm_build_1060 float64) []byte {
	return Dm_build_1059.Dm_build_1070(math.Float64bits(dm_build_1060))
}

func (Dm_build_1062 *dm_build_867) Dm_build_1061(dm_build_1063 uint8) []byte {
	return []byte{byte(dm_build_1063)}
}

func (Dm_build_1065 *dm_build_867) Dm_build_1064(dm_build_1066 uint16) []byte {
	return []byte{byte(dm_build_1066), byte(dm_build_1066 >> 8)}
}

func (Dm_build_1068 *dm_build_867) Dm_build_1067(dm_build_1069 uint32) []byte {
	return []byte{byte(dm_build_1069), byte(dm_build_1069 >> 8), byte(dm_build_1069 >> 16), byte(dm_build_1069 >> 24)}
}

func (Dm_build_1071 *dm_build_867) Dm_build_1070(dm_build_1072 uint64) []byte {
	return []byte{byte(dm_build_1072), byte(dm_build_1072 >> 8), byte(dm_build_1072 >> 16), byte(dm_build_1072 >> 24), byte(dm_build_1072 >> 32), byte(dm_build_1072 >> 40), byte(dm_build_1072 >> 48), byte(dm_build_1072 >> 56)}
}

func (Dm_build_1074 *dm_build_867) Dm_build_1073(dm_build_1075 []byte, dm_build_1076 string, dm_build_1077 *DmConnection) []byte {
	if dm_build_1076 == "UTF-8" {
		return dm_build_1075
	}

	if dm_build_1077 == nil {
		if e := dm_build_1120(dm_build_1076); e != nil {
			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1075), e.NewEncoder()),
			)
			if err != nil {
				panic("UTF8 To Charset error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1077.encodeBuffer == nil {
		dm_build_1077.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1077.encode = dm_build_1120(dm_build_1077.getServerEncoding())
		dm_build_1077.transformReaderDst = make([]byte, 4096)
		dm_build_1077.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1077.encode; e != nil {

		dm_build_1077.encodeBuffer.Reset()

		n, err := dm_build_1077.encodeBuffer.ReadFrom(
			Dm_build_1134(bytes.NewReader(dm_build_1075), e.NewEncoder(), dm_build_1077.transformReaderDst, dm_build_1077.transformReaderSrc),
		)
		if err != nil {
			panic("UTF8 To Charset error!")
		}
		var tmp = make([]byte, n)
		if _, err = dm_build_1077.encodeBuffer.Read(tmp); err != nil {
			panic("UTF8 To Charset error!")
		}
		return tmp
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1079 *dm_build_867) Dm_build_1078(dm_build_1080 string, dm_build_1081 string, dm_build_1082 *DmConnection) []byte {
	return Dm_build_1079.Dm_build_1073([]byte(dm_build_1080), dm_build_1081, dm_build_1082)
}

func (Dm_build_1084 *dm_build_867) Dm_build_1083(dm_build_1085 []byte) byte {
	return Dm_build_1084.Dm_build_961(dm_build_1085, 0)
}

func (Dm_build_1087 *dm_build_867) Dm_build_1086(dm_build_1088 []byte) int16 {
	return Dm_build_1087.Dm_build_965(dm_build_1088, 0)
}

func (Dm_build_1090 *dm_build_867) Dm_build_1089(dm_build_1091 []byte) int32 {
	return Dm_build_1090.Dm_build_970(dm_build_1091, 0)
}

func (Dm_build_1093 *dm_build_867) Dm_build_1092(dm_build_1094 []byte) int64 {
	return Dm_build_1093.Dm_build_975(dm_build_1094, 0)
}

func (Dm_build_1096 *dm_build_867) Dm_build_1095(dm_build_1097 []byte) float32 {
	return Dm_build_1096.Dm_build_980(dm_build_1097, 0)
}

func (Dm_build_1099 *dm_build_867) Dm_build_1098(dm_build_1100 []byte) float64 {
	return Dm_build_1099.Dm_build_984(dm_build_1100, 0)
}

func (Dm_build_1102 *dm_build_867) Dm_build_1101(dm_build_1103 []byte) uint8 {
	return Dm_build_1102.Dm_build_988(dm_build_1103, 0)
}

func (Dm_build_1105 *dm_build_867) Dm_build_1104(dm_build_1106 []byte) uint16 {
	return Dm_build_1105.Dm_build_992(dm_build_1106, 0)
}

func (Dm_build_1108 *dm_build_867) Dm_build_1107(dm_build_1109 []byte) uint32 {
	return Dm_build_1108.Dm_build_997(dm_build_1109, 0)
}

func (Dm_build_1111 *dm_build_867) Dm_build_1110(dm_build_1112 []byte, dm_build_1113 string, dm_build_1114 *DmConnection) []byte {
	if dm_build_1113 == "UTF-8" {
		return dm_build_1112
	}

	if dm_build_1114 == nil {
		if e := dm_build_1120(dm_build_1113); e != nil {

			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1112), e.NewDecoder()),
			)
			if err != nil {

				panic("Charset To UTF8 error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1114.encodeBuffer == nil {
		dm_build_1114.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1114.encode = dm_build_1120(dm_build_1114.getServerEncoding())
		dm_build_1114.transformReaderDst = make([]byte, 4096)
		dm_build_1114.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1114.encode; e != nil {

		dm_build_1114.encodeBuffer.Reset()

		n, err := dm_build_1114.encodeBuffer.ReadFrom(
			Dm_build_1134(bytes.NewReader(dm_build_1112), e.NewDecoder(), dm_build_1114.transformReaderDst, dm_build_1114.transformReaderSrc),
		)
		if err != nil {

			panic("Charset To UTF8 error!")
		}

		return dm_build_1114.encodeBuffer.Next(int(n))
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1116 *dm_build_867) Dm_build_1115(dm_build_1117 []byte, dm_build_1118 string, dm_build_1119 *DmConnection) string {
	return string(Dm_build_1116.Dm_build_1110(dm_build_1117, dm_build_1118, dm_build_1119))
}

func dm_build_1120(dm_build_1121 string) encoding.Encoding {
	if e, err := ianaindex.MIB.Encoding(dm_build_1121); err == nil && e != nil {
		return e
	}
	return nil
}

type Dm_build_1122 struct {
	dm_build_1123 io.Reader
	dm_build_1124 transform.Transformer
	dm_build_1125 error

	dm_build_1126                []byte
	dm_build_1127, dm_build_1128 int

	dm_build_1129                []byte
	dm_build_1130, dm_build_1131 int

	dm_build_1132 bool
}

const dm_build_1133 = 4096

func Dm_build_1134(dm_build_1135 io.Reader, dm_build_1136 transform.Transformer, dm_build_1137 []byte, dm_build_1138 []byte) *Dm_build_1122 {
	dm_build_1136.Reset()
	return &Dm_build_1122{
		dm_build_1123: dm_build_1135,
		dm_build_1124: dm_build_1136,
		dm_build_1126: dm_build_1137,
		dm_build_1129: dm_build_1138,
	}
}

func (dm_build_1140 *Dm_build_1122) Read(dm_build_1141 []byte) (int, error) {
	dm_build_1142, dm_build_1143 := 0, error(nil)
	for {

		if dm_build_1140.dm_build_1127 != dm_build_1140.dm_build_1128 {
			dm_build_1142 = copy(dm_build_1141, dm_build_1140.dm_build_1126[dm_build_1140.dm_build_1127:dm_build_1140.dm_build_1128])
			dm_build_1140.dm_build_1127 += dm_build_1142
			if dm_build_1140.dm_build_1127 == dm_build_1140.dm_build_1128 && dm_build_1140.dm_build_1132 {
				return dm_build_1142, dm_build_1140.dm_build_1125
			}
			return dm_build_1142, nil
		} else if dm_build_1140.dm_build_1132 {
			return 0, dm_build_1140.dm_build_1125
		}

		if dm_build_1140.dm_build_1130 != dm_build_1140.dm_build_1131 || dm_build_1140.dm_build_1125 != nil {
			dm_build_1140.dm_build_1127 = 0
			dm_build_1140.dm_build_1128, dm_build_1142, dm_build_1143 = dm_build_1140.dm_build_1124.Transform(dm_build_1140.dm_build_1126, dm_build_1140.dm_build_1129[dm_build_1140.dm_build_1130:dm_build_1140.dm_build_1131], dm_build_1140.dm_build_1125 == io.EOF)
			dm_build_1140.dm_build_1130 += dm_build_1142

			switch {
			case dm_build_1143 == nil:
				if dm_build_1140.dm_build_1130 != dm_build_1140.dm_build_1131 {
					dm_build_1140.dm_build_1125 = nil
				}

				dm_build_1140.dm_build_1132 = dm_build_1140.dm_build_1125 != nil
				continue
			case dm_build_1143 == transform.ErrShortDst && (dm_build_1140.dm_build_1128 != 0 || dm_build_1142 != 0):

				continue
			case dm_build_1143 == transform.ErrShortSrc && dm_build_1140.dm_build_1131-dm_build_1140.dm_build_1130 != len(dm_build_1140.dm_build_1129) && dm_build_1140.dm_build_1125 == nil:

			default:
				dm_build_1140.dm_build_1132 = true

				if dm_build_1140.dm_build_1125 == nil || dm_build_1140.dm_build_1125 == io.EOF {
					dm_build_1140.dm_build_1125 = dm_build_1143
				}
				continue
			}
		}

		if dm_build_1140.dm_build_1130 != 0 {
			dm_build_1140.dm_build_1130, dm_build_1140.dm_build_1131 = 0, copy(dm_build_1140.dm_build_1129, dm_build_1140.dm_build_1129[dm_build_1140.dm_build_1130:dm_build_1140.dm_build_1131])
		}
		dm_build_1142, dm_build_1140.dm_build_1125 = dm_build_1140.dm_build_1123.Read(dm_build_1140.dm_build_1129[dm_build_1140.dm_build_1131:])
		dm_build_1140.dm_build_1131 += dm_build_1142
	}
}
