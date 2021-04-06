package debug_logger

type debugLogger struct{} 

func (*debugLogger) Debugf(f string, args... interface{}) { 
	fmt.Printf(f, args...)
	fmt.Println()
}

func (*debugLogger) Infof(f string, args... interface{}) { 
	fmt.Printf(f, args...) 
	fmt.Println() 
}

func (*debugLogger) Errorf(f string, args... interface{}) { 
	fmt.Printf(f, args...)
	fmt.Println()
}