/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"container/list"
	"io"
)

type Dm_build_1144 struct {
	dm_build_1145 *list.List
	dm_build_1146 *dm_build_1198
	dm_build_1147 int
}

func Dm_build_1148() *Dm_build_1144 {
	return &Dm_build_1144{
		dm_build_1145: list.New(),
		dm_build_1147: 0,
	}
}

func (dm_build_1150 *Dm_build_1144) Dm_build_1149() int {
	return dm_build_1150.dm_build_1147
}

func (dm_build_1152 *Dm_build_1144) Dm_build_1151(dm_build_1153 *Dm_build_1222, dm_build_1154 int) int {
	var dm_build_1155 = 0
	var dm_build_1156 = 0
	for dm_build_1155 < dm_build_1154 && dm_build_1152.dm_build_1146 != nil {
		dm_build_1156 = dm_build_1152.dm_build_1146.dm_build_1206(dm_build_1153, dm_build_1154-dm_build_1155)
		if dm_build_1152.dm_build_1146.dm_build_1201 == 0 {
			dm_build_1152.dm_build_1188()
		}
		dm_build_1155 += dm_build_1156
		dm_build_1152.dm_build_1147 -= dm_build_1156
	}
	return dm_build_1155
}

func (dm_build_1158 *Dm_build_1144) Dm_build_1157(dm_build_1159 []byte, dm_build_1160 int, dm_build_1161 int) int {
	var dm_build_1162 = 0
	var dm_build_1163 = 0
	for dm_build_1162 < dm_build_1161 && dm_build_1158.dm_build_1146 != nil {
		dm_build_1163 = dm_build_1158.dm_build_1146.dm_build_1210(dm_build_1159, dm_build_1160, dm_build_1161-dm_build_1162)
		if dm_build_1158.dm_build_1146.dm_build_1201 == 0 {
			dm_build_1158.dm_build_1188()
		}
		dm_build_1162 += dm_build_1163
		dm_build_1158.dm_build_1147 -= dm_build_1163
		dm_build_1160 += dm_build_1163
	}
	return dm_build_1162
}

func (dm_build_1165 *Dm_build_1144) Dm_build_1164(dm_build_1166 io.Writer, dm_build_1167 int) int {
	var dm_build_1168 = 0
	var dm_build_1169 = 0
	for dm_build_1168 < dm_build_1167 && dm_build_1165.dm_build_1146 != nil {
		dm_build_1169 = dm_build_1165.dm_build_1146.dm_build_1215(dm_build_1166, dm_build_1167-dm_build_1168)
		if dm_build_1165.dm_build_1146.dm_build_1201 == 0 {
			dm_build_1165.dm_build_1188()
		}
		dm_build_1168 += dm_build_1169
		dm_build_1165.dm_build_1147 -= dm_build_1169
	}
	return dm_build_1168
}

func (dm_build_1171 *Dm_build_1144) Dm_build_1170(dm_build_1172 []byte, dm_build_1173 int, dm_build_1174 int) {
	if dm_build_1174 == 0 {
		return
	}
	var dm_build_1175 = dm_build_1202(dm_build_1172, dm_build_1173, dm_build_1174)
	if dm_build_1171.dm_build_1146 == nil {
		dm_build_1171.dm_build_1146 = dm_build_1175
	} else {
		dm_build_1171.dm_build_1145.PushBack(dm_build_1175)
	}
	dm_build_1171.dm_build_1147 += dm_build_1174
}

func (dm_build_1177 *Dm_build_1144) dm_build_1176(dm_build_1178 int) byte {
	var dm_build_1179 = dm_build_1178
	var dm_build_1180 = dm_build_1177.dm_build_1146
	for dm_build_1179 > 0 && dm_build_1180 != nil {
		if dm_build_1180.dm_build_1201 == 0 {
			continue
		}
		if dm_build_1179 > dm_build_1180.dm_build_1201-1 {
			dm_build_1179 -= dm_build_1180.dm_build_1201
			dm_build_1180 = dm_build_1177.dm_build_1145.Front().Value.(*dm_build_1198)
		} else {
			break
		}
	}
	return dm_build_1180.dm_build_1219(dm_build_1179)
}
func (dm_build_1182 *Dm_build_1144) Dm_build_1181(dm_build_1183 *Dm_build_1144) {
	if dm_build_1183.dm_build_1147 == 0 {
		return
	}
	var dm_build_1184 = dm_build_1183.dm_build_1146
	for dm_build_1184 != nil {
		dm_build_1182.dm_build_1185(dm_build_1184)
		dm_build_1183.dm_build_1188()
		dm_build_1184 = dm_build_1183.dm_build_1146
	}
	dm_build_1183.dm_build_1147 = 0
}
func (dm_build_1186 *Dm_build_1144) dm_build_1185(dm_build_1187 *dm_build_1198) {
	if dm_build_1187.dm_build_1201 == 0 {
		return
	}
	if dm_build_1186.dm_build_1146 == nil {
		dm_build_1186.dm_build_1146 = dm_build_1187
	} else {
		dm_build_1186.dm_build_1145.PushBack(dm_build_1187)
	}
	dm_build_1186.dm_build_1147 += dm_build_1187.dm_build_1201
}

func (dm_build_1189 *Dm_build_1144) dm_build_1188() {
	var dm_build_1190 = dm_build_1189.dm_build_1145.Front()
	if dm_build_1190 == nil {
		dm_build_1189.dm_build_1146 = nil
	} else {
		dm_build_1189.dm_build_1146 = dm_build_1190.Value.(*dm_build_1198)
		dm_build_1189.dm_build_1145.Remove(dm_build_1190)
	}
}

func (dm_build_1192 *Dm_build_1144) Dm_build_1191() []byte {
	var dm_build_1193 = make([]byte, dm_build_1192.dm_build_1147)
	var dm_build_1194 = dm_build_1192.dm_build_1146
	var dm_build_1195 = 0
	var dm_build_1196 = len(dm_build_1193)
	var dm_build_1197 = 0
	for dm_build_1194 != nil {
		if dm_build_1194.dm_build_1201 > 0 {
			if dm_build_1196 > dm_build_1194.dm_build_1201 {
				dm_build_1197 = dm_build_1194.dm_build_1201
			} else {
				dm_build_1197 = dm_build_1196
			}
			copy(dm_build_1193[dm_build_1195:dm_build_1195+dm_build_1197], dm_build_1194.dm_build_1199[dm_build_1194.dm_build_1200:dm_build_1194.dm_build_1200+dm_build_1197])
			dm_build_1195 += dm_build_1197
			dm_build_1196 -= dm_build_1197
		}
		if dm_build_1192.dm_build_1145.Front() == nil {
			dm_build_1194 = nil
		} else {
			dm_build_1194 = dm_build_1192.dm_build_1145.Front().Value.(*dm_build_1198)
		}
	}
	return dm_build_1193
}

type dm_build_1198 struct {
	dm_build_1199 []byte
	dm_build_1200 int
	dm_build_1201 int
}

func dm_build_1202(dm_build_1203 []byte, dm_build_1204 int, dm_build_1205 int) *dm_build_1198 {
	return &dm_build_1198{
		dm_build_1203,
		dm_build_1204,
		dm_build_1205,
	}
}

func (dm_build_1207 *dm_build_1198) dm_build_1206(dm_build_1208 *Dm_build_1222, dm_build_1209 int) int {
	if dm_build_1207.dm_build_1201 <= dm_build_1209 {
		dm_build_1209 = dm_build_1207.dm_build_1201
	}
	dm_build_1208.Dm_build_1301(dm_build_1207.dm_build_1199[dm_build_1207.dm_build_1200 : dm_build_1207.dm_build_1200+dm_build_1209])
	dm_build_1207.dm_build_1200 += dm_build_1209
	dm_build_1207.dm_build_1201 -= dm_build_1209
	return dm_build_1209
}

func (dm_build_1211 *dm_build_1198) dm_build_1210(dm_build_1212 []byte, dm_build_1213 int, dm_build_1214 int) int {
	if dm_build_1211.dm_build_1201 <= dm_build_1214 {
		dm_build_1214 = dm_build_1211.dm_build_1201
	}
	copy(dm_build_1212[dm_build_1213:dm_build_1213+dm_build_1214], dm_build_1211.dm_build_1199[dm_build_1211.dm_build_1200:dm_build_1211.dm_build_1200+dm_build_1214])
	dm_build_1211.dm_build_1200 += dm_build_1214
	dm_build_1211.dm_build_1201 -= dm_build_1214
	return dm_build_1214
}

func (dm_build_1216 *dm_build_1198) dm_build_1215(dm_build_1217 io.Writer, dm_build_1218 int) int {
	if dm_build_1216.dm_build_1201 <= dm_build_1218 {
		dm_build_1218 = dm_build_1216.dm_build_1201
	}
	dm_build_1217.Write(dm_build_1216.dm_build_1199[dm_build_1216.dm_build_1200 : dm_build_1216.dm_build_1200+dm_build_1218])
	dm_build_1216.dm_build_1200 += dm_build_1218
	dm_build_1216.dm_build_1201 -= dm_build_1218
	return dm_build_1218
}
func (dm_build_1220 *dm_build_1198) dm_build_1219(dm_build_1221 int) byte {
	return dm_build_1220.dm_build_1199[dm_build_1220.dm_build_1200+dm_build_1221]
}
