package mix

func NullConfigureOutput(s AudioSpec) {
	go func() {
		for {
			OutNextBytes()
		}
	}()
	// nothing to do
}
