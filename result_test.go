package result

import "testing"

func TestVerification(t *testing.T) {
	r := Ok(10)

	if r.IsError() {
		t.Errorf("It must be an Okay result")
	}

	if !r.IsOkay() {
		t.Error("It must be an Okay result")
	}

	r = Err(10)

	if !r.IsError() {
		t.Errorf("It must be an Err result")
	}

	if r.IsOkay() {
		t.Error("It must be an Err result")
	}
}

func TestAnd(t *testing.T) {
	var v interface{}
	var expected string

	var x Result = Ok(2)
	var y Result = Err("Error")

	expected = "Error"
	v = x.And(y).Value()
	if v != expected {
		t.Errorf("Expected %v, got %v\n", expected, v)
	}

	x = Err("Error")
	y = Ok("Okay")
	expected = "Error"
	v = x.And(y).Value()
	if v != expected {
		t.Errorf("Expected %v, got %v\n", expected, v)
	}

	x = Err("Error 1")
	y = Err("Error 2")
	expected = "Error 1"
	v = x.And(y).Value()
	if v != expected {
		t.Errorf("Expected %v, got %v\n", expected, v)
	}

	x = Ok("Ok 1")
	y = Ok("Ok 2")
	expected = "Ok 2"
	v = x.And(y).Value()
	if v != expected {
		t.Errorf("Expected %v, got %v\n", expected, v)
	}
}

func TestOr(t *testing.T) {
	var v interface{}
	var expected interface{}

	var x Result = Ok(20)
	var y Result = Err("Error")

	v = x.Or(y).Value()
	expected = 20
	if v != expected {
		t.Errorf("Expected %v, got %v\n", expected, v)
	}

	x = Err("Error")
	y = Ok(20)
	expected = 20
	v = x.Or(y).Value()
	if v != expected {
		t.Errorf("Expected %v, got %v\n", expected, v)
	}

	x = Err("Error 1")
	y = Err("Error 2")
	expected = "Error 2"
	v = x.Or(y).Value()
	if v != expected {
		t.Errorf("Expected %v, got %v\n", expected, v)
	}

	x = Ok(10)
	y = Ok(20)
	expected = 10
	v = x.Or(y).Value()
	if v != expected {
		t.Errorf("Expected %v, got %v\n", expected, v)
	}
}

func TestAndThen(t *testing.T) {
	var v interface{}
	var expected int
	var okFn = func(v interface{}) Result { return Ok(v.(int) * v.(int)) }
	var errFn = func(v interface{}) Result { return Err(v) }

	v = Ok(2).AndThen(okFn).AndThen(okFn).Value()
	expected = 16
	if v.(int) != expected {
		t.Errorf("Expected %d, got %v", expected, v)
	}

	v = Ok(2).AndThen(okFn).AndThen(errFn).Value()
	expected = 4
	if v.(int) != expected {
		t.Errorf("Expected %d, got %v", expected, v)
	}

	v = Ok(2).AndThen(errFn).AndThen(okFn).Value()
	expected = 2
	if v.(int) != expected {
		t.Errorf("Expected %d, got %v", expected, v)
	}

	v = Err(3).AndThen(okFn).AndThen(okFn).Value()
	expected = 3
	if v.(int) != expected {
		t.Errorf("Expected %d, got %v", expected, v)
	}
}

func TestOrElse(t *testing.T) {
	var v interface{}
	var expected int
	var okFn = func(v interface{}) Result { return Ok(v.(int) * v.(int)) }
	var errFn = func(v interface{}) Result { return Err(v) }

	v = Ok(2).OrElse(okFn).OrElse(okFn).Value()
	expected = 2
	if v.(int) != expected {
		t.Errorf("Expected %d, got %v", expected, v)
	}

	v = Ok(2).OrElse(errFn).OrElse(okFn).Value()
	expected = 2
	if v.(int) != expected {
		t.Errorf("Expected %d, got %v", 2, v)
	}

	v = Err(3).OrElse(okFn).OrElse(errFn).Value()
	expected = 9
	if v.(int) != expected {
		t.Errorf("Expected %d, got %v", expected, v)
	}

	v = Err(3).OrElse(errFn).OrElse(errFn).Value()
	expected = 3
	if v != expected {
		t.Errorf("Expected %v, got %v", expected, v)
	}
}
