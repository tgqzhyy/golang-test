package main

import "errors"

func f1(arg int)(int,error)  {
	if arg ==42{
		return -1,errors.New("can't work with 42")
	}
	return arg +3,nil
}

type argError struct {
	arg int
	prob string
}

func (e *argError) Error() string  {

}
