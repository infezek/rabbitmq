package pkg

func init() {
	RootCmd.AddCommand(
		createQueueCmd,
		consumerQueueCmd,
		publishQueueCmd(),
		deleteQueueCmd(),
	)
}
