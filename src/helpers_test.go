// These tests verify the inner workings of the helper methods associated
// with gocheck.T.

package gocheck_test

import (
    "gocheck"
    "os"
)


var helpersS = gocheck.Suite(&HelpersS{})

type HelpersS struct{}

func (s *HelpersS) TestCountSuite(t *gocheck.T) {
    suitesRun += 1
}

func (s *HelpersS) TestCheckEqualSucceeding(t *gocheck.T) {
    testHelperSuccess(t, "CheckEqual(10, 10)", true, func() interface{} {
        return t.CheckEqual(10, 10)
    })
}

func (s *HelpersS) TestCheckEqualFailing(t *gocheck.T) {
    log := "helpers_test.go:[0-9]+:\n" +
           "\\.+ CheckEqual\\(obtained, expected\\):\n" +
           "\\.+ Obtained \\(int\\): 10\n" +
           "\\.+ Expected \\(int\\): 20\n\n"
    testHelperFailure(t, "CheckEqual(10, 20)", false, false, log,
                      func() interface{} {
        return t.CheckEqual(10, 20)
    })
}

func (s *HelpersS) TestCheckEqualFailingWithDiffTypes(t *gocheck.T) {
    log := "helpers_test.go:[0-9]+:\n" +
           "\\.+ CheckEqual\\(obtained, expected\\):\n" +
           "\\.+ Obtained \\(int\\): 10\n" +
           "\\.+ Expected \\(uint\\): 0xa\n\n"
    testHelperFailure(t, "CheckEqual(10, uint(10))", false, false, log,
                      func() interface{} {
        return t.CheckEqual(10, uint(10))
    })
}

func (s *HelpersS) TestCheckEqualFailingWithNil(t *gocheck.T) {
    log := "helpers_test.go:[0-9]+:\n" +
           "\\.+ CheckEqual\\(obtained, expected\\):\n" +
           "\\.+ Obtained \\(int\\): 10\n" +
           "\\.+ Expected \\(nil\\): nil\n\n"
    testHelperFailure(t, "CheckEqual(10, nil)", false, false, log,
                      func() interface{} {
        return t.CheckEqual(10, nil)
    })
}

func (s *HelpersS) TestCheckEqualWithMessage(t *gocheck.T) {
    log := "helpers_test.go:[0-9]+:\n" +
           "\\.+ CheckEqual\\(obtained, expected\\):\n" +
           "\\.+ Obtained \\(int\\): 10\n" +
           "\\.+ Expected \\(int\\): 20\n" +
           "\\.+ That's clearly WRONG!\n\n"
    testHelperFailure(t, "CheckEqual(10, 20, issue)", false, false, log,
                      func() interface{} {
        return t.CheckEqual(10, 20, "That's clearly ", "WRONG!")
    })
}

func (s *HelpersS) TestCheckNotEqualSucceeding(t *gocheck.T) {
    testHelperSuccess(t, "CheckNotEqual(10, 20)", true, func() interface{} {
        return t.CheckNotEqual(10, 20)
    })
}

func (s *HelpersS) TestCheckNotEqualSucceedingWithNil(t *gocheck.T) {
    testHelperSuccess(t, "CheckNotEqual(10, nil)", true, func() interface{} {
        return t.CheckNotEqual(10, nil)
    })
}

func (s *HelpersS) TestCheckNotEqualFailing(t *gocheck.T) {
    log := "helpers_test.go:[0-9]+:\n" +
           "\\.+ CheckNotEqual\\(obtained, unexpected\\):\n" +
           "\\.+ Both \\(int\\): 10\n\n"
    testHelperFailure(t, "CheckNotEqual(10, 10)", false, false, log,
                      func() interface{} {
        return t.CheckNotEqual(10, 10)
    })
}

func (s *HelpersS) TestCheckNotEqualWithMessage(t *gocheck.T) {
    log := "helpers_test.go:[0-9]+:\n" +
           "\\.+ CheckNotEqual\\(obtained, unexpected\\):\n" +
           "\\.+ Both \\(int\\): 10\n" +
           "\\.+ That's clearly WRONG!\n\n"
    testHelperFailure(t, "CheckNotEqual(10, 10, issue)", false, false, log,
                      func() interface{} {
        return t.CheckNotEqual(10, 10, "That's clearly ", "WRONG!")
    })
}

func (s *HelpersS) TestAssertEqualSucceeding(t *gocheck.T) {
    testHelperSuccess(t, "AssertEqual(10, 10)", nil, func() interface{} {
        t.AssertEqual(10, 10)
        return nil
    })
}

func (s *HelpersS) TestAssertEqualFailing(t *gocheck.T) {
    log := "helpers_test.go:[0-9]+:\n" +
           "\\.+ AssertEqual\\(obtained, expected\\):\n" +
           "\\.+ Obtained \\(int\\): 10\n" +
           "\\.+ Expected \\(int\\): 20\n\n"
    testHelperFailure(t, "AssertEqual(10, 20)", nil, true, log,
                      func() interface{} {
        t.AssertEqual(10, 20)
        return nil
    })
}

func (s *HelpersS) TestAssertEqualWithMessage(t *gocheck.T) {
    log := "helpers_test.go:[0-9]+:\n" +
           "\\.+ AssertEqual\\(obtained, expected\\):\n" +
           "\\.+ Obtained \\(int\\): 10\n" +
           "\\.+ Expected \\(int\\): 20\n" +
           "\\.+ That's clearly WRONG!\n\n"
    testHelperFailure(t, "AssertEqual(10, 20, issue)", nil, true, log,
                      func() interface{} {
        t.AssertEqual(10, 20, "That's clearly ", "WRONG!")
        return nil
    })
}

func (s *HelpersS) TestAssertNotEqualSucceeding(t *gocheck.T) {
    testHelperSuccess(t, "AssertNotEqual(10, 20)", nil, func() interface{} {
        t.AssertNotEqual(10, 20)
        return nil
    })
}

func (s *HelpersS) TestAssertNotEqualFailing(t *gocheck.T) {
    log := "helpers_test.go:[0-9]+:\n" +
           "\\.+ AssertNotEqual\\(obtained, unexpected\\):\n" +
           "\\.+ Both \\(int\\): 10\n\n"
    testHelperFailure(t, "AssertNotEqual(10, 10)", nil, true, log,
                      func() interface{} {
        t.AssertNotEqual(10, 10)
        return nil
    })
}

func (s *HelpersS) TestAssertNotEqualWithMessage(t *gocheck.T) {
    log := "helpers_test.go:[0-9]+:\n" +
           "\\.+ AssertNotEqual\\(obtained, unexpected\\):\n" +
           "\\.+ Both \\(int\\): 10\n" +
           "\\.+ That's clearly WRONG!\n\n"
    testHelperFailure(t, "AssertNotEqual(10, 10, issue)", nil, true, log,
                      func() interface{} {
        t.AssertNotEqual(10, 10, "That's clearly ", "WRONG!")
        return nil
    })
}


func (s *HelpersS) TestCheckEqualArraySucceeding(t *gocheck.T) {
    testHelperSuccess(t, "CheckEqual([]byte, []byte)", true, func() interface{} {
        return t.CheckEqual([]byte{1,2}, []byte{1,2})
    })
}

func (s *HelpersS) TestCheckEqualArrayFailing(t *gocheck.T) {
    log := "helpers_test.go:[0-9]+:\n" +
           "\\.+ CheckEqual\\(obtained, expected\\):\n" +
           "\\.+ Obtained \\(\\[\\]uint8\\): \\[\\]byte{0x1, 0x2}\n" +
           "\\.+ Expected \\(\\[\\]uint8\\): \\[\\]byte{0x1, 0x3}\n\n"
    testHelperFailure(t, "CheckEqual([]byte{2}, []byte{3})", false, false, log,
                      func() interface{} {
        return t.CheckEqual([]byte{1,2}, []byte{1,3})
    })
}

func (s *HelpersS) TestCheckNotEqualArraySucceeding(t *gocheck.T) {
    testHelperSuccess(t, "CheckNotEqual([]byte, []byte)", true,
                      func() interface{} {
        return t.CheckNotEqual([]byte{1,2}, []byte{1,3})
    })
}

func (s *HelpersS) TestCheckNotEqualArrayFailing(t *gocheck.T) {
    log := "helpers_test.go:[0-9]+:\n" +
           "\\.+ CheckNotEqual\\(obtained, unexpected\\):\n" +
           "\\.+ Both \\(\\[\\]uint8\\): \\[\\]byte{0x1, 0x2}\n\n"
    testHelperFailure(t, "CheckNotEqual([]byte{2}, []byte{3})", false, false,
                      log, func() interface{} {
        return t.CheckNotEqual([]byte{1,2}, []byte{1,2})
    })
}


func (s *HelpersS) TestAssertEqualArraySucceeding(t *gocheck.T) {
    testHelperSuccess(t, "AssertEqual([]byte, []byte)", nil,
                      func() interface{} {
        t.AssertEqual([]byte{1,2}, []byte{1,2})
        return nil
    })
}

func (s *HelpersS) TestAssertEqualArrayFailing(t *gocheck.T) {
    log := "helpers_test.go:[0-9]+:\n" +
           "\\.+ AssertEqual\\(obtained, expected\\):\n" +
           "\\.+ Obtained \\(\\[\\]uint8\\): \\[\\]byte{0x1, 0x2}\n" +
           "\\.+ Expected \\(\\[\\]uint8\\): \\[\\]byte{0x1, 0x3}\n\n"
    testHelperFailure(t, "AssertEqual([]byte{2}, []byte{3})", nil, true, log,
                      func() interface{} {
        t.AssertEqual([]byte{1,2}, []byte{1,3})
        return nil
    })
}

func (s *HelpersS) TestAssertNotEqualArraySucceeding(t *gocheck.T) {
    testHelperSuccess(t, "AssertNotEqual([]byte, []byte)", nil,
                      func() interface{} {
        t.AssertNotEqual([]byte{1,2}, []byte{1,3})
        return nil
    })
}

func (s *HelpersS) TestAssertNotEqualArrayFailing(t *gocheck.T) {
    log := "helpers_test.go:[0-9]+:\n" +
           "\\.+ AssertNotEqual\\(obtained, unexpected\\):\n" +
           "\\.+ Both \\(\\[\\]uint8\\): \\[\\]byte{0x1, 0x2}\n\n"
    testHelperFailure(t, "AssertNotEqual([]byte{2}, []byte{3})", nil, true,
                      log, func() interface{} {
        t.AssertNotEqual([]byte{1,2}, []byte{1,2})
        return nil
    })
}

func (s *HelpersS) TestCheckErrSucceeding(t *gocheck.T) {
    testHelperSuccess(t, "CheckErr(nil, nil)", true, func() interface{} {
        return t.CheckErr(nil, nil)
    })
}

func (s *HelpersS) TestCheckErrFailing(t *gocheck.T) {
    log := "helpers_test.go:[0-9]+:\n" +
           "\\.+ CheckErr\\(error, nil\\):\n" +
           "\\.+ Error \\(os\\.Errno\\): 13 \\(permission denied\\)\n\n"
    testHelperFailure(t, "CheckErr(error, nil)", false, false, log,
                      func() interface{} {
        return t.CheckErr(os.Errno(13), nil)
    })
}

func (s *HelpersS) TestAssertErrSucceeding(t *gocheck.T) {
    testHelperSuccess(t, "AssertErr(nil, nil)", nil, func() interface{} {
        t.AssertErr(nil, nil)
        return nil
    })
}

func (s *HelpersS) TestAssertErrSucceedingWithValue(t *gocheck.T) {
    testHelperSuccess(t, "AssertErr(nil, nil)", nil, func() interface{} {
        t.AssertErr(os.Errno(13), os.Errno(13))
        return nil
    })
}

func (s *HelpersS) TestAssertErrSucceedingWithStrMatch(t *gocheck.T) {
    testHelperSuccess(t, "AssertErr(nil, nil)", nil, func() interface{} {
        t.AssertErr(os.Errno(13), "perm.*denied")
        return nil
    })
}

func (s *HelpersS) TestAssertErrSucceedingWithPureStrMatch(t *gocheck.T) {
    testHelperSuccess(t, "AssertErr(nil, nil)", nil, func() interface{} {
        t.AssertErr("str error", "str.*")
        return nil
    })
}

func (s *HelpersS) TestAssertErrFailing(t *gocheck.T) {
    log := "helpers_test.go:[0-9]+:\n" +
           "\\.+ AssertErr\\(error, nil\\):\n" +
           "\\.+ Error \\(os\\.Errno\\): 13 \\(permission denied\\)\n\n"
    testHelperFailure(t, "AssertErr(error, nil)", nil, true, log,
                      func() interface{} {
        t.AssertErr(os.Errno(13), nil)
        return nil
    })
}

func (s *HelpersS) TestAssertErrFailingWithValue(t *gocheck.T) {
    log := "helpers_test.go:[0-9]+:\n" +
           "\\.+ AssertErr\\(error, expected\\):\n" +
           "\\.+ Error \\(os\\.Errno\\): 13 \\(permission denied\\)\n" +
           "\\.+ Expected \\(os\\.Errno\\): 12 \\(cannot allocate.+\\)\n\n"
    testHelperFailure(t, "AssertErr(error, nil)", nil, true, log,
                      func() interface{} {
        t.AssertErr(os.Errno(13), os.Errno(12))
        return nil
    })
}

func (s *HelpersS) TestAssertErrFailingWithStrMatch(t *gocheck.T) {
    log := "helpers_test.go:[0-9]+:\n" +
           "\\.+ AssertErr\\(error, expected\\):\n" +
           "\\.+ Error \\(os\\.Errno\\): 13 \\(permission denied\\)\n" +
           "\\.+ Expected to match expression: \"foobar\"\n\n"
    testHelperFailure(t, "AssertErr(error, nil)", nil, true, log,
                      func() interface{} {
        t.AssertErr(os.Errno(13), "foobar")
        return nil
    })
}

func (s *HelpersS) TestAssertErrFailingWithPureStrMatch(t *gocheck.T) {
    log := "helpers_test.go:[0-9]+:\n" +
           "\\.+ AssertErr\\(error, expected\\):\n" +
           "\\.+ Error \\(string\\): \"foobar\"\n" +
           "\\.+ Expected to match expression: \"foobaz\"\n\n"
    testHelperFailure(t, "AssertErr(error, nil)", nil, true, log,
                      func() interface{} {
        t.AssertErr("foobar", "foobaz")
        return nil
    })
}

func (s *HelpersS) TestAssertErrWithMessage(t *gocheck.T) {
    log := "helpers_test.go:[0-9]+:\n" +
           "\\.+ AssertErr\\(error, nil\\):\n" +
           "\\.+ Error \\(os\\.Errno\\): 13 \\(permission denied\\)\n" +
           "\\.+ That's clearly WRONG!\n\n"
    testHelperFailure(t, "AssertErr(error, nil)", nil, true, log,
                      func() interface{} {
        t.AssertErr(os.Errno(13), nil, "That's clearly ", "WRONG!")
        return nil
    })
}


// -----------------------------------------------------------------------
// MakeDir() tests.

type MkDirHelper struct {
    path1 string
    path2 string
    isDir1 bool
    isDir2 bool
    isDir3 bool
    isDir4 bool
}

func (s *MkDirHelper) SetUpSuite(f *gocheck.F) {
    s.path1 = f.MkDir()
    s.isDir1 = isDir(s.path1)
}

func (s *MkDirHelper) Test(t *gocheck.T) {
    s.path2 = t.MkDir()
    s.isDir2 = isDir(s.path2)
}

func (s *MkDirHelper) TearDownSuite(f *gocheck.F) {
    s.isDir3 = isDir(s.path1)
    s.isDir4 = isDir(s.path2)
}


func (s *HelpersS) TestMkDir(t *gocheck.T) {
    helper := MkDirHelper{}
    output := String{}
    gocheck.RunWithWriter(&helper, &output)
    t.AssertEqual(output.value, "")
    t.CheckEqual(helper.isDir1, true)
    t.CheckEqual(helper.isDir2, true)
    t.CheckEqual(helper.isDir3, true)
    t.CheckEqual(helper.isDir4, true)
    t.CheckNotEqual(helper.path1, helper.path2)
    t.CheckEqual(isDir(helper.path1), false)
    t.CheckEqual(isDir(helper.path2), false)
}

func isDir(path string) bool {
    if stat, err := os.Stat(path); err == nil {
        return stat.IsDirectory()
    }
    return false
}


// -----------------------------------------------------------------------
// A couple of helper functions to test helper functions. :-)

func testHelperSuccess(t *gocheck.T, name string,
                       expectedResult interface{},
                       closure func() interface{}) {
    var result interface{}
    defer (func() {
        if err := recover(); err != nil {
            panic(err)
        }
        checkState(t, result,
                   &expectedState{
                        name: name,
                        result: expectedResult,
                        failed: false,
                        log: "",
                   })
    })()
    result = closure()
}

func testHelperFailure(t *gocheck.T, name string,
                       expectedResult interface{},
                       shouldStop bool, log string,
                       closure func() interface{}) {
    var result interface{}
    defer (func() {
        if err := recover(); err != nil {
            panic(err)
        }
        checkState(t, result,
                   &expectedState{
                        name: name,
                        result: expectedResult,
                        failed: true,
                        log: log,
                   })
    })()
    result = closure()
    if shouldStop {
        t.Logf("%s didn't stop when it should", name)
    }
}
