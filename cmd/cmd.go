package cmd

// Execute executes the commands.
func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}
