package debug_logger

import "fmt"

type DebugLogger struct{} 

func (*DebugLogger) Debugf(f string, args... interface{}) { 
	fmt.Printf(f + "\n", args...)
}

func (*DebugLogger) Infof(f string, args... interface{}) { 
	fmt.Printf(f + "\n", args...) 
}

func (*DebugLogger) Errorf(f string, args... interface{}) { 
	fmt.Printf(f + "\n", args...)
}