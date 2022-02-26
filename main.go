package main

var (
	version = "dev"
	helpMsg = `whack - a better process killer
						
	Usage
		$ whack <pid> <name> <port>

	Options
		--force								Force kill a process
		--verbose							Show process arguments

	Examples
	$ whack 1445
	$ whack discord
	$ whack :3000

	To terminate a specific port, add a colon. e.g: :3000.`
)
