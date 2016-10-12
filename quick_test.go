package quick_test

import ("goblin"
	"testing"
        "testing/quick"
	"fmt"
	"math"
	"strconv"
)

// TODO: install github.com/stretchr/testify
// and make these unit tests use something other than just ==


func TestRoundTripFloat(t *testing.T) {
	f := func(flt float64) bool {
		flt = math.Abs(flt)
		needed := fmt.Sprintf("%f", flt)
		gotten := goblin.TestExpr(needed)
		result, _ := strconv.ParseFloat(gotten["value"].(string), 64)
		return flt == result
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestIota(t *testing.T) {
	gotten := goblin.TestExpr("iota")
	if gotten["token-kind"] != "IOTA" {
		t.Error("Didn't parse iota as a literal")
	}

}

func TestTrue(t *testing.T) {
	gotten := goblin.TestExpr("true")
	if gotten["token-kind"] != "BOOL" || gotten["value"] != "true" {
		t.Error("Didn't parse 'true' as true")
	}
}

func TestFalse(t *testing.T) {
	gotten := goblin.TestExpr("false")
	if gotten["token-kind"] != "BOOL" || gotten["value"] != "false" {
		t.Error("Didn't parse 'false' as false")
	}
}


func TestProvidedFloat(t *testing.T) {
	gotten := goblin.TestExpr("3.14")
	if gotten["value"].(string) != "3.14" {
		t.Error("Floats not parsing correctly")
	}
}


func TestRoundTripUInt(t *testing.T) {
	f := func(int uint64) bool {
		needed := fmt.Sprintf("%d", int)
		gotten := goblin.TestExpr(needed)
		return needed == gotten["value"]
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
