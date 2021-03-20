/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"io"
	"math"
)

type Dm_build_1222 struct {
	dm_build_1223 []byte
	dm_build_1224 int
}

func Dm_build_1225(dm_build_1226 int) *Dm_build_1222 {
	return &Dm_build_1222{make([]byte, 0, dm_build_1226), 0}
}

func Dm_build_1227(dm_build_1228 []byte) *Dm_build_1222 {
	return &Dm_build_1222{dm_build_1228, 0}
}

func (dm_build_1230 *Dm_build_1222) dm_build_1229(dm_build_1231 int) *Dm_build_1222 {

	dm_build_1232 := len(dm_build_1230.dm_build_1223)
	dm_build_1233 := cap(dm_build_1230.dm_build_1223)

	if dm_build_1232+dm_build_1231 <= dm_build_1233 {
		dm_build_1230.dm_build_1223 = dm_build_1230.dm_build_1223[:dm_build_1232+dm_build_1231]
	} else {
		remain := dm_build_1231 + dm_build_1232 - dm_build_1233
		nbuf := make([]byte, dm_build_1231+dm_build_1232, 2*dm_build_1233+remain)
		copy(nbuf, dm_build_1230.dm_build_1223)
		dm_build_1230.dm_build_1223 = nbuf
	}

	return dm_build_1230
}

func (dm_build_1235 *Dm_build_1222) Dm_build_1234() int {
	return len(dm_build_1235.dm_build_1223)
}

func (dm_build_1237 *Dm_build_1222) Dm_build_1236(dm_build_1238 int) *Dm_build_1222 {
	for i := dm_build_1238; i < len(dm_build_1237.dm_build_1223); i++ {
		dm_build_1237.dm_build_1223[i] = 0
	}
	dm_build_1237.dm_build_1223 = dm_build_1237.dm_build_1223[:dm_build_1238]
	return dm_build_1237
}

func (dm_build_1240 *Dm_build_1222) Dm_build_1239(dm_build_1241 int) *Dm_build_1222 {
	dm_build_1240.dm_build_1224 = dm_build_1241
	return dm_build_1240
}

func (dm_build_1243 *Dm_build_1222) Dm_build_1242() int {
	return dm_build_1243.dm_build_1224
}

func (dm_build_1245 *Dm_build_1222) Dm_build_1244(dm_build_1246 bool) int {
	return len(dm_build_1245.dm_build_1223) - dm_build_1245.dm_build_1224
}

func (dm_build_1248 *Dm_build_1222) Dm_build_1247(dm_build_1249 int, dm_build_1250 bool, dm_build_1251 bool) *Dm_build_1222 {

	if dm_build_1250 {
		if dm_build_1251 {
			dm_build_1248.dm_build_1229(dm_build_1249)
		} else {
			dm_build_1248.dm_build_1223 = dm_build_1248.dm_build_1223[:len(dm_build_1248.dm_build_1223)-dm_build_1249]
		}
	} else {
		if dm_build_1251 {
			dm_build_1248.dm_build_1224 += dm_build_1249
		} else {
			dm_build_1248.dm_build_1224 -= dm_build_1249
		}
	}

	return dm_build_1248
}

func (dm_build_1253 *Dm_build_1222) Dm_build_1252(dm_build_1254 io.Reader, dm_build_1255 int) int {
	dm_build_1256 := len(dm_build_1253.dm_build_1223)
	dm_build_1253.dm_build_1229(dm_build_1255)
	dm_build_1257 := 0
	for dm_build_1255 > 0 {
		n, err := dm_build_1254.Read(dm_build_1253.dm_build_1223[dm_build_1256+dm_build_1257:])
		if n > 0 && err == io.EOF {
			dm_build_1257 += n
			dm_build_1253.dm_build_1223 = dm_build_1253.dm_build_1223[:dm_build_1256+dm_build_1257]
			return dm_build_1257
		} else if n > 0 && err == nil {
			dm_build_1255 -= n
			dm_build_1257 += n
		} else if n == 0 && err != nil {
			panic("load err")
		}
	}

	return dm_build_1257
}

func (dm_build_1259 *Dm_build_1222) Dm_build_1258(dm_build_1260 io.Writer) *Dm_build_1222 {
	dm_build_1260.Write(dm_build_1259.dm_build_1223)
	return dm_build_1259
}

func (dm_build_1262 *Dm_build_1222) Dm_build_1261(dm_build_1263 bool) int {
	dm_build_1264 := len(dm_build_1262.dm_build_1223)
	dm_build_1262.dm_build_1229(1)

	if dm_build_1263 {
		return copy(dm_build_1262.dm_build_1223[dm_build_1264:], []byte{1})
	} else {
		return copy(dm_build_1262.dm_build_1223[dm_build_1264:], []byte{0})
	}
}

func (dm_build_1266 *Dm_build_1222) Dm_build_1265(dm_build_1267 byte) int {
	dm_build_1268 := len(dm_build_1266.dm_build_1223)
	dm_build_1266.dm_build_1229(1)

	return copy(dm_build_1266.dm_build_1223[dm_build_1268:], Dm_build_868.Dm_build_1043(dm_build_1267))
}

func (dm_build_1270 *Dm_build_1222) Dm_build_1269(dm_build_1271 int16) int {
	dm_build_1272 := len(dm_build_1270.dm_build_1223)
	dm_build_1270.dm_build_1229(2)

	return copy(dm_build_1270.dm_build_1223[dm_build_1272:], Dm_build_868.Dm_build_1046(dm_build_1271))
}

func (dm_build_1274 *Dm_build_1222) Dm_build_1273(dm_build_1275 int32) int {
	dm_build_1276 := len(dm_build_1274.dm_build_1223)
	dm_build_1274.dm_build_1229(4)

	return copy(dm_build_1274.dm_build_1223[dm_build_1276:], Dm_build_868.Dm_build_1049(dm_build_1275))
}

func (dm_build_1278 *Dm_build_1222) Dm_build_1277(dm_build_1279 uint8) int {
	dm_build_1280 := len(dm_build_1278.dm_build_1223)
	dm_build_1278.dm_build_1229(1)

	return copy(dm_build_1278.dm_build_1223[dm_build_1280:], Dm_build_868.Dm_build_1061(dm_build_1279))
}

func (dm_build_1282 *Dm_build_1222) Dm_build_1281(dm_build_1283 uint16) int {
	dm_build_1284 := len(dm_build_1282.dm_build_1223)
	dm_build_1282.dm_build_1229(2)

	return copy(dm_build_1282.dm_build_1223[dm_build_1284:], Dm_build_868.Dm_build_1064(dm_build_1283))
}

func (dm_build_1286 *Dm_build_1222) Dm_build_1285(dm_build_1287 uint32) int {
	dm_build_1288 := len(dm_build_1286.dm_build_1223)
	dm_build_1286.dm_build_1229(4)

	return copy(dm_build_1286.dm_build_1223[dm_build_1288:], Dm_build_868.Dm_build_1067(dm_build_1287))
}

func (dm_build_1290 *Dm_build_1222) Dm_build_1289(dm_build_1291 uint64) int {
	dm_build_1292 := len(dm_build_1290.dm_build_1223)
	dm_build_1290.dm_build_1229(8)

	return copy(dm_build_1290.dm_build_1223[dm_build_1292:], Dm_build_868.Dm_build_1070(dm_build_1291))
}

func (dm_build_1294 *Dm_build_1222) Dm_build_1293(dm_build_1295 float32) int {
	dm_build_1296 := len(dm_build_1294.dm_build_1223)
	dm_build_1294.dm_build_1229(4)

	return copy(dm_build_1294.dm_build_1223[dm_build_1296:], Dm_build_868.Dm_build_1067(math.Float32bits(dm_build_1295)))
}

func (dm_build_1298 *Dm_build_1222) Dm_build_1297(dm_build_1299 float64) int {
	dm_build_1300 := len(dm_build_1298.dm_build_1223)
	dm_build_1298.dm_build_1229(8)

	return copy(dm_build_1298.dm_build_1223[dm_build_1300:], Dm_build_868.Dm_build_1070(math.Float64bits(dm_build_1299)))
}

func (dm_build_1302 *Dm_build_1222) Dm_build_1301(dm_build_1303 []byte) int {
	dm_build_1304 := len(dm_build_1302.dm_build_1223)
	dm_build_1302.dm_build_1229(len(dm_build_1303))
	return copy(dm_build_1302.dm_build_1223[dm_build_1304:], dm_build_1303)
}

func (dm_build_1306 *Dm_build_1222) Dm_build_1305(dm_build_1307 []byte) int {
	return dm_build_1306.Dm_build_1273(int32(len(dm_build_1307))) + dm_build_1306.Dm_build_1301(dm_build_1307)
}

func (dm_build_1309 *Dm_build_1222) Dm_build_1308(dm_build_1310 []byte) int {
	return dm_build_1309.Dm_build_1277(uint8(len(dm_build_1310))) + dm_build_1309.Dm_build_1301(dm_build_1310)
}

func (dm_build_1312 *Dm_build_1222) Dm_build_1311(dm_build_1313 []byte) int {
	return dm_build_1312.Dm_build_1281(uint16(len(dm_build_1313))) + dm_build_1312.Dm_build_1301(dm_build_1313)
}

func (dm_build_1315 *Dm_build_1222) Dm_build_1314(dm_build_1316 []byte) int {
	return dm_build_1315.Dm_build_1301(dm_build_1316) + dm_build_1315.Dm_build_1265(0)
}

func (dm_build_1318 *Dm_build_1222) Dm_build_1317(dm_build_1319 string, dm_build_1320 string, dm_build_1321 *DmConnection) int {
	dm_build_1322 := Dm_build_868.Dm_build_1078(dm_build_1319, dm_build_1320, dm_build_1321)
	return dm_build_1318.Dm_build_1305(dm_build_1322)
}

func (dm_build_1324 *Dm_build_1222) Dm_build_1323(dm_build_1325 string, dm_build_1326 string, dm_build_1327 *DmConnection) int {
	dm_build_1328 := Dm_build_868.Dm_build_1078(dm_build_1325, dm_build_1326, dm_build_1327)
	return dm_build_1324.Dm_build_1308(dm_build_1328)
}

func (dm_build_1330 *Dm_build_1222) Dm_build_1329(dm_build_1331 string, dm_build_1332 string, dm_build_1333 *DmConnection) int {
	dm_build_1334 := Dm_build_868.Dm_build_1078(dm_build_1331, dm_build_1332, dm_build_1333)
	return dm_build_1330.Dm_build_1311(dm_build_1334)
}

func (dm_build_1336 *Dm_build_1222) Dm_build_1335(dm_build_1337 string, dm_build_1338 string, dm_build_1339 *DmConnection) int {
	dm_build_1340 := Dm_build_868.Dm_build_1078(dm_build_1337, dm_build_1338, dm_build_1339)
	return dm_build_1336.Dm_build_1314(dm_build_1340)
}

func (dm_build_1342 *Dm_build_1222) Dm_build_1341() byte {
	dm_build_1343 := Dm_build_868.Dm_build_961(dm_build_1342.dm_build_1223, dm_build_1342.dm_build_1224)
	dm_build_1342.dm_build_1224++
	return dm_build_1343
}

func (dm_build_1345 *Dm_build_1222) Dm_build_1344() int16 {
	dm_build_1346 := Dm_build_868.Dm_build_965(dm_build_1345.dm_build_1223, dm_build_1345.dm_build_1224)
	dm_build_1345.dm_build_1224 += 2
	return dm_build_1346
}

func (dm_build_1348 *Dm_build_1222) Dm_build_1347() int32 {
	dm_build_1349 := Dm_build_868.Dm_build_970(dm_build_1348.dm_build_1223, dm_build_1348.dm_build_1224)
	dm_build_1348.dm_build_1224 += 4
	return dm_build_1349
}

func (dm_build_1351 *Dm_build_1222) Dm_build_1350() int64 {
	dm_build_1352 := Dm_build_868.Dm_build_975(dm_build_1351.dm_build_1223, dm_build_1351.dm_build_1224)
	dm_build_1351.dm_build_1224 += 8
	return dm_build_1352
}

func (dm_build_1354 *Dm_build_1222) Dm_build_1353() float32 {
	dm_build_1355 := Dm_build_868.Dm_build_980(dm_build_1354.dm_build_1223, dm_build_1354.dm_build_1224)
	dm_build_1354.dm_build_1224 += 4
	return dm_build_1355
}

func (dm_build_1357 *Dm_build_1222) Dm_build_1356() float64 {
	dm_build_1358 := Dm_build_868.Dm_build_984(dm_build_1357.dm_build_1223, dm_build_1357.dm_build_1224)
	dm_build_1357.dm_build_1224 += 8
	return dm_build_1358
}

func (dm_build_1360 *Dm_build_1222) Dm_build_1359() uint8 {
	dm_build_1361 := Dm_build_868.Dm_build_988(dm_build_1360.dm_build_1223, dm_build_1360.dm_build_1224)
	dm_build_1360.dm_build_1224 += 1
	return dm_build_1361
}

func (dm_build_1363 *Dm_build_1222) Dm_build_1362() uint16 {
	dm_build_1364 := Dm_build_868.Dm_build_992(dm_build_1363.dm_build_1223, dm_build_1363.dm_build_1224)
	dm_build_1363.dm_build_1224 += 2
	return dm_build_1364
}

func (dm_build_1366 *Dm_build_1222) Dm_build_1365() uint32 {
	dm_build_1367 := Dm_build_868.Dm_build_997(dm_build_1366.dm_build_1223, dm_build_1366.dm_build_1224)
	dm_build_1366.dm_build_1224 += 4
	return dm_build_1367
}

func (dm_build_1369 *Dm_build_1222) Dm_build_1368(dm_build_1370 int) []byte {
	dm_build_1371 := Dm_build_868.Dm_build_1017(dm_build_1369.dm_build_1223, dm_build_1369.dm_build_1224, dm_build_1370)
	dm_build_1369.dm_build_1224 += dm_build_1370
	return dm_build_1371
}

func (dm_build_1373 *Dm_build_1222) Dm_build_1372() []byte {
	return dm_build_1373.Dm_build_1368(int(dm_build_1373.Dm_build_1347()))
}

func (dm_build_1375 *Dm_build_1222) Dm_build_1374() []byte {
	return dm_build_1375.Dm_build_1368(int(dm_build_1375.Dm_build_1341()))
}

func (dm_build_1377 *Dm_build_1222) Dm_build_1376() []byte {
	return dm_build_1377.Dm_build_1368(int(dm_build_1377.Dm_build_1344()))
}

func (dm_build_1379 *Dm_build_1222) Dm_build_1378(dm_build_1380 int) []byte {
	return dm_build_1379.Dm_build_1368(dm_build_1380)
}

func (dm_build_1382 *Dm_build_1222) Dm_build_1381() []byte {
	dm_build_1383 := 0
	for dm_build_1382.Dm_build_1341() != 0 {
		dm_build_1383++
	}
	dm_build_1382.Dm_build_1247(dm_build_1383, false, false)
	return dm_build_1382.Dm_build_1368(dm_build_1383)
}

func (dm_build_1385 *Dm_build_1222) Dm_build_1384(dm_build_1386 int, dm_build_1387 string, dm_build_1388 *DmConnection) string {
	return Dm_build_868.Dm_build_1115(dm_build_1385.Dm_build_1368(dm_build_1386), dm_build_1387, dm_build_1388)
}

func (dm_build_1390 *Dm_build_1222) Dm_build_1389(dm_build_1391 string, dm_build_1392 *DmConnection) string {
	return Dm_build_868.Dm_build_1115(dm_build_1390.Dm_build_1372(), dm_build_1391, dm_build_1392)
}

func (dm_build_1394 *Dm_build_1222) Dm_build_1393(dm_build_1395 string, dm_build_1396 *DmConnection) string {
	return Dm_build_868.Dm_build_1115(dm_build_1394.Dm_build_1374(), dm_build_1395, dm_build_1396)
}

func (dm_build_1398 *Dm_build_1222) Dm_build_1397(dm_build_1399 string, dm_build_1400 *DmConnection) string {
	return Dm_build_868.Dm_build_1115(dm_build_1398.Dm_build_1376(), dm_build_1399, dm_build_1400)
}

func (dm_build_1402 *Dm_build_1222) Dm_build_1401(dm_build_1403 string, dm_build_1404 *DmConnection) string {
	return Dm_build_868.Dm_build_1115(dm_build_1402.Dm_build_1381(), dm_build_1403, dm_build_1404)
}

func (dm_build_1406 *Dm_build_1222) Dm_build_1405(dm_build_1407 int, dm_build_1408 byte) int {
	return dm_build_1406.Dm_build_1441(dm_build_1407, Dm_build_868.Dm_build_1043(dm_build_1408))
}

func (dm_build_1410 *Dm_build_1222) Dm_build_1409(dm_build_1411 int, dm_build_1412 int16) int {
	return dm_build_1410.Dm_build_1441(dm_build_1411, Dm_build_868.Dm_build_1046(dm_build_1412))
}

func (dm_build_1414 *Dm_build_1222) Dm_build_1413(dm_build_1415 int, dm_build_1416 int32) int {
	return dm_build_1414.Dm_build_1441(dm_build_1415, Dm_build_868.Dm_build_1049(dm_build_1416))
}

func (dm_build_1418 *Dm_build_1222) Dm_build_1417(dm_build_1419 int, dm_build_1420 int64) int {
	return dm_build_1418.Dm_build_1441(dm_build_1419, Dm_build_868.Dm_build_1052(dm_build_1420))
}

func (dm_build_1422 *Dm_build_1222) Dm_build_1421(dm_build_1423 int, dm_build_1424 float32) int {
	return dm_build_1422.Dm_build_1441(dm_build_1423, Dm_build_868.Dm_build_1055(dm_build_1424))
}

func (dm_build_1426 *Dm_build_1222) Dm_build_1425(dm_build_1427 int, dm_build_1428 float64) int {
	return dm_build_1426.Dm_build_1441(dm_build_1427, Dm_build_868.Dm_build_1058(dm_build_1428))
}

func (dm_build_1430 *Dm_build_1222) Dm_build_1429(dm_build_1431 int, dm_build_1432 uint8) int {
	return dm_build_1430.Dm_build_1441(dm_build_1431, Dm_build_868.Dm_build_1061(dm_build_1432))
}

func (dm_build_1434 *Dm_build_1222) Dm_build_1433(dm_build_1435 int, dm_build_1436 uint16) int {
	return dm_build_1434.Dm_build_1441(dm_build_1435, Dm_build_868.Dm_build_1064(dm_build_1436))
}

func (dm_build_1438 *Dm_build_1222) Dm_build_1437(dm_build_1439 int, dm_build_1440 uint32) int {
	return dm_build_1438.Dm_build_1441(dm_build_1439, Dm_build_868.Dm_build_1067(dm_build_1440))
}

func (dm_build_1442 *Dm_build_1222) Dm_build_1441(dm_build_1443 int, dm_build_1444 []byte) int {
	return copy(dm_build_1442.dm_build_1223[dm_build_1443:], dm_build_1444)
}

func (dm_build_1446 *Dm_build_1222) Dm_build_1445(dm_build_1447 int, dm_build_1448 []byte) int {
	return dm_build_1446.Dm_build_1413(dm_build_1447, int32(len(dm_build_1448))) + dm_build_1446.Dm_build_1441(dm_build_1447+4, dm_build_1448)
}

func (dm_build_1450 *Dm_build_1222) Dm_build_1449(dm_build_1451 int, dm_build_1452 []byte) int {
	return dm_build_1450.Dm_build_1405(dm_build_1451, byte(len(dm_build_1452))) + dm_build_1450.Dm_build_1441(dm_build_1451+1, dm_build_1452)
}

func (dm_build_1454 *Dm_build_1222) Dm_build_1453(dm_build_1455 int, dm_build_1456 []byte) int {
	return dm_build_1454.Dm_build_1409(dm_build_1455, int16(len(dm_build_1456))) + dm_build_1454.Dm_build_1441(dm_build_1455+2, dm_build_1456)
}

func (dm_build_1458 *Dm_build_1222) Dm_build_1457(dm_build_1459 int, dm_build_1460 []byte) int {
	return dm_build_1458.Dm_build_1441(dm_build_1459, dm_build_1460) + dm_build_1458.Dm_build_1405(dm_build_1459+len(dm_build_1460), 0)
}

func (dm_build_1462 *Dm_build_1222) Dm_build_1461(dm_build_1463 int, dm_build_1464 string, dm_build_1465 string, dm_build_1466 *DmConnection) int {
	return dm_build_1462.Dm_build_1445(dm_build_1463, Dm_build_868.Dm_build_1078(dm_build_1464, dm_build_1465, dm_build_1466))
}

func (dm_build_1468 *Dm_build_1222) Dm_build_1467(dm_build_1469 int, dm_build_1470 string, dm_build_1471 string, dm_build_1472 *DmConnection) int {
	return dm_build_1468.Dm_build_1449(dm_build_1469, Dm_build_868.Dm_build_1078(dm_build_1470, dm_build_1471, dm_build_1472))
}

func (dm_build_1474 *Dm_build_1222) Dm_build_1473(dm_build_1475 int, dm_build_1476 string, dm_build_1477 string, dm_build_1478 *DmConnection) int {
	return dm_build_1474.Dm_build_1453(dm_build_1475, Dm_build_868.Dm_build_1078(dm_build_1476, dm_build_1477, dm_build_1478))
}

func (dm_build_1480 *Dm_build_1222) Dm_build_1479(dm_build_1481 int, dm_build_1482 string, dm_build_1483 string, dm_build_1484 *DmConnection) int {
	return dm_build_1480.Dm_build_1457(dm_build_1481, Dm_build_868.Dm_build_1078(dm_build_1482, dm_build_1483, dm_build_1484))
}

func (dm_build_1486 *Dm_build_1222) Dm_build_1485(dm_build_1487 int) byte {
	return Dm_build_868.Dm_build_1083(dm_build_1486.Dm_build_1512(dm_build_1487, 1))
}

func (dm_build_1489 *Dm_build_1222) Dm_build_1488(dm_build_1490 int) int16 {
	return Dm_build_868.Dm_build_1086(dm_build_1489.Dm_build_1512(dm_build_1490, 2))
}

func (dm_build_1492 *Dm_build_1222) Dm_build_1491(dm_build_1493 int) int32 {
	return Dm_build_868.Dm_build_1089(dm_build_1492.Dm_build_1512(dm_build_1493, 4))
}

func (dm_build_1495 *Dm_build_1222) Dm_build_1494(dm_build_1496 int) int64 {
	return Dm_build_868.Dm_build_1092(dm_build_1495.Dm_build_1512(dm_build_1496, 8))
}

func (dm_build_1498 *Dm_build_1222) Dm_build_1497(dm_build_1499 int) float32 {
	return Dm_build_868.Dm_build_1095(dm_build_1498.Dm_build_1512(dm_build_1499, 4))
}

func (dm_build_1501 *Dm_build_1222) Dm_build_1500(dm_build_1502 int) float64 {
	return Dm_build_868.Dm_build_1098(dm_build_1501.Dm_build_1512(dm_build_1502, 8))
}

func (dm_build_1504 *Dm_build_1222) Dm_build_1503(dm_build_1505 int) uint8 {
	return Dm_build_868.Dm_build_1101(dm_build_1504.Dm_build_1512(dm_build_1505, 1))
}

func (dm_build_1507 *Dm_build_1222) Dm_build_1506(dm_build_1508 int) uint16 {
	return Dm_build_868.Dm_build_1104(dm_build_1507.Dm_build_1512(dm_build_1508, 2))
}

func (dm_build_1510 *Dm_build_1222) Dm_build_1509(dm_build_1511 int) uint32 {
	return Dm_build_868.Dm_build_1107(dm_build_1510.Dm_build_1512(dm_build_1511, 4))
}

func (dm_build_1513 *Dm_build_1222) Dm_build_1512(dm_build_1514 int, dm_build_1515 int) []byte {
	return dm_build_1513.dm_build_1223[dm_build_1514 : dm_build_1514+dm_build_1515]
}

func (dm_build_1517 *Dm_build_1222) Dm_build_1516(dm_build_1518 int) []byte {
	dm_build_1519 := dm_build_1517.Dm_build_1491(dm_build_1518)
	return dm_build_1517.Dm_build_1512(dm_build_1518+4, int(dm_build_1519))
}

func (dm_build_1521 *Dm_build_1222) Dm_build_1520(dm_build_1522 int) []byte {
	dm_build_1523 := dm_build_1521.Dm_build_1485(dm_build_1522)
	return dm_build_1521.Dm_build_1512(dm_build_1522+1, int(dm_build_1523))
}

func (dm_build_1525 *Dm_build_1222) Dm_build_1524(dm_build_1526 int) []byte {
	dm_build_1527 := dm_build_1525.Dm_build_1488(dm_build_1526)
	return dm_build_1525.Dm_build_1512(dm_build_1526+2, int(dm_build_1527))
}

func (dm_build_1529 *Dm_build_1222) Dm_build_1528(dm_build_1530 int) []byte {
	dm_build_1531 := 0
	for dm_build_1529.Dm_build_1485(dm_build_1530) != 0 {
		dm_build_1530++
		dm_build_1531++
	}

	return dm_build_1529.Dm_build_1512(dm_build_1530-dm_build_1531, int(dm_build_1531))
}

func (dm_build_1533 *Dm_build_1222) Dm_build_1532(dm_build_1534 int, dm_build_1535 string, dm_build_1536 *DmConnection) string {
	return Dm_build_868.Dm_build_1115(dm_build_1533.Dm_build_1516(dm_build_1534), dm_build_1535, dm_build_1536)
}

func (dm_build_1538 *Dm_build_1222) Dm_build_1537(dm_build_1539 int, dm_build_1540 string, dm_build_1541 *DmConnection) string {
	return Dm_build_868.Dm_build_1115(dm_build_1538.Dm_build_1520(dm_build_1539), dm_build_1540, dm_build_1541)
}

func (dm_build_1543 *Dm_build_1222) Dm_build_1542(dm_build_1544 int, dm_build_1545 string, dm_build_1546 *DmConnection) string {
	return Dm_build_868.Dm_build_1115(dm_build_1543.Dm_build_1524(dm_build_1544), dm_build_1545, dm_build_1546)
}

func (dm_build_1548 *Dm_build_1222) Dm_build_1547(dm_build_1549 int, dm_build_1550 string, dm_build_1551 *DmConnection) string {
	return Dm_build_868.Dm_build_1115(dm_build_1548.Dm_build_1528(dm_build_1549), dm_build_1550, dm_build_1551)
}
