/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"dm/util"
	"os"
	"strconv"
	"strings"
)

const (
	Dm_build_266 = "7.6.0.0"

	Dm_build_267 = "7.0.0.9"

	Dm_build_268 = "8.0.0.73"

	Dm_build_269 = "7.1.2.128"

	Dm_build_270 = "7.1.5.144"

	Dm_build_271 = "7.1.6.123"

	Dm_build_272 = 32768 - 128

	Dm_build_273 int16 = 1

	Dm_build_274 int16 = 2

	Dm_build_275 int16 = 3

	Dm_build_276 int16 = 4

	Dm_build_277 int16 = 5

	Dm_build_278 int16 = 6

	Dm_build_279 int16 = 7

	Dm_build_280 int16 = 8

	Dm_build_281 int16 = 9

	Dm_build_282 int16 = 13

	Dm_build_283 int16 = 14

	Dm_build_284 int16 = 15

	Dm_build_285 int16 = 17

	Dm_build_286 int16 = 21

	Dm_build_287 int16 = 24

	Dm_build_288 int16 = 27

	Dm_build_289 int16 = 29

	Dm_build_290 int16 = 30

	Dm_build_291 int16 = 31

	Dm_build_292 int16 = 32

	Dm_build_293 int16 = 44

	Dm_build_294 int16 = 52

	Dm_build_295 int16 = 60

	Dm_build_296 int16 = 71

	Dm_build_297 int16 = 90

	Dm_build_298 int16 = 91

	Dm_build_299 int16 = 200

	Dm_build_300 = 64

	Dm_build_301 = 20

	Dm_build_302 = 0

	Dm_build_303 = 4

	Dm_build_304 = 6

	Dm_build_305 = 10

	Dm_build_306 = 14

	Dm_build_307 = 18

	Dm_build_308 = 19

	Dm_build_309 = 128

	Dm_build_310 = 256

	Dm_build_311 = 0xffff

	Dm_build_312 int32 = 2

	Dm_build_313 int32 = 5

	Dm_build_314 = -1

	Dm_build_315 uint16 = 0xFFFE

	Dm_build_316 uint16 = uint16(Dm_build_315 - 3)

	Dm_build_317 uint16 = Dm_build_315

	Dm_build_318 int32 = 0xFFFF

	Dm_build_319 int32 = 0x80

	Dm_build_320 byte = 0x60

	Dm_build_321 uint16 = uint16(Dm_build_317)

	Dm_build_322 uint16 = uint16(Dm_build_318)

	Dm_build_323 int16 = 0x00

	Dm_build_324 int16 = 0x03

	Dm_build_325 int32 = 0x80

	Dm_build_326 byte = 0

	Dm_build_327 byte = 1

	Dm_build_328 byte = 2

	Dm_build_329 byte = 3

	Dm_build_330 byte = 4

	Dm_build_331 byte = Dm_build_326

	Dm_build_332 int = 10

	Dm_build_333 int32 = 32

	Dm_build_334 int32 = 65536

	Dm_build_335 byte = 0

	Dm_build_336 byte = 1

	Dm_build_337 int32 = 0x00000000

	Dm_build_338 int32 = 0x00000020

	Dm_build_339 int32 = 0x00000040

	Dm_build_340 int32 = 0x00000FFF

	Dm_build_341 int32 = 0

	Dm_build_342 int32 = 1

	Dm_build_343 int32 = 2

	Dm_build_344 int32 = 3

	Dm_build_345 = 8192

	Dm_build_346 = 1

	Dm_build_347 = 2

	Dm_build_348 = 0

	Dm_build_349 = 0

	Dm_build_350 = 1

	Dm_build_351 = -1

	Dm_build_352 int16 = 0

	Dm_build_353 int16 = 1

	Dm_build_354 int16 = 2

	Dm_build_355 int16 = 3

	Dm_build_356 int16 = 4

	Dm_build_357 int16 = 127

	Dm_build_358 int16 = Dm_build_357 + 20

	Dm_build_359 int16 = Dm_build_357 + 21

	Dm_build_360 int16 = Dm_build_357 + 22

	Dm_build_361 int16 = Dm_build_357 + 24

	Dm_build_362 int16 = Dm_build_357 + 25

	Dm_build_363 int16 = Dm_build_357 + 26

	Dm_build_364 int16 = Dm_build_357 + 30

	Dm_build_365 int16 = Dm_build_357 + 31

	Dm_build_366 int16 = Dm_build_357 + 32

	Dm_build_367 int16 = Dm_build_357 + 33

	Dm_build_368 int16 = Dm_build_357 + 35

	Dm_build_369 int16 = Dm_build_357 + 38

	Dm_build_370 int16 = Dm_build_357 + 39

	Dm_build_371 int16 = Dm_build_357 + 51

	Dm_build_372 int16 = Dm_build_357 + 71

	Dm_build_373 int16 = Dm_build_357 + 124

	Dm_build_374 int16 = Dm_build_357 + 125

	Dm_build_375 int16 = Dm_build_357 + 126

	Dm_build_376 int16 = Dm_build_357 + 127

	Dm_build_377 int16 = Dm_build_357 + 128

	Dm_build_378 int16 = Dm_build_357 + 129

	Dm_build_379 byte = 0

	Dm_build_380 byte = 2

	Dm_build_381 = 2048

	Dm_build_382 = -1

	Dm_build_383 = 0

	Dm_build_384 = 16000

	Dm_build_385 = 32000

	Dm_build_386 = 0x00000000

	Dm_build_387 = 0x00000020

	Dm_build_388 = 0x00000040

	Dm_build_389 = 0x00000FFF
)

type dm_build_390 interface {
	dm_build_391()
	dm_build_392() error
	dm_build_393()
	dm_build_394(imsg dm_build_390) error
	dm_build_395() error
	dm_build_396() (interface{}, error)
	dm_build_397()
	dm_build_398(imsg dm_build_390) (interface{}, error)
	dm_build_399()
	dm_build_400() error
	dm_build_401() byte
	dm_build_402() int32
	dm_build_403(length int32)
	dm_build_404() int16
}

type dm_build_405 struct {
	dm_build_406 *dm_build_2

	dm_build_407 int16

	dm_build_408 int32

	dm_build_409 *DmStatement
}

func (dm_build_411 *dm_build_405) dm_build_410(dm_build_412 *dm_build_2, dm_build_413 int16) *dm_build_405 {
	dm_build_411.dm_build_406 = dm_build_412
	dm_build_411.dm_build_407 = dm_build_413
	return dm_build_411
}

func (dm_build_415 *dm_build_405) dm_build_414(dm_build_416 *dm_build_2, dm_build_417 int16, dm_build_418 *DmStatement) *dm_build_405 {
	dm_build_415.dm_build_410(dm_build_416, dm_build_417).dm_build_409 = dm_build_418
	return dm_build_415
}

func dm_build_419(dm_build_420 *dm_build_2, dm_build_421 int16) *dm_build_405 {
	return new(dm_build_405).dm_build_410(dm_build_420, dm_build_421)
}

func dm_build_422(dm_build_423 *dm_build_2, dm_build_424 int16, dm_build_425 *DmStatement) *dm_build_405 {
	return new(dm_build_405).dm_build_414(dm_build_423, dm_build_424, dm_build_425)
}

func (dm_build_427 *dm_build_405) dm_build_391() {
	dm_build_427.dm_build_406.dm_build_5.Dm_build_1236(0)
	dm_build_427.dm_build_406.dm_build_5.Dm_build_1247(Dm_build_300, true, true)
}

func (dm_build_429 *dm_build_405) dm_build_392() error {
	return nil
}

func (dm_build_431 *dm_build_405) dm_build_393() {
	if dm_build_431.dm_build_409 == nil {
		dm_build_431.dm_build_406.dm_build_5.Dm_build_1413(Dm_build_302, 0)
	} else {
		dm_build_431.dm_build_406.dm_build_5.Dm_build_1413(Dm_build_302, dm_build_431.dm_build_409.id)
	}

	dm_build_431.dm_build_406.dm_build_5.Dm_build_1409(Dm_build_303, dm_build_431.dm_build_407)
	dm_build_431.dm_build_406.dm_build_5.Dm_build_1413(Dm_build_304, int32(dm_build_431.dm_build_406.dm_build_5.Dm_build_1234()-Dm_build_300))
}

func (dm_build_433 *dm_build_405) dm_build_395() error {
	dm_build_433.dm_build_406.dm_build_5.Dm_build_1239(0)
	dm_build_433.dm_build_406.dm_build_5.Dm_build_1247(Dm_build_300, false, true)
	return dm_build_433.dm_build_438()
}

func (dm_build_435 *dm_build_405) dm_build_396() (interface{}, error) {
	return nil, nil
}

func (dm_build_437 *dm_build_405) dm_build_397() {
}

func (dm_build_439 *dm_build_405) dm_build_438() error {
	dm_build_439.dm_build_408 = dm_build_439.dm_build_406.dm_build_5.Dm_build_1491(Dm_build_305)
	if dm_build_439.dm_build_408 < 0 && dm_build_439.dm_build_408 != EC_RN_EXCEED_ROWSET_SIZE.ErrCode {
		return (&DmError{dm_build_439.dm_build_408, dm_build_439.dm_build_440(), nil, ""}).throw()
	} else if dm_build_439.dm_build_408 > 0 {

	} else if dm_build_439.dm_build_407 == Dm_build_299 || dm_build_439.dm_build_407 == Dm_build_273 {
		dm_build_439.dm_build_440()
	}

	return nil
}

func (dm_build_441 *dm_build_405) dm_build_440() string {

	dm_build_442 := dm_build_441.dm_build_406.dm_build_6.getServerEncoding()

	if dm_build_442 != "" && dm_build_442 == ENCODING_EUCKR && Locale != LANGUAGE_EN {
		dm_build_442 = ENCODING_GB18030
	}

	dm_build_441.dm_build_406.dm_build_5.Dm_build_1247(int(dm_build_441.dm_build_406.dm_build_5.Dm_build_1347()), false, true)

	dm_build_441.dm_build_406.dm_build_5.Dm_build_1247(int(dm_build_441.dm_build_406.dm_build_5.Dm_build_1347()), false, true)

	dm_build_441.dm_build_406.dm_build_5.Dm_build_1247(int(dm_build_441.dm_build_406.dm_build_5.Dm_build_1347()), false, true)

	return dm_build_441.dm_build_406.dm_build_5.Dm_build_1389(dm_build_442, dm_build_441.dm_build_406.dm_build_6)
}

func (dm_build_444 *dm_build_405) dm_build_394(dm_build_445 dm_build_390) (dm_build_446 error) {
	dm_build_445.dm_build_391()
	if dm_build_446 = dm_build_445.dm_build_392(); dm_build_446 != nil {
		return dm_build_446
	}
	dm_build_445.dm_build_393()
	return nil
}

func (dm_build_448 *dm_build_405) dm_build_398(dm_build_449 dm_build_390) (dm_build_450 interface{}, dm_build_451 error) {
	dm_build_451 = dm_build_449.dm_build_395()
	if dm_build_451 != nil {
		return nil, dm_build_451
	}
	dm_build_450, dm_build_451 = dm_build_449.dm_build_396()
	if dm_build_451 != nil {
		return nil, dm_build_451
	}
	dm_build_449.dm_build_397()
	return dm_build_450, nil
}

func (dm_build_453 *dm_build_405) dm_build_399() {
	dm_build_453.dm_build_406.dm_build_5.Dm_build_1405(Dm_build_308, dm_build_453.dm_build_401())
}

func (dm_build_455 *dm_build_405) dm_build_400() error {
	dm_build_456 := dm_build_455.dm_build_406.dm_build_5.Dm_build_1485(Dm_build_308)
	dm_build_457 := dm_build_455.dm_build_401()
	if dm_build_456 != dm_build_457 {
		return ECGO_MSG_CHECK_ERROR.throw()
	}
	return nil
}

func (dm_build_459 *dm_build_405) dm_build_401() byte {
	dm_build_460 := dm_build_459.dm_build_406.dm_build_5.Dm_build_1485(0)

	for i := 1; i < Dm_build_308; i++ {
		dm_build_460 ^= dm_build_459.dm_build_406.dm_build_5.Dm_build_1485(i)
	}

	return dm_build_460
}

func (dm_build_462 *dm_build_405) dm_build_402() int32 {
	return dm_build_462.dm_build_406.dm_build_5.Dm_build_1491(Dm_build_304)
}

func (dm_build_464 *dm_build_405) dm_build_403(dm_build_465 int32) {
	dm_build_464.dm_build_406.dm_build_5.Dm_build_1413(Dm_build_304, dm_build_465)
}

func (dm_build_467 *dm_build_405) dm_build_404() int16 {
	return dm_build_467.dm_build_407
}

type dm_build_468 struct {
	dm_build_405
}

func dm_build_469(dm_build_470 *dm_build_2) *dm_build_468 {
	dm_build_471 := new(dm_build_468)
	dm_build_471.dm_build_410(dm_build_470, Dm_build_280)
	return dm_build_471
}

type dm_build_472 struct {
	dm_build_405
	dm_build_473 string
}

func dm_build_474(dm_build_475 *dm_build_2, dm_build_476 *DmStatement, dm_build_477 string) *dm_build_472 {
	dm_build_478 := new(dm_build_472)
	dm_build_478.dm_build_414(dm_build_475, Dm_build_288, dm_build_476)
	dm_build_478.dm_build_473 = dm_build_477
	dm_build_478.dm_build_409.cursorName = dm_build_477
	return dm_build_478
}

func (dm_build_480 *dm_build_472) dm_build_392() error {
	dm_build_480.dm_build_406.dm_build_5.Dm_build_1335(dm_build_480.dm_build_473, dm_build_480.dm_build_406.dm_build_6.getServerEncoding(), dm_build_480.dm_build_406.dm_build_6)
	dm_build_480.dm_build_406.dm_build_5.Dm_build_1273(1)
	return nil
}

type Dm_build_481 struct {
	dm_build_497
	dm_build_482 []OptParameter
}

func dm_build_483(dm_build_484 *dm_build_2, dm_build_485 *DmStatement, dm_build_486 []OptParameter) *Dm_build_481 {
	dm_build_487 := new(Dm_build_481)
	dm_build_487.dm_build_414(dm_build_484, Dm_build_298, dm_build_485)
	dm_build_487.dm_build_482 = dm_build_486
	return dm_build_487
}

func (dm_build_489 *Dm_build_481) dm_build_392() error {
	dm_build_490 := len(dm_build_489.dm_build_482)

	dm_build_489.dm_build_511(int16(dm_build_490), 1)

	dm_build_489.dm_build_406.dm_build_5.Dm_build_1335(dm_build_489.dm_build_409.nativeSql, dm_build_489.dm_build_409.dmConn.getServerEncoding(), dm_build_489.dm_build_409.dmConn)

	for _, param := range dm_build_489.dm_build_482 {
		dm_build_489.dm_build_406.dm_build_5.Dm_build_1265(param.ioType)
		dm_build_489.dm_build_406.dm_build_5.Dm_build_1273(int32(param.tp))
		dm_build_489.dm_build_406.dm_build_5.Dm_build_1273(int32(param.prec))
		dm_build_489.dm_build_406.dm_build_5.Dm_build_1273(int32(param.scale))
	}

	for _, param := range dm_build_489.dm_build_482 {
		if param.bytes == nil {
			dm_build_489.dm_build_406.dm_build_5.Dm_build_1281(Dm_build_317)
		} else {
			dm_build_489.dm_build_406.dm_build_5.Dm_build_1311(param.bytes[:len(param.bytes)])
		}
	}
	return nil
}

func (dm_build_492 *Dm_build_481) dm_build_396() (interface{}, error) {
	return dm_build_492.dm_build_497.dm_build_396()
}

const (
	Dm_build_493 int = 0x01

	Dm_build_494 int = 0x02

	Dm_build_495 int = 0x04

	Dm_build_496 int = 0x08
)

type dm_build_497 struct {
	dm_build_405
	dm_build_498 [][]interface{}
	dm_build_499 []parameter
	dm_build_500 bool
}

func dm_build_501(dm_build_502 *dm_build_2, dm_build_503 int16, dm_build_504 *DmStatement) *dm_build_497 {
	dm_build_505 := new(dm_build_497)
	dm_build_505.dm_build_414(dm_build_502, dm_build_503, dm_build_504)
	dm_build_505.dm_build_500 = true
	return dm_build_505
}

func dm_build_506(dm_build_507 *dm_build_2, dm_build_508 *DmStatement, dm_build_509 [][]interface{}) *dm_build_497 {
	dm_build_510 := new(dm_build_497)

	if dm_build_507.dm_build_6.Execute2 {
		dm_build_510.dm_build_414(dm_build_507, Dm_build_282, dm_build_508)
	} else {
		dm_build_510.dm_build_414(dm_build_507, Dm_build_278, dm_build_508)
	}

	dm_build_510.dm_build_499 = dm_build_508.params
	dm_build_510.dm_build_498 = dm_build_509
	dm_build_510.dm_build_500 = true
	return dm_build_510
}

func (dm_build_512 *dm_build_497) dm_build_511(dm_build_513 int16, dm_build_514 int64) {

	dm_build_515 := Dm_build_301

	if dm_build_512.dm_build_406.dm_build_6.autoCommit {
		dm_build_515 += dm_build_512.dm_build_406.dm_build_5.Dm_build_1405(dm_build_515, 1)
	} else {
		dm_build_515 += dm_build_512.dm_build_406.dm_build_5.Dm_build_1405(dm_build_515, 0)
	}

	dm_build_515 += dm_build_512.dm_build_406.dm_build_5.Dm_build_1409(dm_build_515, dm_build_513)

	dm_build_515 += dm_build_512.dm_build_406.dm_build_5.Dm_build_1405(dm_build_515, 1)

	dm_build_515 += dm_build_512.dm_build_406.dm_build_5.Dm_build_1417(dm_build_515, dm_build_514)

	dm_build_515 += dm_build_512.dm_build_406.dm_build_5.Dm_build_1417(dm_build_515, dm_build_512.dm_build_409.cursorUpdateRow)

	if dm_build_512.dm_build_409.maxRows <= 0 || dm_build_512.dm_build_409.dmConn.dmConnector.enRsCache {
		dm_build_515 += dm_build_512.dm_build_406.dm_build_5.Dm_build_1417(dm_build_515, INT64_MAX)
	} else {
		dm_build_515 += dm_build_512.dm_build_406.dm_build_5.Dm_build_1417(dm_build_515, dm_build_512.dm_build_409.maxRows)
	}

	dm_build_515 += dm_build_512.dm_build_406.dm_build_5.Dm_build_1405(dm_build_515, 1)

	if dm_build_512.dm_build_406.dm_build_6.dmConnector.continueBatchOnError {
		dm_build_515 += dm_build_512.dm_build_406.dm_build_5.Dm_build_1405(dm_build_515, 1)
	} else {
		dm_build_515 += dm_build_512.dm_build_406.dm_build_5.Dm_build_1405(dm_build_515, 0)
	}

	dm_build_515 += dm_build_512.dm_build_406.dm_build_5.Dm_build_1405(dm_build_515, 0)

	dm_build_515 += dm_build_512.dm_build_406.dm_build_5.Dm_build_1405(dm_build_515, 0)

	if dm_build_512.dm_build_409.queryTimeout == 0 {
		dm_build_515 += dm_build_512.dm_build_406.dm_build_5.Dm_build_1413(dm_build_515, -1)
	} else {
		dm_build_515 += dm_build_512.dm_build_406.dm_build_5.Dm_build_1413(dm_build_515, dm_build_512.dm_build_409.queryTimeout)
	}
}

func (dm_build_517 *dm_build_497) dm_build_392() error {
	var dm_build_518 int16
	var dm_build_519 int64

	if dm_build_517.dm_build_499 != nil {
		dm_build_518 = int16(len(dm_build_517.dm_build_499))
	} else {
		dm_build_518 = 0
	}

	if dm_build_517.dm_build_498 != nil {
		dm_build_519 = int64(len(dm_build_517.dm_build_498))
	} else {
		dm_build_519 = 0
	}

	dm_build_517.dm_build_511(dm_build_518, dm_build_519)

	if dm_build_518 > 0 {
		err := dm_build_517.dm_build_520(dm_build_517.dm_build_499)
		if err != nil {
			return err
		}
		if dm_build_517.dm_build_498 != nil && len(dm_build_517.dm_build_498) > 0 {
			for _, paramObject := range dm_build_517.dm_build_498 {
				if err := dm_build_517.dm_build_523(paramObject); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (dm_build_521 *dm_build_497) dm_build_520(dm_build_522 []parameter) error {
	for _, param := range dm_build_522 {

		if param.colType == CURSOR && param.ioType == IO_TYPE_OUT {
			dm_build_521.dm_build_406.dm_build_5.Dm_build_1265(IO_TYPE_INOUT)
		} else {
			dm_build_521.dm_build_406.dm_build_5.Dm_build_1265(param.ioType)
		}

		dm_build_521.dm_build_406.dm_build_5.Dm_build_1273(param.colType)

		lprec := param.prec
		lscale := param.scale
		typeDesc := param.typeDescriptor
		switch param.colType {
		case ARRAY, SARRAY:
			tmp, err := getPackArraySize(typeDesc)
			if err != nil {
				return err
			}
			lprec = int32(tmp)
		case PLTYPE_RECORD:
			tmp, err := getPackRecordSize(typeDesc)
			if err != nil {
				return err
			}
			lprec = int32(tmp)
		case CLASS:
			tmp, err := getPackClassSize(typeDesc)
			if err != nil {
				return err
			}
			lprec = int32(tmp)
		case BLOB:
			if isComplexType(int(param.colType), int(param.scale)) {
				lprec = int32(typeDesc.getObjId())
				if lprec == 4 {
					lprec = int32(typeDesc.getOuterId())
				}
			}
		}

		dm_build_521.dm_build_406.dm_build_5.Dm_build_1273(lprec)

		dm_build_521.dm_build_406.dm_build_5.Dm_build_1273(lscale)

		switch param.colType {
		case ARRAY, SARRAY:
			err := packArray(typeDesc, dm_build_521.dm_build_406.dm_build_5)
			if err != nil {
				return err
			}

		case PLTYPE_RECORD:
			err := packRecord(typeDesc, dm_build_521.dm_build_406.dm_build_5)
			if err != nil {
				return err
			}

		case CLASS:
			err := packClass(typeDesc, dm_build_521.dm_build_406.dm_build_5)
			if err != nil {
				return err
			}

		}
	}

	return nil
}

func (dm_build_524 *dm_build_497) dm_build_523(dm_build_525 []interface{}) error {
	for i := 0; i < len(dm_build_524.dm_build_499); i++ {

		if dm_build_524.dm_build_499[i].colType == CURSOR {
			dm_build_524.dm_build_406.dm_build_5.Dm_build_1269(ULINT_SIZE)
			dm_build_524.dm_build_406.dm_build_5.Dm_build_1273(dm_build_524.dm_build_499[i].cursorStmt.id)
			continue
		}

		if dm_build_524.dm_build_499[i].ioType == IO_TYPE_OUT {
			continue
		}

		switch dm_build_525[i].(type) {
		case []byte:
			if dataBytes, ok := dm_build_525[i].([]byte); ok {
				if len(dataBytes) > Dm_build_311 {
					return ECGO_DATA_TOO_LONG.throw()
				}
				dm_build_524.dm_build_406.dm_build_5.Dm_build_1311(dataBytes)
			}
		case int:
			if dm_build_525[i] == ParamDataEnum_Null {
				dm_build_524.dm_build_406.dm_build_5.Dm_build_1281(Dm_build_317)
			} else if dm_build_525[i] == ParamDataEnum_OFF_ROW {
				dm_build_524.dm_build_406.dm_build_5.Dm_build_1269(0)
			}
		case lobCtl:
			dm_build_524.dm_build_406.dm_build_5.Dm_build_1281(uint16(Dm_build_316))
			dm_build_524.dm_build_406.dm_build_5.Dm_build_1301(dm_build_525[i].(lobCtl).value)
		default:
			panic("Bind param data failed by invalid param data type: ")
		}
	}

	return nil
}

func (dm_build_527 *dm_build_497) dm_build_396() (interface{}, error) {
	dm_build_528 := execInfo{}
	dm_build_529 := dm_build_527.dm_build_409.dmConn

	dm_build_530 := Dm_build_301

	dm_build_528.retSqlType = dm_build_527.dm_build_406.dm_build_5.Dm_build_1488(dm_build_530)
	dm_build_530 += USINT_SIZE

	dm_build_531 := dm_build_527.dm_build_406.dm_build_5.Dm_build_1488(dm_build_530)
	dm_build_530 += USINT_SIZE

	dm_build_528.updateCount = dm_build_527.dm_build_406.dm_build_5.Dm_build_1494(dm_build_530)
	dm_build_530 += DDWORD_SIZE

	dm_build_532 := dm_build_527.dm_build_406.dm_build_5.Dm_build_1488(dm_build_530)
	dm_build_530 += USINT_SIZE

	dm_build_528.rsUpdatable = dm_build_527.dm_build_406.dm_build_5.Dm_build_1485(dm_build_530) != 0
	dm_build_530 += BYTE_SIZE

	dm_build_533 := dm_build_527.dm_build_406.dm_build_5.Dm_build_1488(dm_build_530)
	dm_build_530 += ULINT_SIZE

	dm_build_528.printLen = dm_build_527.dm_build_406.dm_build_5.Dm_build_1491(dm_build_530)
	dm_build_530 += ULINT_SIZE

	var dm_build_534 int16 = -1
	if dm_build_528.retSqlType == Dm_build_367 || dm_build_528.retSqlType == Dm_build_368 {
		dm_build_528.rowid = 0

		dm_build_528.rsBdta = dm_build_527.dm_build_406.dm_build_5.Dm_build_1485(dm_build_530) == Dm_build_380
		dm_build_530 += BYTE_SIZE

		dm_build_534 = dm_build_527.dm_build_406.dm_build_5.Dm_build_1488(dm_build_530)
		dm_build_530 += USINT_SIZE
		dm_build_530 += 5
	} else {
		dm_build_528.rowid = dm_build_527.dm_build_406.dm_build_5.Dm_build_1494(dm_build_530)
		dm_build_530 += DDWORD_SIZE
	}

	dm_build_528.execId = dm_build_527.dm_build_406.dm_build_5.Dm_build_1491(dm_build_530)
	dm_build_530 += ULINT_SIZE

	dm_build_528.rsCacheOffset = dm_build_527.dm_build_406.dm_build_5.Dm_build_1491(dm_build_530)
	dm_build_530 += ULINT_SIZE

	dm_build_535 := dm_build_527.dm_build_406.dm_build_5.Dm_build_1485(dm_build_530)
	dm_build_530 += BYTE_SIZE
	dm_build_536 := (dm_build_535 & 0x01) == 0x01
	dm_build_537 := (dm_build_535 & 0x02) == 0x02

	dm_build_529.TrxStatus = dm_build_527.dm_build_406.dm_build_5.Dm_build_1491(dm_build_530)
	dm_build_529.setTrxFinish(dm_build_529.TrxStatus)
	dm_build_530 += ULINT_SIZE

	if dm_build_528.printLen > 0 {
		bytes := dm_build_527.dm_build_406.dm_build_5.Dm_build_1368(int(dm_build_528.printLen))
		dm_build_528.printMsg = Dm_build_868.Dm_build_1022(bytes, 0, len(bytes), dm_build_529.getServerEncoding(), dm_build_529)
	}

	if dm_build_532 > 0 {
		dm_build_528.outParamDatas = dm_build_527.dm_build_538(int(dm_build_532))
	}

	switch dm_build_528.retSqlType {
	case Dm_build_369:
		dm_build_529.dmConnector.localTimezone = dm_build_527.dm_build_406.dm_build_5.Dm_build_1344()
	case Dm_build_367:
		dm_build_528.hasResultSet = true
		if dm_build_531 > 0 {
			dm_build_527.dm_build_409.columns = dm_build_527.dm_build_547(int(dm_build_531), dm_build_528.rsBdta)
		}
		dm_build_527.dm_build_557(&dm_build_528, len(dm_build_527.dm_build_409.columns), int(dm_build_533), int(dm_build_534))
	case Dm_build_368:
		if dm_build_531 > 0 || dm_build_533 > 0 {
			dm_build_528.hasResultSet = true
		}
		if dm_build_531 > 0 {
			dm_build_527.dm_build_409.columns = dm_build_527.dm_build_547(int(dm_build_531), dm_build_528.rsBdta)
		}
		dm_build_527.dm_build_557(&dm_build_528, len(dm_build_527.dm_build_409.columns), int(dm_build_533), int(dm_build_534))
	case Dm_build_370:
		dm_build_529.IsoLevel = int32(dm_build_527.dm_build_406.dm_build_5.Dm_build_1344())
		dm_build_529.ReadOnly = dm_build_527.dm_build_406.dm_build_5.Dm_build_1341() == 1
	case Dm_build_363:
		dm_build_529.Schema = dm_build_527.dm_build_406.dm_build_5.Dm_build_1389(dm_build_529.getServerEncoding(), dm_build_529)
	case Dm_build_360:
		dm_build_528.explain = dm_build_527.dm_build_406.dm_build_5.Dm_build_1389(dm_build_529.getServerEncoding(), dm_build_529)

	case Dm_build_364, Dm_build_366, Dm_build_365:
		if dm_build_536 {

			counts := dm_build_527.dm_build_406.dm_build_5.Dm_build_1347()
			rowCounts := make([]int64, counts)
			for i := 0; i < int(counts); i++ {
				rowCounts[i] = dm_build_527.dm_build_406.dm_build_5.Dm_build_1350()
			}
			dm_build_528.updateCounts = rowCounts
		}

		if dm_build_537 {
			rows := dm_build_527.dm_build_406.dm_build_5.Dm_build_1347()

			var lastInsertId int64
			for i := 0; i < int(rows); i++ {
				lastInsertId = dm_build_527.dm_build_406.dm_build_5.Dm_build_1350()
			}
			dm_build_528.lastInsertId = lastInsertId

		} else if dm_build_528.updateCount == 1 {
			dm_build_528.lastInsertId = dm_build_528.rowid
		}

		if dm_build_527.dm_build_408 == EC_BP_WITH_ERROR.ErrCode {
			dm_build_527.dm_build_563(dm_build_528.updateCounts)
		}
	case Dm_build_373:
		len := dm_build_527.dm_build_406.dm_build_5.Dm_build_1359()
		dm_build_529.FormatDate = dm_build_527.dm_build_406.dm_build_5.Dm_build_1384(int(len), dm_build_529.getServerEncoding(), dm_build_529)
	case Dm_build_375:

		len := dm_build_527.dm_build_406.dm_build_5.Dm_build_1359()
		dm_build_529.FormatTimestamp = dm_build_527.dm_build_406.dm_build_5.Dm_build_1384(int(len), dm_build_529.getServerEncoding(), dm_build_529)
	case Dm_build_376:

		len := dm_build_527.dm_build_406.dm_build_5.Dm_build_1359()
		dm_build_529.FormatTimestampTZ = dm_build_527.dm_build_406.dm_build_5.Dm_build_1384(int(len), dm_build_529.getServerEncoding(), dm_build_529)
	case Dm_build_374:
		len := dm_build_527.dm_build_406.dm_build_5.Dm_build_1359()
		dm_build_529.FormatTime = dm_build_527.dm_build_406.dm_build_5.Dm_build_1384(int(len), dm_build_529.getServerEncoding(), dm_build_529)
	case Dm_build_377:
		len := dm_build_527.dm_build_406.dm_build_5.Dm_build_1359()
		dm_build_529.FormatTimeTZ = dm_build_527.dm_build_406.dm_build_5.Dm_build_1384(int(len), dm_build_529.getServerEncoding(), dm_build_529)
	case Dm_build_378:
		dm_build_529.OracleDateLanguage = dm_build_527.dm_build_406.dm_build_5.Dm_build_1359()
	}

	return &dm_build_528, nil
}

func (dm_build_539 *dm_build_497) dm_build_538(dm_build_540 int) [][]byte {
	dm_build_541 := make([]int, dm_build_540)

	dm_build_542 := 0
	for i := 0; i < len(dm_build_539.dm_build_499); i++ {
		if dm_build_539.dm_build_499[i].ioType == IO_TYPE_INOUT || dm_build_539.dm_build_499[i].ioType == IO_TYPE_OUT {
			dm_build_541[dm_build_542] = i
			dm_build_542++
		}
	}

	dm_build_543 := make([][]byte, len(dm_build_539.dm_build_499))
	var dm_build_544 int32
	var dm_build_545 bool
	var dm_build_546 []byte = nil
	for i := 0; i < dm_build_540; i++ {
		dm_build_545 = false
		dm_build_544 = int32(dm_build_539.dm_build_406.dm_build_5.Dm_build_1362())

		if dm_build_544 == int32(Dm_build_317) {
			dm_build_544 = 0
			dm_build_545 = true
		} else if dm_build_544 == int32(Dm_build_318) {
			dm_build_544 = dm_build_539.dm_build_406.dm_build_5.Dm_build_1347()
		}

		if dm_build_545 {
			dm_build_543[dm_build_541[i]] = nil
		} else {
			dm_build_546 = dm_build_539.dm_build_406.dm_build_5.Dm_build_1368(int(dm_build_544))
			dm_build_543[dm_build_541[i]] = dm_build_546
		}
	}

	return dm_build_543
}

func (dm_build_548 *dm_build_497) dm_build_547(dm_build_549 int, dm_build_550 bool) []column {
	dm_build_551 := dm_build_548.dm_build_406.dm_build_6.getServerEncoding()
	var dm_build_552, dm_build_553, dm_build_554, dm_build_555 int16
	dm_build_556 := make([]column, dm_build_549)
	for i := 0; i < dm_build_549; i++ {
		dm_build_556[i].InitColumn()

		dm_build_556[i].colType = dm_build_548.dm_build_406.dm_build_5.Dm_build_1347()

		dm_build_556[i].prec = dm_build_548.dm_build_406.dm_build_5.Dm_build_1347()

		dm_build_556[i].scale = dm_build_548.dm_build_406.dm_build_5.Dm_build_1347()

		dm_build_556[i].nullable = dm_build_548.dm_build_406.dm_build_5.Dm_build_1347() != 0

		itemFlag := dm_build_548.dm_build_406.dm_build_5.Dm_build_1344()
		dm_build_556[i].lob = int(itemFlag)&Dm_build_494 != 0
		dm_build_556[i].identity = int(itemFlag)&Dm_build_493 != 0
		dm_build_556[i].readonly = int(itemFlag)&Dm_build_495 != 0

		dm_build_548.dm_build_406.dm_build_5.Dm_build_1247(4, false, true)

		dm_build_548.dm_build_406.dm_build_5.Dm_build_1247(2, false, true)

		dm_build_552 = dm_build_548.dm_build_406.dm_build_5.Dm_build_1344()

		dm_build_553 = dm_build_548.dm_build_406.dm_build_5.Dm_build_1344()

		dm_build_554 = dm_build_548.dm_build_406.dm_build_5.Dm_build_1344()

		dm_build_555 = dm_build_548.dm_build_406.dm_build_5.Dm_build_1344()
		dm_build_556[i].name = dm_build_548.dm_build_406.dm_build_5.Dm_build_1384(int(dm_build_552), dm_build_551, dm_build_548.dm_build_406.dm_build_6)
		dm_build_556[i].typeName = dm_build_548.dm_build_406.dm_build_5.Dm_build_1384(int(dm_build_553), dm_build_551, dm_build_548.dm_build_406.dm_build_6)
		dm_build_556[i].tableName = dm_build_548.dm_build_406.dm_build_5.Dm_build_1384(int(dm_build_554), dm_build_551, dm_build_548.dm_build_406.dm_build_6)
		dm_build_556[i].schemaName = dm_build_548.dm_build_406.dm_build_5.Dm_build_1384(int(dm_build_555), dm_build_551, dm_build_548.dm_build_406.dm_build_6)

		if dm_build_548.dm_build_409.readBaseColName {
			dm_build_556[i].baseName = dm_build_548.dm_build_406.dm_build_5.Dm_build_1397(dm_build_551, dm_build_548.dm_build_406.dm_build_6)
		}

		if dm_build_556[i].lob {
			dm_build_556[i].lobTabId = dm_build_548.dm_build_406.dm_build_5.Dm_build_1347()
			dm_build_556[i].lobColId = dm_build_548.dm_build_406.dm_build_5.Dm_build_1344()
		}

	}

	for i := 0; i < dm_build_549; i++ {

		if isComplexType(int(dm_build_556[i].colType), int(dm_build_556[i].scale)) {
			strDesc := newTypeDescriptor(dm_build_548.dm_build_406.dm_build_6)
			strDesc.unpack(dm_build_548.dm_build_406.dm_build_5)
			dm_build_556[i].typeDescriptor = strDesc
		}
	}

	return dm_build_556
}

func (dm_build_558 *dm_build_497) dm_build_557(dm_build_559 *execInfo, dm_build_560 int, dm_build_561 int, dm_build_562 int) {
	if dm_build_561 > 0 {
		startOffset := dm_build_558.dm_build_406.dm_build_5.Dm_build_1242()
		if dm_build_559.rsBdta {
			dm_build_559.rsDatas = dm_build_558.dm_build_576(dm_build_558.dm_build_409.columns, dm_build_562)
		} else {
			datas := make([][][]byte, dm_build_561)

			for i := 0; i < dm_build_561; i++ {

				datas[i] = make([][]byte, dm_build_560+1)

				dm_build_558.dm_build_406.dm_build_5.Dm_build_1247(2, false, true)

				datas[i][0] = dm_build_558.dm_build_406.dm_build_5.Dm_build_1368(LINT64_SIZE)

				dm_build_558.dm_build_406.dm_build_5.Dm_build_1247(2*dm_build_560, false, true)

				for j := 1; j < dm_build_560+1; j++ {

					colLen := dm_build_558.dm_build_406.dm_build_5.Dm_build_1362()
					if colLen == Dm_build_321 {
						datas[i][j] = nil
					} else if colLen != Dm_build_322 {
						datas[i][j] = dm_build_558.dm_build_406.dm_build_5.Dm_build_1368(int(colLen))
					} else {
						datas[i][j] = dm_build_558.dm_build_406.dm_build_5.Dm_build_1372()
					}
				}
			}

			dm_build_559.rsDatas = datas
		}
		dm_build_559.rsSizeof = dm_build_558.dm_build_406.dm_build_5.Dm_build_1242() - startOffset
	}

	if dm_build_559.rsCacheOffset > 0 {
		tbCount := dm_build_558.dm_build_406.dm_build_5.Dm_build_1344()

		ids := make([]int32, tbCount)
		tss := make([]int64, tbCount)

		for i := 0; i < int(tbCount); i++ {
			ids[i] = dm_build_558.dm_build_406.dm_build_5.Dm_build_1347()
			tss[i] = dm_build_558.dm_build_406.dm_build_5.Dm_build_1350()
		}

		dm_build_559.tbIds = ids[:]
		dm_build_559.tbTss = tss[:]
	}
}

func (dm_build_564 *dm_build_497) dm_build_563(dm_build_565 []int64) error {

	dm_build_564.dm_build_406.dm_build_5.Dm_build_1247(4, false, true)

	dm_build_566 := dm_build_564.dm_build_406.dm_build_5.Dm_build_1347()

	dm_build_567 := make([]string, 0, 8)
	for i := 0; i < int(dm_build_566); i++ {
		irow := dm_build_564.dm_build_406.dm_build_5.Dm_build_1347()

		dm_build_565[irow] = -3

		code := dm_build_564.dm_build_406.dm_build_5.Dm_build_1347()

		errStr := dm_build_564.dm_build_406.dm_build_5.Dm_build_1397(dm_build_564.dm_build_406.dm_build_6.getServerEncoding(), dm_build_564.dm_build_406.dm_build_6)

		dm_build_567 = append(dm_build_567, "row["+strconv.Itoa(int(irow))+"]:"+strconv.Itoa(int(code))+", "+errStr)
	}

	if len(dm_build_567) > 0 {
		builder := &strings.Builder{}
		for _, str := range dm_build_567 {
			builder.WriteString(util.LINE_SEPARATOR)
			builder.WriteString(str)
		}
		EC_BP_WITH_ERROR.ErrText += builder.String()
		return EC_BP_WITH_ERROR.throw()
	}
	return nil
}

const (
	Dm_build_568 = 0

	Dm_build_569 = Dm_build_568 + ULINT_SIZE

	Dm_build_570 = Dm_build_569 + USINT_SIZE

	Dm_build_571 = Dm_build_570 + ULINT_SIZE

	Dm_build_572 = Dm_build_571 + ULINT_SIZE

	Dm_build_573 = Dm_build_572 + BYTE_SIZE

	Dm_build_574 = -2

	Dm_build_575 = -3
)

func (dm_build_577 *dm_build_497) dm_build_576(dm_build_578 []column, dm_build_579 int) [][][]byte {

	dm_build_580 := dm_build_577.dm_build_406.dm_build_5.Dm_build_1365()

	dm_build_581 := dm_build_577.dm_build_406.dm_build_5.Dm_build_1362()

	var dm_build_582 bool

	if dm_build_579 >= 0 && int(dm_build_581) == len(dm_build_578)+1 {
		dm_build_582 = true
	} else {
		dm_build_582 = false
	}

	dm_build_577.dm_build_406.dm_build_5.Dm_build_1247(ULINT_SIZE, false, true)

	dm_build_577.dm_build_406.dm_build_5.Dm_build_1247(ULINT_SIZE, false, true)

	dm_build_577.dm_build_406.dm_build_5.Dm_build_1247(BYTE_SIZE, false, true)

	dm_build_583 := make([]uint16, dm_build_581)
	for icol := 0; icol < int(dm_build_581); icol++ {
		dm_build_583[icol] = dm_build_577.dm_build_406.dm_build_5.Dm_build_1362()
	}

	dm_build_584 := make([]uint32, dm_build_581)
	dm_build_585 := make([][][]byte, dm_build_580)

	for i := uint32(0); i < dm_build_580; i++ {
		dm_build_585[i] = make([][]byte, len(dm_build_578)+1)
	}

	for icol := 0; icol < int(dm_build_581); icol++ {
		dm_build_584[icol] = dm_build_577.dm_build_406.dm_build_5.Dm_build_1365()
	}

	for icol := 0; icol < int(dm_build_581); icol++ {

		dataCol := icol + 1
		if dm_build_582 && icol == dm_build_579 {
			dataCol = 0
		} else if dm_build_582 && icol > dm_build_579 {
			dataCol = icol
		}

		allNotNull := dm_build_577.dm_build_406.dm_build_5.Dm_build_1347() == 1
		var isNull []bool = nil
		if !allNotNull {
			isNull = make([]bool, dm_build_580)
			for irow := uint32(0); irow < dm_build_580; irow++ {
				isNull[irow] = dm_build_577.dm_build_406.dm_build_5.Dm_build_1341() == 0
			}
		}

		for irow := uint32(0); irow < dm_build_580; irow++ {
			if allNotNull || !isNull[irow] {
				dm_build_585[irow][dataCol] = dm_build_577.dm_build_586(int(dm_build_583[icol]))
			}
		}
	}

	if !dm_build_582 && dm_build_579 >= 0 {
		for irow := uint32(0); irow < dm_build_580; irow++ {
			dm_build_585[irow][0] = dm_build_585[irow][dm_build_579+1]
		}
	}

	return dm_build_585
}

func (dm_build_587 *dm_build_497) dm_build_586(dm_build_588 int) []byte {

	dm_build_589 := dm_build_587.dm_build_592(dm_build_588)

	dm_build_590 := int32(0)
	if dm_build_589 == Dm_build_574 {
		dm_build_590 = dm_build_587.dm_build_406.dm_build_5.Dm_build_1347()
		dm_build_589 = int(dm_build_587.dm_build_406.dm_build_5.Dm_build_1347())
	} else if dm_build_589 == Dm_build_575 {
		dm_build_589 = int(dm_build_587.dm_build_406.dm_build_5.Dm_build_1347())
	}

	dm_build_591 := dm_build_587.dm_build_406.dm_build_5.Dm_build_1368(dm_build_589 + int(dm_build_590))
	if dm_build_590 == 0 {
		return dm_build_591
	}

	for i := dm_build_589; i < len(dm_build_591); i++ {
		dm_build_591[i] = ' '
	}
	return dm_build_591
}

func (dm_build_593 *dm_build_497) dm_build_592(dm_build_594 int) int {

	dm_build_595 := 0
	switch dm_build_594 {
	case INT:
	case BIT:
	case TINYINT:
	case SMALLINT:
	case BOOLEAN:
	case NULL:
		dm_build_595 = 4

	case BIGINT:

		dm_build_595 = 8

	case CHAR:
	case VARCHAR2:
	case VARCHAR:
	case BINARY:
	case VARBINARY:
	case BLOB:
	case CLOB:
		dm_build_595 = Dm_build_574

	case DECIMAL:
		dm_build_595 = Dm_build_575

	case REAL:
		dm_build_595 = 4

	case DOUBLE:
		dm_build_595 = 8

	case DATE:
	case TIME:
	case DATETIME:
	case TIME_TZ:
	case DATETIME_TZ:
		dm_build_595 = 12

	case INTERVAL_YM:
		dm_build_595 = 12

	case INTERVAL_DT:
		dm_build_595 = 24

	default:
		panic(ECGO_INVALID_COLUMN_TYPE)
	}
	return dm_build_595
}

const (
	Dm_build_596 = Dm_build_301

	Dm_build_597 = Dm_build_596 + DDWORD_SIZE

	Dm_build_598 = Dm_build_597 + LINT64_SIZE

	Dm_build_599 = Dm_build_598 + USINT_SIZE

	Dm_build_600 = Dm_build_301

	Dm_build_601 = Dm_build_600 + DDWORD_SIZE
)

type dm_build_602 struct {
	dm_build_497
	dm_build_603 *innerRows
	dm_build_604 int64
	dm_build_605 int64
}

func dm_build_606(dm_build_607 *dm_build_2, dm_build_608 *innerRows, dm_build_609 int64, dm_build_610 int64) *dm_build_602 {
	dm_build_611 := new(dm_build_602)
	dm_build_611.dm_build_414(dm_build_607, Dm_build_279, dm_build_608.dmStmt)
	dm_build_611.dm_build_603 = dm_build_608
	dm_build_611.dm_build_604 = dm_build_609
	dm_build_611.dm_build_605 = dm_build_610
	return dm_build_611
}

func (dm_build_613 *dm_build_602) dm_build_392() error {

	dm_build_613.dm_build_406.dm_build_5.Dm_build_1417(Dm_build_596, dm_build_613.dm_build_604)

	dm_build_613.dm_build_406.dm_build_5.Dm_build_1417(Dm_build_597, dm_build_613.dm_build_605)

	dm_build_613.dm_build_406.dm_build_5.Dm_build_1409(Dm_build_598, dm_build_613.dm_build_603.id)

	dm_build_614 := dm_build_613.dm_build_603.dmStmt.dmConn.dmConnector.bufPrefetch
	var dm_build_615 int32
	if dm_build_613.dm_build_603.sizeOfRow != 0 && dm_build_613.dm_build_603.fetchSize != 0 {
		if dm_build_613.dm_build_603.sizeOfRow*dm_build_613.dm_build_603.fetchSize > int(INT32_MAX) {
			dm_build_615 = INT32_MAX
		} else {
			dm_build_615 = int32(dm_build_613.dm_build_603.sizeOfRow * dm_build_613.dm_build_603.fetchSize)
		}

		if dm_build_615 < Dm_build_333 {
			dm_build_614 = int(Dm_build_333)
		} else if dm_build_615 > Dm_build_334 {
			dm_build_614 = int(Dm_build_334)
		} else {
			dm_build_614 = int(dm_build_615)
		}

		dm_build_613.dm_build_406.dm_build_5.Dm_build_1413(Dm_build_599, int32(dm_build_614))
	}

	return nil
}

func (dm_build_617 *dm_build_602) dm_build_396() (interface{}, error) {
	dm_build_618 := execInfo{}
	dm_build_618.rsBdta = dm_build_617.dm_build_603.isBdta

	dm_build_618.updateCount = dm_build_617.dm_build_406.dm_build_5.Dm_build_1494(Dm_build_600)
	dm_build_619 := dm_build_617.dm_build_406.dm_build_5.Dm_build_1491(Dm_build_601)

	dm_build_617.dm_build_557(&dm_build_618, len(dm_build_617.dm_build_603.columns), int(dm_build_619), -1)

	return &dm_build_618, nil
}

type dm_build_620 struct {
	dm_build_405
	dm_build_621 *lob
	dm_build_622 int
	dm_build_623 int
}

func dm_build_624(dm_build_625 *dm_build_2, dm_build_626 *lob, dm_build_627 int, dm_build_628 int) *dm_build_620 {
	dm_build_629 := new(dm_build_620)
	dm_build_629.dm_build_410(dm_build_625, Dm_build_292)
	dm_build_629.dm_build_621 = dm_build_626
	dm_build_629.dm_build_622 = dm_build_627
	dm_build_629.dm_build_623 = dm_build_628
	return dm_build_629
}

func (dm_build_631 *dm_build_620) dm_build_392() error {

	dm_build_631.dm_build_406.dm_build_5.Dm_build_1265(byte(dm_build_631.dm_build_621.lobFlag))

	dm_build_631.dm_build_406.dm_build_5.Dm_build_1273(dm_build_631.dm_build_621.tabId)

	dm_build_631.dm_build_406.dm_build_5.Dm_build_1269(dm_build_631.dm_build_621.colId)

	dm_build_631.dm_build_406.dm_build_5.Dm_build_1289(uint64(dm_build_631.dm_build_621.blobId))

	dm_build_631.dm_build_406.dm_build_5.Dm_build_1269(dm_build_631.dm_build_621.groupId)

	dm_build_631.dm_build_406.dm_build_5.Dm_build_1269(dm_build_631.dm_build_621.fileId)

	dm_build_631.dm_build_406.dm_build_5.Dm_build_1273(dm_build_631.dm_build_621.pageNo)

	dm_build_631.dm_build_406.dm_build_5.Dm_build_1269(dm_build_631.dm_build_621.curFileId)

	dm_build_631.dm_build_406.dm_build_5.Dm_build_1273(dm_build_631.dm_build_621.curPageNo)

	dm_build_631.dm_build_406.dm_build_5.Dm_build_1273(dm_build_631.dm_build_621.totalOffset)

	dm_build_631.dm_build_406.dm_build_5.Dm_build_1273(int32(dm_build_631.dm_build_622))

	dm_build_631.dm_build_406.dm_build_5.Dm_build_1273(int32(dm_build_631.dm_build_623))

	if dm_build_631.dm_build_406.dm_build_6.NewLobFlag {
		dm_build_631.dm_build_406.dm_build_5.Dm_build_1289(uint64(dm_build_631.dm_build_621.rowId))
		dm_build_631.dm_build_406.dm_build_5.Dm_build_1269(dm_build_631.dm_build_621.exGroupId)
		dm_build_631.dm_build_406.dm_build_5.Dm_build_1269(dm_build_631.dm_build_621.exFileId)
		dm_build_631.dm_build_406.dm_build_5.Dm_build_1273(dm_build_631.dm_build_621.exPageNo)
	}

	return nil
}

func (dm_build_633 *dm_build_620) dm_build_396() (interface{}, error) {

	dm_build_633.dm_build_621.readOver = dm_build_633.dm_build_406.dm_build_5.Dm_build_1341() == 1
	var dm_build_634 = dm_build_633.dm_build_406.dm_build_5.Dm_build_1365()
	if dm_build_634 <= 0 {
		return []byte(nil), nil
	}
	dm_build_633.dm_build_621.curFileId = dm_build_633.dm_build_406.dm_build_5.Dm_build_1344()
	dm_build_633.dm_build_621.curPageNo = dm_build_633.dm_build_406.dm_build_5.Dm_build_1347()
	dm_build_633.dm_build_621.totalOffset = dm_build_633.dm_build_406.dm_build_5.Dm_build_1347()

	return dm_build_633.dm_build_406.dm_build_5.Dm_build_1378(int(dm_build_634)), nil
}

type dm_build_635 struct {
	dm_build_405
	dm_build_636 *lob
}

func dm_build_637(dm_build_638 *dm_build_2, dm_build_639 *lob) *dm_build_635 {
	dm_build_640 := new(dm_build_635)
	dm_build_640.dm_build_410(dm_build_638, Dm_build_289)
	dm_build_640.dm_build_636 = dm_build_639
	return dm_build_640
}

func (dm_build_642 *dm_build_635) dm_build_392() error {

	dm_build_642.dm_build_406.dm_build_5.Dm_build_1265(byte(dm_build_642.dm_build_636.lobFlag))

	dm_build_642.dm_build_406.dm_build_5.Dm_build_1289(uint64(dm_build_642.dm_build_636.blobId))

	dm_build_642.dm_build_406.dm_build_5.Dm_build_1269(dm_build_642.dm_build_636.groupId)

	dm_build_642.dm_build_406.dm_build_5.Dm_build_1269(dm_build_642.dm_build_636.fileId)

	dm_build_642.dm_build_406.dm_build_5.Dm_build_1273(dm_build_642.dm_build_636.pageNo)

	if dm_build_642.dm_build_406.dm_build_6.NewLobFlag {
		dm_build_642.dm_build_406.dm_build_5.Dm_build_1273(dm_build_642.dm_build_636.tabId)
		dm_build_642.dm_build_406.dm_build_5.Dm_build_1269(dm_build_642.dm_build_636.colId)
		dm_build_642.dm_build_406.dm_build_5.Dm_build_1289(uint64(dm_build_642.dm_build_636.rowId))

		dm_build_642.dm_build_406.dm_build_5.Dm_build_1269(dm_build_642.dm_build_636.exGroupId)
		dm_build_642.dm_build_406.dm_build_5.Dm_build_1269(dm_build_642.dm_build_636.exFileId)
		dm_build_642.dm_build_406.dm_build_5.Dm_build_1273(dm_build_642.dm_build_636.exPageNo)
	}

	return nil
}

func (dm_build_644 *dm_build_635) dm_build_396() (interface{}, error) {

	return int64(dm_build_644.dm_build_406.dm_build_5.Dm_build_1347()), nil
}

type dm_build_645 struct {
	dm_build_405
	dm_build_646 *lob
	dm_build_647 int
}

func dm_build_648(dm_build_649 *dm_build_2, dm_build_650 *lob, dm_build_651 int) *dm_build_645 {
	dm_build_652 := new(dm_build_645)
	dm_build_652.dm_build_410(dm_build_649, Dm_build_291)
	dm_build_652.dm_build_646 = dm_build_650
	dm_build_652.dm_build_647 = dm_build_651
	return dm_build_652
}

func (dm_build_654 *dm_build_645) dm_build_392() error {

	dm_build_654.dm_build_406.dm_build_5.Dm_build_1265(byte(dm_build_654.dm_build_646.lobFlag))

	dm_build_654.dm_build_406.dm_build_5.Dm_build_1289(uint64(dm_build_654.dm_build_646.blobId))

	dm_build_654.dm_build_406.dm_build_5.Dm_build_1269(dm_build_654.dm_build_646.groupId)

	dm_build_654.dm_build_406.dm_build_5.Dm_build_1269(dm_build_654.dm_build_646.fileId)

	dm_build_654.dm_build_406.dm_build_5.Dm_build_1273(dm_build_654.dm_build_646.pageNo)

	dm_build_654.dm_build_406.dm_build_5.Dm_build_1273(dm_build_654.dm_build_646.tabId)
	dm_build_654.dm_build_406.dm_build_5.Dm_build_1269(dm_build_654.dm_build_646.colId)
	dm_build_654.dm_build_406.dm_build_5.Dm_build_1289(uint64(dm_build_654.dm_build_646.rowId))
	dm_build_654.dm_build_406.dm_build_5.Dm_build_1301(Dm_build_868.Dm_build_1067(uint32(dm_build_654.dm_build_647)))

	if dm_build_654.dm_build_406.dm_build_6.NewLobFlag {
		dm_build_654.dm_build_406.dm_build_5.Dm_build_1269(dm_build_654.dm_build_646.exGroupId)
		dm_build_654.dm_build_406.dm_build_5.Dm_build_1269(dm_build_654.dm_build_646.exFileId)
		dm_build_654.dm_build_406.dm_build_5.Dm_build_1273(dm_build_654.dm_build_646.exPageNo)
	}
	return nil
}

func (dm_build_656 *dm_build_645) dm_build_396() (interface{}, error) {

	dm_build_657 := dm_build_656.dm_build_406.dm_build_5.Dm_build_1365()
	dm_build_656.dm_build_646.blobId = dm_build_656.dm_build_406.dm_build_5.Dm_build_1350()
	dm_build_656.dm_build_646.resetCurrentInfo()
	return int64(dm_build_657), nil
}

const (
	Dm_build_658 = Dm_build_301

	Dm_build_659 = Dm_build_658 + ULINT_SIZE

	Dm_build_660 = Dm_build_659 + ULINT_SIZE

	Dm_build_661 = Dm_build_660 + ULINT_SIZE

	Dm_build_662 = Dm_build_661 + BYTE_SIZE

	Dm_build_663 = Dm_build_662 + USINT_SIZE

	Dm_build_664 = Dm_build_663 + ULINT_SIZE

	Dm_build_665 = Dm_build_664 + BYTE_SIZE

	Dm_build_666 = Dm_build_665 + BYTE_SIZE

	Dm_build_667 = Dm_build_666 + BYTE_SIZE

	Dm_build_668 = Dm_build_301

	Dm_build_669 = Dm_build_668 + ULINT_SIZE

	Dm_build_670 = Dm_build_669 + ULINT_SIZE

	Dm_build_671 = Dm_build_670 + BYTE_SIZE

	Dm_build_672 = Dm_build_671 + ULINT_SIZE

	Dm_build_673 = Dm_build_672 + BYTE_SIZE

	Dm_build_674 = Dm_build_673 + BYTE_SIZE

	Dm_build_675 = Dm_build_674 + USINT_SIZE

	Dm_build_676 = Dm_build_675 + USINT_SIZE

	Dm_build_677 = Dm_build_676 + BYTE_SIZE

	Dm_build_678 = Dm_build_677 + USINT_SIZE

	Dm_build_679 = Dm_build_678 + BYTE_SIZE

	Dm_build_680 = Dm_build_679 + BYTE_SIZE

	Dm_build_681 = Dm_build_680 + ULINT_SIZE
)

type dm_build_682 struct {
	dm_build_405

	dm_build_683 *DmConnection

	dm_build_684 bool
}

func dm_build_685(dm_build_686 *dm_build_2) *dm_build_682 {
	dm_build_687 := new(dm_build_682)
	dm_build_687.dm_build_410(dm_build_686, Dm_build_273)
	dm_build_687.dm_build_683 = dm_build_686.dm_build_6
	return dm_build_687
}

func (dm_build_689 *dm_build_682) dm_build_392() error {

	if dm_build_689.dm_build_683.dmConnector.newClientType {
		dm_build_689.dm_build_406.dm_build_5.Dm_build_1413(Dm_build_658, Dm_build_313)
	} else {
		dm_build_689.dm_build_406.dm_build_5.Dm_build_1413(Dm_build_658, Dm_build_312)
	}

	dm_build_689.dm_build_406.dm_build_5.Dm_build_1413(Dm_build_659, g2dbIsoLevel(dm_build_689.dm_build_683.IsoLevel))
	dm_build_689.dm_build_406.dm_build_5.Dm_build_1413(Dm_build_660, int32(Locale))
	dm_build_689.dm_build_406.dm_build_5.Dm_build_1409(Dm_build_662, dm_build_689.dm_build_683.dmConnector.localTimezone)

	if dm_build_689.dm_build_683.ReadOnly {
		dm_build_689.dm_build_406.dm_build_5.Dm_build_1405(Dm_build_661, Dm_build_336)
	} else {
		dm_build_689.dm_build_406.dm_build_5.Dm_build_1405(Dm_build_661, Dm_build_335)
	}

	dm_build_689.dm_build_406.dm_build_5.Dm_build_1413(Dm_build_663, int32(dm_build_689.dm_build_683.dmConnector.sessionTimeout))

	if dm_build_689.dm_build_683.dmConnector.mppLocal {
		dm_build_689.dm_build_406.dm_build_5.Dm_build_1405(Dm_build_664, 1)
	} else {
		dm_build_689.dm_build_406.dm_build_5.Dm_build_1405(Dm_build_664, 0)
	}

	if dm_build_689.dm_build_683.dmConnector.rwSeparate {
		dm_build_689.dm_build_406.dm_build_5.Dm_build_1405(Dm_build_665, 1)
	} else {
		dm_build_689.dm_build_406.dm_build_5.Dm_build_1405(Dm_build_665, 0)
	}

	if dm_build_689.dm_build_683.NewLobFlag {
		dm_build_689.dm_build_406.dm_build_5.Dm_build_1405(Dm_build_666, 1)
	} else {
		dm_build_689.dm_build_406.dm_build_5.Dm_build_1405(Dm_build_666, 0)
	}

	dm_build_689.dm_build_406.dm_build_5.Dm_build_1405(Dm_build_667, dm_build_689.dm_build_683.dmConnector.osAuthType)

	dm_build_690 := dm_build_689.dm_build_683.getServerEncoding()

	if dm_build_689.dm_build_406.dm_build_11 != "" {

	}

	dm_build_691 := Dm_build_868.Dm_build_1078(dm_build_689.dm_build_683.dmConnector.user, dm_build_690, dm_build_689.dm_build_406.dm_build_6)
	dm_build_692 := Dm_build_868.Dm_build_1078(dm_build_689.dm_build_683.dmConnector.password, dm_build_690, dm_build_689.dm_build_406.dm_build_6)
	if len(dm_build_691) > Dm_build_309 {
		return ECGO_USERNAME_TOO_LONG.throw()
	}
	if len(dm_build_692) > Dm_build_309 {
		return ECGO_PASSWORD_TOO_LONG.throw()
	}

	if dm_build_689.dm_build_406.dm_build_8 && dm_build_689.dm_build_683.dmConnector.loginCertificate != "" {

	} else if dm_build_689.dm_build_406.dm_build_8 {
		dm_build_691 = dm_build_689.dm_build_406.dm_build_7.Encrypt(dm_build_691, false)
		dm_build_692 = dm_build_689.dm_build_406.dm_build_7.Encrypt(dm_build_692, false)
	}

	dm_build_689.dm_build_406.dm_build_5.Dm_build_1305(dm_build_691)
	dm_build_689.dm_build_406.dm_build_5.Dm_build_1305(dm_build_692)

	dm_build_689.dm_build_406.dm_build_5.Dm_build_1317(dm_build_689.dm_build_683.dmConnector.appName, dm_build_690, dm_build_689.dm_build_406.dm_build_6)
	dm_build_689.dm_build_406.dm_build_5.Dm_build_1317(dm_build_689.dm_build_683.dmConnector.osName, dm_build_690, dm_build_689.dm_build_406.dm_build_6)

	if hostName, err := os.Hostname(); err != nil {
		dm_build_689.dm_build_406.dm_build_5.Dm_build_1317(hostName, dm_build_690, dm_build_689.dm_build_406.dm_build_6)
	} else {
		dm_build_689.dm_build_406.dm_build_5.Dm_build_1317("", dm_build_690, dm_build_689.dm_build_406.dm_build_6)
	}

	if dm_build_689.dm_build_683.dmConnector.rwStandby {
		dm_build_689.dm_build_406.dm_build_5.Dm_build_1265(1)
	} else {
		dm_build_689.dm_build_406.dm_build_5.Dm_build_1265(0)
	}

	return nil
}

func (dm_build_694 *dm_build_682) dm_build_396() (interface{}, error) {

	dm_build_694.dm_build_683.MaxRowSize = dm_build_694.dm_build_406.dm_build_5.Dm_build_1491(Dm_build_668)
	dm_build_694.dm_build_683.DDLAutoCommit = dm_build_694.dm_build_406.dm_build_5.Dm_build_1485(Dm_build_670) == 1
	dm_build_694.dm_build_683.IsoLevel = dm_build_694.dm_build_406.dm_build_5.Dm_build_1491(Dm_build_671)
	dm_build_694.dm_build_683.dmConnector.caseSensitive = dm_build_694.dm_build_406.dm_build_5.Dm_build_1485(Dm_build_672) == 1
	dm_build_694.dm_build_683.BackslashEscape = dm_build_694.dm_build_406.dm_build_5.Dm_build_1485(Dm_build_673) == 1
	dm_build_694.dm_build_683.SvrStat = dm_build_694.dm_build_406.dm_build_5.Dm_build_1488(Dm_build_675)
	dm_build_694.dm_build_683.SvrMode = dm_build_694.dm_build_406.dm_build_5.Dm_build_1488(Dm_build_674)
	dm_build_694.dm_build_683.ConstParaOpt = dm_build_694.dm_build_406.dm_build_5.Dm_build_1485(Dm_build_676) == 1
	dm_build_694.dm_build_683.DbTimezone = dm_build_694.dm_build_406.dm_build_5.Dm_build_1488(Dm_build_677)
	dm_build_694.dm_build_683.NewLobFlag = dm_build_694.dm_build_406.dm_build_5.Dm_build_1485(Dm_build_679) == 1

	if dm_build_694.dm_build_683.dmConnector.bufPrefetch == 0 {
		dm_build_694.dm_build_683.dmConnector.bufPrefetch = int(dm_build_694.dm_build_406.dm_build_5.Dm_build_1491(Dm_build_680))
	}

	dm_build_694.dm_build_683.LifeTimeRemainder = dm_build_694.dm_build_406.dm_build_5.Dm_build_1488(Dm_build_681)

	dm_build_695 := dm_build_694.dm_build_683.getServerEncoding()

	dm_build_694.dm_build_683.InstanceName = dm_build_694.dm_build_406.dm_build_5.Dm_build_1389(dm_build_695, dm_build_694.dm_build_406.dm_build_6)
	dm_build_694.dm_build_683.Schema = dm_build_694.dm_build_406.dm_build_5.Dm_build_1389(dm_build_695, dm_build_694.dm_build_406.dm_build_6)
	dm_build_694.dm_build_683.LastLoginIP = dm_build_694.dm_build_406.dm_build_5.Dm_build_1389(dm_build_695, dm_build_694.dm_build_406.dm_build_6)
	dm_build_694.dm_build_683.LastLoginTime = dm_build_694.dm_build_406.dm_build_5.Dm_build_1389(dm_build_695, dm_build_694.dm_build_406.dm_build_6)
	dm_build_694.dm_build_683.FailedAttempts = dm_build_694.dm_build_406.dm_build_5.Dm_build_1347()
	dm_build_694.dm_build_683.LoginWarningID = dm_build_694.dm_build_406.dm_build_5.Dm_build_1347()
	dm_build_694.dm_build_683.GraceTimeRemainder = dm_build_694.dm_build_406.dm_build_5.Dm_build_1347()
	dm_build_694.dm_build_683.Guid = dm_build_694.dm_build_406.dm_build_5.Dm_build_1389(dm_build_695, dm_build_694.dm_build_406.dm_build_6)
	dm_build_694.dm_build_683.DbName = dm_build_694.dm_build_406.dm_build_5.Dm_build_1389(dm_build_695, dm_build_694.dm_build_406.dm_build_6)

	if dm_build_694.dm_build_406.dm_build_5.Dm_build_1485(Dm_build_678) == 1 {
		dm_build_694.dm_build_683.StandbyHost = dm_build_694.dm_build_406.dm_build_5.Dm_build_1389(dm_build_695, dm_build_694.dm_build_406.dm_build_6)
		dm_build_694.dm_build_683.StandbyPort = dm_build_694.dm_build_406.dm_build_5.Dm_build_1347()
		dm_build_694.dm_build_683.StandbyCount = dm_build_694.dm_build_406.dm_build_5.Dm_build_1362()
	}

	if dm_build_694.dm_build_406.dm_build_5.Dm_build_1244(false) > 0 {
		dm_build_694.dm_build_683.SessionID = dm_build_694.dm_build_406.dm_build_5.Dm_build_1350()
	}

	if dm_build_694.dm_build_406.dm_build_5.Dm_build_1244(false) > 0 {
		if dm_build_694.dm_build_406.dm_build_5.Dm_build_1341() == 1 {

			dm_build_694.dm_build_683.FormatDate = "DD-MON-YY"

			dm_build_694.dm_build_683.FormatTime = "HH12.MI.SS.FF6 AM"

			dm_build_694.dm_build_683.FormatTimestamp = "DD-MON-YY HH12.MI.SS.FF6 AM"

			dm_build_694.dm_build_683.FormatTimestampTZ = "DD-MON-YY HH12.MI.SS.FF6 AM +TZH:TZM"

			dm_build_694.dm_build_683.FormatTimeTZ = "HH12.MI.SS.FF6 AM +TZH:TZM"
		}
	}

	if dm_build_694.dm_build_406.dm_build_5.Dm_build_1244(false) > 0 {

		format := dm_build_694.dm_build_406.dm_build_5.Dm_build_1393(dm_build_695, dm_build_694.dm_build_406.dm_build_6)
		if format != "" {
			dm_build_694.dm_build_683.FormatDate = format
		}
		format = dm_build_694.dm_build_406.dm_build_5.Dm_build_1393(dm_build_695, dm_build_694.dm_build_406.dm_build_6)
		if format != "" {
			dm_build_694.dm_build_683.FormatTime = format
		}
		format = dm_build_694.dm_build_406.dm_build_5.Dm_build_1393(dm_build_695, dm_build_694.dm_build_406.dm_build_6)
		if format != "" {
			dm_build_694.dm_build_683.FormatTimestamp = format
		}
		format = dm_build_694.dm_build_406.dm_build_5.Dm_build_1393(dm_build_695, dm_build_694.dm_build_406.dm_build_6)
		if format != "" {
			dm_build_694.dm_build_683.FormatTimestampTZ = format
		}
		format = dm_build_694.dm_build_406.dm_build_5.Dm_build_1393(dm_build_695, dm_build_694.dm_build_406.dm_build_6)
		if format != "" {
			dm_build_694.dm_build_683.FormatTimeTZ = format
		}
	}

	return nil, nil
}

const (
	Dm_build_696 = Dm_build_301
)

type dm_build_697 struct {
	dm_build_497
	dm_build_698 int16
}

func dm_build_699(dm_build_700 *dm_build_2, dm_build_701 *DmStatement, dm_build_702 int16) *dm_build_697 {
	dm_build_703 := new(dm_build_697)
	dm_build_703.dm_build_414(dm_build_700, Dm_build_293, dm_build_701)
	dm_build_703.dm_build_698 = dm_build_702
	return dm_build_703
}

func (dm_build_705 *dm_build_697) dm_build_392() error {
	dm_build_705.dm_build_406.dm_build_5.Dm_build_1409(Dm_build_696, dm_build_705.dm_build_698)
	return nil
}

func (dm_build_707 *dm_build_697) dm_build_396() (interface{}, error) {
	return dm_build_707.dm_build_497.dm_build_396()
}

const (
	Dm_build_708 = Dm_build_301
	Dm_build_709 = Dm_build_708 + USINT_SIZE
)

type dm_build_710 struct {
	dm_build_497
	dm_build_711 []parameter
}

func dm_build_712(dm_build_713 *dm_build_2, dm_build_714 *DmStatement, dm_build_715 []parameter) *dm_build_710 {
	dm_build_716 := new(dm_build_710)
	dm_build_716.dm_build_414(dm_build_713, Dm_build_297, dm_build_714)
	dm_build_716.dm_build_711 = dm_build_715
	return dm_build_716
}

func (dm_build_718 *dm_build_710) dm_build_392() error {

	if dm_build_718.dm_build_711 == nil {
		dm_build_718.dm_build_406.dm_build_5.Dm_build_1409(Dm_build_708, 0)
	} else {
		dm_build_718.dm_build_406.dm_build_5.Dm_build_1409(Dm_build_708, int16(len(dm_build_718.dm_build_711)))
	}

	dm_build_718.dm_build_406.dm_build_5.Dm_build_1405(Dm_build_709, 0)

	return dm_build_718.dm_build_520(dm_build_718.dm_build_711)
}

type dm_build_719 struct {
	dm_build_497
	dm_build_720 bool
	dm_build_721 int16
}

func dm_build_722(dm_build_723 *dm_build_2, dm_build_724 *DmStatement, dm_build_725 bool, dm_build_726 int16) *dm_build_719 {
	dm_build_727 := new(dm_build_719)
	dm_build_727.dm_build_414(dm_build_723, Dm_build_277, dm_build_724)
	dm_build_727.dm_build_720 = dm_build_725
	dm_build_727.dm_build_721 = dm_build_726
	return dm_build_727
}

func (dm_build_729 *dm_build_719) dm_build_392() error {

	dm_build_730 := Dm_build_301

	if dm_build_729.dm_build_406.dm_build_6.autoCommit {
		dm_build_730 += dm_build_729.dm_build_406.dm_build_5.Dm_build_1405(dm_build_730, 1)
	} else {
		dm_build_730 += dm_build_729.dm_build_406.dm_build_5.Dm_build_1405(dm_build_730, 0)
	}

	if dm_build_729.dm_build_720 {
		dm_build_730 += dm_build_729.dm_build_406.dm_build_5.Dm_build_1405(dm_build_730, 1)
	} else {
		dm_build_730 += dm_build_729.dm_build_406.dm_build_5.Dm_build_1405(dm_build_730, 0)
	}

	dm_build_730 += dm_build_729.dm_build_406.dm_build_5.Dm_build_1405(dm_build_730, 0)

	dm_build_730 += dm_build_729.dm_build_406.dm_build_5.Dm_build_1405(dm_build_730, 1)

	if dm_build_729.dm_build_406.dm_build_6.CompatibleOracle() {
		dm_build_730 += dm_build_729.dm_build_406.dm_build_5.Dm_build_1405(dm_build_730, 0)
	} else {
		dm_build_730 += dm_build_729.dm_build_406.dm_build_5.Dm_build_1405(dm_build_730, 1)
	}

	dm_build_730 += dm_build_729.dm_build_406.dm_build_5.Dm_build_1409(dm_build_730, dm_build_729.dm_build_721)

	if dm_build_729.dm_build_409.maxRows <= 0 || dm_build_729.dm_build_406.dm_build_6.dmConnector.enRsCache {
		dm_build_730 += dm_build_729.dm_build_406.dm_build_5.Dm_build_1417(dm_build_730, INT64_MAX)
	} else {
		dm_build_730 += dm_build_729.dm_build_406.dm_build_5.Dm_build_1417(dm_build_730, dm_build_729.dm_build_409.maxRows)
	}

	if dm_build_729.dm_build_406.dm_build_6.dmConnector.isBdtaRS {
		dm_build_730 += dm_build_729.dm_build_406.dm_build_5.Dm_build_1405(dm_build_730, Dm_build_380)
	} else {
		dm_build_730 += dm_build_729.dm_build_406.dm_build_5.Dm_build_1405(dm_build_730, Dm_build_379)
	}

	dm_build_730 += dm_build_729.dm_build_406.dm_build_5.Dm_build_1409(dm_build_730, 0)

	dm_build_730 += dm_build_729.dm_build_406.dm_build_5.Dm_build_1405(dm_build_730, 1)

	dm_build_730 += dm_build_729.dm_build_406.dm_build_5.Dm_build_1405(dm_build_730, 0)

	dm_build_730 += dm_build_729.dm_build_406.dm_build_5.Dm_build_1405(dm_build_730, 0)

	dm_build_730 += dm_build_729.dm_build_406.dm_build_5.Dm_build_1413(dm_build_730, dm_build_729.dm_build_409.queryTimeout)

	dm_build_729.dm_build_406.dm_build_5.Dm_build_1335(dm_build_729.dm_build_409.nativeSql, dm_build_729.dm_build_406.dm_build_6.getServerEncoding(), dm_build_729.dm_build_406.dm_build_6)

	return nil
}

func (dm_build_732 *dm_build_719) dm_build_396() (interface{}, error) {

	if dm_build_732.dm_build_720 {
		return dm_build_732.dm_build_497.dm_build_396()
	}

	dm_build_733 := NewExceInfo()
	dm_build_734 := Dm_build_301

	dm_build_733.retSqlType = dm_build_732.dm_build_406.dm_build_5.Dm_build_1488(dm_build_734)
	dm_build_734 += USINT_SIZE

	dm_build_735 := dm_build_732.dm_build_406.dm_build_5.Dm_build_1488(dm_build_734)
	dm_build_734 += USINT_SIZE

	dm_build_736 := dm_build_732.dm_build_406.dm_build_5.Dm_build_1488(dm_build_734)
	dm_build_734 += USINT_SIZE

	dm_build_732.dm_build_406.dm_build_5.Dm_build_1494(dm_build_734)
	dm_build_734 += DDWORD_SIZE

	dm_build_732.dm_build_406.dm_build_6.TrxStatus = dm_build_732.dm_build_406.dm_build_5.Dm_build_1491(dm_build_734)
	dm_build_734 += ULINT_SIZE

	if dm_build_735 > 0 {
		dm_build_732.dm_build_409.params = dm_build_732.dm_build_737(int(dm_build_735))
		dm_build_732.dm_build_409.paramCount = dm_build_735
	} else {
		dm_build_732.dm_build_409.params = make([]parameter, 0)
		dm_build_732.dm_build_409.paramCount = 0
	}

	if dm_build_736 > 0 {
		dm_build_732.dm_build_409.columns = dm_build_732.dm_build_547(int(dm_build_736), dm_build_733.rsBdta)
	} else {
		dm_build_732.dm_build_409.columns = make([]column, 0)
	}

	return dm_build_733, nil
}

func (dm_build_738 *dm_build_719) dm_build_737(dm_build_739 int) []parameter {

	var dm_build_740, dm_build_741, dm_build_742, dm_build_743 int16

	dm_build_744 := make([]parameter, dm_build_739)
	for i := 0; i < dm_build_739; i++ {

		dm_build_744[i].InitParameter()

		dm_build_744[i].colType = dm_build_738.dm_build_406.dm_build_5.Dm_build_1347()

		dm_build_744[i].prec = dm_build_738.dm_build_406.dm_build_5.Dm_build_1347()

		dm_build_744[i].scale = dm_build_738.dm_build_406.dm_build_5.Dm_build_1347()

		dm_build_744[i].nullable = dm_build_738.dm_build_406.dm_build_5.Dm_build_1347() != 0

		itemFlag := dm_build_738.dm_build_406.dm_build_5.Dm_build_1344()

		if int(itemFlag)&Dm_build_496 != 0 {
			dm_build_744[i].typeFlag = TYPE_FLAG_RECOMMEND
		} else {
			dm_build_744[i].typeFlag = TYPE_FLAG_EXACT
		}

		dm_build_744[i].lob = int(itemFlag)&Dm_build_494 != 0

		dm_build_738.dm_build_406.dm_build_5.Dm_build_1347()

		dm_build_744[i].ioType = byte(dm_build_738.dm_build_406.dm_build_5.Dm_build_1344())

		dm_build_740 = dm_build_738.dm_build_406.dm_build_5.Dm_build_1344()

		dm_build_741 = dm_build_738.dm_build_406.dm_build_5.Dm_build_1344()

		dm_build_742 = dm_build_738.dm_build_406.dm_build_5.Dm_build_1344()

		dm_build_743 = dm_build_738.dm_build_406.dm_build_5.Dm_build_1344()
		dm_build_744[i].name = dm_build_738.dm_build_406.dm_build_5.Dm_build_1384(int(dm_build_740), dm_build_738.dm_build_406.dm_build_6.getServerEncoding(), dm_build_738.dm_build_406.dm_build_6)
		dm_build_744[i].typeName = dm_build_738.dm_build_406.dm_build_5.Dm_build_1384(int(dm_build_741), dm_build_738.dm_build_406.dm_build_6.getServerEncoding(), dm_build_738.dm_build_406.dm_build_6)
		dm_build_744[i].tableName = dm_build_738.dm_build_406.dm_build_5.Dm_build_1384(int(dm_build_742), dm_build_738.dm_build_406.dm_build_6.getServerEncoding(), dm_build_738.dm_build_406.dm_build_6)
		dm_build_744[i].schemaName = dm_build_738.dm_build_406.dm_build_5.Dm_build_1384(int(dm_build_743), dm_build_738.dm_build_406.dm_build_6.getServerEncoding(), dm_build_738.dm_build_406.dm_build_6)

		if dm_build_744[i].lob {
			dm_build_744[i].lobTabId = dm_build_738.dm_build_406.dm_build_5.Dm_build_1347()
			dm_build_744[i].lobColId = dm_build_738.dm_build_406.dm_build_5.Dm_build_1344()
		}
	}

	for i := 0; i < dm_build_739; i++ {

		if isComplexType(int(dm_build_744[i].colType), int(dm_build_744[i].scale)) {

			strDesc := newTypeDescriptor(dm_build_738.dm_build_406.dm_build_6)
			strDesc.unpack(dm_build_738.dm_build_406.dm_build_5)
			dm_build_744[i].typeDescriptor = strDesc
		}
	}

	return dm_build_744
}

const (
	Dm_build_745 = Dm_build_301
)

type dm_build_746 struct {
	dm_build_405
	dm_build_747 int16
	dm_build_748 *Dm_build_1144
	dm_build_749 int32
}

func dm_build_750(dm_build_751 *dm_build_2, dm_build_752 *DmStatement, dm_build_753 int16, dm_build_754 *Dm_build_1144, dm_build_755 int32) *dm_build_746 {
	dm_build_756 := new(dm_build_746)
	dm_build_756.dm_build_414(dm_build_751, Dm_build_283, dm_build_752)
	dm_build_756.dm_build_747 = dm_build_753
	dm_build_756.dm_build_748 = dm_build_754
	dm_build_756.dm_build_749 = dm_build_755
	return dm_build_756
}

func (dm_build_758 *dm_build_746) dm_build_392() error {
	dm_build_758.dm_build_406.dm_build_5.Dm_build_1409(Dm_build_745, dm_build_758.dm_build_747)

	dm_build_758.dm_build_406.dm_build_5.Dm_build_1273(dm_build_758.dm_build_749)

	if dm_build_758.dm_build_406.dm_build_6.NewLobFlag {
		dm_build_758.dm_build_406.dm_build_5.Dm_build_1273(-1)
	}
	dm_build_758.dm_build_748.Dm_build_1151(dm_build_758.dm_build_406.dm_build_5, int(dm_build_758.dm_build_749))
	return nil
}

type dm_build_759 struct {
	dm_build_405
}

func dm_build_760(dm_build_761 *dm_build_2) *dm_build_759 {
	dm_build_762 := new(dm_build_759)
	dm_build_762.dm_build_410(dm_build_761, Dm_build_281)
	return dm_build_762
}

type dm_build_763 struct {
	dm_build_405
	dm_build_764 int32
}

func dm_build_765(dm_build_766 *dm_build_2, dm_build_767 int32) *dm_build_763 {
	dm_build_768 := new(dm_build_763)
	dm_build_768.dm_build_410(dm_build_766, Dm_build_294)
	dm_build_768.dm_build_764 = dm_build_767
	return dm_build_768
}

func (dm_build_770 *dm_build_763) dm_build_392() error {

	dm_build_771 := Dm_build_301
	dm_build_771 += dm_build_770.dm_build_406.dm_build_5.Dm_build_1413(dm_build_771, g2dbIsoLevel(dm_build_770.dm_build_764))
	return nil
}

type dm_build_772 struct {
	dm_build_405
	dm_build_773 *lob
	dm_build_774 byte
	dm_build_775 int
	dm_build_776 []byte
	dm_build_777 int
	dm_build_778 int
}

func dm_build_779(dm_build_780 *dm_build_2, dm_build_781 *lob, dm_build_782 byte, dm_build_783 int, dm_build_784 []byte,
	dm_build_785 int, dm_build_786 int) *dm_build_772 {
	dm_build_787 := new(dm_build_772)
	dm_build_787.dm_build_410(dm_build_780, Dm_build_290)
	dm_build_787.dm_build_773 = dm_build_781
	dm_build_787.dm_build_774 = dm_build_782
	dm_build_787.dm_build_775 = dm_build_783
	dm_build_787.dm_build_776 = dm_build_784
	dm_build_787.dm_build_777 = dm_build_785
	dm_build_787.dm_build_778 = dm_build_786
	return dm_build_787
}

func (dm_build_789 *dm_build_772) dm_build_392() error {

	dm_build_789.dm_build_406.dm_build_5.Dm_build_1265(byte(dm_build_789.dm_build_773.lobFlag))
	dm_build_789.dm_build_406.dm_build_5.Dm_build_1265(dm_build_789.dm_build_774)
	dm_build_789.dm_build_406.dm_build_5.Dm_build_1289(uint64(dm_build_789.dm_build_773.blobId))
	dm_build_789.dm_build_406.dm_build_5.Dm_build_1269(dm_build_789.dm_build_773.groupId)
	dm_build_789.dm_build_406.dm_build_5.Dm_build_1269(dm_build_789.dm_build_773.fileId)
	dm_build_789.dm_build_406.dm_build_5.Dm_build_1273(dm_build_789.dm_build_773.pageNo)
	dm_build_789.dm_build_406.dm_build_5.Dm_build_1269(dm_build_789.dm_build_773.curFileId)
	dm_build_789.dm_build_406.dm_build_5.Dm_build_1273(dm_build_789.dm_build_773.curPageNo)
	dm_build_789.dm_build_406.dm_build_5.Dm_build_1273(dm_build_789.dm_build_773.totalOffset)
	dm_build_789.dm_build_406.dm_build_5.Dm_build_1273(dm_build_789.dm_build_773.tabId)
	dm_build_789.dm_build_406.dm_build_5.Dm_build_1269(dm_build_789.dm_build_773.colId)
	dm_build_789.dm_build_406.dm_build_5.Dm_build_1289(uint64(dm_build_789.dm_build_773.rowId))

	dm_build_789.dm_build_406.dm_build_5.Dm_build_1273(int32(dm_build_789.dm_build_775))
	dm_build_789.dm_build_406.dm_build_5.Dm_build_1273(int32(dm_build_789.dm_build_778))
	dm_build_789.dm_build_406.dm_build_5.Dm_build_1301(dm_build_789.dm_build_776)

	if dm_build_789.dm_build_406.dm_build_6.NewLobFlag {
		dm_build_789.dm_build_406.dm_build_5.Dm_build_1269(dm_build_789.dm_build_773.exGroupId)
		dm_build_789.dm_build_406.dm_build_5.Dm_build_1269(dm_build_789.dm_build_773.exFileId)
		dm_build_789.dm_build_406.dm_build_5.Dm_build_1273(dm_build_789.dm_build_773.exPageNo)
	}
	return nil
}

func (dm_build_791 *dm_build_772) dm_build_396() (interface{}, error) {

	var dm_build_792 = dm_build_791.dm_build_406.dm_build_5.Dm_build_1347()
	dm_build_791.dm_build_773.blobId = dm_build_791.dm_build_406.dm_build_5.Dm_build_1350()
	dm_build_791.dm_build_773.fileId = dm_build_791.dm_build_406.dm_build_5.Dm_build_1344()
	dm_build_791.dm_build_773.pageNo = dm_build_791.dm_build_406.dm_build_5.Dm_build_1347()
	dm_build_791.dm_build_773.curFileId = dm_build_791.dm_build_406.dm_build_5.Dm_build_1344()
	dm_build_791.dm_build_773.curPageNo = dm_build_791.dm_build_406.dm_build_5.Dm_build_1347()
	dm_build_791.dm_build_773.totalOffset = dm_build_791.dm_build_406.dm_build_5.Dm_build_1347()
	return dm_build_792, nil
}

const (
	Dm_build_793 = Dm_build_301

	Dm_build_794 = Dm_build_793 + ULINT_SIZE

	Dm_build_795 = Dm_build_794 + ULINT_SIZE

	Dm_build_796 = Dm_build_795 + BYTE_SIZE

	Dm_build_797 = Dm_build_796 + BYTE_SIZE

	Dm_build_798 = Dm_build_797 + BYTE_SIZE

	Dm_build_799 = Dm_build_798 + BYTE_SIZE

	Dm_build_800 = Dm_build_301

	Dm_build_801 = Dm_build_800 + ULINT_SIZE

	Dm_build_802 = Dm_build_801 + ULINT_SIZE

	Dm_build_803 = Dm_build_802 + ULINT_SIZE

	Dm_build_804 = Dm_build_803 + ULINT_SIZE

	Dm_build_805 = Dm_build_804 + ULINT_SIZE

	Dm_build_806 = Dm_build_805 + BYTE_SIZE

	Dm_build_807 = Dm_build_806 + BYTE_SIZE

	Dm_build_808 = Dm_build_807 + BYTE_SIZE

	Dm_build_809 = Dm_build_808 + ULINT_SIZE
)

type dm_build_810 struct {
	dm_build_405
	dm_build_811 *DmConnection
	dm_build_812 int
	Dm_build_813 int32
	Dm_build_814 []byte
	dm_build_815 byte
}

func dm_build_816(dm_build_817 *dm_build_2) *dm_build_810 {
	dm_build_818 := new(dm_build_810)
	dm_build_818.dm_build_410(dm_build_817, Dm_build_299)
	dm_build_818.dm_build_811 = dm_build_817.dm_build_6
	return dm_build_818
}

func dm_build_819(dm_build_820 string, dm_build_821 string) int {
	dm_build_822 := strings.Split(dm_build_820, ".")
	dm_build_823 := strings.Split(dm_build_821, ".")

	for i, serStr := range dm_build_822 {
		ser, _ := strconv.ParseInt(serStr, 10, 32)
		global, _ := strconv.ParseInt(dm_build_823[i], 10, 32)
		if ser < global {
			return -1
		} else if ser == global {
			continue
		} else {
			return 1
		}
	}

	return 0
}

func (dm_build_825 *dm_build_810) dm_build_392() error {

	dm_build_825.dm_build_406.dm_build_5.Dm_build_1413(Dm_build_793, int32(0))
	dm_build_825.dm_build_406.dm_build_5.Dm_build_1413(Dm_build_794, int32(dm_build_825.dm_build_811.dmConnector.compress))

	if dm_build_825.dm_build_811.dmConnector.loginEncrypt {
		dm_build_825.dm_build_406.dm_build_5.Dm_build_1405(Dm_build_796, 2)
		dm_build_825.dm_build_406.dm_build_5.Dm_build_1405(Dm_build_795, 1)
	} else {
		dm_build_825.dm_build_406.dm_build_5.Dm_build_1405(Dm_build_796, 0)
		dm_build_825.dm_build_406.dm_build_5.Dm_build_1405(Dm_build_795, 0)
	}

	if dm_build_825.dm_build_811.dmConnector.isBdtaRS {
		dm_build_825.dm_build_406.dm_build_5.Dm_build_1405(Dm_build_797, Dm_build_380)
	} else {
		dm_build_825.dm_build_406.dm_build_5.Dm_build_1405(Dm_build_797, Dm_build_379)
	}

	dm_build_825.dm_build_406.dm_build_5.Dm_build_1405(Dm_build_798, byte(dm_build_825.dm_build_811.dmConnector.compressID))

	if dm_build_825.dm_build_811.dmConnector.loginCertificate != "" {
		dm_build_825.dm_build_406.dm_build_5.Dm_build_1405(Dm_build_799, 1)
	} else {
		dm_build_825.dm_build_406.dm_build_5.Dm_build_1405(Dm_build_799, 0)
	}

	dm_build_826 := dm_build_825.dm_build_811.getServerEncoding()
	dm_build_825.dm_build_406.dm_build_5.Dm_build_1317(Dm_build_266, dm_build_826, dm_build_825.dm_build_406.dm_build_6)

	var dm_build_827 byte
	if dm_build_825.dm_build_811.dmConnector.uKeyName != "" {
		dm_build_827 = 1
	} else {
		dm_build_827 = 0
	}

	dm_build_825.dm_build_406.dm_build_5.Dm_build_1265(0)

	if dm_build_827 == 1 {

	}

	if dm_build_825.dm_build_811.dmConnector.loginEncrypt {
		clientPubKey, err := dm_build_825.dm_build_406.dm_build_246()
		if err != nil {
			return err
		}
		dm_build_825.dm_build_406.dm_build_5.Dm_build_1305(clientPubKey)
	}

	return nil
}

func (dm_build_829 *dm_build_810) dm_build_395() error {
	dm_build_829.dm_build_406.dm_build_5.Dm_build_1239(0)
	dm_build_829.dm_build_406.dm_build_5.Dm_build_1247(Dm_build_300, false, true)
	return nil
}

func (dm_build_831 *dm_build_810) dm_build_396() (interface{}, error) {

	dm_build_831.dm_build_811.sslEncrypt = int(dm_build_831.dm_build_406.dm_build_5.Dm_build_1491(Dm_build_800))
	dm_build_831.dm_build_811.GlobalServerSeries = int(dm_build_831.dm_build_406.dm_build_5.Dm_build_1491(Dm_build_801))

	switch dm_build_831.dm_build_406.dm_build_5.Dm_build_1491(Dm_build_802) {
	case 1:
		dm_build_831.dm_build_811.serverEncoding = ENCODING_UTF8
	case 2:
		dm_build_831.dm_build_811.serverEncoding = ENCODING_EUCKR
	default:
		dm_build_831.dm_build_811.serverEncoding = ENCODING_GB18030
	}

	dm_build_831.dm_build_811.dmConnector.compress = int(dm_build_831.dm_build_406.dm_build_5.Dm_build_1491(Dm_build_803))
	dm_build_832 := dm_build_831.dm_build_406.dm_build_5.Dm_build_1485(Dm_build_805)
	dm_build_833 := dm_build_831.dm_build_406.dm_build_5.Dm_build_1485(Dm_build_806)
	dm_build_831.dm_build_811.dmConnector.isBdtaRS = dm_build_831.dm_build_406.dm_build_5.Dm_build_1485(Dm_build_807) > 0
	dm_build_831.dm_build_811.dmConnector.compressID = int8(dm_build_831.dm_build_406.dm_build_5.Dm_build_1485(Dm_build_808))
	dm_build_831.dm_build_811.dmConnector.newClientType = dm_build_831.dm_build_406.dm_build_5.Dm_build_1485(Dm_build_809) == 1

	dm_build_834 := dm_build_831.dm_build_438()
	if dm_build_834 != nil {
		return nil, dm_build_834
	}

	dm_build_835 := dm_build_831.dm_build_406.dm_build_5.Dm_build_1389(dm_build_831.dm_build_811.getServerEncoding(), dm_build_831.dm_build_406.dm_build_6)
	if dm_build_819(dm_build_835, Dm_build_267) < 0 {
		return nil, ECGO_ERROR_SERVER_VERSION.throw()
	}

	dm_build_831.dm_build_811.ServerVersion = dm_build_835
	dm_build_831.dm_build_811.Malini2 = dm_build_819(dm_build_835, Dm_build_268) > 0
	dm_build_831.dm_build_811.Execute2 = dm_build_819(dm_build_835, Dm_build_269) > 0
	dm_build_831.dm_build_811.LobEmptyCompOrcl = dm_build_819(dm_build_835, Dm_build_270) > 0

	if dm_build_831.dm_build_406.dm_build_6.dmConnector.uKeyName != "" {
		dm_build_831.dm_build_815 = 1
	} else {
		dm_build_831.dm_build_815 = 0
	}

	if dm_build_831.dm_build_815 == 1 {
		dm_build_831.dm_build_406.dm_build_11 = dm_build_831.dm_build_406.dm_build_5.Dm_build_1384(16, dm_build_831.dm_build_811.getServerEncoding(), dm_build_831.dm_build_406.dm_build_6)
	}

	dm_build_831.dm_build_812 = -1
	dm_build_836 := false
	dm_build_837 := false
	dm_build_831.Dm_build_813 = -1
	if dm_build_833 > 0 {
		dm_build_831.dm_build_812 = int(dm_build_831.dm_build_406.dm_build_5.Dm_build_1347())
	}

	if dm_build_832 > 0 {

		if dm_build_831.dm_build_812 == -1 {
			dm_build_836 = true
		} else {
			dm_build_837 = true
		}

		dm_build_831.Dm_build_814 = dm_build_831.dm_build_406.dm_build_5.Dm_build_1372()
	}

	if dm_build_833 == 2 {
		dm_build_831.Dm_build_813 = dm_build_831.dm_build_406.dm_build_5.Dm_build_1347()
	}
	dm_build_831.dm_build_406.dm_build_8 = dm_build_836
	dm_build_831.dm_build_406.dm_build_9 = dm_build_837

	return nil, nil
}

type dm_build_838 struct {
	dm_build_405
}

func dm_build_839(dm_build_840 *dm_build_2, dm_build_841 *DmStatement) *dm_build_838 {
	dm_build_842 := new(dm_build_838)
	dm_build_842.dm_build_414(dm_build_840, Dm_build_275, dm_build_841)
	return dm_build_842
}

func (dm_build_844 *dm_build_838) dm_build_392() error {

	dm_build_844.dm_build_406.dm_build_5.Dm_build_1405(Dm_build_301, 1)
	return nil
}

func (dm_build_846 *dm_build_838) dm_build_396() (interface{}, error) {

	dm_build_846.dm_build_409.id = dm_build_846.dm_build_406.dm_build_5.Dm_build_1491(Dm_build_302)

	dm_build_846.dm_build_409.readBaseColName = dm_build_846.dm_build_406.dm_build_5.Dm_build_1485(Dm_build_301) == 1
	return nil, nil
}

type dm_build_847 struct {
	dm_build_405
	dm_build_848 int32
}

func dm_build_849(dm_build_850 *dm_build_2, dm_build_851 int32) *dm_build_847 {
	dm_build_852 := new(dm_build_847)
	dm_build_852.dm_build_410(dm_build_850, Dm_build_276)
	dm_build_852.dm_build_848 = dm_build_851
	return dm_build_852
}

func (dm_build_854 *dm_build_847) dm_build_393() {
	dm_build_854.dm_build_405.dm_build_393()
	dm_build_854.dm_build_406.dm_build_5.Dm_build_1413(Dm_build_302, dm_build_854.dm_build_848)
}

type dm_build_855 struct {
	dm_build_405
	dm_build_856 []uint32
}

func dm_build_857(dm_build_858 *dm_build_2, dm_build_859 []uint32) *dm_build_855 {
	dm_build_860 := new(dm_build_855)
	dm_build_860.dm_build_410(dm_build_858, Dm_build_296)
	dm_build_860.dm_build_856 = dm_build_859
	return dm_build_860
}

func (dm_build_862 *dm_build_855) dm_build_392() error {

	dm_build_862.dm_build_406.dm_build_5.Dm_build_1433(Dm_build_301, uint16(len(dm_build_862.dm_build_856)))

	for _, tableID := range dm_build_862.dm_build_856 {
		dm_build_862.dm_build_406.dm_build_5.Dm_build_1285(uint32(tableID))
	}

	return nil
}

func (dm_build_864 *dm_build_855) dm_build_396() (interface{}, error) {
	dm_build_865 := dm_build_864.dm_build_406.dm_build_5.Dm_build_1506(Dm_build_301)
	if dm_build_865 <= 0 {
		return nil, nil
	}

	dm_build_866 := make([]int64, dm_build_865)
	for i := 0; i < int(dm_build_865); i++ {
		dm_build_866[i] = dm_build_864.dm_build_406.dm_build_5.Dm_build_1350()
	}

	return dm_build_866, nil
}
